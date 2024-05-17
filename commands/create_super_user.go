package commands

import (
	"fmt"
	"os"
	"syscall"

	"github.com/henrieto/jax/command"
	"golang.org/x/term"
)

func getPassword(display string) (string, error) {
	fmt.Println(display)
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		os.Exit(1)
	}
	return string(password), nil
}

var CreateSuper = command.Command{
	Use:   "createsuperuser",
	Short: "create a super user",
	Long:  "A command for creating a super admin",
	Run: func(c *command.Cmd, s []string) {
		email, _ := c.GetFlag("email")
		password, _ := getPassword("Enter Password : ")
		confirm_password, _ := getPassword("Confirm Password : ")
		if password != confirm_password {
			fmt.Println("passwords doesn't match")
			return
		}
		fmt.Println(email, password)
	},
	Flags: []*command.Flag{
		{
			Requird:    false,
			Name:       "email",
			Short_Name: "e",
			Help_Text:  "user email",
		},
	},
}
