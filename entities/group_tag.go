package entities

type GroupTagUsecase interface {
	DeleteGroup(*GroupDeleteReq) error
	CreateGroup(*GroupCreateReq) error

	GroupTagMasterData(*GroupTagReq) (*GroupTagRes, error)
	UpdateGroupName(*GroupNameUpdateReq) error

	GetGroupSite(*GroupSiteReq) ([]GroupSiteRes, error)
	RemoveGroupSite(*RemoveGroupSiteReq) error
	UpdateGroupSite(*GroupSiteUpdateReq) error

	AllZoneName(string) ([]ZoneName, error)
}

type GroupTagRepository interface {
	DeleteGroup(*GroupDeleteReq) error
	CreateGroup(*GroupCreateReq) error

	GetGroupByName(*GroupReq) (*GroupRes, error)
	UpdateGroupName(*GroupNameUpdateReq) error

	GroupMasterData(*GroupTagReq) ([]GroupRes, error)
	TagMasterData(*GroupTagReq) ([]TagRes, error)

	GetGroupSite(*GroupSiteReq) ([]GroupSiteRes, error)
	RemoveGroupSite(*SiteGroupReq) error

	AllZoneName(string) ([]ZoneName, error)
}

type ZoneName struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type GroupSiteUpdateReq struct {
	GroupId   string   `json:"group_id" db:"group_id"`
	SiteId    []string `json:"site_id" db:"site_id"`
	CompanyId string   `json:"company_id" db:"company_id"`
	User      string
}

type RemoveGroupSiteReq struct {
	SiteId  []string `json:"site_id" db:"site_id"`
	GroupId string   `json:"group_id" db:"group_id"`
	User    string
}

type GroupSiteReq struct {
	GroupId   string `json:"group_id" db:"group_id"`
	CompanyId string `json:"company_id" db:"company_id"`
}

type GroupSiteRes struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type GroupReq struct {
	Name string `json:"name" db:"name"`
}

type GroupNameUpdateReq struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CompanyId string `json:"company_id" db:"company_id"`
	User      string
}

type GroupDeleteReq struct {
	Id        string `json:"id" db:"id"`
	CompanyId string `json:"company_id" db:"company_id"`
	User      string
}

type GroupCreateReq struct {
	Name      string `json:"name" db:"name"`
	CompanyId string `json:"company_id" db:"company_id"`
	User      string
}

type GroupRes struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type TagReq struct {
	Name string `json:"name" db:"name"`
}

type TagRes struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type GroupTagReq struct {
	CompanyId string `json:"id" db:"id"`
}
type GroupTagRes struct {
	GroupRes []GroupRes `json:"group_res"`
	TagRes   []TagRes   `json:"tag_res"`
}
