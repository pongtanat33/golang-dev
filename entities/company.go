package entities

import "time"

type CompanyUsecase interface {
	CompanyCreate(*CompanyReq) (*CompanyRes, error)
	CompanyValue(string, string) (*CompanyRes, error)
	//TestCompanyValue(string, string) (*CompanyRes, error)
	CompanyUpdate(*CompanyReq) error
	CompanyDelete(*CompanyReq) error

	CompanyGroupsByName(*CompanyGroupByNameReq) ([]CompanyGroupRes, int, error)
	CompanyGroups(*CompanyGroupReq) ([]CompanyGroupRes, int, error)
	CompanySites(*CompanySiteReq) ([]CompanySiteRes, int, error)
	CompanySelectedSites(*CompanySiteReq) ([]CompanySiteRes, int, error)
	CompanyLists(string) ([]CompanyList, *CompanyRes, error)

	CompanyTargetBaseAll(*CompanyTargetBaseReq) ([]CompanyTargetBaseRes, error)
	CompanyTargetBaseById(*CompanyTargetBaseByIdReq) (*CompanyTargetBaseRes, error)
	CompanyTargetBaseUpsert(*CompanyTargetBaseUpsertReq) (*CompanyTargetBaseUpsertRes, error)
	CompanyTargetBaseDelete(*CompanyTargetBaseDeleteReq) error

	CompanyTargetAll(*CompanyTargetBaseReq) ([]CompanyTargetRes, error)
	CompanyTargetById(*CompanyTargetByIdReq) (*CompanyTargetRes, error)
	CompanyTargetUpsert(*CompanyTargetUpsertReq) (*CompanyTargetUpsertRes, error)
	CompanyTargetDelete(*CompanyTargetDeleteReq) error
	CompanyMainTargetUpdate(*CompanyMainTargetUpdateReq) error

	CompanyMember(*CompanyMemberReq) ([]TeamMemberRes, error)
	CompanyMemberDelete(*MemberDeleteReq) error
	CompanyMemberChange(*UpsertUserPermissionReq) error
	CompanyGroupSites(*CompanySiteReq) ([]CompanyGroupSiteRes, int, error)
}

type CompanyRepository interface {
	CompanyCreate(*CompanyReq) (*CompanyRes, error)
	CompanyUpdate(*CompanyReq) (*CompanyUpdateRes, error)
	CompanyDelete(*CompanyReq) error

	CheckCompanyByName(string) (bool, error) //bool
	CompanyGroupsByName(*CompanyGroupByNameReq) ([]CompanyGroupRes, int, error)
	CompanyGroups(*CompanyGroupReq) ([]CompanyGroupRes, int, error)
	CompanySites(*CompanySiteReq) ([]CompanySiteRes, int, error)
	CompanySelectedSites(*CompanySiteReq) ([]CompanySiteRes, int, error)

	CompanyTargetBaseAll(*CompanyTargetBaseReq) ([]CompanyTargetBaseRes, error)
	CompanyTargetBaseById(*CompanyTargetBaseByIdReq) (*CompanyTargetBaseRes, error)
	CompanyTargetBaseUpsert(*CompanyTargetBaseUpsertReq) error
	CompanyTargetBaseDelete(*CompanyTargetBaseDeleteReq) (*CompanyTargetBaseDeleteRes, error)

	CompanyTargetAll(*CompanyTargetBaseReq) ([]CompanyTarget, error)
	CompanyTargetById(*CompanyTargetByIdReq) (*CompanyTarget, error)
	CompanyTargetUpsert(*CompanyTargetUpsertReq) error
	CompanyTargetDelete(*CompanyTargetDeleteReq) error
	CompanyMainTargetUpdate(*CompanyMainTargetUpdateReq) error
	CompanyGroupSites(*CompanySiteReq) ([]CompanyGroupSiteRes, int, error)
}

type CompanyKey struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

type CompanyGroupByNameReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
	Name      string `json:"name" db:"name"`
	Offset    int    `json:"offset" db:"offset"`
	PageSize  int    `json:"page_size" db:"page_size"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
	User      string `json:"user" db:"user"`
}

type CompanyGroupReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
	Offset    int    `json:"offset" db:"offset"`
	PageSize  int    `json:"page_size" db:"page_size"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
	User      string `json:"user" db:"user"`
}

type CompanySiteKey struct {
	Key       string `db:"key"`
	Value     string `db:"value"`
	TotalRows int    `db:"total_rows"`
}

type CompanyGroupKey struct {
	Key       string `db:"key"`        // ต้องตรงกับชื่อคอลัมน์ในฐานข้อมูล
	Value     string `db:"value"`      // ต้องตรงกับชื่อคอลัมน์ในฐานข้อมูล
	TotalRows int    `db:"total_rows"` // ต้องตรงกับชื่อคอลัมน์ในฐานข้อมูล
}

