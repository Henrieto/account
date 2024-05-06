package commands

import (
	"fmt"

	"github.com/henrieto/jax/command"
)

var CreateSuper = command.Command{
	Use:   "createsuperuser",
	Short: "create a super user",
	Long:  "A command for creating a super admin",
	Run: func(c *command.Cmd, s []string) {
		name, _ := c.GetFlag("email")
		password, _ := c.GetFlag("password")
		fmt.Println(name, password)
	},
}
