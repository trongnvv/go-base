package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	Mode       string      `json:"mode"`
	Address    *Address    `json:"address"`
	Key        *Key        `json:"key"`
	Vib        *Vib        `json:"vib"`
	Log        *Log        `json:"log"`
	Token      *Token      `json:"token"`
	Postgresql *Postgresql `json:"pgsql"`
}

type Address struct {
	Server  *ServerAddress  `json:"server"`
	Clients *ClientsAddress `json:"clients"`
}

type ServerAddress struct {
	Restful  string `json:"restful"`
	Internal string `json:"internal"`
}

type ClientsAddress struct {
	TransactionController string `json:"transaction_controller"`
}

type Key struct {
	Private      string `json:"private"`
	Public       string `json:"public"`
	VibPublicKey string `json:"vib_public_key"`
}

type Vib struct {
	TimeOut   time.Duration `json:"time_out"`
	UserName  string        `json:"username"`
	Password  string        `json:"password"`
	Url       *VibUrl       `json:"url"`
	LinkType  string        `json:"link_type"`
	PartnerId string        `json:"partner_id"`
	Token     string        `json:"token"`
}

type VibUrl struct {
	GetToken          string `json:"get_token"`
	RefreshToken      string `json:"refresh_token"`
	Register          string `json:"register"`
	RegisterConfirm   string `json:"register_confirm"`
	Unregister        string `json:"unregister"`
	UnregisterConfirm string `json:"unregister_confirm"`
}

type Log struct {
	FileName   string `json:"file_name"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
	Compress   bool   `json:"compress"`
}

type Token struct {
	AccessSecret  string `json:"access_secret"`
	ExpireAccess  string `json:"expire_access"`
	RefreshSecret string `json:"refresh_secret"`
	ExpireRefresh string `json:"expire_refresh"`
}

type Postgresql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	DbName   string `json:"db_name"`
	SslMode  string `json:"ssl_mode"`
	Password string `json:"password"`
}

var common *Config

func Get() *Config {
	return common
}

func LoadConfig(path string) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(configFile)
	if err = json.Unmarshal(byteValue, &common); err != nil {
		panic(err)
	}
}
