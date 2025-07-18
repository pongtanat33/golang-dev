package entities

type RoleUsecase interface {
	DeleteUserFromRole(*DeleteUserFromRole) error
	GetUserListByRoleId(*RoleReq) ([]UsersRoleRes, error)
	UpdateSelectedRoleToUser(*UpdateRoleToUserRes) error
	DeleteRoleById(*RolePermissionReq) error
	UpdateRoleAndFeatureById(*UpsertRoleReq) error
	CreateNewRoleAndFeature(*UpsertRoleReq) error
	GetAllFeature(string) ([]FeatureStandardRes, error)
	GetRoleManagementList(*RoleReq) ([]RoleManagementRes, error)
	RoleMasterData(*RoleReq) ([]RoleMasterRes, error)
	GetTableUsersRoles(*TableUsersRolesReq) ([]TableUsersRolesRes, error)
	UpsertRoleFeatureByAdmin(*UpsertRoleFeatureReq) error
	GetRoleFeatureMappedToFrontEnd(*CheckRoleFeatureReq) ([]MergeRoleFeatureRes, error)
}

type RoleRepository interface {
	DeleteUserFromRole(*DeleteUserFromRole) error
	GetUserListByRoleId(*RoleReq) ([]UsersRoleRes, error)
	GetRoleNameById(string) (string, error)
	GetRoleIdsExistInPermission(*RolePermissionReq) ([]string, error)
	DeletePermissionOnUserByRoleLists([]string, string, string, string) error
	DeletePermissionByRoleId(*RolePermissionReq) error
	DeleteRoleById(*RolePermissionReq) error
	IsExistRolePermissionWithoutUserId(*RolePermissionReq) (bool, error)
	IsExistRoleById(string) (bool, error)
	IsExistRolePermissionByName(string, string) (bool, error)
	IsExistRolePermissionByNameAndType(string, string, string) (bool, error)
	IsExistRoleInAnyPermission(*RolePermissionReq) (bool, error)
	UpdateRoleName(*HandleRoleReq) error
	CreateNewRole(*HandleRoleReq) (string, error)
	GetAllFeature(string) ([]FeatureStandardRes, error)
	GetRoleManagementList(*RoleReq) ([]RoleManagementRes, error)
	RoleStandardData(string) ([]RoleMasterRes, error)
	RoleMasterData(*RoleReq) ([]RoleMasterRes, error)
	GetTableUsersRoles(*TableUsersRolesReq) ([]TableUsersRolesRes, error)
	UpsertRoleFeature(*RoleFeatureReq) error
	GetRoleFeatureMappedToFrontEnd(*CheckRoleFeatureReq) ([]MergeRoleFeatureRes, error)
	GetRoleFeatureStandard(string) ([]RoleFeatureStandard, error)
	CreateRole(*RoleCreateReq) (string, error)
	GetRoleStandard(string) ([]RoleRes, error)
}

type DeleteUserFromRole struct {
	PermissionId string `json:"permission_id" db:"permission_id"`
	User         string `json:"user" db:"user"`
}
type UsersRoleRes struct {
	Id           string `json:"user_id" db:"user_id"`
	PermissionId string `json:"permission_id" db:"permission_id"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	Email        string `json:"email" db:"email"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	RoleName     string `json:"role_name" db:"role_name"`
	LastActive   string `json:"last_active" db:"last_active"`
}

type UpsertRoleFeatureReq struct {
	RoleFeatureList []RoleFeatureList `json:"role_feature_lists"`
}

type RoleFeatureList struct {
	RoleName          string            `json:"role_name"`
	RoleType          string            `json:"role_type"`
	FeatureName       string            `json:"feature_name"`
	PermissionSetting PermissionSetting `json:"permission_setting"`
}

type RoleFeatureReq struct {
	RoleId    string
	FeatureId string
	Create    bool
	Read      bool
	Update    bool
	Delete    bool
	User      string
}

type EmailRole struct {
	Email  string `json:"email" db:"email"`
	RoleId string `json:"role_id" db:"role_id"`
}

type RoleReq struct {
	TableId  string `db:"table_id"`
	RoleType string `json:"role_type" db:"role_type"`
	RoleName string `json:"role_name" db:"role_name"`
}

type RoleRes struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	RoleType string `json:"role_type" db:"role_type"`
}

type RoleMasterRes struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type RoleMemberRes struct {
	Id   *string `json:"id" db:"id"`
	Name *string `json:"name" db:"name"`
}

type RoleKey struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

type TableUsersRolesReq struct {
	TableId  string `json:"table_id" db:"table_id"`
	RoleType string `json:"role_type" db:"role_type"`
}

type TableUsersRolesRes struct {
	Id         string `json:"id" db:"id"`
	FullName   string `json:"fullname" db:"fullname"`
	LastActive string `json:"last_active" db:"last_active"`
	Role       string `json:"role" db:"role"`
}

type CheckRoleFeatureRes struct {
	//Id          string `json:"id" db:"id"`
	//RoleName    string `json:"role_name" db:"role_name"`
	//RoleType    string `json:"role_type" db:"role_type"`
	FeatureName string `json:"feature_name" db:"feature_name"`
	Create      bool   `json:"create" db:"create"`
	Read        bool   `json:"read" db:"read"`
	Update      bool   `json:"update" db:"update"`
	Delete      bool   `json:"delete" db:"delete"`
	Route       string `json:"route" db:"route"`
}

