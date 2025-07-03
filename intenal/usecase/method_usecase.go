package usecase

import (
	"fmt"

	"github.com/dwincahya/payment-be/intenal/entity"
	"github.com/dwincahya/payment-be/intenal/repository"
)

type PaymentMethodUseCase struct {
	PaymentMethodRepo repository.PaymentMethodRepository
}

func NewPaymentMethodUseCase(repo repository.PaymentMethodRepository) *PaymentMethodUseCase {
	return &PaymentMethodUseCase{
		PaymentMethodRepo: repo,
	}
}

func (uc *PaymentMethodUseCase) CreatePaymentMethod(code, name, desc string, orderNum int, userAction string) (*entity.PaymentMethod, error) {
	newPaymentMethod := &entity.PaymentMethod{
		Code:       code,
		Name:       name,
		Desc:       desc,
		OrderNum:   orderNum,
		UserAction: userAction,
	}
	err := uc.PaymentMethodRepo.Create(newPaymentMethod)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment method: %w", err)
	}
	return newPaymentMethod, nil
}

func (uc *PaymentMethodUseCase) GetPaymentMethodByID(id uint) (*entity.PaymentMethod, error) {
	if id == 0 {
		return nil, fmt.Errorf("payment method ID cannot be zero")
	}

	pm, err := uc.PaymentMethodRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve payment method: %w", err)
	}
	return pm, nil
}

func (uc *PaymentMethodUseCase) GetAllPaymentMethods() ([]entity.PaymentMethod, error) {
	pms, err := uc.PaymentMethodRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all payment methods: %w", err)
	}
	return pms, nil
}

func (uc *PaymentMethodUseCase) UpdatePaymentMethod(id uint, code, name, desc string, orderNum int, userAction string) (*entity.PaymentMethod, error) {
	if id == 0 {
		return nil, fmt.Errorf("payment method ID cannot be zero for update")
	}
	if code == "" || len(code) > 25 {
		return nil, fmt.Errorf("invalid payment method code: cannot be empty or exceed 25 characters")
	}
	if name == "" || len(name) > 50 {
		return nil, fmt.Errorf("invalid payment method name: cannot be empty or exceed 50 characters")
	}
	if desc == "" {
		return nil, fmt.Errorf("payment method description cannot be empty")
	}
	if orderNum <= 0 {
		return nil, fmt.Errorf("order number must be positive")
	}
	if userAction == "" || len(userAction) > 25 {
		return nil, fmt.Errorf("invalid user action: cannot be empty or exceed 25 characters")
	}

	existingPm, err := uc.PaymentMethodRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("payment method with ID %d not found for update: %w", id, err)
	}

	existingPm.Code = code
	existingPm.Name = name
	existingPm.Desc = desc
	existingPm.OrderNum = orderNum
	existingPm.UserAction = userAction

	err = uc.PaymentMethodRepo.Update(existingPm)
	if err != nil {
		return nil, fmt.Errorf("failed to update payment method: %w", err)
	}

	return existingPm, nil
}

func (uc *PaymentMethodUseCase) DeletePaymentMethod(id uint) error {
	if id == 0 {
		return fmt.Errorf("payment method ID cannot be zero for deletion")
	}

	err := uc.PaymentMethodRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete payment method with ID %d: %w", id, err)
	}
	return nil
}
