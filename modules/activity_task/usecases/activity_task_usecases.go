package usecases

type activityTaskUse struct {
	ActivityTaskRepo entities.ActivityTaskRepository
	AttachmantRepo   entities.AttachmentRepository
}

func NewActivityTaskUsecase(activityTaskRepo entities.ActivityTaskRepository, attachmantRepo entities.AttachmentRepository) entities.ActivityTaskUsecase {
	return &activityTaskUse{
		ActivityTaskRepo: activityTaskRepo,
		AttachmantRepo:   attachmantRepo,
	}
}

func (u *activityTaskUse) ActivityMemberDelete(req *entities.MemberDeleteByRoleReq) error {
	chkOwner, err := utils.CheckUserOwner(req.Email, req.TablePermissionId, "Owner")
	if err != nil {
		return err
	}

	if chkOwner {
		return fmt.Errorf("cannot delete owner")
	}

	for _, role := range req.RoleIds {
		req.RoleId = role
		if err := utils.DeleteMemberByRole(req); err != nil {
			logs.Error(err)
			return err
		}
	}

	return nil
}

func (u *activityTaskUse) ActivityMember(req *entities.ActivityMemberReq) ([]entities.TeamMemberRes, error) {
	return utils.GetTeamMembers(req.ActivityId, req.RoleType)
}

func (u *activityTaskUse) VerifyActivitySubTask(req *entities.ActivitySubTaskReq) ([]entities.ActivitySubTaskRes, error) {
	// Get Sub Task List
	activitySubTaskList, err := u.ActivityTaskRepo.ActivitySubTaskLists(req)
	if err != nil {
		return nil, err
	}

	// Get Sub Task Verify Fail list
	failList, err := u.ActivityTaskRepo.CheckActivitySubTaskDuplicated(activitySubTaskList)
	if err != nil {
		return nil, err
	}

	if failList != nil {
		for i, subTask := range activitySubTaskList {
			verified := "Passed"
			for _, fail := range failList {
				if subTask.SubTaskId == fail {
					verified = "Duplicated"
					break
				}
			}
			activitySubTaskList[i].Verify = verified
		}
		return activitySubTaskList, nil

	} else {
		// Verified PASS
		for i := range activitySubTaskList {
			activitySubTaskList[i].Verify = "Passed"
		}
		return activitySubTaskList, nil

	}
}

func (u *activityTaskUse) ActivityTaskHistoryLists(req *entities.ActivityTaskStatusReq) ([]entities.ActivityTaskHistoryRes, error) {
	return u.ActivityTaskRepo.ActivityTaskHistoryLists(req)
}

func (u *activityTaskUse) ActivityTaskById(TaskId string) (entities.ActivityTaskRes, error) {
	activityTask, err := u.ActivityTaskRepo.ActivityTaskById(TaskId)
	if err != nil {
		return entities.ActivityTaskRes{}, err
	}
	return activityTask, nil
}

func (u *activityTaskUse) ActivitySubTaskLists(req *entities.ActivitySubTaskReq) ([]entities.ActivitySubTaskRes, *string, error) {
	res, err := u.ActivityTaskRepo.ActivitySubTaskLists(req)
	if err != nil {
		return nil, nil, err
	}
	if len(res) == 0 {
		return res, nil, nil
	}

	// Response without reason
	if res[0].Status != constants.RejectedTask {
		return res, nil, nil
	}

	// Response with reason
	hisData, err := u.ActivityTaskRepo.GetLatestTaskHistoryByTaskId(res[0].TaskId)
	if err != nil {
		return nil, nil, err
	}

	return res, hisData.Reason, nil
}

func (u *activityTaskUse) DeleteActivityTask(req *entities.DeleteTaskReq) error {
	if err := u.ActivityTaskRepo.DeleteActivityAllSubTaskByTaskId(req); err != nil {
		return err
	}

	return u.ActivityTaskRepo.DeleteActivityTask(req)
}

func (u *activityTaskUse) SubmitActivityTask(req *entities.SubmitTaskReq) error {
	// check all is valid
	taskFound, err := u.ActivityTaskRepo.ActivityTaskById(req.TaskId)
	if err != nil {
		return err
	}
	if taskFound.TransactionQty == 0 {
		logs.Error("task_has_no_transaction")
		return fmt.Errorf("task_has_no_transaction")
	}

	// For re-submit when checker rejected
	if taskFound.SubmittedDate != nil && taskFound.Status != constants.RejectedTask {
		logs.Error("task_already_submitted")
		return fmt.Errorf("task_already_submitted")
	}
	if taskFound.Status != constants.NewTask && taskFound.Status != constants.RejectedTask {
		logs.Error("task_status_is_not_new")
		return fmt.Errorf("task_status_is_not_new")
	}

	// If re-submit
	if taskFound.Status == constants.RejectedTask {
		newReq := &entities.UpdateTaskStatusReq{
			TaskId:  req.TaskId,
			Status:  constants.NewTask,
			User:    req.User,
			Message: "",
		}

		if err = u.ActivityTaskRepo.UpdateActivityTaskStatus(newReq); err != nil {
			return err
		}
	}

	// Do submit
	if err := u.ActivityTaskRepo.SubmitActivityTask(req); err != nil {
		return err
	}

	return nil
}

func (u *activityTaskUse) DeleteActivitySubTask(req *entities.DeleteSubTaskReq) error {
	return u.ActivityTaskRepo.DeleteActivitySubTask(req)
}

