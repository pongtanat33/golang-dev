package entities

import (
	"database/sql"
	"time"
)

type SubscriptionUsecase interface {
	UpsertSelectedSiteById(*SelectedSiteReq) error
	DoVerifyTicket(*TicketReq) (string, error)
	CheckVerifyTicket(*TicketKeyReq) (bool, *TicketKeyVerifyRes, error)
	GenTicketByAdmin(*TicketReq) (string, error)
	GetSubscription(*SubscriptionReq) (*SubscriptionRes, error)
	Unsubscription(*UnsubscriptionReq) (string, error)
	GetPackageFeatureByCompany(*PackageFeatureReq) ([]PackageFeatureRes, error)
	GetPackageList() ([]PackageRes, error)
	CreateNewSubscription(*TicketReq) (string, error)
	SubscriptionChangePackage(req *TicketReq) (string, error)
	CheckExistTicket(req *TicketKeyReq) (*TicketKeyRes, error)
}

type SubscriptionRepository interface {
	ClearAllSelectedSite(*NewCompanySiteReq) error
	UpsertSelectedSiteById(*UpsertCompanySiteReq) error
	GenTicketByAdmin(*TicketReq) (string, error)
	GetPackageFeatureByCompany(*PackageFeatureReq) ([]PackageFeatureRes, error)
	GetPackageList() ([]PackageRes, error)
}

type TicketKeyReq struct {
	TicketKey string `json:"ticket_key" db:"ticket_key"`
}

type TicketKeyRes struct {
	TicketId    string `json:"ticket_id" db:"ticket_id"`
	PackageId   string `json:"package_id" db:"package_id"`
	PackageName string `json:"package_name" db:"package_name"`
	IsUsed      bool   `json:"is_used" db:"is_used"`
	IsExpired   bool   `json:"is_expired" db:"is_expired"`
}

type TicketKeyVerifyRes struct {
	TicketKey   string `json:"ticket_key" db:"ticket_key"`
	PackageId   string `json:"package_id" db:"package_id"`
	PackageName string `json:"package_name" db:"package_name"`
	IsVerify    bool   `json:"is_verify" db:"is_verify"`
}

type CompanySiteSelectedReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
	TicketId  string `json:"ticket_id" db:"ticket_id"`
	User      string `json:"user" db:"user"`
}

type NewCompanySiteReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
	SiteId    string `json:"site_id" db:"site_id"`
	User      string `json:"user" db:"user"`
}

type UpsertCompanySiteReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
	SiteId    string `json:"site_id" db:"site_id"`
	User      string `json:"user" db:"user"`
	IsUsed    bool   `json:"is_used" db:"is_used"`
}

type TicketReq struct {
	Id        string `json:"id" db:"id"`
	CompanyId string `json:"company_id" db:"company_id"`
	TicketKey string `json:"ticket_key" db:"ticket_key"`
	PackageId string `json:"package_id" db:"package_id"`
	User      string `json:"user" db:"user"`
	Type      string `json:"type" db:"type"`
}

type TicketStatusReq struct {
	TicketId string `json:"ticket_id" db:"ticket_id"`
	User     string `json:"user" db:"user"`
}

type CompanyTicketReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
	TicketId  string `json:"ticket_id" db:"ticket_id"`
	User      string `json:"user" db:"user"`
}

type PackageUserReq struct {
	PackageId string `json:"package_id" db:"package_id"`
	User      string `json:"user" db:"user"`
}

type PackageUserRes struct {
	TicketId    string    `json:"ticket_id" db:"ticket_id"`
	IsExpired   bool   `json:"is_expired" db:"is_expired"`
	IsActive    bool   `json:"is_active" db:"is_active"`
	IsUsed      bool   `json:"is_used" db:"is_used"`
	ExpiredDate time.Time `json:"expired_date" db:"expired_date"`
}

type UnsubscriptionReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
	TicketId  string `json:"ticket_id" db:"ticket_id"`
	User      string `json:"user" db:"user"`
}

type SubscriptionReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
}

type SubscriptionRes struct {
	TicketId        string         `json:"ticket_id" db:"ticket_id"`
	TicketKey       string         `json:"ticket_key" db:"ticket_key"`
	PackageId       string         `json:"package_id" db:"package_id"`
	PackageName     string         `json:"package_name" db:"package_name"`
	ExpiredDate     time.Time      `json:"expired_date" db:"expired_date"`
	DayRemain       int            `json:"day_remain" db:"day_remain"`
	SelectedSite    int            `json:"selected_site" db:"selected_site"`
	ActualSite      int            `json:"actual_site" db:"actual_site"`
	StandardSite    int            `json:"standard_site" db:"standard_site"`
	IsLimit         bool           `json:"is_site_limit" db:"is_site_limit"`
	PrevPackageName sql.NullString `json:"prev_package_name" db:"prev_package_name"`
	PrevExpiredDate sql.NullTime   `json:"prev_expired_date" db:"prev_expired_date"`
	IsExpired       bool           `json:"is_expired" db:"is_expired"`
	Storage       	int            `json:"limit_storage" db:"limit_storage"`
}

type PackageFeatureReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
	Route     string `json:"route" db:"route"`
}

type PackageFeatureRes struct {
	Id          string `json:"id" db:"id"`
	FeatureName string `json:"feature_name" db:"feature_name"`
	Route       string `json:"route" db:"route"`
	PackageId   string `json:"package_id" db:"package_id"`
	PackageName string `json:"package_name" db:"package_name"`
	Visible     bool   `json:"visible" db:"visible"`
}

type SitePackageLimitReq struct {
	CompanyId string `json:"company_id" db:"company_id"`
	PackageId string `json:"package_id" db:"package_id"`
}

type SitePackageLimitRes struct {
	IsLimit      bool `json:"is_limit" db:"is_limit"`
	SelectedSite int  `json:"selected_site" db:"selected_site"`
	StandardSite int  `json:"standard_site" db:"standard_site"`
	ActualSite   int  `json:"actual_site" db:"actual_site"`
}

type SelectedSiteReq struct {
	CompanyId string   `json:"company_id" db:"company_id"`
	TicketKey string   `json:"ticket_key" db:"ticket_key"`
	User      string   `json:"user" db:"user"`
	SiteList  []string `json:"site_lists"`
}

type PackageRes struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Detail    string `json:"detail" db:"detail"`
	StdSite   int    `json:"std_qty_site" db:"std_qty_site"`
	DaysOfUse int    `json:"days_of_use" db:"days_of_use"`
	Price     int    `json:"price" db:"price"`
	IsBypass  bool   `json:"is_bypass" db:"is_bypass"`
}
