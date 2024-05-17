package account

import (
	"github.com/henrieto/account/commands"
	"github.com/henrieto/account/config"
	"github.com/henrieto/account/models/database/db"
	"github.com/henrieto/account/models/repository"
	"github.com/henrieto/account/notification"
	"github.com/henrieto/account/notification/notifiers"
	"github.com/henrieto/account/storage"
	"github.com/henrieto/jax"
	"github.com/henrieto/jax/command"
)

var PluginConfig = jax.Plugin{
	Routes: Routes,
	ModelActions: map[string]any{
		"user": repository.User,
	},
	Commands: []*command.Command{
		&commands.CreateSuper,
	},
}

func Plugin(config *AccountConfig) jax.Plugin {
	SetSecretKet(config.SecretKey)
	SetRepositories(config.DatabaseConnection)
	SetDefaultEmailNotifierParameters(config.DefaultSmtpConfig)
	SetNotifier(config.Notifiers)
	return PluginConfig
}

func SetRepositories(database_connection db.DBTX) {
	querier := db.New(database_connection)
	repository.User = storage.NewUserStorage(querier)
	repository.Group = storage.NewGroupStorage(querier)
	repository.Permission = storage.NewPermissionStorage(querier)
	repository.VerifyIdentityData = storage.NewVerifyIdentityData(querier)
}

func SetDefaultEmailNotifierParameters(default_config *DefaultSmtpConfig) {
	config.SMTP_EMAILSENDER = default_config.EmailSender
	config.SMTP_NOTIFIER_USERNAME = default_config.UserName
	config.SMTP_NOTIFIER_PASSWORD = default_config.Password
	config.SMTP_NOTIFIER_HOST = default_config.Host
	config.SMTP_NOTIFIER_SERVER_ADDRESS = default_config.ServerAddress
}

func SetNotifier(notifierList map[string]notification.Notifier) {
	if notifierList != nil || len(notifierList) == 0 {
		notifier := &notifiers.SmtpNotifier{
			Username:       config.SMTP_NOTIFIER_USERNAME,
			Password:       config.SMTP_NOTIFIER_PASSWORD,
			Host:           config.SMTP_NOTIFIER_HOST,
			Server_Address: config.SMTP_NOTIFIER_SERVER_ADDRESS,
		}
		notification.Manager.Register("email", notifier)
		return
	}
	for key, notifier := range notifierList {
		notification.Manager.Register(key, notifier)
	}
}

func SetSecretKet(secret string) {
	config.SECRET = secret
}

type AccountConfig struct {
	SecretKey          string
	DatabaseConnection db.DBTX
	Notifiers          map[string]notification.Notifier
	DefaultSmtpConfig  *DefaultSmtpConfig
}

type DefaultSmtpConfig struct {
	EmailSender   string
	UserName      string
	Password      string
	Host          string
	ServerAddress string
}
