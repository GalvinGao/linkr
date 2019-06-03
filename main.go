package main

import (
	"github.com/GalvinGao/linkr/notify"
	"github.com/GalvinGao/linkr/notify/gotify"
	"github.com/GalvinGao/linkr/notify/server_chan"
	"github.com/GalvinGao/linkr/notify/telegram"
	"github.com/GalvinGao/linkr/notify/webhook"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"reflect"
)

var DB *gorm.DB

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

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

	if config.Logging.Sentry.Enabled {
		if err := raven.SetDSN(config.Logging.Sentry.DSN); err != nil {
			panic(err)
		}
	}

	DB.AutoMigrate(&Token{}, &Link{}, &LinkRecord{})

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

	e.Validator = &CustomValidator{validator: validator.New()}

	// homepage
	e.GET("/", func(c echo.Context) error {
		return c.File("home.html")
	})

	e.POST("/__link", func(c echo.Context) error {
		form := new(NewLinkForm)
		if err := c.Bind(form); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad form")
		}
		if err = c.Validate(form); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad form")
		}

		var attemptToken Token
		if err := DB.Where("token = ?", form.Key).Find(&attemptToken).Error; err != nil {
			return err
		}

		if attemptToken.Token != form.Key {
			return echo.NewHTTPError(http.StatusForbidden, "bad key")
		}

		return DB.Create(&Link{
			ShortURL:      form.ShortURL,
			LongURL:       form.LongURL,
			NotifyOnVisit: form.Notify,
		}).Error
	})

	// short link
	e.GET("/:link", func(c echo.Context) error {
		var result Link
		link := c.Param("link")
		chk := DB.Where("short_url = ?", link).Find(&result)
		if chk.Error != nil {
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		err := recordLinkVisit(c.Request(), result.LinkID)
		if err != nil {
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

	e.Logger.Fatal(e.Start(config.Server.Address))
}