type CompanyGroupRes struct {
	Id             string  `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	SiteQty        int     `json:"site" db:"site"`
	ActivityQty    int     `json:"activities" db:"activities"`
	ResponsibleQty int     `json:"responsible" db:"responsible"`
	TotalEmission  float64 `json:"total_emission" db:"total_emission"`
	Scope1Qty      float64 `json:"scope1" db:"scope1"`
	Scope2Qty      float64 `json:"scope2" db:"scope2"`
	Scope3Qty      float64 `json:"scope3" db:"scope3"`
}

type CompanyGroupSiteRes struct {
	Id             string  `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	ActivityQty    int     `json:"activities" db:"activities"`
	ResponsibleQty int     `json:"responsible" db:"responsible"`
	TotalEmission  float64 `json:"total_emission" db:"total_emission"`
}

type CompanySiteReq struct {
	CompanyId string  `json:"company_id" db:"company_id"`
	GroupId   *string `json:"group_id,omitempty" db:"group_id,omitempty"`
	Offset    int     `json:"offset" db:"offset"`
	PageSize  int     `json:"page_size" db:"page_size"`
	FromDate  string  `json:"from_date" db:"from_date"`
	ToDate    string  `json:"to_date" db:"to_date"`
	User      string  `json:"user" db:"user"`
}

type CompanySiteRes struct {
	Id                  string      `json:"id"`
	Name                string      `json:"name" `
	Responsible         int         `json:"responsible"`
	TotalEmission       float64     `json:"total_emission" db:"total_emission"`
	Activity            int         `json:"activity"`
	ActivityTransaction int         `json:"activity_transaction"`
	Groups              *[]GroupRes `json:"groups,omitempty"`
}

type CompanyValue struct {
	Id          string `db:"id"`
	Name        string `json:"name" db:"name"`
	Category    int    `json:"category" db:"category"`
	Country     string `json:"country" db:"country"`
	Address     string `json:"address" db:"address"`
	Postcode    string `json:"postcode" db:"postcode"`
	District    string `json:"district" db:"district"`
	SubDistrict string `json:"subdistrict" db:"subdistrict"`
	Province    string `json:"province" db:"province"`
	CreatedBy   string `json:"created_by" db:"created_by"`
}

type CompanyList struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}
type CompanyReq struct {
	Id          string `db:"id"`
	Name        string `json:"name" db:"name"`
	Category    string `json:"category" db:"category"`
	Country     string `json:"country" db:"country"`
	Address     string `json:"address" db:"address"`
	Postcode    string `json:"postcode" db:"postcode"`
	District    string `json:"district" db:"district"`
	SubDistrict string `json:"subdistrict" db:"subdistrict"`
	Province    string `json:"province" db:"province"`
	User        string `json:"user" db:"user"`
}

type CompanyRes struct {
	Id          string `db:"id"`
	Name        string `json:"name" db:"name"`
	CategoryId  string `json:"category_id" db:"category_id"`
	Category_TH string `json:"business_category_th" db:"business_category_th"`
	Category_EN string `json:"business_category_en" db:"business_category_en"`
	Country     string `json:"country" db:"country"`
	CountryName string `json:"country_name" db:"country_name"`
	Address     string `json:"address" db:"address"`
	Postcode    string `json:"postcode" db:"postcode"`
	District    string `json:"district" db:"district"`
	SubDistrict string `json:"subdistrict" db:"subdistrict"`
	Province    string `json:"province" db:"province"`
	CreatedBy   string `json:"created_by" db:"created_by"`
	SectorId    string `json:"sector_id" db:"sector_id"`
	Sector_TH   string `json:"business_sector_th" db:"business_sector_th"`
	Sector_EN   string `json:"business_sector_en" db:"business_sector_en"`
	IsActive    bool   `json:"is_active" db:"is_active"`
	SiteQty     int    `json:"site_qty" db:"site_qty"`
}

type CompanyUpdateRes struct {
	Result string `db:"result"`
}

type CompanyTargetBaseReq struct {
	CompanyId string
}

