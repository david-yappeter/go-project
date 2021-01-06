package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	config "github.com/davidyap2002/user-go/config"
	model "github.com/davidyap2002/user-go/graph/model"
	tools "github.com/davidyap2002/user-go/tools"
)

//Generated By github.com/davidyap2002/GormCrudGenerator

//UserCreate Create
func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	timeNow := time.Now().UTC().Format("2006-01-02 15:04:05")

	user := model.User{
		Address:         input.Address,
		Email:           input.Email,
		Name:            input.Name,
		Password:        tools.HashPassword(input.Password),
		TelephoneNumber: input.TelephoneNumber,
		CreatedAt:       timeNow,
	}

	err := db.Table("user").Create(&user).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

//UserCreateBatch Create Batch
func UserCreateBatch(ctx context.Context, input []*model.NewUser) ([]*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var userBatch []*model.User
	timeNow := time.Now().UTC().Format("2006-01-02 15:04:05")

	for _, val := range input {
		user := model.User{
			Address:         val.Address,
			Email:           val.Email,
			Name:            val.Name,
			Password:        tools.HashPassword(val.Password),
			TelephoneNumber: val.TelephoneNumber,
			CreatedAt:       timeNow,
		}

		userBatch = append(userBatch, &user)
	}

	err := db.Table("user").Create(&userBatch).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return userBatch, nil
}

//UserUpdate Update
func UserUpdate(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	timeNow := time.Now().UTC().Format("2006-01-02 15:04:05")
	user := model.User{
		Address:         input.Address,
		AuthDigit:       input.AuthDigit,
		Email:           input.Email,
		ID:              input.ID,
		Name:            input.Name,
		Password:        input.Password,
		TelephoneNumber: input.TelephoneNumber,
		UpdatedAt:       &timeNow,
	}

	err := db.Table("user").Where("id = ?", input.ID).Updates(map[string]interface{}{
		"address":          input.Address,
		"auth_digit":       input.AuthDigit,
		"email":            input.Email,
		"id":               input.ID,
		"name":             input.Name,
		"password":         input.Password,
		"telephone_number": input.TelephoneNumber,
		"updated_at":       &timeNow,
	}).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

//UserSoftDelete Soft Delete
func UserSoftDelete(ctx context.Context, id int) (string, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	timeNow := time.Now().UTC().Format("2006-01-02 15:04:05")
	err := db.Table("user").Where("id = ?", id).Update("deleted_at", timeNow).Error

	if err != nil {
		fmt.Println(err)
		return "Fail", err
	}

	return "Success", nil
}

//UserRestoreDelete Restore Soft Delete
func UserRestoreDelete(ctx context.Context, id int) (string, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	err := db.Table("user").Where("id = ?", id).Update("deleted_at", nil).Error

	if err != nil {
		fmt.Println(err)
		return "Fail", err
	}

	return "Success", nil
}

//UserHardDelete Hard Delete
func UserHardDelete(ctx context.Context, id int) (string, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	err := db.Table("user").Where("id = ?", id).Delete(&model.User{}).Error

	if err != nil {
		fmt.Println(err)
		return "Fail", err
	}

	return "Success", nil
}

//UserGetByID Get By ID
func UserGetByID(ctx context.Context, id int, scopes *bool) (*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user model.User

	query := db.Table("user")

	tools.DeletedAt(query, scopes)

	err := query.Where("id = ?", id).First(&user).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

//UserGetByEmail Get By Email
func UserGetByEmail(ctx context.Context, email string, scopes *bool) (*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user model.User

	query := db.Table("user")

	tools.DeletedAt(query, scopes)

	err := query.Where("lower(email) = ?", strings.ToLower(email)).First(&user).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

//UserGetAll GetAll
func UserGetAll(ctx context.Context, scopes *bool) ([]*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user []*model.User

	query := db.Table("user")

	tools.DeletedAt(query, scopes)

	err := query.Find(&user).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}

//UserPagination Pagination
func UserPagination(ctx context.Context, limit int, page int, ascending bool, sortBy string, filter []int, scopes *bool) ([]*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user []*model.User

	query := db.Table("user")

	tools.QueryMaker(query, limit, page, ascending, sortBy, filter)

	tools.DeletedAt(query, scopes)

	err := query.Find(&user).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}

//UserTotalDataPagination  Total Data Pagination
func UserTotalDataPagination(ctx context.Context, limit int, page int, ascending bool, sortBy string, filter []int, scopes *bool) (int, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	query := db.Table("user")

	tools.QueryMaker(query, limit, page, ascending, sortBy, filter)

	tools.DeletedAt(query, scopes)

	err := query.Count(&count).Error

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int(count), nil
}