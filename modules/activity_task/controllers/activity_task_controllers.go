package controllers

import (
	"bufferbox_backend_go/constants"
	"bufferbox_backend_go/entities"
	"bufferbox_backend_go/logs"
	"bufferbox_backend_go/middlewares"
	"bufferbox_backend_go/pkg/utils"
	"strings"
	"context"

	"github.com/gofiber/fiber/v2"
)

type activityTaskController struct {
	ActivityTaskUsecase entities.ActivityTaskUsecase
}

func NewActivityTaskController(group *fiber.Group, activityTaskUsecase entities.ActivityTaskUsecase) {
	activityTask := &activityTaskController{
		ActivityTaskUsecase: activityTaskUsecase,
	}
	group.Use(middlewares.VerifyToken)

	// 1. Assingment Page
	assingment := group.Group("/assingment").(*fiber.Group)
	assingment.Get("/company/:company_id/site/:site_id/task-status", activityTask.ActivityTaskListBySite)
	assingment.Get("/company/:company_id/site/:site_id/task-assign", activityTask.ActivityTaskAssignBySite)
	assingment.Get("/company/:company_id/site/:site_id/task-history", activityTask.ActivityTaskHistoryLists)
	assingment.Get("/company/:company_id/site/:site_id/activity/:activity_id/activity-member", activityTask.ActivityMember)
	assingment.Delete("/company/:company_id/site/:site_id/activity/:activity_id/activity-member", activityTask.ActivityMemberDelete)

	// 2. Issue Side || Importer
	issueSide := constants.TaskIssueSide
	issue := group.Group("/" + issueSide).(*fiber.Group)
	issue.Use(middlewares.SetTaskSideBySubGroup(issueSide))
	issue.Get("/", activityTask.AssignmentListByUser)
	issue.Get("/company/:company_id/site/:site_id/activities", activityTask.ActivityAssignmentListBySite)
	issue.Get("/company/:company_id/site/:site_id/activity/:activity_id/tasks", activityTask.ActivityTaskListByActivity)
	issue.Get("/company/:company_id/site/:site_id/activity/:activity_id/task/:task_id/sub-tasks", activityTask.ActivitySubTaskLists)
	issue.Post("/company/:company_id/site/:site_id/activity/:activity_id/create-task", activityTask.CreateNewActivityTask)
	issue.Delete("/company/:company_id/site/:site_id/activity/:activity_id/task/:task_id/delete-task", activityTask.DeleteActivityTask)
	issue.Post("/company/:company_id/site/:site_id/activity/:activity_id/task/:task_id/create-sub-task", activityTask.CreateNewActivitySubTask)
	issue.Delete("/company/:company_id/site/:site_id/activity/:activity_id/task/:task_id/sub-task/:sub_task_id/delete-sub-task", activityTask.DeleteActivitySubTask)
	issue.Patch("/company/:company_id/site/:site_id/activity/:activity_id/task/:task_id/sub-task/:sub_task_id/edit-sub-task", activityTask.UpdateActivitySubTask)
	issue.Patch("/company/:company_id/site/:site_id/activity/:activity_id/task/:task_id/submit-task", activityTask.SubmitActivityTask)

	// 3. Receive Side || Checker
	receiveSide := constants.TaskReceiveSide
	receive := group.Group("/" + receiveSide).(*fiber.Group)
	receive.Use(middlewares.SetTaskSideBySubGroup(receiveSide))
	receive.Get("/", activityTask.AllActivityTaskListByUser)
	receive.Get("/company/:company_id/site/:site_id/task/:task_id/sub-tasks", activityTask.ActivitySubTaskLists)
	receive.Get("/company/:company_id/site/:site_id/task/:task_id/sub-tasks-verify", activityTask.VerifyActivitySubTask)
	receive.Delete("/company/:company_id/site/:site_id/task/:task_id/sub-task/:sub_task_id/delete-sub-task", activityTask.DeleteActivitySubTask)
	receive.Patch("/company/:company_id/site/:site_id/task/:task_id/edit-task-status", activityTask.UpdateActivityTaskStatus)
	receive.Patch("/company/:company_id/site/:site_id/task/:task_id/approve-task-status", middlewares.RoleActivityVerify, activityTask.ApproveActivityTaskStatus)
	receive.Patch("/company/:company_id/site/:site_id/task/:task_id/close-task-status", middlewares.RoleActivityVerify, activityTask.CloseActivityTaskStatus)
}

func (a *activityTaskController) ActivityMemberDelete(c *fiber.Ctx) error {
	req := new(entities.MemberDeleteByRoleReq)
	req.TablePermissionId = c.Params("activity_id")
	req.Email = c.Query("email")
	req.IsActive = c.Query("is_active")
	req.User = c.Locals("usr_id").(string)
	req.TableType = constants.A_ROLE_TYPE
	req.RoleIds = strings.Split(c.Query("role_ids"), ",")

	err := a.ActivityTaskUsecase.ActivityMemberDelete(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "activity_member_delete_successfully", nil)
}