type CompanyTargetBaseRes struct {
	Id        string    `json:"id" db:"id"`
	Year      int       `json:"year" db:"year"`
	CompanyId string    `json:"company" db:"company"`
	IsCustom  bool      `json:"isCustom" db:"is_custom"`
	Scope1    float64   `json:"scope1" db:"scope_1"`
	Scope2    float64   `json:"scope2" db:"scope_2"`
	Scope3    float64   `json:"scope3" db:"scope_3"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type CompanyTargetBaseByIdReq struct {
	CompanyId string
	BaseId    string
}

type CompanyTargetByIdReq struct {
	CompanyId string
	TargetId  string
}

type CompanyTarget struct {
	Id                     string    `json:"id" db:"id"`
	Year                   int       `json:"year" db:"year"`
	CompanyId              string    `json:"company" db:"company"`
	Base                   *string   `json:"base" db:"base"`
	Scope1Max              float64   `json:"scope1Max" db:"scope_1_max"`
	Scope2Max              float64   `json:"scope2Max" db:"scope_2_max"`
	Scope3Max              float64   `json:"scope3Max" db:"scope_3_max"`
	Scope1ReductionPercent float64   `json:"scope1ReductionPercent" db:"scope_1_reduction_percent"`
	Scope2ReductionPercent float64   `json:"scope2ReductionPercent" db:"scope_2_reduction_percent"`
	Scope3ReductionPercent float64   `json:"scope3ReductionPercent" db:"scope_3_reduction_percent"`
	MainTarget             bool      `json:"mainTarget" db:"main_target"`
	CreatedAt              time.Time `json:"createdAt" db:"created_at"`
}

type CompanyTargetRes struct {
	Id                     string                `json:"id" db:"id"`
	Year                   int                   `json:"year" db:"year"`
	CompanyId              string                `json:"company" db:"company"`
	Base                   *CompanyTargetBaseRes `json:"base"`
	Scope1Max              float64               `json:"scope1Max" db:"scope_1_max"`
	Scope2Max              float64               `json:"scope2Max" db:"scope_2_max"`
	Scope3Max              float64               `json:"scope3Max" db:"scope_3_max"`
	Scope1ReductionPercent float64               `json:"scope1ReductionPercent" db:"scope_1_reduction_percent"`
	Scope2ReductionPercent float64               `json:"scope2ReductionPercent" db:"scope_2_reduction_percent"`
	Scope3ReductionPercent float64               `json:"scope3ReductionPercent" db:"scope_3_reduction_percent"`
	MainTarget             bool                  `json:"mainTarget" db:"main_target"`
	CreatedAt              time.Time             `json:"createdAt" db:"created_at"`
}

type CompanyMainTargetUpdateReq struct {
	Id   string `json:"id" db:"id"`
	User string
}

type CompanyTargetDeleteReq struct {
	Id        string `json:"id" db:"id"`
	CompanyId string `json:"company" db:"company"`
	User      string
}

type CompanyTargetBaseDeleteReq struct {
	Id        string `json:"id" db:"id"`
	CompanyId string `json:"company" db:"company"`
	User      string
}

type CompanyTargetBaseDeleteRes struct {
	Result bool `json:"result" db:"result"`
}

type CompanyTargetBaseUpsertReq struct {
	Id        *string `json:"id" db:"id"`
	Year      int     `json:"year" db:"year"`
	IsCustom  bool    `json:"isCustom" db:"is_custom"`
	Scope1    float64 `json:"scope1" db:"scope_1"`
	Scope2    float64 `json:"scope2" db:"scope_2"`
	Scope3    float64 `json:"scope3" db:"scope_3"`
	CompanyId string  `json:"company" db:"company"`
	User      string
}

type CompanyTargetBaseUpsertRes struct {
	Id        *string `json:"id" db:"id"`
	Year      int     `json:"year" db:"year"`
	IsCustom  bool    `json:"isCustom" db:"is_custom"`
	Scope1    float64 `json:"scope1" db:"scope_1"`
	Scope2    float64 `json:"scope2" db:"scope_2"`
	Scope3    float64 `json:"scope3" db:"scope_3"`
	CompanyId string  `json:"company" db:"company"`
}

type CompanyTargetUpsertReq struct {
	Id                     *string `json:"id" db:"id"`
	Year                   int     `json:"year" db:"year"`
	Base                   *string `json:"base" db:"base"`
	Scope1Max              float64 `json:"scope1Max" db:"scope_1_max"`
	Scope2Max              float64 `json:"scope2Max" db:"scope_2_max"`
	Scope3Max              float64 `json:"scope3Max" db:"scope_3_max"`
	Scope1ReductionPercent float64 `json:"scope1ReductionPercent" db:"scope_1_reduction_percent"`
	Scope2ReductionPercent float64 `json:"scope2ReductionPercent" db:"scope_2_reduction_percent"`
	Scope3ReductionPercent float64 `json:"scope3ReductionPercent" db:"scope_3_reduction_percent"`
	CompanyId              string  `json:"company" db:"company"`
	User                   string
}

type CompanyTargetUpsertRes struct {
	Id                     *string               `json:"id" db:"id"`
	Year                   int                   `json:"year" db:"year"`
	Base                   *CompanyTargetBaseRes `json:"base" db:"base"`
	Scope1Max              float64               `json:"scope1Max" db:"scope_1_max"`
	Scope2Max              float64               `json:"scope2Max" db:"scope_2_max"`
	Scope3Max              float64               `json:"scope3Max" db:"scope_3_max"`
	Scope1ReductionPercent float64               `json:"scope1ReductionPercent" db:"scope_1_reduction_percent"`
	Scope2ReductionPercent float64               `json:"scope2ReductionPercent" db:"scope_2_reduction_percent"`
	Scope3ReductionPercent float64               `json:"scope3ReductionPercent" db:"scope_3_reduction_percent"`
	CompanyId              string                `json:"company" db:"company"`
}
