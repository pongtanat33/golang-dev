package entities

type InviteDeleteByRoleReq struct {
	Email             string
	TablePermissionId string
	TableType         string
	User              string
	RoleId            string
}

type InviteDeleteReq struct {
	Email             string
	TablePermissionId string
	TableType         string
	User              string
}

type InviteUserByEmailReq struct {
	EmailRole []EmailRole `json:"users"`
	TableId   string      `json:"table_id" db:"table_id"`
	User      string
	TableType string
	Lang      string
}

type InviteUserByEmailRes struct {
	Email       string `json:"email"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type InvitationUpsertReq struct {
	RoleId    string `json:"role_id" db:"role_id"`
	Email     string `json:"email" db:"email"`
	TableId   string `json:"table_id" db:"table_id"`
	StatusId  int
	User      string
	NewUser   bool
	TableType string
}

type InvitationByActivityUpdateReq struct {
	InvitationUpsertReq
	InviteId string `json:"invite_id" db:"invite_id"`
}

type InvitationUpsertRes struct {
	Id string `json:"id" db:"id"`
}

type InvitationCheckReq struct {
	Token string
}

type InvitationCheckRes struct {
	RoleId    *string `json:"role" db:"role"`
	Email     *string `json:"user_email" db:"user_email"`
	TableId   *string `json:"table_permission_id" db:"table_permission_id"`
	NewUser   *bool   `json:"new_user" db:"new_user"`
	TableType *string `json:"table_type" db:"table_type"`
	Result    bool    `json:"result" db:"result"`
	StatusId  int     `json:"status_id" db:"status_id"`
}
