package repository

import (
	"middleware-project/database"
	"middleware-project/models"
)

type UserRepo struct{}

// get all users
func (u *UserRepo) GetUsers() ([]models.User, error) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// get user by id
func (u *UserRepo) GetUserById(id int) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// create new user
func (u *UserRepo) CreateUser(input models.UserInput) (*models.User, error) {

	var newUser models.User = models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	if err := database.DB.Save(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

// update user by id
func (u *UserRepo) UpdateUser(id int, input models.UserInput) (*models.User, error) {
 user, err := u.GetUserById(id)
	if err != nil {
		return nil, err
	}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	if err := database.DB.Save(&user).Error ; err != nil {
		return nil, err
	}
	return user, nil
}

// delete user by id
func (u *UserRepo) DeleteUser(id int) (*models.User, error) {
	user, err := u.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if err := database.DB.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