func (u *activityTaskUse) UpdateActivityTaskStatus(req *entities.UpdateTaskStatusReq) error {
	// check all is valid
	taskFound, err := u.ActivityTaskRepo.ActivityTaskById(req.TaskId)
	if err != nil {
		return err
	}
	if taskFound.TransactionQty == 0 {
		logs.Error("task_has_no_transaction")
		return fmt.Errorf("task_has_no_transaction")
	}
	if taskFound.Status == constants.RejectedTask || taskFound.Status == constants.ClosedTask {
		logs.Error("task_already_closed_or_rejected")
		return fmt.Errorf("task_already_closed_or_rejected")
	}
	if taskFound.Status == req.Status {
		logs.Error("request_task_same_status")
		return fmt.Errorf("request_task_same_status")
	}

	err = u.ActivityTaskRepo.UpdateActivityTaskStatus(req)
	if err != nil {
		return err
	}
	return nil
}

func (u *activityTaskUse) CloseActivityTask(req *entities.UpdateTaskStatusReq) error {
	// check closed or rejected status
	taskFound, err := u.ActivityTaskRepo.ActivityTaskById(req.TaskId)
	if err != nil {
		return err
	}
	if taskFound.Status != constants.ApprovedTask {
		logs.Error("only_approved_can_close")
		return fmt.Errorf("only_approved_can_close")
	}

	// delete all sub-task
	delReq := &entities.DeleteTaskReq{
		ActivityId: taskFound.ActivityId,
		TaskId:     req.TaskId,
		User:       req.User,
	}
	if err := u.ActivityTaskRepo.DeleteActivityAllSubTaskByTaskId(delReq); err != nil {
		logs.Error(err)
		return err
	}

	// update task status to CLOSED
	err = u.ActivityTaskRepo.UpdateActivityTaskStatus(req)
	if err != nil {
		return err
	}

	return nil
}

func (u *activityTaskUse) ApproveActivityTask(req *entities.UpdateTaskStatusReq) ([]entities.ApproveTaskStatusRes, error) {
	// check closed or rejected status
	taskFound, err := u.ActivityTaskRepo.ActivityTaskById(req.TaskId)
	if err != nil {
		return nil, err
	}
	if taskFound.TransactionQty == 0 {
		logs.Error("task_has_no_transaction")
		return nil, fmt.Errorf("task_has_no_transaction")
	}
	if taskFound.Status == constants.RejectedTask || taskFound.Status == constants.ClosedTask {
		logs.Error("task_already_closed_or_rejected")
		return nil, fmt.Errorf("task_already_closed_or_rejected")
	}

	// write sub-task to transaction table using transaction/rollback in sql
	tasksRes, err := u.ActivityTaskRepo.WriteActivitySubTaskToTransactionAndActivitySync(req)
	if err != nil {
		return nil, err
	}

	// insert only transaction_id to attachment (not edit file_path, will be updated by minio)
	for _, task := range tasksRes {
		req := &entities.UpdateAttachmentSubTaskReq{
			SubTaskId:     task.SubTaskID,
			TransactionId: task.TransactionID,
			User:          req.User,
		}
		if err := u.AttachmantRepo.InsertTransactionIdToAttachment(req); err != nil {
			return nil, err
		}
	}

	// update task status to APPROVE
	err = u.ActivityTaskRepo.UpdateActivityTaskStatus(req)
	if err != nil {
		return nil, err
	}

	return tasksRes, nil
}

func (u *activityTaskUse) CreateNewActivityTask(req *entities.CreateTaskReq) error {
	return u.ActivityTaskRepo.CreateNewActivityTask(req)
}

func (u *activityTaskUse) CreateNewActivitySubTask(req *entities.UpsertSubTaskReq) (string, error) {
	// check closed or rejected status
	taskFound, err := u.ActivityTaskRepo.ActivityTaskById(req.TaskId)
	if err != nil {
		return "", err
	}
	if taskFound.Status == constants.RejectedTask || taskFound.Status == constants.ClosedTask {
		logs.Error("task_already_closed_or_rejected")
		return "", fmt.Errorf("task_already_closed_or_rejected")
	}

	subTaskId, err := u.ActivityTaskRepo.CreateNewActivitySubTask(req)
	if err != nil {
		return "", err
	}
	return subTaskId, nil
}

func (u *activityTaskUse) AssignmentListByUser(req *entities.AssignmentListReq) ([]entities.AssignmentListRes, error) {
	return u.ActivityTaskRepo.AssignmentListByUser(req)
}

func (u *activityTaskUse) ActivityAssignmentListBySite(req *entities.ActivityAssignmentListReq) ([]entities.ActivityAssignmentListRes, error) {
	return u.ActivityTaskRepo.ActivityAssignmentListBySite(req)
}

func (u *activityTaskUse) ActivityTaskAssignBySite(req *entities.ActivityTaskListReq) ([]entities.ActivityTaskAssignRes, error) {
	return u.ActivityTaskRepo.ActivityTaskAssignBySite(req)
}

func (u *activityTaskUse) ActivityTaskListBySite(req *entities.ActivityTaskListReq) ([]entities.ActivityTaskListRes, error) {
	return u.ActivityTaskRepo.ActivityTaskListBySite(req)
}

func (u *activityTaskUse) ActivityTaskListByActivity(req *entities.ActivityTaskListReq) ([]entities.ActivityTaskListByActivityRes, error) {
	return u.ActivityTaskRepo.ActivityTaskListByActivity(req)
}

func (u *activityTaskUse) UpdateActivitySubTask(req *entities.UpsertSubTaskReq) error {
	return u.ActivityTaskRepo.UpdateActivitySubTask(req)
}

func (u *activityTaskUse) AllActivityTaskListByUser(userId string) ([]entities.ActivityTaskListRes, error) {
	return u.ActivityTaskRepo.AllActivityTaskListByUser(userId)
}
