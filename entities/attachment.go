package entities

type AttachmentUsecase interface {
	AttachmentValueById(string) (*AttachmentRes, error)
	CreateAttachment(*CreateAttachmentReq) (string, error)
	UpdateAttachment(*UpdateAttachmentReq) error
	DeleteAttachment(*DeleteAttachmentReq) error
	CreateAttachmentSubTask(*CreateAttachmentSubTaskReq) (string, error)
	UpdateAttachmentSubTask(*UpdateAttachmentSubTaskReq) error
	DeleteAttachmentSubTask(*DeleteAttachmentSubTaskReq) error
	DeleteAttachmentByTask(*DeleteAttachmentByTaskReq) error
}

type AttachmentRepository interface {
	AttachmentValueById(string) (*AttachmentRes, error)
	CreateAttachment(*CreateAttachmentReq) (string, error)
	UpdateAttachment(*UpdateAttachmentReq) error
	DeleteAttachment(*DeleteAttachmentReq) error
	CreateAttachmentSubTask(*CreateAttachmentSubTaskReq) (string, error)
	UpdateAttachmentSubTask(*UpdateAttachmentSubTaskReq) error
	DeleteAttachmentSubTask(*DeleteAttachmentSubTaskReq) error
	DeleteAttachmentByTask(*DeleteAttachmentByTaskReq) error
	InsertTransactionIdToAttachment(*UpdateAttachmentSubTaskReq) error
}

type Attachment struct {
	Id string `json:"id" `
}

type AttachmentRes struct {
	Id                    string  `json:"id" db:"id"`
	ActivityTransactionId *string `json:"activity_transaction_id" db:"activity_transaction_id"`
	FileName              string  `json:"file_name" db:"file_name"`
	FilePath              string  `json:"file_path" db:"file_path"`
	CreatedBy             string  `json:"created_by" db:"created_by"`
	SubTaskId             *string `json:"sub_task_id" db:"sub_task_id"`
}

type CreateAttachmentReq struct {
	AttachmentId          string `json:"attachment_id" db:"attachment_id"`
	ActivityTransactionId string `json:"activity_transaction_id" db:"activity_transaction_id"`
	FileName              string `json:"file_name" db:"file_name"`
	FilePath              string `json:"file_path" db:"file_path"`
	User                  string `json:"user" db:"user"`
}

type UpdateAttachmentReq struct {
	Id                    string `json:"id" db:"id"`
	ActivityTransactionId string `json:"activity_transaction_id" db:"activity_transaction_id"`
	FileName              string `json:"file_name" db:"file_name"`
	FilePath              string `json:"file_path" db:"file_path"`
	User                  string `json:"user" db:"user"`
}

type DeleteAttachmentReq struct {
	ActivityTransactionId string `json:"activity_transaction_id" db:"activity_transaction_id"`
	User                  string `json:"user" db:"user"`
}

type CreateAttachmentSubTaskReq struct {
	AttachmentId string `json:"attachment_id" db:"attachment_id"`
	SubTaskId    string `json:"sub_task_id" db:"sub_task_id"`
	FileName     string `json:"file_name" db:"file_name"`
	FilePath     string `json:"file_path" db:"file_path"`
	User         string `json:"user" db:"user"`
}

type UpdateAttachmentSubTaskReq struct {
	Id            string `json:"id" db:"id"`
	SubTaskId     string `json:"sub_task_id" db:"sub_task_id"`
	FileName      string `json:"file_name" db:"file_name"`
	FilePath      string `json:"file_path" db:"file_path"`
	User          string `json:"user" db:"user"`
	TransactionId string `json:"activity_transaction_id" db:"activity_transaction_id"`
}

type DeleteAttachmentSubTaskReq struct {
	SubTaskId string `json:"sub_task_id" db:"sub_task_id"`
	User      string `json:"user" db:"user"`
}

type DeleteAttachmentByTaskReq struct {
	TaskId string `json:"task_id" db:"task_id"`
	User   string `json:"user" db:"user"`
}
