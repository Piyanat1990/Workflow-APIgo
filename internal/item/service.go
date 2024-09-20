package item

import (
	"github.com/Piyanat1990/workflow/internal/constant"
	"github.com/Piyanat1990/workflow/internal/model"

	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
}

func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db),
	}
}

func (service Service) Create(req model.RequestItem) (model.Item, error) {
	item := model.Item{
		Title:    req.Title,
		Amount:    req.Amount,
		Quantity: req.Quantity,
		Status:   constant.ItemPendingStatus,

		

	}

	if err := service.Repository.Create(&item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (service Service) Find(query model.RequestFindItem) ([]model.Item, error) {
	return service.Repository.Find(query)
}

func (service Service) UpdateStatus(id uint, status constant.ItemStatus) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err
	}

	// Fill data
	item.Status = status

	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}
/////////////////////////////// get by id//////
func (service Service) FindId(id uint ) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}


func (service Service) DeleteId(id uint) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err
	}



	if err := service.Repository.Delete(item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}


func (service Service) UpdateAll(id uint, title string,amount float64,quantity uint) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err
	}

	// Fill data
	item.Title = title
	item.Amount = amount
	item.Quantity = quantity

	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}