type MergeRoleFeatureRes struct {
	FeatureName string `json:"feature_name" db:"feature_name"`
	DisplayName string `json:"display_name" db:"display_name"`
	Route       string `json:"route" db:"route"`
	Create      *bool  `json:"create" db:"create"`
	Read        *bool  `json:"read" db:"read"`
	Update      *bool  `json:"update" db:"update"`
	Delete      *bool  `json:"delete" db:"delete"`
}

type CheckRoleFeatureReq struct {
	TableId  string `json:"table_id" db:"table_id"`
	User     string `json:"user" db:"user"`
	Route    string `json:"route"`
	Method   string `json:"method"`
	RoleType string `json:"role_types" db:"role_types"`
}

type RoleFeatureStandard struct {
	RoleId      string `json:"role_id" db:"role_id"`
	RoleName    string `json:"role_name" db:"role_name"`
	DisplayName string `json:"display_name" db:"display_name"`
	FeatureId   string `json:"feature_id" db:"feature_id"`
	Create      *bool  `json:"create" db:"create"`
	Read        *bool  `json:"read" db:"read"`
	Update      *bool  `json:"update" db:"update"`
	Delete      *bool  `json:"delete" db:"delete"`
}

type RoleCreateReq struct {
	RoleName string `json:"name" db:"name"`
	RoleType string `json:"role_type" db:"role_type"`
	User     string `db:"created_by"`
}

type RoleManagementDbRes struct {
	RoleName  string  `db:"role_name"`
	UserCount int     `db:"user_count"`
	RoleTypes *string `db:"role_types"`
}

type RoleManagementRes struct {
	RoleName  string        `json:"role_name"`
	UserCount int           `json:"user_count"`
	RoleTypes []RoleTypeRes `json:"role_types"`
}

type RoleTypeRes struct {
	RoleId   string       `json:"role_id"`
	RoleType string       `json:"role_type"`
	Features []FeatureRes `json:"features"`
}

type RoleTypeDbRes struct {
	RoleId   string         `json:"role_id"`
	RoleType string         `json:"role_type"`
	Features []FeatureDbRes `json:"features"`
}

type FeatureDbRes struct {
	FeatureId   string `json:"feature_id"`
	DisplayName string `json:"display_name"`
	Create      bool   `json:"create"`
	Read        bool   `json:"read"`
	Update      bool   `json:"update"`
	Delete      bool   `json:"delete"`
}

type FeatureRes struct {
	FeatureId   string        `json:"feature_id"`
	DisplayName string        `json:"display_name"`
	Create      FeatureValues `json:"create"`
	Read        FeatureValues `json:"read"`
	Update      FeatureValues `json:"update"`
	Delete      FeatureValues `json:"delete"`
}

type FeatureValues struct {
	Value bool   `json:"value"`
	Name  string `json:"name"`
}

type FeatureStandardDbRes struct {
	FeatureId   string  `db:"feature_id"`
	DisplayName string  `db:"display_name"`
	FeatureName string  `db:"feature_name"`
	Create      *bool   `db:"create,omitempty"`
	Read        *bool   `db:"read,omitempty"`
	Update      *bool   `db:"update,omitempty"`
	Delete      *bool   `db:"delete,omitempty"`
	Icon        *string `db:"icon"`
}

type FeatureStandardRes struct {
	FeatureId   string         `json:"feature_id"`
	DisplayName string         `json:"display_name"`
	FeatureName string         `json:"feature_name"`
	Create      *FeatureValues `json:"create,omitempty"`
	Read        *FeatureValues `json:"read,omitempty"`
	Update      *FeatureValues `json:"update,omitempty"`
	Delete      *FeatureValues `json:"delete,omitempty"`
	Icon        *string        `json:"icon"`
}

type HandleRoleReq struct {
	RoleId   string `json:"role_id" db:"role_id"`
	RoleType string `json:"role_type" db:"role_type"`
	RoleName string `json:"role_name" db:"role_name"`
	User     string `json:"user" db:"user"`
}

type UpsertRoleReq struct {
	TableId   string        `json:"table_id" db:"table_id"`
	RoleName  string        `json:"role_name" db:"role_name"`
	RoleTypes []RoleTypeReq `json:"role_types" db:"role_types"`
	User      string        `db:"user"`
}

type RoleTypeReq struct {
	RoleId   string             `json:"role_id" db:"role_id"`
	RoleType string             `json:"role_type" db:"role_type"`
	Features []CreateFeatureReq `json:"features"`
}

type CreateFeatureReq struct {
	FeatureId string `json:"feature_id" db:"feature_id"`
	Create    bool   `json:"create" db:"create"`
	Read      bool   `json:"read" db:"read"`
	Update    bool   `json:"update" db:"update"`
	Delete    bool   `json:"delete" db:"delete"`
}

type RolePermissionReq struct {
	RoleId    string `json:"role_id" db:"role_id"`
	TableId   string `json:"table_id" db:"table_id"`
	User      string `json:"user" db:"user"`
	UpdatedBy string `db:"updated_by"`
}

type UpdateRoleToUserRes struct {
	TableId     string   `json:"table_id" db:"table_id"`
	UserId      string   `json:"user_id" db:"user_id"`
	RoleIds     []string `json:"role_ids" db:"role_ids"`
	UpdatedUser string   `json:"updated_by" db:"updated_by"`
}
