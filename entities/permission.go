package entities

type CheckPermissionReq struct {
	TableId string
	User    string
}

type CheckPermissionRes struct {
	RoleId string `json:"role_id" db:"role_id"`
	Result bool   `json:"result" db:"result"`
}

type UpsertUserPermissionReq struct {
	Id                *string `json:"id" db:"id"`
	TablePermissionId string  `json:"table_permission_id" db:"table_permission_id"`
	UserId            string  `json:"user_id" db:"user_id"`
	RoleName          string  `json:"role_name" db:"role_name"`
	RoleType          string  `json:"role_type" db:"role_type"`
	UserUpdate        string  `json:"user_update" db:"user_update"`
}

type PermissionSetting struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type DeleteUserPermissionByRoleReq struct {
	UserId            string
	TablePermissionId string
	UserDelete        string
	RoleId            string
}

type DeleteUserPermissionReq struct {
	UserId            string
	TablePermissionId string
	RoleType          string
	UserDelete        string
}
