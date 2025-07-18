package entities

import "github.com/jmoiron/sqlx"

type UsersUsecase interface {
	DeleteUser(username string) error
	UpdateUser(*UpdateUserReq) error
	UpdateUserPassword(*UpdateUserPasswordReq) error
	UsersCreate(*UsersReq) (*UsersRes, error)
	VerifyEmail(*UsersReq) error
	UpdateLang(*UpdateLangReq) error
	InviteUserByEmail(*InviteUserByEmailReq) ([]InviteUserByEmailRes, error)
	InvitationCheck(*InvitationCheckReq, bool) error
	SendEmailVerify(*UsersReq) SendEmailStatusRes
	SendEmailResetPassword(*ResetPasswordReq) error
	UserCreateViaInviteEmail(*UsersCreateViaInviteEmailReq) error
	SendEmailVerifyByEmail(*UsersReq) SendEmailStatusRes
	GetUserInfoByUsername(*UsernameReq) (*FullUsersRes, error)
}

type UsersRepository interface {
	DeleteUser(username string) error
	DeleteCompanyByUserId(tx *sqlx.Tx, id string) error
	DeleteUserById(tx *sqlx.Tx, id string) error
	UpdateUserById(*UpdateUserReq) error
	UsersCreate(*UsersReq) (*UsersRes, error)
	SearchUser(*UsersReq) (*UsersRes, error)
	VerifyEmail(string) error
	UpdateLang(*UpdateLangReq) error
	UpsertInvitation(*InvitationUpsertReq) (*InvitationUpsertRes, error)
	InvitationCheck(*InvitationCheckReq) (*InvitationCheckRes, error)
	UpdateInvitationByActivity(*InvitationByActivityUpdateReq) error
}

type UsersCreateViaInviteEmailReq struct {
	Username    string `json:"username" db:"username"`
	Firstname   string `json:"first_name" db:"first_name"`
	Lastname    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Verify_code string `json:"code" db:"verify_code"`
	Token       string `json:"token" db:"token"`
}

type UsernameReq struct {
	Username string `json:"username" db:"username"`
}

type UsersValueRes struct {
	Id                 string  `json:"id" db:"id"`
	Username           string  `json:"username" db:"username"`
	Firstname          string  `json:"firstname" db:"first_name"`
	Lastname           string  `json:"lastname" db:"last_name"`
	Email              string  `json:"email" db:"email"`
	IsRegisterd        string  `json:"is_registered" db:"is_registered"`
	Verify_code        string  `json:"verifycode" db:"verify_code"`
	LastVisitedCompany *string `json:"last_visited_company" db:"last_visited_company"`
	Lang               string  `json:"lang" db:"lang"`
	Result             bool    `json:"result" db:"result"`
}

type UpdateUserPasswordReq struct {
	Username    string `json:"username" db:"username"`
	NewPassword string `json:"new_password" db:"new_password"`
}

type UpdateUserReq struct {
	UserID    *string `json:"user_id" db:"user_id"`
	Username  string  `json:"username" db:"username"`
	Firstname string  `json:"first_name" db:"first_name"`
	Lastname  string  `json:"last_name" db:"last_name"`
}

type UsersReq struct {
	Username    string `json:"username" db:"username"`
	Firstname   string `json:"first_name" db:"first_name"`
	Lastname    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Verify_code string `json:"code" db:"verify_code"`
}

type UsersRes struct {
	Username    string `json:"username" db:"username"`
	Firstname   string `json:"firstname" db:"first_name"`
	Lastname    string `json:"lastname" db:"last_name"`
	Email       string `json:"email" db:"email"`
	IsRegisterd string `json:"is_registered" db:"is_registered"`
	Verify_code string `json:"verifycode" db:"verify_code"`
}

type FullUsersRes struct {
	Id                 string  `json:"id" db:"id"`
	Username           string  `json:"username" db:"username"`
	Firstname          string  `json:"firstname" db:"first_name"`
	Lastname           string  `json:"lastname" db:"last_name"`
	Email              string  `json:"email" db:"email"`
	IsRegisterd        string  `json:"is_registered" db:"is_registered"`
	Verify_code        string  `json:"verifycode" db:"verify_code"`
	LastVisitedCompany *string `json:"last_visited_company" db:"last_visited_company"`
	Lang               string  `json:"lang" db:"lang"`
}

// type UsersInfoRes struct {
// 	Id                 string  `json:"id" db:"id"`
// 	Username           string  `json:"username" db:"username"`
// 	Firstname          string  `json:"firstname" db:"first_name"`
// 	Lastname           string  `json:"lastname" db:"last_name"`
// 	Email              string  `json:"email" db:"email"`
// 	IsRegisterd        string  `json:"is_registered" db:"is_registered"`
// 	Verify_code        string  `json:"verifycode" db:"verify_code"`
// 	LastVisitedCompany *string `json:"last_visited_company" db:"last_visited_company"`
// 	Lang               string  `json:"lang" db:"lang"`
// }

type UpdateLangReq struct {
	Id   string `json:"id" db:"id"`
	Lang string `json:"lang" db:"lang"`
}

type MemberUpdatePermissionReq struct {
	Id                string `json:"id" db:"id"`
	TablePermissionId string `json:"table_permission_id" db:"table_permission_id"`
	UserId            string `json:"user_id" db:"user_id"`
	RoleName          string `json:"role_name" db:"role_name"`
	RoleType          string `json:"role_type" db:"role_type"`
	User              string `json:"user" db:"user"`
}

type FirebaseError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
	} `json:"error"`
}

type UsersVerifyByEmailReq struct {
	Username    string `json:"username" db:"username"`
	Firstname   string `json:"first_name" db:"first_name"`
	Lastname    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Verify_code string `json:"code" db:"verify_code"`
}

type TeamMemberDbRes struct {
	Id         string  `db:"id"`
	FirstName  string  `db:"first_name"`
	LastName   string  `db:"last_name"`
	Email      string  `db:"email"`
	LastActive *string `db:"last_active"`
	AcceptedAt *string `db:"accepted_at"` // Nullable field
	Role       *string `db:"roles"`
	IsActive   int     `db:"is_active"`
}

type TeamMemberRes struct {
	Id         string          `json:"id"`
	FirstName  string          `json:"first_name"`
	LastName   string          `json:"last_name"`
	Email      string          `json:"email"`
	LastActive *string         `json:"last_active"`
	AcceptedAt *string         `json:"accepted_at"` // Nullable field
	Role       []RoleMemberRes `json:"roles"`       // Parsed JSON array
	IsActive   int             `json:"is_active"`
}

type CompanyMemberReq struct {
	CompanyId string
	RoleType  string
}

type MemberDeleteByRoleReq struct {
	Email             string   `json:"email" db:"email"`
	TablePermissionId string   `json:"table_permission_id" db:"table_permission_id"`
	TableType         string   `json:"table_type" db:"table_type"`
	IsActive          string   `json:"is_active" db:"is_active"`
	User              string   `json:"user" db:"user"`
	RoleIds           []string `json:"role_ids"`
	RoleId            string   `json:"role_id" db:"role_id"`
}

type MemberDeleteReq struct {
	Email             string `json:"email" db:"email"`
	TablePermissionId string `json:"table_permission_id" db:"table_permission_id"`
	TableType         string `json:"table_type" db:"table_type"`
	IsActive          string `json:"is_active" db:"is_active"`
	User              string `json:"user" db:"user"`
}

type SiteMemberReq struct {
	SiteId   string
	RoleType string
}

type ActivityMemberReq struct {
	ActivityId string
	RoleType   string
}
