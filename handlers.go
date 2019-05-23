package main

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/labstack/echo"
	"net/http"
)

func preflightHandler(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func adminPanelHandler(c echo.Context) error {
	return c.String(http.StatusNotFound, "not implemented")
}

func adminLoginHandler(c echo.Context) error {
	var form AdminLoginForm
	var attemptLoginUser User
	if err := c.Bind(&form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	if len(form.PasswordKey) != 36 {
		return echo.NewHTTPError(http.StatusBadRequest, "key should be a valid uuid")
	}
	userDb := DB.Where("username = ?", form.Username)
	if userDb.Error != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "incorrect username or password")
	}
	if err := userDb.First(&attemptLoginUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "incorrect username or password")
	}

	expected := blakeHash(form.PasswordKey, attemptLoginUser.Password)
	if expected == form.EncryptedPassword {
		return c.JSON(http.StatusOK, ResponseLogin{
			Username: attemptLoginUser.Username,
			WebToken: attemptLoginUser.WebToken,
		})
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "incorrect username or password")
}

func queryLinkHandler(c echo.Context) error {
	var links []Link
	var params QueryLinkParams
	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	sortField := params.SortField
	sortOrderMapping := map[string]string{
		"ascend":  "ASC",
		"descend": "DESC",
	}
	var sortOrder string
	var present bool
	sortOrder, present = sortOrderMapping[params.SortOrder]
	if !present {
		sortField = "link_id"
		sortOrder = "DESC"
	}
	sortString := fmt.Sprintf("%s %s", sortField, sortOrder)

	paginator := pagination.Paging(&pagination.Param{
		DB:      DB,
		Page:    params.Page,
		Limit:   params.Limit,
		OrderBy: []string{sortString},
		ShowSQL: true,
	}, &links)
	return c.JSON(http.StatusOK, paginator)
}

func createLinkHandler(c echo.Context) error {
	return c.String(http.StatusNotFound, "not implemented")
}

func updateLinkHandler(c echo.Context) error {
	return c.String(http.StatusNotFound, "not implemented")
}

func deleteLinkHandler(c echo.Context) error {
	return c.String(http.StatusNotFound, "not implemented")
}
