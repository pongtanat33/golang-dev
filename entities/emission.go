package entities

import (
	"database/sql"
	"time"
)

type EmissionUsecase interface {
	GetEmissionListsByType(string, string, string, string) ([]EmissionListRes, error)
	GetEmissionListsExistOnSite(string, string, string) ([]EmissionListRes, error)
	GetEmissionTypeBySite(string, string) ([]EmissionTypeRes, error)
	GetEmissionTypeExistOnSite(string, string) ([]EmissionTypeRes, error)
	GetEmissionTypeExistOnSiteByCriteria(string, string, string, string) ([]EmissionTypeRes, error)
	GetUnitsList(*UnitsListReq) ([]Unit, error)
	GetSummaryResultEmission(*SummaryResultEmissionReq) ([]SummaryResultEmissionRes, error)
	GetFactorType() ([]FactorTypeRes, error)
	GetGWP() ([]GWPRes, error)
	GetEmissionSourceBySite(*EmissionSourceReq) ([]EmissionSourceRes, int, error)
	GetEmissionSourceBySiteByName(*EmissionSourceByNameReq) ([]EmissionSourceRes, int, error)
	GetEmissionSourceByCompany(*EmissionSourceReq) ([]EmissionSourceRes, int, error)
	GetEmissionSourceByCompanyByName(*EmissionSourceByNameReq) ([]EmissionSourceRes, int, error)
	DelelteEmission(*DeleteEmissionReq) error
	//GetEmissionSourceByEmission(*EmissionSourceReq) ([]EmissionSourceRes, error)
	GetEmissionSourceByEmission(*EmissionSourceByEmissionReq) ([]EmissionSourceByEmissionRes, int, error)
	UpdateEmissionSourceBySite(*UpdateEmissionSourceBySiteReq) error
	UpdateEmissionSourceByCompany(*UpdateEmissionSourceByCompanyReq) error
}

type EmissionRepository interface {
	GetEmissionListsByType(string, string, string, string) ([]EmissionListRes, error)
	GetEmissionListsExistOnSite(string, string, string) ([]EmissionListRes, error)
	GetEmissionType(string) ([]EmissionTypeRes, error)
	GetEmissionTypeExistOnSite(string, string) ([]EmissionTypeRes, error)
	GetEmissionTypeExistOnSiteByCriteria(string, string, string, string) ([]EmissionTypeRes, error)
	GetUnitsList(*UnitsListReq) ([]Unit, error)
	GetSummaryResultEmission(*SummaryResultEmissionReq) ([]SummaryResultEmissionRes, error)
	GetEmissionSourceByCompany(*EmissionSourceReq) ([]EmissionSourceRes, error)
	EmissionSourceByCompanyRowCount(*EmissionSourceReq) (int, error)
	GetEmissionSourceByCompanyByName(*EmissionSourceByNameReq) ([]EmissionSourceRes, error)
	EmissionSourceByCompanyByNameRowCount(*EmissionSourceByNameReq) (int, error)
	GetEmissionSourceBySite(*EmissionSourceReq) ([]EmissionSourceRes, error)
	EmissionSourceBySiteRowCount(*EmissionSourceReq) (int, error)
	GetEmissionSourceBySiteByName(*EmissionSourceByNameReq) ([]EmissionSourceRes, error)
	EmissionSourceBySiteByNameRowCount(*EmissionSourceByNameReq) (int, error)

	//GetEmissionSourceByEmission(*EmissionSourceReq) ([]EmissionSourceRes, error)
	GetEmissionSourceByEmission(*EmissionSourceByEmissionReq) ([]EmissionSourceByEmissionRes, int, error)
	UpdateEmissionSourceBySite(*UpdateEmissionSourceBySiteReq) error
	UpdateEmissionSourceByCompany(*UpdateEmissionSourceByCompanyReq) error
	UpdateFactorEmissionSource(*UpdateEmissionSourceReq) error
}

type EmissionKey struct {
	Key   string         `db:"key"`
	Value sql.NullString `db:"value"`
}

type EmissionListReq struct {
	TypeId string `json:"site_id" db:"type_id"`
}

type EmissionListRes struct {
	Id           string  `json:"emission_factor_id" `
	EmissionName string  `json:"emission_name" `
	Description  *string `json:"description" `
}

type EmissionType struct {
	Id               string `json:"id" `
	EmissionTypeName string `json:"name" `
	Icon             string `json:"icon" `
}

type EmissionTypeRes struct {
	Scope         string         `json:"scope"`
	EmissionTypes []EmissionType `json:"emission_type,omitempty"`
}

type ScopeListRes struct {
	Id        string `json:"id" db:"id"`
	ScopeName string `json:"name" db:"name"`
}

