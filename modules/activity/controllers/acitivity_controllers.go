package controllers

import (
	"bufferbox_backend_go/entities"
	"bufferbox_backend_go/logs"
	"bufferbox_backend_go/middlewares"
	"bufferbox_backend_go/pkg/utils"
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type activityController struct {
	ActivityUsecase entities.ActivityUsecase
}

func NewActivityController(group *fiber.Group, activityUsecase entities.ActivityUsecase) {
	activity := &activityController{
		ActivityUsecase: activityUsecase,
	}
	group.Get("/activity-lists", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.ActivityLists)

	group.Get("/activity", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.ActivityValueById)
	group.Post("/activity", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.CreateActivity)
	group.Patch("/activity", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.UpdateActivity)
	group.Delete("/activity", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.DeleteActivity)

	// Activity Sync to scope3
	group.Get("/activity-available-sync", middlewares.VerifyToken, activity.GetActivityAvailableSyncBySite)
	group.Get("/activity-sync-connection", middlewares.VerifyToken, activity.CheckActivitySyncConnection)
	group.Delete("/activity-sync-by-source", middlewares.VerifyToken, activity.DeleteActivitySyncBySource)

	group.Post("/activity-sync", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.CreateActivitySync)
	group.Get("/activity-sync", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.GetActivitySyncBySite)
	group.Delete("/activity-sync", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.DeleteActivitySync)

	group.Post("/duplicate-activity", middlewares.VerifyToken, activity.DuplicateActivity)
	group.Patch("/restore-activity", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.RestoreActivity)

	group.Post("/activity-by-scope", middlewares.VerifyToken, activity.ActivityByScope)
	// group.Post("/activity-by-scope-emission", middlewares.VerifyToken, activity.ActivityByScopeEmission)

	group.Get("/activity-transaction", middlewares.VerifyToken, activity.ActivityTransaction)
	group.Post("/activity-transaction", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.CreateActivityTransaction)
	group.Patch("/activity-transaction", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.UpdateActivityTransaction)
	group.Delete("/activity-transaction", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.DeleteActivityTransaction)
	group.Get("/activity-transaction-by-date", middlewares.VerifyToken, activity.ActivityTransactionByDate)
	group.Post("/activity-with-custom-emission", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.CreateActivityWithCustomizedEmission)
	group.Patch("/activity-with-custom-emission", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.UpdateActivityWithCustomizedEmission)
	group.Get("/scope-lists", middlewares.VerifyToken, activity.Scopelist)

	//Package Feature Verify + Role Company Verify
	group.Post("/check-billing-info", middlewares.VerifyToken, middlewares.PackageFeatureVerify, activity.CheckBillingInfo)
	group.Get("/activity-auto-scope2", middlewares.VerifyToken, middlewares.PackageFeatureVerify, activity.ActivityPeaEmissionValueById)
	group.Post("/activity-auto-scope2", middlewares.VerifyToken, middlewares.RoleSiteVerify, middlewares.PackageFeatureVerify, activity.CreateActivityAutoScope2)
	group.Patch("/activity-auto-scope2", middlewares.VerifyToken, middlewares.RoleSiteVerify, middlewares.PackageFeatureVerify, activity.UpdateActivityAutoScope2)
	group.Delete("/disconnect-bill", middlewares.VerifyToken, middlewares.RoleSiteVerify, middlewares.PackageFeatureVerify, activity.DisconnectBill)
	group.Patch("/bill-connect", middlewares.VerifyToken, activity.UpdateBillConnect)
	group.Post("/activity-with-recepi-emission", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.CreateActivityWithRecepiEmission)
	group.Patch("/activity-with-recepi-emission", middlewares.VerifyToken, middlewares.RoleSiteVerify, activity.UpdateActivityWithRecepiEmission)

}

func (a *activityController) GetActivityAvailableSyncBySite(c *fiber.Ctx) error {
	req := new(entities.ActivityListsReq)
	req.Lang = c.Locals("lang").(string)
	req.SiteId = c.Query("site_id")
	req.StartDate = c.Query("start_date")
	req.EndDate = c.Query("end_date")
	activityLists, err := a.ActivityUsecase.GetActivityAvailableSyncBySite(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_available_sync_lists_failed", nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "activity_available_sync_lists_successfully", fiber.Map{
		"status":         "OK",
		"status_code":    fiber.StatusOK,
		"message":        "activity_available_sync_lists_successfully",
		"activity_lists": activityLists,
	})
}

