package main

import (
	"github.com/henrieto/account"
	"github.com/henrieto/account/models/repository"
	mock_storage "github.com/henrieto/account/storage/mock"
	"github.com/henrieto/jax"
	"github.com/henrieto/jax/command"
)

var config = jax.Config{
	Plugins: []jax.Plugin{
		account.PluginConfig,
	},
	Server: &jax.ServerConfig{
		Port: "8000",
	},
	Router: &jax.RouterConfig{
		Version: 1,
	},
}

func init() {
	command.Register(account.PluginConfig.Commands...)
}
func main() {
	repository.UserRepository = mock_storage.NewUserStorage()
	command.LoadCommands()
	command.Execute()
	Jax := jax.New(&config)
	Jax.Initialize()
	Jax.Server.Address(":9000")
	Jax.Server.Listen()
}