func (a *activityTaskController) ActivityMember(c *fiber.Ctx) error {
	req := new(entities.ActivityMemberReq)
	req.ActivityId = c.Params("activity_id")
	req.RoleType = constants.A_ROLE_TYPE
	res, err := a.ActivityTaskUsecase.ActivityMember(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "activity_member_successfully", fiber.Map{
		"result": res,
	})
}

func (a *activityTaskController) VerifyActivitySubTask(c *fiber.Ctx) error {
	req := new(entities.ActivitySubTaskReq)
	req.CompanyId = c.Params("company_id")
	req.SiteId = c.Params("site_id")
	req.TaskId = c.Params("task_id")
	req.Lang = c.Query("lang")

	res, err := a.ActivityTaskUsecase.VerifyActivitySubTask(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "verify_activity_sub_task_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "verify_activity_sub_task_successfully",
		"result":      res,
		"row_count":   len(res),
	})
}

func (a *activityTaskController) ActivityTaskHistoryLists(c *fiber.Ctx) error {
	req := new(entities.ActivityTaskStatusReq)
	req.CompanyId = c.Params("company_id")
	req.SiteId = c.Params("site_id")
	req.StartDate = c.Query("start_date")
	req.EndDate = c.Query("end_date")

	res, err := a.ActivityTaskUsecase.ActivityTaskHistoryLists(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "activity_task_logs_lists_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "activity_task_logs_lists_successfully",
		"result":      res,
		"row_count":   len(res),
	})
}

func (a *activityTaskController) CloseActivityTaskStatus(c *fiber.Ctx) error {
	req := new(entities.UpdateTaskStatusReq)
	req.TaskId = c.Params("task_id")
	req.User = c.Locals("usr_id").(string)
	if err := c.BodyParser(req); err != nil {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	if !utils.IsValidTaskStatus(req.Status) || req.Status != constants.ClosedTask {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	if err := a.ActivityTaskUsecase.CloseActivityTask(req); err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "close_activity_task_status_successfully", nil)
}

func (a *activityTaskController) ApproveActivityTaskStatus(c *fiber.Ctx) error {
	req := new(entities.UpdateTaskStatusReq)
	req.TaskId = c.Params("task_id")
	req.User = c.Locals("usr_id").(string)
	if err := c.BodyParser(req); err != nil {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	if !utils.IsValidTaskStatus(req.Status) || req.Status != constants.ApprovedTask {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	mappingList, err := a.ActivityTaskUsecase.ApproveActivityTask(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	// Remove all data from Redis hash
	if _, removeErr := utils.RemoveAllDataFromRedisByCompany(context.Background(), c.Locals("company_id").(string)); removeErr != nil {
		logs.Error(removeErr)
	}
	
	return utils.HandleResponse(c, fiber.StatusOK, "update_activity_task_status_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"result":      mappingList,
	})
}

func (a *activityTaskController) UpdateActivityTaskStatus(c *fiber.Ctx) error {
	req := new(entities.UpdateTaskStatusReq)
	req.TaskId = c.Params("task_id")
	req.User = c.Locals("usr_id").(string)
	if err := c.BodyParser(req); err != nil {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	if !utils.IsValidTaskStatus(req.Status) || req.Status == constants.ApprovedTask {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	if err := a.ActivityTaskUsecase.UpdateActivityTaskStatus(req); err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "update_activity_task_status_successfully", nil)
}

func (a *activityTaskController) AssignmentListByUser(c *fiber.Ctx) error {
	req := new(entities.AssignmentListReq)
	req.User = c.Locals("usr_id").(string)
	req.TaskSide = c.Locals("task_side").(string)
	if req.TaskSide == constants.TaskIssueSide {
		req.Role = constants.ActivityImporter
	}

	res, err := a.ActivityTaskUsecase.AssignmentListByUser(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "assignment_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "assignment_successfully",
		"result":      res,
		"row_count":   len(res),
	})
}

func (a *activityTaskController) ActivityAssignmentListBySite(c *fiber.Ctx) error {
	req := new(entities.ActivityAssignmentListReq)
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Params("company_id")
	req.SiteId = c.Params("site_id")
	req.Lang = c.Query("lang")
	req.TaskSide = c.Locals("task_side").(string)
	if req.TaskSide == constants.TaskIssueSide {
		req.Role = constants.ActivityImporter
	}

	res, err := a.ActivityTaskUsecase.ActivityAssignmentListBySite(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_assignment_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "activity_assignment_successfully",
		"result":      res,
		"row_count":   len(res),
	})
}

func (a *activityTaskController) ActivityTaskAssignBySite(c *fiber.Ctx) error {
	req := new(entities.ActivityTaskListReq)
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Params("company_id")
	req.SiteId = c.Params("site_id")
	req.Lang = c.Query("lang")

	res, err := a.ActivityTaskUsecase.ActivityTaskAssignBySite(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_tasks_assign_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "activity_tasks_assign_successfully",
		"result":      res,
		"row_count":   len(res),
	})
}

func (a *activityTaskController) ActivityTaskListByActivity(c *fiber.Ctx) error {
	req := new(entities.ActivityTaskListReq)
	req.User = c.Locals("usr_id").(string)
	req.Lang = c.Query("lang")
	req.CompanyId = c.Params("company_id")
	req.SiteId = c.Params("site_id")
	req.ActivityId = c.Params("activity_id")
	req.TaskSide = c.Locals("task_side").(string)
	if req.TaskSide == constants.TaskIssueSide {
		req.Role = constants.ActivityImporter
	}

	res, err := a.ActivityTaskUsecase.ActivityTaskListByActivity(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_tasks_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "activity_tasks_successfully",
		"result":      res,
		"row_count":   len(res),
	})
}

