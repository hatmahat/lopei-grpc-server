package repo

import "lopei-grpc-server/model"

type LopeiRepo interface {
	RetriveById(id int32) (model.Customer, error)
}

type lopeiRepo struct {
	db []model.Customer
}

func (l *lopeiRepo) RetriveById(id int32) (model.Customer, error) {
	for _, customer := range l.db {
		if customer.LopeiId == id {
			return customer, nil
		}
	}
	return model.Customer{}, nil
}

func NewLopeiRepo() LopeiRepo {
	repo := new(lopeiRepo)
	repo.db = []model.Customer{
		{LopeiId: 1, Balance: 5000},
		{LopeiId: 2, Balance: 1000},
		{LopeiId: 3, Balance: 15000},
	}
	return repo
}
