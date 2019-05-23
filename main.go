package main

import (
	"github.com/GalvinGao/linkr/notify"
	"github.com/GalvinGao/linkr/notify/gotify"
	"github.com/GalvinGao/linkr/notify/server_chan"
	"github.com/GalvinGao/linkr/notify/telegram"
	"github.com/GalvinGao/linkr/notify/webhook"
	"github.com/dchest/uniuri"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"reflect"
)

var DB *gorm.DB

func main() {
	var config Config
	err := configor.New(&configor.Config{ENVPrefix: "LINKR"}).Load(&config, "config.yml")
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(config.Database.Type, config.Database.DSN)
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&User{}, &Token{}, &Link{}, &LinkRecord{})

	// process the service providers
	v := reflect.ValueOf(config.Notification)
	t := reflect.TypeOf(config.Notification)
	var providers []notify.ServiceProvider

	// enumerates through the notification providers in the config file
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		provider := v.Field(i).Interface()
		status := reflect.ValueOf(provider).FieldByName("Enabled").Interface().(bool)

		// see if the provider should be enabled or not
		if status {
			// provider intended to be initialized. here's the bindings
			switch field.Tag.Get("provider") {
			case "server_chan":
				config := provider.(ServerChanConfig)
				providers = append(providers, server_chan.New(config.ApiKey))
			case "telegram":
				config := provider.(TelegramConfig)
				providers = append(providers, telegram.New(config.BotToken, config.ChatID))
			case "gotify":
				config := provider.(GotifyConfig)
				providers = append(providers, gotify.New(config.Endpoint, config.ApplicationKey))
			case "webhooks":
				config := provider.(WebhooksConfig)
				for _, webhookUrl := range config.URLs {
					providers = append(providers, webhook.New(webhookUrl))
				}
			}
		}
	}

	// get the notifier
	notifier := NewNotifier(providers)

	// router bindings
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowHeaders: []string{
			echo.HeaderAuthorization,
			echo.HeaderAccept,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	// homepage
	e.GET("/", func(c echo.Context) error {
		return c.File("home.html")
	})

	// admin control panel
	admin := e.Group("/admin")
	admin.GET("", adminPanelHandler)

	// admin control panel apis
	adminApi := admin.Group("/api")
	adminApi.POST("/login", adminLoginHandler)

	// public api group
	api := e.Group("/api")
	api.Use(middleware.KeyAuth(func(s string, c echo.Context) (bool, error) {
		var token Token
		DB.Where("token = ?", s).First(&token)
		if s == token.Token {
			return true, nil
		}
		var webTokenUser User
		DB.Where("web_token = ?", s).First(&webTokenUser)
		if s == webTokenUser.WebToken {
			return true, nil
		}
		return false, echo.NewHTTPError(http.StatusUnauthorized, "bad token")
	}))

	api.GET("/link", queryLinkHandler)
	api.POST("/link", createLinkHandler)
	api.PUT("/link/:id", updateLinkHandler)
	api.DELETE("/link/:id", deleteLinkHandler)

	// short link
	e.GET("/:link", func(c echo.Context) error {
		var result Link
		link := c.Param("link")
		chk := DB.Where("short_url = ?", link).Find(&result)
		if chk.Error != nil {
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		if result.NotifyOnVisit {
			go func() {
				notifier.notify(c.Request(), notify.Extras{
					ShortUrl: result.ShortURL,
					LongUrl:  result.LongURL,
				})
			}()
			return c.Redirect(http.StatusTemporaryRedirect, result.LongURL)
		} else {
			return c.Redirect(http.StatusPermanentRedirect, result.LongURL)
		}
	})

	var user User
	var count uint
	DB.First(&user).Count(&count)
	if count == 0 {
		DB.Create(&User{
			Username: "a",
			Password: "a",
			WebToken: uniuri.NewLen(32),
		})
	}
	/*
		for i := 0; i < 500; i++ {
			DB.Create(&Link{
				ParentUserID:  1,
				ShortURL:      uniuri.NewLen(5),
				LongURL:       "https://example.com/link/" + uniuri.NewLen(32),
				NotifyOnVisit: false,
			})
		}
	*/

	e.Logger.Fatal(e.Start(config.Server.Address))
}
