package entities

import (
	"bufferbox_backend_go/constants"
	"time"
)

type ActivityTaskUsecase interface {
	// Issue
	AssignmentListByUser(*AssignmentListReq) ([]AssignmentListRes, error)
	ActivityAssignmentListBySite(*ActivityAssignmentListReq) ([]ActivityAssignmentListRes, error)
	ActivityTaskListBySite(*ActivityTaskListReq) ([]ActivityTaskListRes, error)
	ActivityTaskListByActivity(*ActivityTaskListReq) ([]ActivityTaskListByActivityRes, error)
	AllActivityTaskListByUser(string) ([]ActivityTaskListRes, error)
	UpdateActivitySubTask(*UpsertSubTaskReq) error
	DeleteActivityTask(*DeleteTaskReq) error
	SubmitActivityTask(*SubmitTaskReq) error
	CreateNewActivityTask(*CreateTaskReq) error
	CreateNewActivitySubTask(*UpsertSubTaskReq) (string, error)
	// Receive
	ActivityTaskById(string) (ActivityTaskRes, error)
	UpdateActivityTaskStatus(*UpdateTaskStatusReq) error
	CloseActivityTask(*UpdateTaskStatusReq) error
	ApproveActivityTask(*UpdateTaskStatusReq) ([]ApproveTaskStatusRes, error)
	VerifyActivitySubTask(*ActivitySubTaskReq) ([]ActivitySubTaskRes, error)
	ActivityTaskHistoryLists(*ActivityTaskStatusReq) ([]ActivityTaskHistoryRes, error)
	// Both
	ActivitySubTaskLists(*ActivitySubTaskReq) ([]ActivitySubTaskRes, *string, error)
	DeleteActivitySubTask(*DeleteSubTaskReq) error
	// Main Platform
	ActivityTaskAssignBySite(*ActivityTaskListReq) ([]ActivityTaskAssignRes, error)
	ActivityMember(*ActivityMemberReq) ([]TeamMemberRes, error)
	ActivityMemberDelete(*MemberDeleteByRoleReq) error
}

type ActivityTaskRepository interface {
	// Issue
	AssignmentListByUser(*AssignmentListReq) ([]AssignmentListRes, error)
	ActivityAssignmentListBySite(*ActivityAssignmentListReq) ([]ActivityAssignmentListRes, error)
	ActivityTaskListBySite(*ActivityTaskListReq) ([]ActivityTaskListRes, error)
	ActivityTaskListByActivity(*ActivityTaskListReq) ([]ActivityTaskListByActivityRes, error)
	AllActivityTaskListByUser(string) ([]ActivityTaskListRes, error)
	UpdateActivitySubTask(*UpsertSubTaskReq) error
	DeleteActivityTask(*DeleteTaskReq) error
	SubmitActivityTask(*SubmitTaskReq) error
	CreateNewActivityTask(*CreateTaskReq) error
	CreateNewActivitySubTask(*UpsertSubTaskReq) (string, error)
	DeleteActivityAllSubTaskByTaskId(*DeleteTaskReq) error
	// Receive
	ActivityTaskById(string) (ActivityTaskRes, error)
	UpdateActivityTaskStatus(*UpdateTaskStatusReq) error
	CheckActivitySubTaskDuplicated([]ActivitySubTaskRes) ([]string, error)
	ActivityTaskHistoryLists(*ActivityTaskStatusReq) ([]ActivityTaskHistoryRes, error)
	WriteActivitySubTaskToTransactionAndActivitySync(*UpdateTaskStatusReq) ([]ApproveTaskStatusRes, error)
	// Both
	ActivitySubTaskLists(*ActivitySubTaskReq) ([]ActivitySubTaskRes, error)
	DeleteActivitySubTask(*DeleteSubTaskReq) error
	// Utils
	GetLatestTaskHistoryByTaskId(string) (LatestTaskHistoryRes, error)
	// Main Platform
	ActivityTaskAssignBySite(*ActivityTaskListReq) ([]ActivityTaskAssignRes, error)
}