type CreateEmissionReq struct {
	Id        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Version   string    `json:"version" db:"version"`
	CountryId string    `json:"country_id" db:"country_id"`
	Source    string    `json:"source" db:"source"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	TypeId    string    `json:"type_id" db:"type_id"`
	User      string    `json:"user" db:"user"`
}

type CreateEmissionFactorReq struct {
	EmissionId string `json:"emission_id" db:"emission_id"`
	FactorId   string `json:"factor_id" db:"factor_id"`
	SiteId     string `json:"site_id" db:"site_id"`
	CompanyId  string `json:"company_id" db:"company_id"`
	User       string `json:"user" db:"user"`
}

type CreateEmissionFactorByCompanyReq struct {
	EmissionId string `json:"emission_id" db:"emission_id"`
	FactorId   string `json:"factor_id" db:"factor_id"`
	CompanyId  string `json:"company_id" db:"company_id"`
	User       string `json:"user" db:"user"`
}

type UpdateEmissionFactorReq struct {
	Id         string `json:"id" db:"id"`
	EmissionId string `json:"emission_id" db:"emission_id"`
	FactorId   string `json:"factor_id" db:"factor_id"`
	CompanyId  string `json:"company_id" db:"company_id"`
	User       string `json:"user" db:"user"`
}

type SummaryResultEmissionReq struct {
	Year      int    `json:"year" db:"year"`
	CompanyId string `json:"company" db:"company"`
}

type SummaryResultEmissionRes struct {
	Name           string  `json:"name" db:"name"`
	ResultEmission float64 `json:"result_emission" db:"result_emission"`
}

type EmissionSourceRes struct {
	SiteId     *string `json:"site_id" db:"site_id"`
	EmissionId string  `json:"emission_id" db:"emission_id"`
	Name       string  `json:"name" db:"name"`
	Factor     float64 `json:"factor" db:"factor"`
	Unit       string  `json:"unit" db:"unit"`
	Scope      string  `json:"scope" db:"scope"`
	TypeName   string  `json:"type_name" db:"type_name"`
	Source     string  `json:"source" db:"source"`
	Icon       string  `json:"icon" db:"icon"`
}

type EmissionSourceByNameReq struct {
	CompanyId    string `json:"company_id" db:"company_id"`
	SiteId       string `json:"site_id" db:"site_id"`
	EmissionName string `json:"emission_name" db:"emission_name"`
	Lang         string `db:"lang"`
	Offset       int    `json:"offset" db:"offset"`
	PageSize     int    `json:"page_size" db:"page_size"`
}

type EmissionSourceReq struct {
	CompanyId  string `json:"company_id" db:"company_id"`
	SiteId     string `json:"site_id" db:"site_id"`
	EmissionId string `json:"emission_id" db:"emission_id"`
	Lang       string `db:"lang"`
	Offset     int    `json:"offset" db:"offset"`
	PageSize   int    `json:"page_size" db:"page_size"`
}

type DeleteEmissionReq struct {
	Id        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	CompanyId string  `json:"company_id" db:"company_id"`
	SiteId    *string `json:"site_id" db:"site_id"`
	User      string
}

type EmissionSourceByEmissionReq struct {
	CompanyId  string `json:"company_id" db:"company_id"`
	SiteId     string `json:"site_id" db:"site_id"`
	EmissionId string `json:"emission_id" db:"emission_id"`
}

type EmissionSourceByEmissionRes struct {
	EmissionId   string       `json:"emission_id" db:"emission_id"`
	EmissionName string       `json:"emission_name" db:"emission_name"`
	UnitId       string       `json:"unit_id" db:"unit_id"`
	Source       string       `json:"source" db:"source"`
	FactorType   []FactorType `json:"factortype"`
}

type FactorType struct {
	EmissionFactorId string  `json:"emission_factor_id" db:"emission_factor_id"`
	Id               *string `json:"id" db:"id"`
	Name             string  `json:"name" db:"name"`
	Factor           float64 `json:"factor" db:"factor"`
	GWP_Id           string  `json:"gwp_id" db:"gwp_id"`
	FactorId         string  `json:"factor_id" db:"factor_id"`
}

type UpdateEmissionSourceBySiteReq struct {
	EmissionFactorId string       `json:"emission_factor_id" db:"emission_factor_id"`
	EmissionId       string       `json:"emission_id" db:"emission_id"`
	EmissionName     string       `json:"emission_name" db:"emission_name"`
	FactorType       []FactorType `json:"factortype"`
	UnitId           string       `json:"unit_id" db:"unit_id"`
	Source           string       `json:"source" db:"source"`
	CompanyId        string       `json:"company_id" db:"company_id"`
	SiteId           string       `json:"site_id" db:"site_id"`
	User             string
}

type UpdateEmissionSourceByCompanyReq struct {
	EmissionFactorId string       `json:"emission_factor_id" db:"emission_factor_id"`
	EmissionId       string       `json:"emission_id" db:"emission_id"`
	EmissionName     string       `json:"emission_name" db:"emission_name"`
	FactorType       []FactorType `json:"factortype"`
	UnitId           string       `json:"unit_id" db:"unit_id"`
	Source           string       `json:"source" db:"source"`
	CompanyId        string       `json:"company_id" db:"company_id"`
	User             string
}

type UpdateEmissionSourceReq struct {
	EmissionFactorId string  `json:"emission_factor_id" db:"emission_factor_id"`
	Factor           float64 `json:"factor" db:"factor"`
	UnitId           string  `json:"unit_id" db:"unit_id"`
	CompanyId        string  `json:"company_id" db:"company_id"`
	User             string
}