func (a *activityController) DeleteActivitySync(c *fiber.Ctx) error {
	req := new(entities.ActivitySyncReq)
	req.SiteId = c.Query("site_id")
	req.TargetActivityId = c.Query("target_activity_id")
	req.User = c.Locals("usr_id").(string)

	if err := a.ActivityUsecase.DeleteActivitySync(req); err != nil {
		logs.Error(err)
		errResult := err.Error()
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "delete_activity_sync_failed", fiber.Map{
			"result": errResult,
		})
	}
	return utils.HandleResponse(c, fiber.StatusOK, "delete_activity_sync_successfully", nil)
}

func (a *activityController) DeleteActivitySyncBySource(c *fiber.Ctx) error {
	req := new(entities.ActivitySyncReq)
	req.SiteId = c.Query("site_id")
	req.SourceActivityId = c.Query("source_activity_id")
	req.TargetActivityId = c.Query("target_activity_id")
	req.User = c.Locals("usr_id").(string)

	if err := a.ActivityUsecase.DeleteActivitySyncBySource(req); err != nil {
		logs.Error(err)
		errResult := err.Error()
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "delete_activity_sync_by_source_failed", fiber.Map{
			"result": errResult,
		})
	}
	return utils.HandleResponse(c, fiber.StatusOK, "delete_activity_sync_by_source_successfully", nil)
}

func (a *activityController) GetActivitySyncBySite(c *fiber.Ctx) error {
	req := new(entities.ActivitySyncBySiteReq)
	req.Lang = c.Locals("lang").(string)
	req.CompanyId = c.Locals("company_id").(string)
	req.SiteId = c.Query("site_id")
	req.StartDate = c.Query("start_date")
	req.EndDate = c.Query("end_date")

	res, err := a.ActivityUsecase.GetActivitySyncBySite(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_sync_list_failed", nil)

	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_sync_list_successfully", fiber.Map{
		"result":     res,
		"total_rows": len(res),
	})
}

func (a *activityController) CheckActivitySyncConnection(c *fiber.Ctx) error {
	req := new(entities.ActivitySyncReq)
	req.SiteId = c.Query("site_id")
	req.SourceActivityId = c.Query("source_activity_id")
	req.TargetActivityId = c.Query("target_activity_id")

	res, err := a.ActivityUsecase.CheckActivitySyncConnection(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_sync_connection_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_sync_connection_successfully", fiber.Map{
		"is_connected": res,
	})
}

func (a *activityController) CreateActivitySync(c *fiber.Ctx) error {
	req := new(entities.ActivitySyncReq)
	req.User = c.Locals("usr_id").(string)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	if err := a.ActivityUsecase.CreateActivitySync(req); err != nil {
		logs.Error(err)
		errResult := err.Error()
		return utils.HandleResponse(c, fiber.ErrConflict.Code, "create_activity_sync_failed", fiber.Map{
			"result": errResult,
		})
	}
	return utils.HandleResponse(c, fiber.StatusOK, "create_activity_sync_successfully", nil)
}

func (a *activityController) UpdateActivityAutoScope2(c *fiber.Ctx) error {
	req := new(entities.UpdateActivityAutoScope2Req)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Locals("company_id").(string)
	err := a.ActivityUsecase.UpdateActivityAutoScope2(req)
	if err != nil {
		if err.Error() == "bill_already_used" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "update_activity_auto_scope2_failed", fiber.Map{
				"result": "bill_already_used",
			})
		} else {
			logs.Error(err)
			return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "update_activity_auto_scope2_failed", nil)
		}
	}
	return utils.HandleResponse(c, fiber.StatusOK, "update_activity_auto_scope2_successfully", nil)

}

func (a *activityController) CreateActivityAutoScope2(c *fiber.Ctx) error {
	req := new(entities.CreateActivityAutoScope2Req)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Locals("company_id").(string)
	err := a.ActivityUsecase.CreateActivityAutoScope2(req)
	if err != nil {
		if err.Error() == "bill_already_used" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "create_activity_auto_scope2_failed", fiber.Map{
				"result": "bill_already_used",
			})
		} else {
			logs.Error(err)
			return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "create_activity_auto_scope2_failed", nil)
		}
	}
	return utils.HandleResponse(c, fiber.StatusOK, "create_activity_auto_scope2_successfully", nil)

}

