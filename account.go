package account

import (
	"github.com/henrieto/jax"
	"github.com/henrieto/jax/command"
	"github.com/henrieto/plugins/commands"
)

var PluginConfig = jax.Plugin{
	Routes: Routes,
	Commands: []*command.Command{
		&commands.CreateSuper,
	},
}
