package main

import (
	"github.com/henrieto/account"
	"github.com/henrieto/account/utils/test_utils"
	"github.com/henrieto/jax"
	"github.com/jackc/pgx/v5"
)

var DBCONN pgx.Conn

var config = &jax.Config{
	Server: jax.DefaultServerConfig(),
}

func main() {
	// initialize db connection
	conn, err := test_utils.DatabaseConnection(test_utils.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// set the global db connction variable
	DBCONN = *conn

	// add plugins to jax config

	config.Plugins = []jax.Plugin{
		account.Plugin(&AccountConfg),
	}

	jx := jax.New(config)
	jx.Initialize()
}

var AccountConfg = account.AccountConfig{
	SecretKey:          "",
	DatabaseConnection: &DBCONN,
	DefaultSmtpConfig: &account.DefaultSmtpConfig{
		EmailSender:   "evihsltd@gmail.com",
		UserName:      "evihsltd@gmail.com",
		Password:      "bdizhilunqhmmqmm",
		Host:          "smtp.gmail.com",
		ServerAddress: "smtp.gmail.com:587",
	},
}