func (a *activityController) CheckBillingInfo(c *fiber.Ctx) error {
	req := new(entities.CustomerInfoReq)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}
	result, err := a.ActivityUsecase.CheckBillingInfo(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "check_billing_info_failed", nil)
	}

	if result.BillNotUsed {
		return utils.HandleResponse(c, fiber.StatusOK, "check_billing_info_failed", fiber.Map{
			"result": "bill_already_used",
		})
	}

	if !result.BillInfoValid {
		return utils.HandleResponse(c, fiber.StatusOK, "check_billing_info_failed", fiber.Map{
			"result": "bill_not_valid",
		})
	}

	return utils.HandleResponse(c, fiber.StatusOK, "check_billing_info_successfully", fiber.Map{
		"result": "bill_valid",
	})
}

func (a *activityController) RestoreActivity(c *fiber.Ctx) error {
	req := new(entities.UpdateActivityStatusReq)
	req.User = c.Locals("usr_id").(string)
	req.IsActive = true
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}
	err := a.ActivityUsecase.UpdateActivityStatus(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "restore_activity_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "restore_activity_successfully", nil)
}

func (a *activityController) Scopelist(c *fiber.Ctx) error {
	result, err := utils.GetScopeList()
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "scope_lists_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "scope_lists_successfully", fiber.Map{
		"result": result,
	})

}

func (a *activityController) ActivityValueById(c *fiber.Ctx) error {
	activity, err := a.ActivityUsecase.GetActivityValueById(c.Query("activity_id"), c.Locals("lang").(string))

	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_value_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_value_successfully", fiber.Map{
		"result": activity,
	})

}

func (a *activityController) ActivityPeaEmissionValueById(c *fiber.Ctx) error {
	activity, err := a.ActivityUsecase.GetActivityPeaEmissionValueById(c.Query("activity_id"), c.Locals("lang").(string))
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_pea_emission_value_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_pea_emission_value_successfully", fiber.Map{
		"result": activity,
	})

}

func (a *activityController) CreateActivityWithCustomizedEmission(c *fiber.Ctx) error {
	req := new(entities.CreateActivityWithCustomizedEmissionReq)
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Locals("company_id").(string)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	err := a.ActivityUsecase.CreateActivityWithCustomizedEmission(req)
	if err != nil {
		if err.Error() == "mssql: duplicate emission" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "duplicate_emission_found", nil)
		} else if err.Error() == "mssql: duplicate emission in site" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "duplicate_emission_found_in_site", nil)
		} else {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "activity_with_custom_emission_created_failed", nil)
		}
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_with_custom_emission_created_successfully", nil)
}

func (a *activityController) UpdateActivityWithCustomizedEmission(c *fiber.Ctx) error {
	req := new(entities.UpdateActivityWithCustomizedEmissionReq)
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Locals("company_id").(string)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	err := a.ActivityUsecase.UpdateActivityWithCustomizedEmission(req)
	if err != nil {
		if err.Error() == "mssql: duplicate emission" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "duplicate_emission_found", nil)
		} else if err.Error() == "mssql: duplicate emission in site" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "duplicate_emission_found_in_site", nil)
		} else {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "activity_with_custom_emission_created_failed", nil)
		}
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_with_custom_emission_created_successfully", nil)
}

func (a *activityController) DeleteActivityTransaction(c *fiber.Ctx) error {
	req := new(entities.DeleteActivityTransactionReq)
	req.Id = c.Query("id")
	req.ActivityId = c.Query("activity_id")
	req.DeleteUser = c.Locals("usr_id").(string)
	err := a.ActivityUsecase.DeleteActivityTransaction(req)
	if err != nil {
		logs.Error(err)
		errResult := err.Error()
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "delete_activity_transaction_failed", fiber.Map{
			"result": errResult,
		})
	}

	// Remove all data from Redis hash
	if _, removeErr := utils.RemoveAllDataFromRedisByCompany(context.Background(), c.Locals("company_id").(string)); removeErr != nil {
		logs.Error(removeErr)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "Delete_activity_transaction_successfully", nil)
}

func (a *activityController) DeleteActivity(c *fiber.Ctx) error {
	req := new(entities.DeleteActivityReq)
	req.Id = c.Query("id")
	req.SiteId = c.Query("site_id")
	//req.ActivityName = c.Query("activity_name")
	req.DeleteUser = c.Locals("usr_id").(string)

	err := a.ActivityUsecase.DeleteActivity(req)
	if err != nil {
		logs.Error(err)
		errResult := err.Error()
		if err.Error() == "mssql: name_not_match" {
			errResult = "name_not_match"
		}
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "delete_activity_failed", fiber.Map{
			"result": errResult,
		})
	}

	// Remove all data from Redis hash
	if _, removeErr := utils.RemoveAllDataFromRedisByCompany(context.Background(), c.Locals("company_id").(string)); removeErr != nil {
		logs.Error(removeErr)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "Delete_activity_successfully", nil)
}

