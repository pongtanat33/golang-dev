package entities

type SiteUsecase interface {
	CreateSite(*SiteReq) (string, error)
	SiteValue(string, string) (*SiteValueScopeConfigGroupTag, error)
	SiteUpdate(*SiteUpdateReq) error
	SiteDelete(*SiteDeleteReq) error

	SiteNameUpdate(*SiteNameUpdateReq) error

	SiteConfigMasterData(string) ([]ConfigsTemplateRes, error)
	AllSiteName(string) ([]SiteName, error)

	UpdateSiteGroup(*SiteGroupUpdateReq) error

	SiteMember(*SiteMemberReq) ([]TeamMemberRes, error)
	SiteMemberDelete(*MemberDeleteReq) error
	SiteMemberChange(*UpsertUserPermissionReq) error

	DuplicateSite(*DuplicateSiteReq) (string, error)

	DuplicateTemplateSite(*DuplicateSiteReq) (string, error)
	SiteTemplate(*TemplateReq) ([]TemplateRes, error)
}

type SiteRepository interface {
	CreateSite(*SiteReq) (*SiteRes, error)
	SiteUpdate(*SiteUpdateReq) (*SiteUpdateRes, error)
	SiteDelete(*SiteDeleteReq) error

	SiteNameUpdate(*SiteNameUpdateReq) (*SiteUpdateRes, error)

	// CreateSiteGroup(*SiteGroupReq) error
	CreateSiteTag(*SiteTagReq) error

	CheckSiteByName(string, string) bool
	SiteConfigMasterData() ([]ConfigsTemplate, error)

	CreateSiteConfig(string, string) (string, error)
	GetSiteConfig(string) (*SiteConfig, error)
	UpdateSiteConfig(*SiteConfigUpdate) error

	AllSiteName(string) ([]SiteName, error)

	GetSiteGroups(string) ([]GroupRes, error)
	// UpdateSiteGroup(*SiteGroupReq) (*SiteGroupRes, error)

	GetSiteTags(string) ([]TagRes, error)
	UpdateSiteTag(*SiteTagReq) (*SiteTagRes, error)

	DuplicateSite(*DuplicateSiteReq) (string, error)

	DuplicateTemplateSite(*DuplicateSiteReq) (string, error)
}

type SiteGroupReq struct {
	SiteId  string `db:"site_id"`
	GroupId string `db:"group_id"`
	User    string
}

type SiteGroupUpdateReq struct {
	SiteId     string   `json:"site_id" db:"site_id"`
	GroupsName []string `json:"group_name" db:"group_name"`
	CompanyId  string   `json:"company_id" db:"company_id"`
	User       string
}

type SiteGroupRes struct {
	Found bool `db:"found"`
}

type SiteTagRes struct {
	Found bool `db:"found"`
}

type SiteUpdateRes struct {
	Result string `db:"result"`
}

type SiteTagReq struct {
	SiteId string `db:"site_id"`
	TagId  string `db:"tag_id"`
	User   string
}

type SiteNameUpdateReq struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CompanyId string `json:"company_id" db:"company_id"`
	User      string
}

type SiteReq struct {
	Id          string            `json:"id" db:"id"`
	Name        string            `json:"name" db:"name"`
	Address     string            `json:"address" db:"address"`
	PostCode    string            `json:"postcode" db:"postcode"`
	District    string            `json:"district" db:"district"`
	SubDistrict string            `json:"subdistrict" db:"subdistrict"`
	Province    string            `json:"province" db:"province"`
	CountryId   string            `json:"country_id" db:"country_id"`
	CompanyId   string            `json:"company_id" db:"company_id"`
	ConfigId    string            `db:"config_id"`
	CategoryId  string            `json:"category_id" db:"category_id"`
	Owner       string            `db:"group_owner"`
	TagsName    []string          `json:"tags_name"`
	GroupsName  []string          `json:"groups_name"`
	Employee    string            `json:"employee"`
	SiteConfigs []SiteConfigValue `json:"site_configs"`
}

type SiteRes struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Address     string `json:"address" db:"address"`
	PostCode    string `json:"postcode" db:"postcode"`
	District    string `json:"district" db:"district"`
	SubDistrict string `json:"subdistrict" db:"subdistrict"`
	Province    string `json:"province" db:"province"`
	CountryId   string `json:"country_id" db:"country_id"`
	CompanyId   string `json:"company_id" db:"company_id"`
	ConfigId    string `json:"config_id" db:"config_id"`
	CategoryId  string `json:"category_id" db:"category_id"`
	Owner       string `db:"group_owner"`
}

type SiteConfig struct {
	Config string `json:"configs" db:"configs"`
}

type SiteConfigValue struct {
	EmissionTypeId string `json:"emission_type_id"`
	Value          int    `json:"value"`
}