type ActivityTaskHistoryRes struct {
	//Id            string               `json:"id" db:"id"`
	//SiteId        string               `json:"site_id" db:"site_id"`
	//TaskId        string               `json:"task_id" db:"task_id"`
	TaskCode string               `json:"task_code" db:"task_code"`
	ActionAt time.Time            `json:"action_at" db:"action_at"`
	Activity string               `json:"activity" db:"activity"`
	Status   constants.TaskStatus `json:"status" db:"status"`
	//ChangedType   string               `json:"changed_type" db:"changed_type"`
	Reason    string `json:"reason" db:"reason"`
	CreatedBy string `json:"created_by" db:"created_by"`
	//CreatedAt time.Time `json:"created_at" db:"created_at"`
	//InvitedBy     string    `json:"invited_by" db:"invited_by"`
	//InvitedAt     time.Time `json:"invited_at" db:"invited_at"`
	//SubmittedDate *string `json:"submitted_date" db:"submitted_date"`
	//UpdatedAt *string `json:"updated_at" db:"updated_at"`
	//UpdatedBy *string `json:"updated_by" db:"updated_by"`
}

type ActivityTaskHistoryReq struct {
	SiteId   string `json:"site_id" db:"site_id"`
	Offset   int    `json:"offset"`
	PageSize int    `json:"page_size"`
}

type ActivityTaskReq struct {
	SiteId   string `json:"site_id" db:"site_id"`
	Lang     string
	Offset   int `json:"offset"`
	PageSize int `json:"page_size"`
}

type ActivityTaskRes struct {
	CompanyId      string               `json:"company_id" db:"company_id"`
	SiteId         string               `json:"site_id" db:"site_id"`
	ActivityId     string               `json:"activity_id" db:"activity_id"`
	TaskId         string               `json:"task_id" db:"task_id"`
	TaskCode       string               `json:"task_code" db:"task_code"`
	Name           string               `json:"name" db:"name"`
	Description    *string              `json:"description" db:"description"`
	Status         constants.TaskStatus `json:"status" db:"status"`
	TransactionQty int                  `json:"transaction_qty" db:"transaction_qty"`
	CreatedBy      string               `json:"created_by" db:"created_by"`
	CreatedAt      time.Time            `json:"created_at" db:"created_at"`
	SubmittedDate  *string              `json:"submitted_date" db:"submitted_date"`
}

type ActivitySubTaskReq struct {
	CompanyId string `db:"company_id"`
	SiteId    string `db:"site_id"`
	TaskId    string `db:"task_id"`
	Lang      string `db:"lang"`
	Role      string `db:"role"`
	TaskSide  string `db:"task_side"`
}

type VerifySubTaskReq struct {
	CompanyId string `db:"company_id"`
	SiteId    string `db:"site_id"`
	TaskId    string `db:"task_id"`
	Lang      string `db:"lang"`
}

type ActivityTaskAssignKey struct {
	TaskId       string `db:"task_id"`
	Name         string `db:"name"`
	ScopeName    string `db:"scope_name"`
	CategoryName string `db:"category_name"`
	Users        string `db:"users"`
}

// dev
type ActivityTaskAssignReq struct {
	TaskId       string        `json:"task_id"`
	Name         string        `json:"name"`
	ScopeName    string        `json:"scope_name"`
	CategoryName string        `json:"category_name"`
	Users        []AssignUsers `json:"users"`
}

// dev
type AssignUsers struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Status   string `json:"additional"`
}

type ActivitySubTaskRes struct {
	TaskId          string               `json:"task_id" db:"task_id"`
	TaskCode        string               `json:"task_code" db:"task_code"`
	Company         string               `json:"company" db:"company"`
	Site            string               `json:"site" db:"site"`
	ActivityId      string               `json:"activity_id" db:"activity_id"`
	Activity        string               `json:"activity" db:"activity"`
	Status          constants.TaskStatus `json:"task_status" db:"task_status"`
	TaskSubmittedAt *string              `json:"task_submitted_at" db:"task_submitted_at"`
	TaskCreatedAt   time.Time            `json:"task_created_at" db:"task_created_at"`
	TaskResult      float64              `json:"task_result_emission" db:"task_result_emission"`
	Factor          float64              `json:"factor" db:"factor"`
	EmissionSource  string               `json:"emission_source" db:"emission_source"`
	Unit            string               `json:"unit" db:"unit"`
	SubTaskId       string               `json:"sub_task_id" db:"sub_task_id"`
	ActionDate      time.Time            `json:"action_date" db:"action_date"`
	Additional      string               `json:"additional" db:"additional"`
	Evidence        string               `json:"evidence" db:"evidence"`
	Amount          float64              `json:"amount" db:"amount"`
	Result          float64              `json:"result_emission" db:"result_emission"`
	IssueBy         string               `json:"issue_by" db:"issue_by"`
	UpdatedAt       *string              `json:"updated_at" db:"updated_at"`
	Verify          string               `json:"verify" db:"verify"`
	FileName        *string              `json:"file_name" db:"file_name"`
	AttachmentId    *string              `json:"attachment_id" db:"attachment_id"`
}