func (a *activityController) UpdateActivity(c *fiber.Ctx) error {
	req := new(entities.UpdateActivityReq)
	req.UpdateUser = c.Locals("usr_id").(string)
	req.CompanyId = c.Locals("company_id").(string)

	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	err := a.ActivityUsecase.UpdateActivity(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, err.Error(), nil)

	}
	return utils.HandleResponse(c, fiber.StatusOK, "update_activity_successfully", nil)
}

func (a *activityController) UpdateActivityTransaction(c *fiber.Ctx) error {
	req := new(entities.UpdateActivityTransactionReq)
	req.User = c.Locals("usr_id").(string)

	companyId, ok := c.Locals("company_id").(string)
	if ok {
		req.CompanyId = companyId
	} else {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "last_visited_company_invalid", nil)
	}

	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)

	}
	err := a.ActivityUsecase.UpdateActivityTransaction(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "update_activity_transaction_failed", nil)
	}

	// Remove all data from Redis hash
	if _, removeErr := utils.RemoveAllDataFromRedisByCompany(context.Background(), c.Locals("company_id").(string)); removeErr != nil {
		logs.Error(removeErr)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "Update_activity_successfully", nil)
}

func (a *activityController) CreateActivityTransaction(c *fiber.Ctx) error {
	req := new(entities.CreateActivityTransactionReq)
	req.User = c.Locals("usr_id").(string)

	companyId, ok := c.Locals("company_id").(string)
	if ok {
		req.CompanyId = companyId
	} else {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "last_visited_company_invalid", nil)
	}

	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)

	}

	id, err := a.ActivityUsecase.CreateActivityTransaction(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "create_activity_transaction_failed", nil)

	}

	// Remove all data from Redis hash
	if _, removeErr := utils.RemoveAllDataFromRedisByCompany(context.Background(), c.Locals("company_id").(string)); removeErr != nil {
		logs.Error(removeErr)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "create_activity_successfully", fiber.Map{
		"result": id,
	})

}

func (a *activityController) ActivityTransaction(c *fiber.Ctx) error {
	req := new(entities.ActivityTransactionReq)
	req.Id = c.Query("activity_id")
	req.Lang = c.Query("lang")
	if req.Lang == "" {
		req.Lang = c.Locals("lang").(string)
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "offset_not_number", nil)
	}
	req.Offset = offset
	pagesize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "page_size_not_number", nil)
	}
	req.PageSize = pagesize
	activityTransactions, rowCount, err := a.ActivityUsecase.ActivityTransaction(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_transaction_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_transaction_successfully", fiber.Map{
		"result":     activityTransactions,
		"total_rows": rowCount,
	})
}

func (a *activityController) ActivityTransactionByDate(c *fiber.Ctx) error {
	req := new(entities.ActivityTransactionByDateReq)
	req.Id = c.Query("activity_id")
	req.Lang = c.Locals("lang").(string)
	req.Date = c.Query("date")
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "offset_not_number", nil)
	}
	req.Offset = offset
	pagesize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "page_size_not_number", nil)
	}
	req.PageSize = pagesize
	res, rowCount, err := a.ActivityUsecase.ActivityTransactionByDate(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_transaction_by_date_failed", nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "activity__tracnsaction_by_date_successfully", fiber.Map{
		"result":     res,
		"total_rows": rowCount,
	})
}

func (a *activityController) ActivityByScope(c *fiber.Ctx) error {
	req := new(entities.ActivityByScopeReq)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)

	}
	req.Lang = c.Locals("lang").(string)

	// Set default time zone if not provided
	if req.TimeZone == nil || *req.TimeZone == "" {
		defaultTZ := "Asia/Bangkok"
		req.TimeZone = &defaultTZ
	}

	res, rowCount, err := a.ActivityUsecase.ActivityByScope(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_by_scope_failed", nil)

	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_by_scope_successfully", fiber.Map{
		"result":     res,
		"total_rows": rowCount,
	})
}
func (a *activityController) ActivityByScopeEmission(c *fiber.Ctx) error {
	req := new(entities.ActivityByScopeEmissionReq)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}
	req.Lang = c.Locals("lang").(string)

	// Set default time zone if not provided
	if req.TimeZone == nil || *req.TimeZone == "" {
		defaultTZ := "Asia/Bangkok"
		req.TimeZone = &defaultTZ
	}

	res, rowCount, err := a.ActivityUsecase.ActivityByScopeEmission(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_by_scope_emission_failed", nil)

	}
	return utils.HandleResponse(c, fiber.StatusOK, "activity_by_scope_emission_successfully", fiber.Map{
		"result":     res,
		"total_rows": rowCount,
	})
}