type SiteValue struct {
	ID               string `json:"id" db:"id"`
	Name             string `json:"name" db:"name"`
	Address          string `json:"address" db:"address"`
	Postcode         string `json:"postcode" db:"postcode"`
	District         string `json:"district" db:"district"`
	Subdistrict      string `json:"subdistrict" db:"subdistrict"`
	Province         string `json:"province" db:"province"`
	CountryID        string `json:"country_id" db:"country_id"`
	Employee         string `json:"employee" db:"employee"`
	CountryName      string `json:"country_name" db:"country_name"`
	CompanyID        string `json:"company_id" db:"company_id"`
	ConfigID         string `json:"config_id" db:"config_id"`
	CategoryID       string `json:"category_id" db:"category_id"`
	CreatedBy        string `json:"created_by" db:"created_by"`
	BusinessCatTH    string `json:"business_category_th" db:"business_category_th"`
	BusinessCatEN    string `json:"business_category_en" db:"business_category_en"`
	BusinessSectorTH string `json:"business_sector_th" db:"business_sector_th"`
	BusinessSectorEN string `json:"business_sector_en" db:"business_sector_en"`
}

type SiteValueScopeConfigGroupTag struct {
	SiteValue   SiteValue          `json:"site_value"`
	ScopeList   []ScopeListRes     `json:"scope_list"`
	SiteConfigs []ConfigsValueFull `json:"site_configs"`
	Groups      []GroupRes         `json:"groups"`
	Tags        []TagRes           `json:"tags"`
}

type SiteConfigMasterDataRaw struct {
	Id              string `json:"id" db:"id"`
	ConfigsTemplate string `db:"configs_template"`
}

type SiteConfigOption struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ConfigsTemplate struct {
	No             int                `json:"no"`
	EmissionTypeId string             `json:"emission_type_id"`
	Type           string             `json:"type"`
	TitleTH        string             `json:"title_th"`
	OptionsTH      []SiteConfigOption `json:"options_th"`
	TitleEN        string             `json:"title_en"`
	OptionsEN      []SiteConfigOption `json:"options_en"`
}

type ConfigsTemplateRes struct {
	No             int                `json:"no"`
	EmissionTypeId string             `json:"emission_type_id"`
	Type           string             `json:"type"`
	Title          string             `json:"title"`
	Options        []SiteConfigOption `json:"options"`
}

type ConfigsValueFull struct {
	No             int                `json:"no"`
	EmissionTypeId string             `json:"emission_type_id"`
	Type           string             `json:"type"`
	Title          string             `json:"title"`
	Value          int                `json:"value"`
	Options        []SiteConfigOption `json:"options"`
}

type SiteName struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name"`
}

type SiteUpdateValue struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Address     string `json:"address" db:"address"`
	Postcode    string `json:"postcode" db:"postcode"`
	District    string `json:"district" db:"district"`
	Subdistrict string `json:"subdistrict" db:"subdistrict"`
	Province    string `json:"province" db:"province"`
	CountryID   string `json:"country_id" db:"country_id"`
	Employee    string `json:"employee" db:"employee"`
	CompanyID   string `json:"company_id" db:"company_id"`
	ConfigID    string `json:"config_id" db:"config_id"`
	CategoryID  string `json:"category_id" db:"category_id"`
}

type SiteUpdateReq struct {
	SiteValue   SiteUpdateValue   `json:"site_value"`
	SiteConfigs []SiteConfigValue `json:"site_configs"`
	GroupsName  []string          `json:"groups_name"`
	TagsName    []string          `json:"tags_name"`
	User        string
}

type SiteConfigUpdate struct {
	Id          string
	SiteConfigs string
	User        string
}

type SiteDeleteReq struct {
	Id        string `json:"id" db:"id"`
	CompanyId string `json:"company_id" db:"company_id"`
	User      string
}

type DuplicateSiteReq struct {
	SiteId    string `json:"site_id" db:"site_id"`
	CompanyID string `json:"company_id" db:"company_id"`
	Name      string `json:"name" db:"name"`
	User      string
}

type TemplateReq struct {
	BusinessSectorId *string `json:"business_sector_id" db:"business_sector_id"`
}

type TemplateRes struct {
	SiteId                      string  `json:"site_id" db:"site_id"`
	Name                        string  `json:"name" db:"name"`
	BusinessSectorId            string  `json:"business_sector_id" db:"business_sector_id"`
	BusinessSectorEN            string  `json:"business_sector_en" db:"business_sector_en"`
	BusinessCategoryEN          string  `json:"business_category_en" db:"business_category_en"`
	BusinessCategoryDescription *string `json:"business_category_description" db:"business_category_description"`
}
