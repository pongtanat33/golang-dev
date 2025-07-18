package entities

import (
	"database/sql"
	"time"
)

type ActivityUsecase interface {
	GetActivityAvailableSyncBySite(*ActivityListsReq) ([]ActivityListRes, error)
	CreateActivitySync(*ActivitySyncReq) error
	DeleteActivitySync(*ActivitySyncReq) error
	DeleteActivitySyncBySource(*ActivitySyncReq) error
	CheckActivitySyncConnection(*ActivitySyncReq) (bool, error)
	GetActivitySyncBySite(*ActivitySyncBySiteReq) ([]ActivitySyncBySiteRes, error)

	ActivityLists(*ActivityListsReq) ([]ActivityListRes, error)
	CreateActivity(*ActivityCreateReq) (string, error)
	UpdateActivity(*UpdateActivityReq) error
	DeleteActivity(*DeleteActivityReq) error
	ActivityByScope(*ActivityByScopeReq) ([]ActivityByScopeRes, int, error)
	ActivityByScopeEmission(*ActivityByScopeEmissionReq) ([]ActivityByScopeRes, int, error)
	ActivityTransaction(*ActivityTransactionReq) ([]ActivityTransactionRes, int, error)
	ActivityTransactionByDate(*ActivityTransactionByDateReq) ([]ActivityTransactionRes, int, error)
	CreateActivityTransaction(*CreateActivityTransactionReq) (string, error)
	UpdateActivityTransaction(*UpdateActivityTransactionReq) error
	DeleteActivityTransaction(*DeleteActivityTransactionReq) error
	CreateActivityWithCustomizedEmission(*CreateActivityWithCustomizedEmissionReq) error
	UpdateActivityWithCustomizedEmission(*UpdateActivityWithCustomizedEmissionReq) error
	GetActivityValueById(string, string) (*ActivityValueRes, error)
	UpdateActivityStatus(*UpdateActivityStatusReq) error
	CheckBillingInfo(*CustomerInfoReq) (*CheckBillingInfoRes, error)
	CreateActivityAutoScope2(*CreateActivityAutoScope2Req) error
	GetActivityPeaEmissionValueById(string, string) (*ActivityPeaEmissionValueRes, error)
	DisconnectBill(*DisConnectBillReq) error
	UpdateActivityAutoScope2(*UpdateActivityAutoScope2Req) error
	CreateActivityWithRecepiEmission(*CreateActivityWithRecepiEmissionReq) error
	UpdateActivityWithRecepiEmission(*UpdateActivityWithRecepiEmissionReq) error
	//ActivityByName(*ActivityByNameReq) ([]ActivityByNameRes, int, error)
	DuplicateActivity(*DuplicateActivityReq) error
	UpdateBillConnect(*UpsertBillingConnectReq) error
}

type ActivityRepository interface {
	GetActivityAvailableSyncBySite(*ActivityListsReq) ([]ActivityListRes, error)
	GetActivitySyncBySourceId(string) []ActivitySyncRes
	GetActivitySyncByTargetId(string) []ActivitySyncRes
	CreateActivitySync(*ActivitySyncReq) error
	DeleteAllActivitySyncById(*DeleteActivityReq) error
	DeleteActivitySync(*ActivitySyncReq) error
	DeleteActivitySyncBySource(*ActivitySyncReq) error
	CheckActivitySyncConnection(*ActivitySyncReq) (bool, error)
	GetActivitySyncBySite(*ActivitySyncBySiteReq) ([]ActivitySyncBySiteRes, error)
	GetReferenceTransactionById(string) (ReferenceTransactionRes, error)

	ActivityLists(*ActivityListsReq) ([]ActivityListRes, error)
	CreateActivity(*ActivityCreateReq) (string, error)
	ActivityByScope(*ActivityByScopeReq) ([]ActivityByScopeRes, error)
	ActivityByScopeEmission(*ActivityByScopeEmissionReq) ([]ActivityByScopeRes, error)
	ActivityByScopeEmissionRowCount(*ActivityByScopeEmissionReq) (int, error)
	ActivityTransaction(*ActivityTransactionReq) ([]ActivityTransactionRes, error)
	ActivityTransactionByDate(*ActivityTransactionByDateReq) ([]ActivityTransactionRes, error)
	ActivityTransactionByDateRowCount(string, string) (int, error)
	ActivityTransactionRowCount(string) (int, error)
	UpdateActivity(*UpdateActivityReq) error
	DeleteActivity(*DeleteActivityReq) error
	CreateActivityTransaction(*CreateActivityTransactionReq) (string, error)
	UpdateActivityTransaction(*UpdateActivityTransactionReq) (string, error)
	DeleteActivityTransaction(*DeleteActivityTransactionReq) error
	UpdateActivityStatus(*UpdateActivityStatusReq) error
	UpdateActivityAutoScope2(*UpdateActivityAutoScope2Req) error
	//ActivityByName(*ActivityByNameReq) ([]ActivityByNameRes, error)
	//ActivityByNameRowCount(string, string, string) (int, error)
}

