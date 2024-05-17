package account

import (
	"net/http"

	"github.com/henrieto/account/auth"
	"github.com/henrieto/account/decorators"
	"github.com/henrieto/account/handlers"
	"github.com/henrieto/jax"
)

var Routes = []jax.Route{
	// profile routes
	// user signup route
	jax.Route{
		Path:    "/signup",
		Handler: handlers.Signup,
		Method:  http.MethodPost,
		Name:    "signup",
	},
	// user login route
	jax.Route{
		Path:    "/login",
		Handler: handlers.Login,
		Method:  http.MethodPost,
		Name:    "login",
	},
	// get all user details
	jax.Route{
		Path: "/profile",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.Profile, "user-can_update_own"),
		),
		Method: http.MethodGet,
		Name:   "profile",
	},
	// group routes
	// create user group
	jax.Route{
		Prefix: "/group",
		Path:   "/create",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.CreateGroup, "group-can_create"),
		),
		Method: http.MethodPost,
		Name:   "create-group",
	},
	// update user group
	jax.Route{
		Prefix: "/group",
		Path:   "/update",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.UpdateGroup, "group-can_update"),
		),
		Method: http.MethodPut,
		Name:   "update-group",
	},
	// list groups
	jax.Route{
		Prefix: "/group",
		Path:   "/list",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.ListGroups, "group-can_view"),
		),
		Method: http.MethodGet,
		Name:   "list-groups",
	},
	// list groups by pagination
	jax.Route{
		Prefix: "/group",
		Path:   "/paginate",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.PaginateGroups, "group-can_view"),
		),
		Method: http.MethodGet,
		Name:   "paginate-groups",
	},
	// get group
	jax.Route{
		Prefix: "/group",
		Path:   "/get/{name}",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.GetGroup, "group-can_view"),
		),
		Method: http.MethodGet,
		Name:   "get-group",
	},
	jax.Route{
		Prefix: "/group",
		Path:   "/delete/{id}",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.DeleteGroup, "group-can_delete"),
		),
		Method: http.MethodDelete,
		Name:   "delete-group",
	},
	// delete group
	jax.Route{
		Prefix: "/group",
		Path:   "/user/add",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.AddUserToGroup, "group-can_edit"),
		),
		Method: http.MethodPost,
		Name:   "add-user-to-group",
	},
	//permission routes
	// create permission
	jax.Route{
		Prefix: "/permission",
		Path:   "/create",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.CreatePermission, "permission-can_create"),
		),
		Method: http.MethodPost,
		Name:   "create-permission",
	},
	// update permission
	jax.Route{
		Prefix: "/permission",
		Path:   "/update/{id}",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.UpdatePermission, "permission-can_edit"),
		),
		Method: http.MethodPut,
		Name:   "update-permission",
	},
	// list permission
	jax.Route{
		Prefix: "/permission",
		Path:   "/list",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.ListPermission, "permission-can_view"),
		),
		Method: http.MethodGet,
		Name:   "list-permissions",
	},
	// get permission
	jax.Route{
		Prefix: "/permission",
		Path:   "/get/{id}",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.GetPermission, "permission-can_view"),
		),
		Method: http.MethodGet,
		Name:   "get-permission",
	},
	// delete permission
	jax.Route{
		Prefix: "/permission",
		Path:   "/delete/{id}",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.DeletePermission, "permission-can_delete"),
		),
		Method: http.MethodDelete,
		Name:   "delete-permission",
	},
	// add permission to group
	jax.Route{
		Prefix: "/permission",
		Path:   "/group/add",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.AddPermissionsToGroup, "permission-can_edit"),
		),
		Method: http.MethodPost,
		Name:   "add-permission-to-group",
	},
	// remove permission from group
	jax.Route{
		Prefix: "/permission",
		Path:   "/group/remove",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.RemoveGroupPermisions, "permission-can_edit"),
		),
		Method: http.MethodPost,
		Name:   "remove-group-permission",
	},
	// get group permissions
	jax.Route{
		Prefix: "/permission",
		Path:   "/group/{id}",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.GetGroupPermissions, "permission-can_view"),
		),
		Method: http.MethodGet,
		Name:   "get-group-permissions",
	},
	// add permission to user
	jax.Route{
		Prefix: "/permission",
		Path:   "/user/add",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.AddPermissionsToUser, "permission-can_edit"),
		),
		Method: http.MethodPost,
		Name:   "add-permissions-to-user",
	},
	//  get user permissions
	jax.Route{
		Prefix: "/permission",
		Path:   "/user/{id}",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.GetUserPermissions, "permission-can_view"),
		),
		Method: http.MethodGet,
		Name:   "get-user-permissions",
	},
	// remove permission from user
	jax.Route{
		Prefix: "/permission",
		Path:   "/user/remove",
		Handler: auth.Protected(
			decorators.HasPermission(handlers.RemoveUserPermisions, "permission-can_edit"),
		),
		Method: http.MethodPost,
		Name:   "remove-user-permissions",
	},
}
