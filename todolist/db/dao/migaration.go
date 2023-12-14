package dao

import (
	"todolist/db/model"
)

func migaration() {

	err := _db.Set("gorm:table_options", "charset=utf8mb4")
	AutoMigarate(&model.User{}, &model.Task{})
	if err != nil {
		return
	}
}

func AutoMigarate(user *model.User, task *model.Task) {
	panic("unimplemented")
}