type LatestTaskHistoryRes struct {
	TaskId      string               `json:"task_id" db:"task_id"`
	Activity    string               `json:"activity" db:"activity"`
	ChangedType string               `json:"changed_type" db:"changed_type"`
	Reason      *string              `json:"reason" db:"reason"`
	Status      constants.TaskStatus `json:"status" db:"status"`
	ActionAt    time.Time            `json:"action_at" db:"action_at"`
}

type VerifySubTaskRes struct {
	Id string `json:"id"`
}

type DeleteSubTaskReq struct {
	SubTaskId string `db:"id"`
	TaskId    string `db:"task_id"`
	User      string `db:"user"`
}

type DeleteTaskReq struct {
	ActivityId string `db:"activity_id"`
	TaskId     string `db:"task_id"`
	User       string `db:"user"`
}

type SubmitTaskReq struct {
	ActivityId string `db:"activity_id"`
	TaskId     string `db:"task_id"`
	User       string `db:"user"`
}

type ApproveTaskStatusRes struct {
	SubTaskID     string `db:"sub_task_id" json:"sub_task_id"`
	TransactionID string `db:"transaction_id" json:"transaction_id"`
}

type UpdateTaskStatusReq struct {
	TaskId  string               `db:"task_id"`
	Status  constants.TaskStatus `json:"status" db:"status"`
	User    string               `db:"user"`
	Message string               `json:"message,omitempty" db:"message"`
}

type CreateTaskReq struct {
	CompanyId  string `db:"company_id"`
	SiteId     string `db:"site_id"`
	ActivityId string `db:"activity_id"`
	User       string `db:"user"`
}

type UpsertSubTaskReq struct {
	TaskId     string    `db:"task_id"`
	SubTaskId  string    `db:"sub_task_id"`
	Amount     float64   `json:"amount" db:"amount"`
	ActionDate time.Time `json:"action_date" db:"action_date"`
	Evidence   string    `json:"evidence" db:"evidence"`
	Additional string    `json:"additional" db:"additional"`
	User       string    `db:"user"`
}

type InviteActivityStatusByActivityRes struct {
	CompanyId    string  `json:"company_id" db:"company_id"`
	SiteId       string  `json:"site_id" db:"site_id"`
	ActivityId   string  `json:"activity_id" db:"activity_id"`
	Activity     string  `json:"activity" db:"activity"`
	InviteStatus string  `json:"invite_status" db:"invite_status"`
	RoleId       *string `json:"role_id" db:"role_id"`
	InviteRole   *string `json:"invite_role" db:"invite_role"`
	InviteTo     string  `json:"invited_to" db:"invited_to"`
	AcceptedDate *string `json:"accepted_date" db:"accepted_date"`
	LastActive   *string `json:"last_active" db:"last_active"`
}

type ActivityTaskStatusReq struct {
	CompanyId string `db:"company_id"`
	SiteId    string `db:"site_id"`
	Lang      string `db:"lang"`
	StartDate string `json:"start_date" db:"start_date"`
	EndDate   string `json:"end_date" db:"end_date"`
}

type ActivityTaskAssignRes struct {
	ActivityId   string `json:"activity_id" db:"activity_id"`
	Activity     string `json:"activity" db:"activity"`
	Scope        string `json:"scope" db:"scope"`
	EmissionType string `json:"emission_type" db:"emission_type"`
	Icon         string `json:"icon" db:"icon"`
	Accepted     int    `json:"accepted" db:"accepted"`
	Waiting      int    `json:"waiting" db:"waiting"`
}