func (a *activityController) ActivityLists(c *fiber.Ctx) error {
	req := new(entities.ActivityListsReq)
	req.Lang = c.Locals("lang").(string)
	req.SiteId = c.Query("site_id")
	req.StartDate = c.Query("start_date")
	req.EndDate = c.Query("end_date")
	activityLists, err := a.ActivityUsecase.ActivityLists(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "activity_lists_failed", nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "activity_lists_successfully", fiber.Map{
		"status":         "OK",
		"status_code":    fiber.StatusOK,
		"message":        "activity_lists_successfully",
		"activity_lists": activityLists,
	})
}

func (a *activityController) CreateActivity(c *fiber.Ctx) error {
	req := new(entities.ActivityCreateReq)
	req.User = c.Locals("usr_id").(string)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	_, err := a.ActivityUsecase.CreateActivity(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "create_activity_failed", nil)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "create_activity_successfully", nil)
}

func (a *activityController) DuplicateActivity(c *fiber.Ctx) error {
	req := new(entities.DuplicateActivityReq)
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Locals("company_id").(string)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	err := a.ActivityUsecase.DuplicateActivity(req)
	if err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "duplicate_activity_failed", nil)
	}

	// Remove all data from Redis hash
	if _, removeErr := utils.RemoveAllDataFromRedisByCompany(context.Background(), c.Locals("company_id").(string)); removeErr != nil {
		logs.Error(removeErr)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "duplicate_activity_successfully", nil)
}

func (a *activityController) UpdateBillConnect(c *fiber.Ctx) error {
	req := new(entities.UpsertBillingConnectReq)
	req.Id = c.Query("id")
	req.Status = c.Query("status")
	err := a.ActivityUsecase.UpdateBillConnect(req)
	if err != nil {
		fmt.Println(err.Error())
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "update_status_bill_failed", nil)
	}
	return utils.HandleResponse(c, fiber.StatusOK, "update_status_bill_successfully", nil)
}

func (a *activityController) DisconnectBill(c *fiber.Ctx) error {
	req := new(entities.DisConnectBillReq)
	req.Id = c.Query("id")
	req.User = c.Locals("usr_id").(string)
	err := a.ActivityUsecase.DisconnectBill(req)
	if err != nil {
		fmt.Println(err.Error())
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "disconnect_bill_failed", nil)

	}
	return utils.HandleResponse(c, fiber.StatusOK, "disconnect_bill_successfully", nil)
}

func (a *activityController) CreateActivityWithRecepiEmission(c *fiber.Ctx) error {
	req := new(entities.CreateActivityWithRecepiEmissionReq)
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Locals("company_id").(string)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	err := a.ActivityUsecase.CreateActivityWithRecepiEmission(req)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "mssql: duplicate emission" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "duplicate_emission_found", nil)
		} else if err.Error() == "mssql: duplicate emission in site" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "duplicate_emission_found_in_site", nil)
		} else {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "activity_with_custom_emission_created_failed", nil)
		}
	}

	return utils.HandleResponse(c, fiber.StatusOK, "activity_with_recepi_emission_created_successfully", nil)
}

func (a *activityController) UpdateActivityWithRecepiEmission(c *fiber.Ctx) error {
	req := new(entities.UpdateActivityWithRecepiEmissionReq)
	req.User = c.Locals("usr_id").(string)
	req.CompanyId = c.Locals("company_id").(string)
	if err := c.BodyParser(req); err != nil {
		logs.Error(err)
		return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "body_not_match", nil)
	}

	err := a.ActivityUsecase.UpdateActivityWithRecepiEmission(req)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "mssql: duplicate emission" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "duplicate_emission_found", nil)
		} else if err.Error() == "mssql: duplicate emission in site" {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "duplicate_emission_found_in_site", nil)
		} else {
			return utils.HandleResponse(c, fiber.ErrBadRequest.Code, "activity_with_custom_emission_created_failed", nil)
		}
	}

	return utils.HandleResponse(c, fiber.StatusOK, "activity_with_recepi_emission_created_successfully", nil)
}
