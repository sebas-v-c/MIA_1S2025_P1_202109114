package Structs

import "fmt"

type User struct {
	Id       string
	Group    string
	Name     string
	Password string
}

func (u *User) ToString() string {
	return fmt.Sprintf("%s,U,%s,%s,%s", u.Id, u.Group, u.Name, u.Password)
}
