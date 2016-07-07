package models

import (
	"errors"
	//"strconv"
	//"time"
)

var (
	UserList map[int]*User
)

func init() {

}

type User struct {
	Id       int `orm:"pk"`
	Username string
	Password string
	Profile  *Profile `orm:"rel(one)"`
}

type Profile struct {
	Id	int `orm:"pk"`
	Gender  string
	Age     int
	Address string
	Email   string

}

func AddUser(u User) int {
	//u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	//UserList[u.Id] = &user_11111
	u.Id = 1 // mock data
	return u.Id
}

func GetUser(uid int) (u *User, err error) {

	return nil, errors.New("User not exists")
}

func GetAllUsers() map[int]*User {
	return UserList
}

func UpdateUser(uid int, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid int) {
	delete(UserList, uid)
}
