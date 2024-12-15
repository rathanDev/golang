package model

import "svc/constant"

type User struct {
	ID       string        `json:"id"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Role     constant.Role `json:"role"`
}