type ActivityKey struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

type Activity struct {
	Id           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	EmissionName string `json:"emission_name" db:"emission_name"`
	Description  string `json:"description" db:"description"`
	Source       string `json:"source" db:"source"`
	SyncToScope  string `json:"sync_to_scope"`
}

type ActivityValueRes struct {
	ActivityName      string     `json:"activity_name" db:"activity_name"`
	Description       *string    `json:"description" db:"description"`
	EmissionTypeName  string     `json:"emission_type_name" db:"emission_type_name"`
	EmissionFactorId  string     `json:"emission_factor_id" db:"emission_factor_id"`
	EmissionId        string     `json:"emission_id" db:"emission_id"`
	EmissionName      string     `json:"emission_name" db:"emission_name"`
	Factor            float64    `json:"factor" db:"factor"`
	UnitId            string     `json:"unit_id" db:"unit_id"`
	Unit              string     `json:"unit" db:"unit"`
	IsCalculated      bool       `json:"is_calculated" db:"is_calculated"`
	IsCustomized      bool       `json:"is_customized" db:"is_customized"`
	SiteId            string     `json:"site_id" db:"site_id"`
	EmissionSource    *string    `json:"emission_source" db:"emission_source"`
	ActivityFrequency *string    `json:"activity_frequency" db:"activity_frequency"`
	FrequencyDate     *time.Time `json:"frequency_date" db:"frequency_date"`
	TaxId             *string    `json:"tax_id" db:"tax_id"`
	CaNo              *string    `json:"ca_no" db:"ca_no"`
	PeaNO             *string    `json:"pea_no" db:"pea_no"`
	IsConnected       *bool      `json:"is_connected" db:"is_connected"`
}

type ActivityPeaEmissionValueRes struct {
	Id               string  `json:"activity_id" db:"id"`
	BillId           *string `json:"bill_id" db:"bill_id"`
	ActivityName     string  `json:"activity_name" db:"activity_name"`
	EmissionTypeName string  `json:"emission_type_name" db:"emission_type_name"`
	EmissionName     string  `json:"emission_name" db:"emission_name"`
	Unit             string  `json:"unit" db:"unit"`
	IsConnected      *bool   `json:"is_connected" db:"is_connected"`
	IsAuto           bool    `json:"is_auto"`
}

type ActivityListReq struct {
	SiteId string `json:"site_id" db:"site_id"`
}

type ActivityEmissionTypeRes struct {
	EmissionTypeName string     `json:"emission_type_name" `
	Icon             string     `json:"icon" `
	ActivityList     []Activity `json:"activities,omitempty"`
	Count            string     `json:"count" `
}

type ActivityListRes struct {
	ScopeName     string                    `json:"scope_name" `
	EmissionTypes []ActivityEmissionTypeRes `json:"emission_types,omitempty"`
}

type ActivityCreateReq struct {
	Name                  string  `json:"name" db:"name"`
	Description           *string `json:"description" db:"description"`
	SiteId                string  `json:"site_id" db:"site_id"`
	CompanyId             string  `json:"company_id" db:"company_id"`
	EmissionFactorId      string  `json:"emission_factor_id" db:"emission_factor_id"`
	User                  string
	FrequencyType         *string    `json:"frequency_type"`
	FrequencySelectedDate *time.Time `json:"frequency_selected_date"`
}

