package entity

import "time"

type User struct {
	Id        int       `json:"id" form:"id"`
	Division  string    `json:"division" form:"division"`
	Role      string    `json:"role" form:"role"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Phone     string    `json:"phone" form:"phone"`
	Password  string    `json:"password" form:"password"`
	Gender    string    `json:"gender" form:"gender"`
	Address   string    `json:"address" form:"address"`
	Avatar    string    `json:"avatar" form:"avatar"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}

type Login struct {
	Input    string `json:"input" form:"input"`
	Password string `json:"password" form:"password"`
}

type UpdateUser struct {
	Id       int    `json:"id" form:"id"`
	Division string `json:"division" form:"division"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
	Address  string `json:"address" form:"address"`
}

type UserSimplified struct {
	Id       int    `json:"id" form:"id"`
	Division string `json:"division" form:"division"`
	Role     string `json:"role" form:"role"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Gender   string `json:"gender" form:"gender"`
	Address  string `json:"address" form:"address"`
	Avatar   string `json:"avatar" form:"avatar"`
}
