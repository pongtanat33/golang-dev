package constants

type TaskStatus string

const (
	NewTask      TaskStatus = "New"
	CheckedTask  TaskStatus = "Checked"
	RejectedTask TaskStatus = "Rejected"
	ApprovedTask TaskStatus = "Approved"
	ClosedTask   TaskStatus = "Closed"
)

type AssignmentStatus string

const (
	NewAssignment    AssignmentStatus = "Waiting"
	AcceptAssignment AssignmentStatus = "Accepted"
	RejectAssignment AssignmentStatus = "Rejected"
)

const (
	TaskIssueSide   = "issue"
	TaskReceiveSide = "receive"
)