type ActivitySyncReq struct {
	SiteId           string   `json:"site_id"`
	SourceActivityId string   `json:"source_activity_id"`
	TargetActivityId string   `json:"target_activity_id"`
	ConversionRate   *float64 `json:"conversion_rate"`
	User             string
}
type DuplicateActivityReq struct {
	ActivityId   string `json:"activity_id" db:"activity_id"`
	ActivityName string `json:"activity_name" db:"activity_name"`
	CompanyId    string `json:"company_id" db:"company_id"`
	User         string `json:"user" db:"user"`
}

type UpdateActivityReq struct {
	Id                    string     `json:"id" db:"id"`
	Name                  string     `json:"name" db:"name"`
	Description           *string    `json:"description" db:"description"`
	CompanyId             string     `json:"company_id" db:"company_id"`
	SiteId                string     `json:"site_id" db:"site_id"`
	EmissionFactorId      string     `json:"emission_factor_id" db:"emission_factor_id"`
	FrequencySelectedDate *time.Time `json:"frequency_selected_date" db:"frequency_selected_date"`
	FrequencyType         string     `json:"frequency_type" db:"frequency_type"`
	UpdateUser            string
}

type ActivityByScopeReq struct {
	SiteId    string  `json:"site_id" db:"site_id"`
	ScopeId   string  `json:"scope_id" db:"scope_id"`
	Lang      string  `json:"lang" db:"lang"`
	StartDate string  `json:"start_date" db:"start_date"`
	EndDate   string  `json:"end_date" db:"end_date"`
	TimeZone  *string `json:"timezone" db:"timezone"`
}

type ActivityByScopeEmissionReq struct {
	SiteId           string  `json:"site_id" db:"site_id"`
	ScopeId          string  `json:"scope_id" db:"scope_id"`
	Offset           int     `json:"offset" db:"offset"`
	PageSize         int     `json:"page_size" db:"page_size"`
	Lang             string  `json:"lang" db:"lang"`
	EmissionTypeId   *string `json:"emission_type_id" db:"emission_type_id"`
	EmissionFactorId *string `json:"emission_factor_id" db:"emission_factor_id"`
	StartDate        string  `json:"start_date" db:"start_date"`
	EndDate          string  `json:"end_date" db:"end_date"`
	TimeZone         *string `json:"timezone" db:"timezone"`
}

type ActivityByScopeRes struct {
	Id                      string    `json:"id" db:"id"`
	EmissionFactorId        string    `json:"emission_factor_id" db:"emission_factor_id"`
	Name                    string    `json:"name" db:"name"`
	Description             *string   `json:"activity_description" db:"activity_description"`
	EmissionTypeName        string    `json:"emission_type_name" db:"emission_type_name"`
	EmissionName            string    `json:"emission_name" db:"emission_name"`
	EmissionDescription     *string   `json:"emission_description" db:"emission_description"`
	Source                  string    `json:"source" db:"source"`
	EmissionFactor          string    `json:"emission_factor" db:"factor"`
	Unit                    string    `json:"unit" db:"unit"`
	Amount                  string    `json:"amount" db:"amount"`
	TotalEmission           string    `json:"total_emission" db:"total_emission"`
	Icon                    string    `json:"icon" db:"icon"`
	Latest                  time.Time `json:"latest" db:"latest"`
	IsCalculated            bool      `json:"is_calculated" db:"is_calculated"`
	IsCustomized            bool      `json:"is_customized" db:"is_customized"`
	Frequency               *string   `json:"frequency" db:"frequency"`
	SyncToScope             *string   `json:"sync_to_scope" db:"sync_to_scope"`
	ActualCount             int       `json:"actual_count" db:"actual_count"`
	ExpectedCount           int       `json:"expected_count" db:"expected_count"`
	PercentComplete         float64   `json:"percent_complete" db:"percent_complete"`
	PercentMissing          float64   `json:"percent_missing" db:"percent_missing"`
	PercentExpectedUntilNow float64   `json:"percent_expected_until_now" db:"percent_expected_until_now"`
	PercentMissingUntilNow  float64   `json:"percent_missing_until_now" db:"percent_missing_until_now"`
	MissingEntries          []string  `json:"missing_entries" db:"missing_entries"`
}

