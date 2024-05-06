package account

import (
	"github.com/henrieto/account/commands"
	"github.com/henrieto/jax"
	"github.com/henrieto/jax/command"
)

var PluginConfig = jax.Plugin{
	Routes: Routes,
	Commands: []*command.Command{
		&commands.CreateSuper,
	},
}
