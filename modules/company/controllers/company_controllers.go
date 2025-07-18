package controllers

import (
	"bufferbox_backend_go/constants"
	"bufferbox_backend_go/entities"
	"bufferbox_backend_go/logs"
	"bufferbox_backend_go/middlewares"
	"bufferbox_backend_go/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type companyController struct {
	CompanyUsecase entities.CompanyUsecase
}

func NewCompanyController(group *fiber.Group, companyUsecase entities.CompanyUsecase) {
	company := &companyController{
		CompanyUsecase: companyUsecase,
	}
	group.Use(middlewares.VerifyToken)

	group.Get("/company", middlewares.TicketExpiredVerify, company.CompanyValue)
	group.Post("/company", company.CompanyCreate)
	group.Patch("/company", middlewares.RoleCompanyVerify, company.CompanyUpdate)
	group.Delete("/company", middlewares.RoleCompanyVerify, company.CompanyDelete)

	group.Get("/company-group-by-name", company.CompanyGroupsByName)
	group.Get("/company-group-lists", company.CompanyGroups)
	group.Get("/company-site-lists", company.CompanySites)
	group.Get("/company-selected-site-lists", company.CompanySelectedSites)
	group.Get("/company-lists", company.CompanyLists)
	group.Get("/company-group-site", company.CompanyGroupSites)

	group.Get("/:company/target_base/", company.CompanyTargetBase)
	group.Get("/:company/target_base/:base", company.CompanyTargetBaseById)
	group.Post("/:company/target_base/", company.CompanyTargetBaseUpsert)
	group.Patch("/:company/target_base/", company.CompanyTargetBaseUpsert)
	group.Delete("/:company/target_base/:base", company.CompanyTargetBaseDelete)

	group.Get("/:company/target/", company.CompanyTargetAll)
	group.Get("/:company/target/:target", company.CompanyTargetById)

	//Package Feature Verify + Role Company Verify
	group.Post("/:company/target/", middlewares.RoleCompanyVerify, middlewares.PackageFeatureVerify, company.CompanyTargetUpsert)
	group.Patch("/:company/target/", middlewares.RoleCompanyVerify, middlewares.PackageFeatureVerify, company.CompanyTargetUpsert)
	group.Delete("/:company/target/:target", middlewares.RoleCompanyVerify, middlewares.PackageFeatureVerify, company.CompanyTargetDelete)
	group.Patch("/company-main-target", middlewares.RoleCompanyVerify, company.CompanyMainTargetUpdate)

	group.Get("/company-member", company.CompanyMember)
	group.Delete("/company-member", middlewares.RoleCompanyVerify, company.CompanyMemberDelete)
	group.Patch("/company-member", middlewares.RoleCompanyVerify, company.CompanyMemberChange)

	//group.Get("/company-test", middlewares.TicketExpiredVerify, company.TestCompanyValue)
}

func (cp *companyController) CompanyMemberDelete(c *fiber.Ctx) error {
	req := new(entities.MemberDeleteReq)
	req.Email = c.Query("email")
	req.TablePermissionId = c.Locals("company_id").(string)
	req.IsActive = c.Query("is_active")
	req.User = c.Locals("usr_id").(string)
	req.TableType = constants.C_ROLE_TYPE
	err := cp.CompanyUsecase.CompanyMemberDelete(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_member_delete_failed", fiber.Map{
			"result": err.Error(),
		})
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_member_delete_successfully", nil)
}

func (cp *companyController) CompanyMemberChange(c *fiber.Ctx) error {

	req := new(entities.UpsertUserPermissionReq)
	if err := c.BodyParser(&req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "body_not_match", nil)
	}
	req.TablePermissionId = c.Locals("company_id").(string)
	req.UserId = c.Locals("usr_id").(string)
	err := cp.CompanyUsecase.CompanyMemberChange(req)

	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_member_update_failed", fiber.Map{
			"result": err.Error(),
		})
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_member_delete_successfully", nil)
}

func (cp *companyController) CompanyMainTargetUpdate(c *fiber.Ctx) error {

	req := new(entities.CompanyMainTargetUpdateReq)
	req.Id = c.Query("id")
	req.User = c.Locals("usr_id").(string)
	err := cp.CompanyUsecase.CompanyMainTargetUpdate(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_main_target_update_failed", fiber.Map{
			"result": err.Error(),
		})
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_main_target_update_successfully", nil)
}

func (cp *companyController) CompanyTargetUpsert(c *fiber.Ctx) error {
	req := new(entities.CompanyTargetUpsertReq)
	if err := c.BodyParser(&req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "body_not_match", nil)
	}
	companyID := c.Params("company")
	req.CompanyId = companyID
	req.User = c.Locals("usr_id").(string)
	result, err := cp.CompanyUsecase.CompanyTargetUpsert(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_target_upsert_failed", fiber.Map{
			"result": err.Error(),
		})
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_target_upsert_successfully", fiber.Map{
		"result": result,
	})
}

