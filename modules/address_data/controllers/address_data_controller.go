package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type addressDataController struct {
	AddressDataUsecase entities.AddressDataUsecase
}

func NewAddressDataController(group *fiber.Group, addressDataUsecase entities.AddressDataUsecase) {
	addressData := &addressDataController{
		AddressDataUsecase: addressDataUsecase,
	}

}

func (a *addressDataController) ProvinceMasterData(c *fiber.Ctx) error {

	postcode := c.Query("postcode")
	req := &entities.ProvinceReq{PostCode: postcode}
	if postcode == "" {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "postcode is required", nil)
	}

	provinces, err := a.AddressDataUsecase.ProvinceMasterData(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "province_master_data_error", nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "province_master_data_success", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "provinces_query_success",
		"result":      provinces,
	})
}

func (a *addressDataController) CountryMasterData(c *fiber.Ctx) error {
	Countrys, err := a.AddressDataUsecase.CountryMasterData()
	if err != nil {
		logs.Error(err)

		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "country_master_data_error", nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "country_master_data_success", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "country_master_data_success",
		"result":      Countrys,
	})

}
