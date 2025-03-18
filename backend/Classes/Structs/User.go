package Structs

import "fmt"

type User struct {
	Id       int
	Group    string
	Name     string
	Password string
}

func (u *User) ToString() string {
	return fmt.Sprintf("%d,U,%s,%s,%s", u.Id, u.Group, u.Name, u.Password)
}