func (cp *companyController) CompanyTargetBaseUpsert(c *fiber.Ctx) error {
	req := new(entities.CompanyTargetBaseUpsertReq)
	if err := c.BodyParser(&req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "body_not_match", nil)
	}
	req.User = c.Locals("usr_id").(string)
	result, err := cp.CompanyUsecase.CompanyTargetBaseUpsert(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_target_base_upsert_failed", fiber.Map{
			"result": err.Error(),
		})
	}

	return utils.HandleResponse(c, fiber.StatusOK, "company_target_base_upsert_successfully", fiber.Map{
		"result": result,
	})
}

func (cp *companyController) CompanyTargetBaseDelete(c *fiber.Ctx) error {
	companyID := c.Params("company")
	BaseId := c.Params("base")

	req := new(entities.CompanyTargetBaseDeleteReq)
	req.Id = BaseId
	req.CompanyId = companyID
	req.User = c.Locals("usr_id").(string)

	err := cp.CompanyUsecase.CompanyTargetBaseDelete(req)
	if err != nil {
		if err.Error() == "company_target_of_this_base_is_active" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "company_target_base_delete_failed", nil)
		} else {
			return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_target_base_delete_failed", nil)
		}
	}

	return utils.HandleResponse(c, fiber.StatusOK, "company_target_base_delete_successfully", nil)
}

func (cp *companyController) CompanyTargetDelete(c *fiber.Ctx) error {
	companyID := c.Params("company")
	TargetId := c.Params("target")

	req := new(entities.CompanyTargetDeleteReq)
	req.Id = TargetId
	req.CompanyId = companyID
	req.User = c.Locals("usr_id").(string)
	err := cp.CompanyUsecase.CompanyTargetDelete(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_target_delete_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_target_delete_successfully", nil)
}

func (cp *companyController) CompanyTargetById(c *fiber.Ctx) error {
	companyID := c.Params("company")
	targetID := c.Params("target")
	req := new(entities.CompanyTargetByIdReq)
	req.CompanyId = companyID
	req.TargetId = targetID
	result, err := cp.CompanyUsecase.CompanyTargetById(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_target_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_target_successfully", fiber.Map{
		"result": result,
	})
}

func (cp *companyController) CompanyTargetAll(c *fiber.Ctx) error {
	companyID := c.Params("company")
	req := new(entities.CompanyTargetBaseReq)
	req.CompanyId = companyID
	result, err := cp.CompanyUsecase.CompanyTargetAll(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_target_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_target_successfully", fiber.Map{
		"result": result,
	})

}

func (cp *companyController) CompanyTargetBaseById(c *fiber.Ctx) error {
	companyID := c.Params("company")
	baseID := c.Params("base")
	req := new(entities.CompanyTargetBaseByIdReq)
	req.CompanyId = companyID
	req.BaseId = baseID
	result, err := cp.CompanyUsecase.CompanyTargetBaseById(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_target_base_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_target_base_successfully", fiber.Map{
		"result": result,
	})
}

func (cp *companyController) CompanyTargetBase(c *fiber.Ctx) error {
	companyID := c.Params("company")
	req := new(entities.CompanyTargetBaseReq)
	req.CompanyId = companyID
	result, err := cp.CompanyUsecase.CompanyTargetBaseAll(req)
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_target_base_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_target_base_successfully", fiber.Map{
		"result": result,
	})
}

func (cp *companyController) CompanyDelete(c *fiber.Ctx) error {
	req := new(entities.CompanyReq)
	req.Id = c.Query("id")
	req.User = c.Locals("usr_id").(string)

	err := cp.CompanyUsecase.CompanyDelete(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_delete_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_delete_successfully", nil)
}

func (cp *companyController) CompanyUpdate(c *fiber.Ctx) error {
	req := new(entities.CompanyReq)
	req.User = c.Locals("usr_id").(string)
	if err := c.BodyParser(&req); err != nil {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "body_not_match", nil)
	}
	err := cp.CompanyUsecase.CompanyUpdate(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_update_failed", fiber.Map{
			"result": err.Error(),
		})
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_update_successfully", nil)
}

