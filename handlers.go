package main

import (
	"fmt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/blake2b"
	"net/http"
)

func adminPanelHandler(c echo.Context) error {
	return c.String(http.StatusNotFound, "not implemented")
}

func adminLoginHandler(c echo.Context) error {
	var form AdminLoginForm
	var attemptLoginUser User
	if err := c.Bind(&form); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := db.Where("username = ?", form.Username).First(&attemptLoginUser).Error; err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	if len(form.PasswordKey) != 36 {
		return c.NoContent(http.StatusBadRequest)
	}
	expectedString := fmt.Sprintf("%s%s%s", form.PasswordKey, "|", attemptLoginUser.Password)
	expectedHashBytes := blake2b.Sum512([]byte(expectedString))
	expectedHex := fmt.Sprintf("%x", expectedHashBytes)
	if expectedHex == form.EncryptedPassword {
		return c.JSON(http.StatusOK, ResponseLogin{
			attemptLoginUser.Username,
			attemptLoginUser.WebToken.Token,
		})
	}
	return c.NoContent(http.StatusUnauthorized)
}

func queryLinkHandler(c echo.Context) error {
	return c.String(http.StatusNotFound, "not implemented")
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
