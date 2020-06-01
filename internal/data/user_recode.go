package data

import (
	"bixpark_server/config"
	"encoding/hex"
	"fmt"
	"github.com/jinzhu/gorm"
	"hash/fnv"
)

type User struct {
	gorm.Model
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	ProfilePic FileRecode `gorm:"foreignkey:ID;",json:"ProfilePic"`
}
type UserLogin struct {
	gorm.Model
	Username  string
	Password  string
	UserRefer uint
	User      User `gorm:"foreignkey:UserRefer;association_autoupdate:false"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserList []User

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
func (credential *Credentials) Hash(config *config.Config) {
	credential.Password = hash(fmt.Sprintf("%s_%s", config.Security.Salt, credential.Password))
}
