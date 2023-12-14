package dao

import (
	"context"
	"todolist/db/model"
	"todolist/types"

	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(c context.Context) *TaskDao {
	if c == nil {
		c = context.Background()
	}
	return &TaskDao{NewDBClient(c)}
}

// 创建task
func (s *TaskDao) CreatTask(task *model.Task) error {
	return s.Model(&model.Task{}).Create(&task).Error

}

// list task
func (s *TaskDao) ListTask(start, limit int, userId uint) (r []*model.Task, total int64, err error) {
	err = s.Model(&model.Task{}).Preload("User").Where("uid = ?", userId).
		Count(&total).
		Limit(limit).Offset((start - 1) * limit).
		Find(&r).Error

	return

}

// find task
func (s *TaskDao) FindTask(uId, id uint) (r *model.Task, err error) {
	err = s.Model(&model.Task{}).Where("id = ? AND uid = ?", id, uId).First(&r).Error
	return

}

// update task
func (s *TaskDao) UpdateTask(uId uint, req *types.UpdateTaskReq) error {
	t := new(model.Task)
	err := s.Model(&model.Task{}).Where("id = ? AND uid=?", req.ID, uId).First(&t).Error
	if err != nil {
		return err
	}

	if req.Status != 0 {
		t.Status = req.Status
	}

	if req.Title != "" {
		t.Title = req.Title
	}

	if req.Content != "" {
		t.Content = req.Content
	}

	return s.Save(t).Error
}

// search task
func (s *TaskDao) SearchTask(uId uint, info string) (tasks []*model.Task, err error) {
	err = s.Where("uid=?", uId).Preload("User").First(&tasks).Error
	if err != nil {
		return
	}

	err = s.Model(&model.Task{}).Where("title LIKE ? OR content LIKE ?",
		"%"+info+"%", "%"+info+"%").Find(&tasks).Error

	return

}

// delete task
func (s *TaskDao) DeleteTaskById(uId, tId uint) error {
	r, err := s.FindTask(uId, tId)
	if err != nil {
		return err
	}
	return s.Delete(&r).Error
}