func (a *activityTaskController) ActivityTaskListBySite(c *fiber.Ctx) error {
	req := new(entities.ActivityTaskListReq)
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Params("company_id")
	req.SiteId = c.Params("site_id")

	res, err := a.ActivityTaskUsecase.ActivityTaskListBySite(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_tasks_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "activity_tasks_successfully",
		"result":      res,
		"row_count":   len(res),
	})
}

func (a *activityTaskController) AllActivityTaskListByUser(c *fiber.Ctx) error {
	res, err := a.ActivityTaskUsecase.AllActivityTaskListByUser(c.Locals("usr_id").(string))
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_tasks_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "activity_tasks_successfully",
		"result":      res,
		"row_count":   len(res),
	})
}

func (a *activityTaskController) CreateNewActivityTask(c *fiber.Ctx) error {
	req := new(entities.CreateTaskReq)
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Params("company_id")
	req.SiteId = c.Params("site_id")
	req.ActivityId = c.Params("activity_id")

	err := a.ActivityTaskUsecase.CreateNewActivityTask(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)

	}
	return utils.HandleResponse(c, fiber.StatusOK, "create_activity_task_successfully", nil)
}

func (a *activityTaskController) DeleteActivityTask(c *fiber.Ctx) error {
	req := new(entities.DeleteTaskReq)
	req.ActivityId = c.Params("activity_id")
	req.TaskId = c.Params("task_id")
	req.User = c.Locals("usr_id").(string)

	err := a.ActivityTaskUsecase.DeleteActivityTask(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "delete_activity_task_successfully", nil)
}

func (a *activityTaskController) SubmitActivityTask(c *fiber.Ctx) error {
	req := new(entities.SubmitTaskReq)
	req.ActivityId = c.Params("activity_id")
	req.TaskId = c.Params("task_id")
	req.User = c.Locals("usr_id").(string)

	err := a.ActivityTaskUsecase.SubmitActivityTask(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "submit_activity_task_successfully", nil)
}

func (a *activityTaskController) CreateNewActivitySubTask(c *fiber.Ctx) error {
	req := new(entities.UpsertSubTaskReq)
	req.TaskId = c.Params("task_id")
	req.User = c.Locals("usr_id").(string)
	if err := c.BodyParser(req); err != nil {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	subTaskId, err := a.ActivityTaskUsecase.CreateNewActivitySubTask(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "create_activity_sub_task_successfully", fiber.Map{
		"sub_task_id": subTaskId,
	})
}

func (a *activityTaskController) ActivitySubTaskLists(c *fiber.Ctx) error {
	req := new(entities.ActivitySubTaskReq)
	req.CompanyId = c.Params("company_id")
	req.SiteId = c.Params("site_id")
	req.TaskId = c.Params("task_id")
	req.Lang = c.Query("lang")
	req.TaskSide = c.Locals("task_side").(string)

	res, reason, err := a.ActivityTaskUsecase.ActivitySubTaskLists(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "activity_sub_task_lists_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "activity_sub_task_lists_successfully",
		"result":      res,
		"reason":      reason,
		"row_count":   len(res),
	})
}

func (a *activityTaskController) DeleteActivitySubTask(c *fiber.Ctx) error {
	req := new(entities.DeleteSubTaskReq)
	req.User = c.Locals("usr_id").(string)
	req.TaskId = c.Params("task_id")
	req.SubTaskId = c.Params("sub_task_id")

	err := a.ActivityTaskUsecase.DeleteActivitySubTask(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)

	}
	return utils.HandleResponse(c, fiber.StatusOK, "delete_activity_sub_task_successfully", nil)
}

func (a *activityTaskController) UpdateActivitySubTask(c *fiber.Ctx) error {
	req := new(entities.UpsertSubTaskReq)
	req.TaskId = c.Params("task_id")
	req.SubTaskId = c.Params("sub_task_id")
	req.User = c.Locals("usr_id").(string)
	if err := c.BodyParser(req); err != nil {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	err := a.ActivityTaskUsecase.UpdateActivitySubTask(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "update_activity_sub_task_successfully", nil)
}
