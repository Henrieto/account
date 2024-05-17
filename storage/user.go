package storage

import (
	"context"

	"github.com/henrieto/account/models"
	"github.com/henrieto/account/models/database/db"
	db_utils "github.com/henrieto/account/utils/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Querier db.Querier
}

func NewUserStorage(querier db.Querier) *User {
	return &User{querier}
}

// (context , model user data) (model user data , error)
func (strg *User) Create(ctx context.Context, model_user *models.User) (*models.User, error) {
	params := db.CreateUserParams{
		Username:     model_user.Username,
		Email:        model_user.Email,
		FirstName:    model_user.FirstName,
		LastName:     model_user.LastName,
		Gender:       model_user.Gender,
		PasswordHash: model_user.PasswordHash,
		Verified:     model_user.Verified,
		Birthday:     model_user.Birthday,
		GroupID:      model_user.GroupID,
		CreatedAt:    model_user.CreatedAt,
		UpdatedAt:    model_user.UpdatedAt,
	}
	user, err := strg.Querier.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	newmodel_user := db_utils.DbUserToModelUser(user)
	return &newmodel_user, nil
}

// (context) (number of users , error )
func (strg *User) Count(ctx context.Context) (int64, error) {
	count, err := strg.Querier.CountUsers(ctx)
	if err != nil {
		return 0, err
	}
	return count, err
}

// (context , model staff user data) (model staff user data , error)
func (strg *User) CreateStaff(ctx context.Context, model_user *models.User) (*models.User, error) {
	params := db.CreateStaffParams{
		Username:     model_user.Username,
		Email:        model_user.Email,
		FirstName:    model_user.FirstName,
		LastName:     model_user.LastName,
		Gender:       model_user.Gender,
		PasswordHash: model_user.PasswordHash,
		Verified:     model_user.Verified,
		Birthday:     model_user.Birthday,
		GroupID:      model_user.GroupID,
		CreatedAt:    model_user.CreatedAt,
		UpdatedAt:    model_user.UpdatedAt,
	}
	user, err := strg.Querier.CreateStaff(ctx, params)
	if err != nil {
		return nil, err
	}
	newmodel_user := db_utils.DbUserToModelUser(user)
	return &newmodel_user, nil
}

// (context , model super user data) (model super user data , error)
func (strg *User) CreateSuperUser(ctx context.Context, model_user *models.User) (*models.User, error) {
	params := db.CreateSuperUserParams{
		Username:     model_user.Username,
		Email:        model_user.Email,
		FirstName:    model_user.FirstName,
		LastName:     model_user.LastName,
		Gender:       model_user.Gender,
		PasswordHash: model_user.PasswordHash,
		Verified:     model_user.Verified,
		Birthday:     model_user.Birthday,
		GroupID:      model_user.GroupID,
		Superuser:    model_user.Superuser,
		CreatedAt:    model_user.CreatedAt,
		UpdatedAt:    model_user.UpdatedAt,
	}
	user, err := strg.Querier.CreateSuperUser(ctx, params)
	if err != nil {
		return nil, err
	}
	newmodel_user := db_utils.DbUserToModelUser(user)
	return &newmodel_user, nil
}

// (context , updated model user data) (updated db user data , error)
func (strg *User) Update(ctx context.Context, model_user *models.User) (*models.User, error) {
	params := db.UpdateUserParams{
		Username:     model_user.Username,
		Email:        model_user.Email,
		FirstName:    model_user.FirstName,
		LastName:     model_user.LastName,
		Gender:       model_user.Gender,
		PasswordHash: model_user.PasswordHash,
		Verified:     model_user.Verified,
		Birthday:     model_user.Birthday,
		Staff:        model_user.Staff,
		Superuser:    model_user.Superuser,
		AuthID:       model_user.AuthID,
		UpdatedAt:    model_user.UpdatedAt,
	}
	user, err := strg.Querier.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	newmodel_user := db_utils.DbUserToModelUser(user)
	return &newmodel_user, nil
}

// (context , order by) (list of user datas , error)
func (strg *User) List(ctx context.Context, order_by string) ([]*models.User, error) {
	users, err := strg.Querier.GetAllUsers(ctx, order_by)
	if err != nil {
		return nil, err
	}
	newmodel_users := []*models.User{}

	for _, user := range users {
		newmodel_user := db_utils.DbUserToModelUser(user)
		newmodel_users = append(newmodel_users, &newmodel_user)
	}
	return newmodel_users, nil
}

// (context , email) (model user data , error)
func (strg *User) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := strg.Querier.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	newmodel_user := db_utils.DbUserToModelUser(user)
	return &newmodel_user, nil
}

// (context , phone number) (model user data , error)
func (strg *User) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	user, err := strg.Querier.GetUserByPhone(ctx, pgtype.Text{String: phone})
	if err != nil {
		return nil, err
	}
	newmodel_user := db_utils.DbUserToModelUser(user)
	return &newmodel_user, nil
}

// (context , auth id) (model user data , error)
func (strg *User) GetByAuthId(ctx context.Context, auth_id string) (*models.User, error) {
	user, err := strg.Querier.GetUserByAuthId(ctx, pgtype.Text{String: auth_id})
	if err != nil {
		return nil, err
	}
	newmodel_user := db_utils.DbUserToModelUser(user)
	return &newmodel_user, nil
}

// (context , field , field value , order by , offset , limit) (list of model user datas , error)
func (strg *User) Filter(ctx context.Context, field string, value any, order string, offset uint, limit uint) ([]*models.User, error) {
	params := db.FilterUsersParams{
		Column1: field,
		Column2: value,
		Column3: order,
		Offset:  int32(offset),
		Limit:   int32(limit),
	}
	users, err := strg.Querier.FilterUsers(ctx, params)
	if err != nil {
		return nil, err
	}
	newmodel_users := []*models.User{}

	for _, user := range users {
		newmodel_user := db_utils.DbUserToModelUser(user)
		newmodel_users = append(newmodel_users, &newmodel_user)
	}
	return newmodel_users, nil

}

// (context , order by , offset , limit) (list of model user datas , error)
func (strg *User) Paginate(ctx context.Context, order string, offset uint, limit uint) ([]*models.User, error) {
	params := db.PaginateUsersParams{
		Column1: order,
		Offset:  int32(offset),
		Limit:   int32(limit),
	}
	users, err := strg.Querier.PaginateUsers(ctx, params)
	if err != nil {
		return nil, err
	}
	newmodel_users := []*models.User{}

	for _, user := range users {
		newmodel_user := db_utils.DbUserToModelUser(user)
		newmodel_users = append(newmodel_users, &newmodel_user)
	}
	return newmodel_users, nil

}

// (context , user id) (model user data , error)
func (strg *User) Get(ctx context.Context, id int32) (*models.User, error) {
	user, err := strg.Querier.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	newmodel_user := db_utils.DbUserToModelUser(user)
	return &newmodel_user, nil
}

// (context , user id) (error)
func (strg *User) Delete(ctx context.Context, id int32) error {
	err := strg.Querier.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// (context , user email) (if user exists)
func (strg *User) EmailExists(ctx context.Context, email string) bool {
	_, err := strg.GetByEmail(ctx, email)
	return err == nil
}

// (context , user phone number) (if user exists)
func (strg *User) PhoneExists(ctx context.Context, phone string) bool {
	_, err := strg.GetByPhone(ctx, phone)
	return err == nil
}
