package account

import (
	"net/http"

	jwt_auth "github.com/henrieto/account/auth/jwt"
	"github.com/henrieto/account/config"
	"github.com/henrieto/account/handlers"
	"github.com/henrieto/account/middlewares"
	"github.com/henrieto/jax"
)

var Routes = []jax.Route{
	{
		Path:    "/signup",
		Handler: handlers.Signup,
		Method:  http.MethodPost,
		Name:    "signup",
	},
	{
		Path:    "/login",
		Handler: handlers.Login,
		Method:  http.MethodPost,
		Name:    "login",
	},
	{
		Path: "/profile",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.Profile, "user-can_update_own"),
		),
		Method: http.MethodGet,
		Name:   "profile",
	},
	{
		Prefix:  "/identiy",
		Path:    "/verifiy",
		Handler: handlers.VerifyIdentity,
		Method:  http.MethodPost,
		Name:    "verify-identity",
	},
	{
		Prefix:  "/password",
		Path:    "/forgort",
		Handler: handlers.ForgortPassword,
		Method:  http.MethodPost,
		Name:    "forgot-password",
	},
	{
		Prefix:  "/password",
		Path:    "/change",
		Handler: handlers.ChangePassword,
		Method:  http.MethodPost,
		Name:    "change-password",
	},
	{
		Prefix: "/group",
		Path:   "/create",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.CreateGroup, "group-can_create"),
		),
		Method: http.MethodPost,
		Name:   "create-group",
	},
	{
		Prefix: "/group",
		Path:   "/update",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.UpdateGroup, "group-can_update"),
		),
		Method: http.MethodPost,
		Name:   "update-group",
	},
	{
		Prefix: "/group",
		Path:   "/list",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.ListGroup, "group-can_view", "can_view"),
		),
		Method: http.MethodGet,
		Name:   "list-group",
	},
	{
		Prefix: "/group",
		Path:   "/get/{id}",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.GetGroup, "group-can_view", "can_view"),
		),
		Method: http.MethodGet,
		Name:   "get-group-bi-id",
	},
	{
		Prefix: "/group",
		Path:   "/get/name/{name}",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.GetGroup, "group-can_view", "can_view"),
		),
		Method: http.MethodGet,
		Name:   "get-group-by-name",
	},
	{
		Prefix: "/group",
		Path:   "/delete/{id}",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.DeleteGroup, "group-can_delete"),
		),
		Method: http.MethodDelete,
		Name:   "delete-group",
	},
	{
		Prefix: "/group",
		Path:   "/add/user",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.AddUserToGroup, "group_user-can_create"),
		),
		Method: http.MethodPost,
		Name:   "add-user-to-group",
	},
	{
		Prefix: "/permission",
		Path:   "/create",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.CreatePermission, "permission-can_create"),
		),
		Method: http.MethodPost,
		Name:   "create-permission",
	},
	{
		Prefix: "/permission",
		Path:   "/update/{id}",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.UpdatePermission, "permission-can_update"),
		),
		Method: http.MethodPut,
		Name:   "update-permission",
	},
	{
		Prefix: "/permission",
		Path:   "/list",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.ListPermission, "permission-can_view"),
		),
		Method: http.MethodGet,
		Name:   "list-permission",
	},
	{
		Prefix: "/permission",
		Path:   "/get/{id}",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.GetPermission, "permission-can_view"),
		),
		Method: http.MethodGet,
		Name:   "get-permission",
	},
	{
		Prefix: "/permission",
		Path:   "/delete/{id}",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.DeletePermission, "permission-can_delete"),
		),
		Method: http.MethodDelete,
		Name:   "delete-permission",
	},
	{
		Prefix: "/permission",
		Path:   "/group/add",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.AddPermissionsToGroup, "permission_group-can_create"),
		),
		Method: http.MethodPost,
		Name:   "add-permissions-to-group",
	},
	{
		Prefix: "/permission",
		Path:   "/group/remove",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.RemoveGroupPermisions, "permission_group-can_delete"),
		),
		Method: http.MethodDelete,
		Name:   "remove-permissions-from-group",
	},
	{
		Prefix: "/permission",
		Path:   "/group/get",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.GetGroupPermissions, "permission_group-can_view"),
		),
		Method: http.MethodGet,
		Name:   "get-group-permissions",
	},
	{
		Prefix: "/permission",
		Path:   "/user/add",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.AddPermissionsToUser, "permission_user-can_create"),
		),
		Method: http.MethodPost,
		Name:   "add-permission-to-user",
	},
	{
		Prefix: "/permission",
		Path:   "/user/get",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.GetUserPermissions, "permission_user-can_view"),
		),
		Method: http.MethodGet,
		Name:   "get-user-permissions",
	},
	{
		Prefix: "/permission",
		Path:   "/user/remove",
		Handler: jwt_auth.Protected(
			config.SECRET,
			middlewares.HasPermission(handlers.RemoveUserPermisions, "permission_user-can_delete"),
		),
		Method: http.MethodDelete,
		Name:   "remove-user-permissions",
	},
}