type ActivityTaskListRes struct {
	CompanyId      string  `json:"company_id" db:"company_id"`
	Company        string  `json:"company" db:"company"`
	SiteId         string  `json:"site_id" db:"site_id"`
	Site           string  `json:"site" db:"site"`
	ActivityId     string  `json:"activity_id" db:"activity_id"`
	Activity       string  `json:"activity" db:"activity"`
	ActivityTaskId string  `json:"task_id" db:"task_id"`
	TaskCode       string  `json:"task_code" db:"task_code"`
	SubTaskCount   int     `json:"sub_task_count" db:"sub_task_count"`
	CreatedBy      string  `json:"created_by" db:"created_by"`
	UpdatedAt      *string `json:"updated_at" db:"updated_at"`
	UpdatedBy      *string `json:"updated_by" db:"updated_by"`
	SubmittedDate  *string `json:"submitted_date" db:"submitted_date"`
	Status         string  `json:"status" db:"status"`
	EmissionSource string  `json:"emission_source" db:"emission_source"`
	Result         float64 `json:"result_emission" db:"result_emission"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
}

type ActivityTaskListByActivityRes struct {
	CompanyId      string  `json:"company_id" db:"company_id"`
	Company        string  `json:"company" db:"company"`
	SiteId         string  `json:"site_id" db:"site_id"`
	Site           string  `json:"site" db:"site"`
	ActivityId     string  `json:"activity_id" db:"activity_id"`
	Activity       string  `json:"activity" db:"activity"`
	ActivityTaskId string  `json:"task_id" db:"task_id"`
	TaskCode       string  `json:"task_code" db:"task_code"`
	SubTaskCount   int     `json:"sub_task_count" db:"sub_task_count"`
	CreatedBy      string  `json:"created_by" db:"created_by"`
	UpdatedAt      *string `json:"updated_at" db:"updated_at"`
	UpdatedBy      *string `json:"updated_by" db:"updated_by"`
	SubmittedDate  *string `json:"submitted_date" db:"submitted_date"`
	Status         string  `json:"status" db:"status"`
	EmissionSource string  `json:"emission_source" db:"emission_source"`
	Result         float64 `json:"result_emission" db:"result_emission"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
	Scope          string  `json:"scope" db:"scope"`
	EmissionType   string  `json:"emission_type" db:"emission_type"`
	Unit           string  `json:"unit" db:"unit"`
	Amount         float64 `json:"amount" db:"amount"`
	Icon           string  `json:"icon" db:"icon"`
}

type AssignmentListReq struct {
	User     string `db:"user"`
	Email    string `db:"invited_to"`
	Role     string `db:"role"`
	TaskSide string `db:"task_side"`
}

type ActivityAssignmentListReq struct {
	AssignmentListReq
	CompanyId string `db:"company_id"`
	SiteId    string `db:"side_id"`
	Lang      string `json:"lang" db:"lang"`
}

type ActivityTaskListReq struct {
	ActivityAssignmentListReq
	ActivityId string `db:"activity_id"`
}

type AssignmentListRes struct {
	CompanyId     string `json:"company_id" db:"company_id"`
	Company       string `json:"company" db:"company"`
	SiteId        string `json:"site_id" db:"site_id"`
	Site          string `json:"site" db:"site"`
	Role          string `json:"role" db:"role"`
	AssignBy      string `json:"assign_by" db:"assign_by"`
	ActivityCount int    `json:"activity_count" db:"activity_count"`
}

type ActivityAssignmentListRes struct {
	CompanyId      string    `json:"company_id" db:"company_id"`
	Company        string    `json:"company" db:"company"`
	SiteId         string    `json:"site_id" db:"site_id"`
	Site           string    `json:"site" db:"site"`
	ActivityId     string    `json:"activity_id" db:"activity_id"`
	Activity       string    `json:"activity" db:"activity"`
	Scope          string    `json:"scope" db:"scope"`
	EmissionType   string    `json:"emission_type" db:"emission_type"`
	EmissionSource string    `json:"emission_source" db:"emission_source"`
	Unit           string    `json:"unit" db:"unit"`
	Icon           string    `json:"icon" db:"icon"`
	Role           string    `json:"role" db:"role"`
	AssignBy       string    `json:"assign_by" db:"assign_by"`
	AssignAt       time.Time `json:"assign_at" db:"assign_at"`
	TaskCount      int       `json:"task_count" db:"task_count"`
	TotalAmount    float64   `json:"total_amount" db:"total_amount"`
}

type UpdateAssignmentStatusReq struct {
	CompanyId  string                     `db:"company_id"`
	SiteId     string                     `db:"site_id"`
	ActivityId string                     `db:"activity_id"`
	User       string                     `db:"user"`
	Status     constants.AssignmentStatus `json:"status" db:"status"`
}
