package service

import (
	"bixpark_server/internal/data"
	"github.com/jinzhu/gorm"
	"log"
)

type UserService struct {
	DB *gorm.DB
}
type UserAccountNotExistsError struct{}
type AuthenticationFailedError struct{}

func (*UserAccountNotExistsError) Error() string {
	return "user account not exists"
}
func (*AuthenticationFailedError) Error() string {
	return "Authentication failed"
}
func (service UserService) Authenticate(credentials *data.Credentials) (user *data.User, err error) {
	var userLogin data.UserLogin
	service.DB.Preload("User").Model(&data.UserLogin{}).Where(&data.UserLogin{
		Username: credentials.Username,
	}).First(&userLogin)
	if userLogin == (data.UserLogin{}) {
		return nil, &UserAccountNotExistsError{}
	}

	log.Println(userLogin.Password, " ", credentials.Password)
	if userLogin.Password == credentials.Password {
		user = &userLogin.User
		return
	} else {
		return nil, &AuthenticationFailedError{}
	}
	return
}

// Basic CURD
func (service UserService) Save(user *data.User, credential *data.Credentials) {
	service.DB.Create(user)
	userLogin := data.UserLogin{
		Username:  credential.Username,
		Password:  credential.Password,
		UserRefer: user.ID,
	}
	service.DB.Create(&userLogin)

}
func (service UserService) Update(updateData *data.User) (content *data.User) {
	if err := service.DB.Preload("ProfilePic").Model(&data.User{}).Where("ID =?", updateData.ID).Update(updateData).Error; err != nil {
		content = nil
	} else {
		content = updateData
	}
	return
}
func (service UserService) Delete(ID int) {
	service.DB.Model(&data.User{}).Where("ID=?", ID).Delete(&data.User{})
}
func (service UserService) Get(ID int) (user *data.User) {
	if err := service.DB.Preload("ProfilePic").Model(&data.User{}).Where("ID =?", ID).First(user).Error; err != nil {
		user = nil
	}
	return
}
func (service UserService) GetAll(limit int, offset int) (users data.UserList) {
	service.DB.Preload("ProfilePic").Limit(limit).Offset(offset).Find(&users)
	return
}
