package user

import (
	_entity "capstone/be/entity"
)

type User interface {
	LoginByEmail(email string) (loginUser _entity.User, code int, err error)
	LoginByPhone(phone string) (loginUser _entity.User, code int, err error)
	GetById(id int) (user _entity.User, code int, err error)
	GetAll() (users []_entity.UserSimplified, code int, err error)
	Update(userData _entity.User) (updatedUser _entity.UserSimplified, code int, err error)
}