type ActivityTransactionByDateReq struct {
	Id       string `json:"id"`
	Date     string `json:"date"`
	Lang     string
	Offset   int `json:"offset"`
	PageSize int `json:"page_size"`
}

type ActivityTransactionReq struct {
	Id       string `json:"id"`
	Lang     string `json:"lang"`
	Offset   int    `json:"offset"`
	PageSize int    `json:"page_size"`
}

type ActivityTransactionRes struct {
	Id             string     `json:"id" db:"id"`
	ActivityName   string     `json:"activity_name" db:"activity_name"`
	EmissionName   string     `json:"emission_name" db:"emission_name"`
	Factor         string     `json:"factor" db:"factor"`
	FactorUnit     *string    `json:"factor_unit" db:"factor_unit"`
	ResultEmission *string    `json:"result_emission" db:"result_emission"`
	Unit           string     `json:"unit" db:"unit"`
	Amount         string     `json:"amount" db:"amount"`
	ActionDate     string     `json:"action_date" db:"action_date"`
	Evidence       *string    `json:"evidence" db:"evidence"`
	Additional     *string    `json:"additional" db:"additional"`
	Attachment_id  *string    `json:"attachment_id" db:"attachment_id"`
	FileName       *string    `json:"file_name" db:"file_name"`
	UpdatedBy      *string    `json:"updated_by" db:"updated_by"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
}

type ReferenceTransactionRes struct {
	Id                     string     `json:"id" db:"id"`
	ActivityId             string     `json:"activity_id" db:"activity_id"`
	ResultEmission         *string    `json:"result_emission" db:"result_emission"`
	Amount                 string     `json:"amount" db:"amount"`
	ActionDate             string     `json:"action_date" db:"action_date"`
	Evidence               *string    `json:"evidence" db:"evidence"`
	Additional             *string    `json:"additional" db:"additional"`
	ReferenceTransactionId *string    `json:"reference_transaction_id" db:"reference_transaction_id"`
	UpdatedBy              *string    `json:"updated_by" db:"updated_by"`
	UpdatedAt              *time.Time `json:"updated_at" db:"updated_at"`
}

type CreateActivityTransactionReq struct {
	ActivityId             string  `json:"activity_id" db:"activity_id"`
	Amount                 float64 `json:"amount" db:"amount"`
	ActionDate             string  `json:"action_date" db:"action_date"`
	Evidence               string  `json:"evidence" db:"evidence"`
	Additional             string  `json:"additional" db:"additional"`
	ReferenceTransactionId string  `json:"reference_transaction_id" db:"reference_transaction_id"`
	User                   string  `json:"user" db:"user"`
	CompanyId              string
}

type UpdateActivityTransactionReq struct {
	Id         string  `json:"id" db:"id"`
	ActivityId string  `json:"activity_id" db:"activity_id"`
	Amount     float64 `json:"amount" db:"amount"`
	ActionDate string  `json:"action_date" db:"action_date"`
	Evidence   string  `json:"evidence" db:"evidence"`
	Additional string  `json:"additional" db:"additional"`
	User       string  `json:"user" db:"user"`
	CompanyId  string
}

type DeleteActivityReq struct {
	Id         string `json:"id" db:"id"`
	SiteId     string `json:"site_id" db:"site_id"`
	DeleteUser string `db:"user"`
}

type DeleteActivityTransactionReq struct {
	Id         string `json:"id" db:"id"`
	ActivityId string `json:"activity_id" db:"activity_id"`
	DeleteUser string `db:"user"`
}

type CreateActivityWithCustomizedEmissionReq struct {
	IsCalculated          bool     `json:"is_calculated" db:"is_calculated"`
	EmissionName          string   `json:"emission_name" db:"emission_name"`
	Factor                *float64 `json:"factor" db:"factor"`
	FactorUnitId          *string  `json:"unit_id" db:"unit_id"`
	ActivityName          string   `json:"activity_name" db:"activity_name"`
	ActivityDescription   *string  `json:"activity_description" db:"activity_description"`
	Source                string   `json:"source" db:"source"`
	SiteId                string   `json:"site_id" db:"site_id"`
	CompanyId             string   `json:"company_id" db:"company_id"`
	TypeId                string   `json:"type_id" db:"type_id"`
	User                  string
	FrequencyType         *string    `json:"frequency_type"`
	FrequencySelectedDate *time.Time `json:"frequency_selected_date"`
}

type UpdateActivityWithCustomizedEmissionReq struct {
	Id                    string   `json:"id" db:"id"`
	IsCalculated          bool     `json:"is_calculated" db:"is_calculated"`
	EmissionName          string   `json:"emission_name" db:"emission_name"`
	Factor                *float64 `json:"factor" db:"factor"`
	FactorUnitId          *string  `json:"unit_id" db:"unit_id"`
	ActivityName          string   `json:"activity_name" db:"activity_name"`
	ActivityDescription   *string  `json:"activity_description" db:"activity_description"`
	Source                string   `json:"source" db:"source"`
	SiteId                string   `json:"site_id" db:"site_id"`
	CompanyId             string   `json:"company_id" db:"company_id"`
	TypeId                string   `json:"type_id" db:"type_id"`
	User                  string
	FrequencyType         *string    `json:"frequency_type"`
	FrequencySelectedDate *time.Time `json:"frequency_selected_date"`
}

type UpdateActivityStatusReq struct {
	Id       string `json:"id" `
	SiteId   string `json:"site_id"`
	User     string
	IsActive bool
}

type CreateActivityAutoScope2Req struct {
	Name                  string          `json:"name"`
	ActivityDescription   *string         `json:"description" db:"description"`
	BillingInfo           CustomerInfoReq `json:"billing_info"`
	FrequencyType         *string         `json:"frequency_type"`
	FrequencySelectedDate *time.Time      `json:"frequency_selected_date"`
	SiteId                string          `json:"site_id" `
	IsAuto                bool            `json:"is_auto" `
	User                  string
	CompanyId             string
}

type UpdateActivityAutoScope2Req struct {
	Id          string          `json:"id" `
	Name        string          `json:"name"`
	BillingInfo CustomerInfoReq `json:"billing_info"`
	SiteId      string          `json:"site_id" `
	CompanyId   string
	User        string
}

type CreateActivityWithRecepiEmissionReq struct {
	IsCalculated          bool         `json:"is_calculated" db:"is_calculated"`
	EmissionName          string       `json:"emission_name" db:"emission_name"`
	FactorType            []FactorType `json:"factortype"`
	FactorUnitId          *string      `json:"unit_id" db:"unit_id"`
	ActivityName          string       `json:"activity_name" db:"activity_name"`
	ActivityDescription   *string      `json:"activity_description" db:"activity_description"`
	Source                string       `json:"source" db:"source"`
	CompanyId             string       `json:"company_id" db:"company_id"`
	SiteId                string       `json:"site_id" db:"site_id"`
	TypeId                string       `json:"type_id" db:"type_id"`
	User                  string
	FrequencyType         *string    `json:"frequency_type"`
	FrequencySelectedDate *time.Time `json:"frequency_selected_date"`
}

type UpdateActivityWithRecepiEmissionReq struct {
	Id                    string       `json:"id" db:"id"`
	IsCalculated          bool         `json:"is_calculated" db:"is_calculated"`
	EmissionName          string       `json:"emission_name" db:"emission_name"`
	FactorType            []FactorType `json:"factortype"`
	FactorUnitId          *string      `json:"unit_id" db:"unit_id"`
	ActivityName          string       `json:"activity_name" db:"activity_name"`
	ActivityDescription   *string      `json:"activity_description" db:"activity_description"`
	Source                string       `json:"source" db:"source"`
	CompanyId             string       `json:"company_id" db:"company_id"`
	SiteId                string       `json:"site_id" db:"site_id"`
	TypeId                string       `json:"type_id" db:"type_id"`
	User                  string
	FrequencyType         *string    `json:"frequency_type"`
	FrequencySelectedDate *time.Time `json:"frequency_selected_date"`
}

type ActivityByNameReq struct {
	SiteId  string `json:"site_id" db:"site_id"`
	Name    string `json:"name" db:"name"`
	ScopeId string `json:"scope_id" db:"scope_id"`
	Lang    string `json:"lang" db:"lang"`
}

type ActivityByNameRes struct {
	Id                  string    `json:"id" db:"id"`
	EmissionFactorId    string    `json:"emission_factor_id" db:"emission_factor_id"`
	Name                string    `json:"name" db:"name"`
	Description         *string   `json:"activity_description" db:"activity_description"`
	EmissionTypeName    string    `json:"emission_type_name" db:"emission_type_name"`
	EmissionName        string    `json:"emission_name" db:"emission_name"`
	EmissionDescription *string   `json:"emission_description" db:"emission_description"`
	Source              string    `json:"source" db:"source"`
	EmissionFactor      string    `json:"emission_factor" db:"factor"`
	Unit                string    `json:"unit" db:"unit"`
	Amount              string    `json:"amount" db:"amount"`
	TotalEmission       string    `json:"total_emission" db:"total_emission"`
	Icon                string    `json:"icon" db:"icon"`
	Latest              time.Time `json:"latest" db:"latest"`
	IsCalculated        bool      `json:"is_calculated" db:"is_calculated"`
	IsCustomized        bool      `json:"is_customized" db:"is_customized"`
	Frequency           string    `json:"frequency" db:"frequency"`
}

type ActivitySyncBySiteReq struct {
	CompanyId string
	SiteId    string `json:"site_id"`
	Lang      string
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ActivityListsReq struct {
	CompanyId string
	SiteId    string `json:"site_id"`
	Lang      string
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ActivitySyncBySiteKey struct {
	ActivityId           string         `db:"activity_id"`
	Activity             string         `db:"activity"`
	Unit                 string         `db:"unit"`
	Scope                string         `db:"scope"`
	EmissionFactor       float64        `db:"emission_factor"`
	SyncTotalTransaction int            `db:"sync_total_transaction"`
	SyncTotalAmount      float64        `db:"sync_total_amount"`
	SyncTotalEmission    float64        `db:"sync_total_emission"`
	SyncScopes           sql.NullString `db:"sync_scopes"`
}

type ActivitySyncBySiteRes struct {
	ActivityId           string         `json:"activity_id"`
	Activity             string         `json:"activity"`
	Unit                 string         `json:"unit"`
	Scope                string         `json:"scope"`
	EmissionFactor       float64        `json:"emission_factor"`
	SyncTotalTransaction int            `json:"sync_total_transaction"`
	SyncTotalAmount      float64        `json:"sync_total_amount"`
	SyncTotalEmission    float64        `json:"sync_total_emission"`
	SyncScopes           []SyncScopeRes `json:"sync_scopes"`
}

type SyncScopeRes struct {
	SubScope       string             `json:"sub_scope"`
	ScopeId        string             `json:"scope_id"`
	SyncActivities []SyncActivityData `json:"sync_activities"`
}

type SyncActivityData struct {
	SubActivityId        string    `json:"sub_activity_id"`
	SubActivity          string    `json:"sub_activity"`
	SubConversionRate    float64   `json:"sub_conversion_rate"`
	SubUnit              string    `json:"sub_unit"`
	SubAmount            float64   `json:"sub_amount"`
	SubTotalEmission     float64   `json:"sub_total_emission"`
	SubTransactions      int       `json:"sub_transactions"`
	TotalSourceTx        int       `json:"total_source_tx"`
	SyncedTx             int       `json:"synced_tx"`
	SyncPercent          float64   `json:"sync_percent"`
	StartSyncAt          time.Time `json:"start_sync_at"`
	SyncedAmount         float64   `json:"synced_amount"`
	SyncedResultEmission float64   `json:"synced_result_emission"`
}

type ActivitySyncRes struct {
	SourceActivityId string    `db:"source_activity_id"`
	TargetActivityId string    `db:"target_activity_id"`
	Scope            string    `db:"scope"`
	StartSyncAt      time.Time `db:"start_sync_at"`
	ConversionRate   float64   `db:"conversion_rate"`
}
