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
	Flags: []*command.Flag{
		{
			Requird:    false,
			Name:       "email",
			Short_Name: "e",
			Help_Text:  "user email",
		},
		{
			Requird:    false,
			Name:       "password",
			Short_Name: "p",
			Help_Text:  "user password",
		},
	},
}
