package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userID uint) ([]model.TaskResponse, error)
	GetTaskByID(userID uint, taskID uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userID uint, taskID uint) (model.TaskResponse, error)
	DeleteTask(userID uint, taskID uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetAllTasks(userID uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userID); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
}