func (cp *companyController) CompanyGroupsByName(c *fiber.Ctx) error {
	companyId := c.Query("company_id")
	name := c.Query("name")
	offset, _ := strconv.Atoi(c.Query("offset"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

	req := &entities.CompanyGroupByNameReq{
		CompanyId: companyId,
		Name:      name,
		Offset:    offset,
		PageSize:  pageSize,
		FromDate:  fromDate,
		ToDate:    toDate,
		User:      c.Locals("usr_id").(string),
	}
	res, totalRows, err := cp.CompanyUsecase.CompanyGroupsByName(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_group_by_name_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_group_by_name_successfully", fiber.Map{
		"result":     res,
		"total_rows": totalRows,
	})
}

func (cp *companyController) CompanyGroups(c *fiber.Ctx) error {
	companyId := c.Query("company_id")
	offset, _ := strconv.Atoi(c.Query("offset"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")
	req := &entities.CompanyGroupReq{
		CompanyId: companyId,
		Offset:    offset,
		PageSize:  pageSize,
		FromDate:  fromDate,
		ToDate:    toDate,
		User:      c.Locals("usr_id").(string),
	}
	res, totalRows, err := cp.CompanyUsecase.CompanyGroups(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_group_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_group_successfully", fiber.Map{
		"result":     res,
		"total_rows": totalRows,
	})
}

func (cp *companyController) CompanySites(c *fiber.Ctx) error {
	companyId := c.Query("company_id")
	var groupId *string
	if groupIDStr := c.Query("group_id"); groupIDStr != "" {
		groupId = &groupIDStr
	}
	offset, _ := strconv.Atoi(c.Query("offset"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")
	req := &entities.CompanySiteReq{
		CompanyId: companyId,
		GroupId:   groupId,
		Offset:    offset,
		PageSize:  pageSize,
		FromDate:  fromDate,
		ToDate:    toDate,
	}
	res, totalRows, err := cp.CompanyUsecase.CompanySites(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_sites_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_site_successfully", fiber.Map{
		"result":     res,
		"total_rows": totalRows,
	})
}

func (cp *companyController) CompanySelectedSites(c *fiber.Ctx) error {
	companyId := c.Query("company_id")
	var groupId *string
	if groupIDStr := c.Query("group_id"); groupIDStr != "" {
		groupId = &groupIDStr
	}
	offset, _ := strconv.Atoi(c.Query("offset"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")
	req := &entities.CompanySiteReq{
		CompanyId: companyId,
		GroupId:   groupId,
		Offset:    offset,
		PageSize:  pageSize,
		FromDate:  fromDate,
		ToDate:    toDate,
	}
	res, totalRows, err := cp.CompanyUsecase.CompanySelectedSites(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_selected_sites_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_selected_site_successfully", fiber.Map{
		"result":     res,
		"total_rows": totalRows,
	})
}

func (cp *companyController) CompanyValue(c *fiber.Ctx) error {
	companyId := c.Query("company_id")

	if companyId == "" {
		companyLists, lastVisitedCompany, err := cp.CompanyUsecase.CompanyLists(c.Locals("usr_id").(string))
		if err != nil {
			logs.Error(err)
			return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_lists_failed", nil)
		}
		if lastVisitedCompany != nil {
			companyId = lastVisitedCompany.Id
		} else {
			companyId = companyLists[0].Id
		}

	}
	res, err := cp.CompanyUsecase.CompanyValue(companyId, c.Locals("usr_id").(string))
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_value_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_value_successfully", fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "company_value_successfully",
		"result":      res,
	})
}

// func (cp *companyController) TestCompanyValue(c *fiber.Ctx) error {
// 	companyId := c.Query("company_id")
// 	res, err := cp.CompanyUsecase.TestCompanyValue(companyId, c.Locals("usr_id").(string))
// 	if err != nil {
// 		logs.Error(err)
// 		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_value_failed", nil)
// 	}
// 	return utils.HandleResponse(c, fiber.StatusOK, "company_value_successfully", fiber.Map{
// 		"status":      "OK",
// 		"status_code": fiber.StatusOK,
// 		"message":     "company_value_successfully",
// 		"result":      res,
// 	})
// }

func (cp *companyController) CompanyLists(c *fiber.Ctx) error {
	companyLists, lastVisitedCompany, err := cp.CompanyUsecase.CompanyLists(c.Locals("usr_id").(string))
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_lists_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_lists_successfully", fiber.Map{
		"status":               "OK",
		"status_code":          fiber.StatusOK,
		"message":              "company_lists_successfully",
		"company_lists":        companyLists,
		"last_visited_company": lastVisitedCompany,
	})
}

func (cp *companyController) CompanyCreate(c *fiber.Ctx) error {
	req := new(entities.CompanyReq)
	req.User = c.Locals("usr_id").(string)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	_, err := cp.CompanyUsecase.CompanyCreate(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_created_successfully", nil)
}

func (cp *companyController) CompanyMember(c *fiber.Ctx) error {
	req := new(entities.CompanyMemberReq)
	req.CompanyId = c.Query("company")
	res, err := cp.CompanyUsecase.CompanyMember(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_member_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_member_successfully", fiber.Map{
		"result": res,
	})
}

func (cp *companyController) CompanyGroupSites(c *fiber.Ctx) error {

	req := new(entities.CompanySiteReq)
	req.CompanyId = c.Locals("company_id").(string)
	groupID := c.Query("group_id")
	if groupID != "" {
		req.GroupId = &groupID
	}
	req.Offset, _ = strconv.Atoi(c.Query("offset"))
	req.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	req.FromDate = c.Query("from_date")
	req.ToDate = c.Query("to_date")
	req.User = c.Locals("usr_id").(string)

	res, totalRows, err := cp.CompanyUsecase.CompanyGroupSites(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "company_group_site_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "company_group_site_successfully", fiber.Map{
		"result":     res,
		"total_rows": totalRows,
	})
}
