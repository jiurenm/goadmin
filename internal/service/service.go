package service

import (
	"admin/internal/dao"
	"admin/internal/model"
)

type Service struct {
	QuestionsService *questionsService
	QuestionService  *questionService
}

type questionsService struct {
	dao *dao.Dao
}

type questionService struct {
	dao *dao.Dao
}

func New(dao *dao.Dao) *Service {
	service := &Service{
		QuestionsService: &questionsService{dao: dao},
		QuestionService:  &questionService{dao: dao},
	}
	return service
}

func (service *questionsService) Count() int {
	if result, err := service.dao.GetAll(); err == nil {
		return len(*result)
	}
	return 0
}

func (service *questionService) Count() int {
	if result1, err := service.dao.GetAllQuestion(); err == nil {
		return len(*result1)
	}
	return 0
}

func (service *questionsService) GetAll() (*[]model.Questions, error) {
	return service.dao.GetAll()
}

func (service *questionService) GetAll() (*[]model.Question, error) {
	return service.dao.GetAllQuestion()
}

func (service *questionsService) Add(questions model.Questions) error {
	return service.dao.Add(questions)
}

func (service *questionService) Add(question model.Question) error {
	return service.dao.AddQuestion(question)
}

func (service *questionsService) Delete(questions model.Questions) error {
	return service.dao.Delete(questions)
}

func (service *questionService) Delete(question model.Question) error {
	return service.dao.DeleteQuestion(question)
}

func (service *questionsService) Update(questions model.Questions) error {
	return service.dao.Update(questions)
}

func (service *questionService) Update(question model.Question) error {
	return service.dao.UpdateQuestion(question)
}

func (service *questionsService) AddBatch(questions []model.Questions) error {
	return service.dao.AddBatch(questions)
}
