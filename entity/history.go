package entity

type UserUsageHistory struct {
	Id       int    `json:"id" form:"id"`
	Division string `json:"division" form:"division"`
	Role     string `json:"role" form:"role"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
	Address  string `json:"address" form:"address"`
	Avatar   string `json:"avatar" form:"avatar"`
}
