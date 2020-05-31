package service

import (
	"bixpark_server/internal/data"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (service UserService) Save(user *data.User) {
	service.DB.Create(user)
}
func (service UserService) Update(updateData *data.User) (content *data.User) {
	if err := service.DB.Model(&data.User{}).Where("ID =?", updateData.ID).Update(updateData).Error; err != nil {
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
	if err := service.DB.Where("ID =?", ID).First(&user).Error; err != nil {
		user = nil
	}
	return
}
func (service UserService) GetAll(limit int, offset int) (users data.UserList) {
	service.DB.Limit(limit).Offset(offset).Find(&users)
	return
}
