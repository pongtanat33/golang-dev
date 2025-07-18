package entities

type IndustryTypeUsecase interface {
	BusinessSectorMasterData(string) ([]BusinessSectorRes, error)
	BusinessCategoryMasterData(*BusinessCategoryReq) ([]BusinessCategoryRes, error)
	AllBusinessCategoryMasterData(string) ([]BusinessCategoryRes, error)
}

type IndustryTypeRepository interface {
	BusinessSectorMasterData(string) ([]BusinessSectorRes, error)
	BusinessCategoryMasterData(*BusinessCategoryReq) ([]BusinessCategoryRes, error)
	AllBusinessCategoryMasterData(string) ([]BusinessCategoryRes, error)
}

type BusinessCategoryRes struct {
	Id                   string `json:"id"`
	Business_category string `json:"business_category"`
	Business_category_description *string `json:"business_category_description"`
	// Business_category_th string `json:"business_category_th" db:"business_category_th"`
}

type BusinessCategoryReq struct {
	Business_sector_id string `json:"business_sector_id" db:"business_sector_id"`
	Lang string 
}

type BusinessSectorRes struct {
	Id                 string `json:"id"`
	Business_sector 	string `json:"business_sector"`
}
