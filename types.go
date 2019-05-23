package main

import (
	"time"
)

type GotifyConfig struct {
	Enabled        bool   `yaml:"enable"`
	Endpoint       string `yaml:"endpoint"`
	ApplicationKey string `yaml:"application_key"`
}

type ServerChanConfig struct {
	Enabled bool   `yaml:"enable"`
	ApiKey  string `yaml:"api_key"`
}

type TelegramConfig struct {
	Enabled  bool   `yaml:"enable"`
	BotToken string `yaml:"bot_token"`
	ChatID   uint32 `yaml:"chat_id"`
}

type WebhooksConfig struct {
	Enabled bool     `default:"" yaml:"enable"`
	URLs    []string `yaml:"urls"`
}

// Config defines struct for the config file
type Config struct {
	Server struct {
		Address string `default:":3000" yaml:"address"`
	}
	Database struct {
		Type string `default:"sqlite" yaml:"type"`
		DSN  string `default:"file:database.db?cache=shared&mode=rwc" yaml:"dsn"`
	} `yaml:"database"`
	Generate struct {
		Length        uint     `default:"6" yaml:"length"`
		AllowedChars  string   `default:"abdfhjkprtuvwxy34569" yaml:"allowed_chars"`
		BannedPhrases []string `yaml:"banned_phrases"`
	} `yaml:"generate"`
	APICompatibility struct {
		YOURLS bool `default:"false" yaml:"yourls"`
	} `yaml:"api_compatibility"`
	Notification struct {
		Gotify     GotifyConfig     `yaml:"gotify" provider:"gotify"`
		ServerChan ServerChanConfig `yaml:"server_chan" provider:"server_chan"`
		Telegram   TelegramConfig   `yaml:"telegram" provider:"telegram"`
		Webhooks   WebhooksConfig   `yaml:"webhooks" provider:"webhooks"`
	} `default:"" yaml:"notification"`
}

// User stores short link users
type User struct {
	UserID    uint       `gorm:"AUTO_INCREMENT;primary_key" json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
	Username  string     `gorm:"type:varchar(64);unique_index" json:"long"`
	Password  string     `gorm:"type:varchar(128)" json:"password"`
	WebToken  string     `gorm:"type:varchar(32);unique_index" json:"web_token"`
	Tokens    []Token    `gorm:"foreignkey:TokenID"`
	Links     []Link     `gorm:"foreignkey:LinkID"`
}

// Token stores user tokens
type Token struct {
	TokenID      uint      `gorm:"AUTO_INCREMENT;primary_key" json:"token_id"`
	CreatedAt    time.Time `json:"created_at"`
	ParentUserID uint      `json:"-"`
	Token        string    `gorm:"type:varchar(32);unique_index" json:"token"`
	Description  string    `gorm:"type:varchar(256)" json:"description"`
}

// Link defines struct for the mapping between the short link and the original (long) link
type Link struct {
	LinkID       uint       `gorm:"AUTO_INCREMENT;primary_key" json:"link_id"`
	ParentUserID uint       `json:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index" json:"-"`
	// ShortURL is the short link id
	ShortURL string `gorm:"type:varchar(20);unique_index" json:"short_url"`
	// LongURL is the original (long) link URI
	// Due to HTTP standards, the maximum length of a URL can be as long as 2048 chars
	// VARCHAR does not allow us to do that. Therefore a `text` field is required
	LongURL       string       `gorm:"type:text(2048)" json:"long_url"`
	NotifyOnVisit bool         `json:"notify_on_visit"`
	Records       []LinkRecord `gorm:"foreignkey:RecordID" json:"-"`
}

// LinkRecord records all visits to the links
type LinkRecord struct {
	RecordID         uint      `gorm:"AUTO_INCREMENT;primary_key" json:"record_id"`
	CreatedAt        time.Time `json:"created_at"`
	ParentLinkID     uint      `json:"-"`
	Referer          string    `json:"referer"`
	EncodedUserAgent uint      `json:"user_agent"`
}

// ResponseLogin returns login response
type ResponseLogin struct {
	Username string `json:"username"`
	WebToken string `json:"token"`
}

// AdminLoginForm indicates an expected admin login form
type AdminLoginForm struct {
	Username          string `form:"username"`
	PasswordKey       string `form:"key"`
	EncryptedPassword string `form:"password"`
}

type QueryLinkParams struct {
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	SortField string `query:"sort_field"`
	SortOrder string `query:"sort_order"`
}
