package service

import (
	"context"
	"encoding/json"
	"lopei-grpc-server/repo"
)

// type LopeiService interface {
// 	CheckBalance(ctx context.Context, in *CheckBalanceMessage) (*ResultMessage, error)
// 	DoPayment(ctx context.Context, in *PaymentMessage) (*ResultMessage, error)
// }

type LopeiService struct {
	repo repo.LopeiRepo
	UnimplementedLopeiPaymentServer
}

func (l *LopeiService) CheckBalance(ctx context.Context, in *CheckBalanceMessage) (*ResultMessage, error) {
	lopeiId := in.LopeiId
	customer, err := l.repo.RetriveById(lopeiId)
	if err != nil {
		return nil, err
	}

	c, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}
	resutlMessage := &ResultMessage{
		Result: string(c),
		Error:  nil,
	}
	return resutlMessage, nil
}

func (l *LopeiService) DoPayment(ctx context.Context, in *PaymentMessage) (*ResultMessage, error) {
	lopeiId := in.LopeiId
	amount := in.Amount
	customer, err := l.repo.RetriveById(lopeiId)
	if err != nil {
		return nil, err
	}
	if customer.Balance < amount {
		return &ResultMessage{
			Result: "FAILED",
			Error: &Error{
				Code:    "X07",
				Message: "Insufficient Balance",
			},
		}, nil
	}
	resultMessage := &ResultMessage{
		Result: "SUCCESS",
		Error:  nil,
	}
	return resultMessage, nil
}

func NewLopeiService(repo repo.LopeiRepo) *LopeiService {
	service := new(LopeiService)
	service.repo = repo
	return service
}
