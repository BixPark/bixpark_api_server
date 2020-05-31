```go
package service

type UserService struct {
	DB *gorm.DB
}

func (service UserService) Save(user *data.User) *data.User {
	return user
}
func (service UserService) Update(user *data.User) *data.User {
	return user
}
func (service UserService) Delete(ID int) bool {
	return true
}
func (service UserService) Get(ID int) (user *data.User) {
	return
}
func (service UserService) GetAll(limit int, offset int) (users *data.UserList) {
	return
}


```