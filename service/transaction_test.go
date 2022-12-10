package service_test

import (
	"errors"
	"pismo/entity"
	mockrepository "pismo/mocks/repository"
	"pismo/repository"
	"pismo/service"
	"reflect"
	"testing"
	"time"
)

func Test_transactiontService_Insert(t *testing.T) {
	account := entity.Account{
		ID:             1,
		DocumentNumber: "12345678900",
	}
	operationType := entity.OperationType{
		ID:          1,
		Description: "Desc",
	}
	response := entity.Transaction{
		ID:              1,
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          123.45,
		CreatedAt:       time.Now(),
	}
	type fields struct {
		TransactionRepository   repository.TransactionInterface
		OperationTypeRepository repository.OperationTypeInterface
		AccountRepository       repository.AccountInterface
	}
	type args struct {
		transaction entity.Transaction
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		want           entity.Transaction
		wantErr        bool
	}{
		{
			name: "Insert transaction",
			fields: fields{
				TransactionRepository:   &mockrepository.TransactionInterface{},
				OperationTypeRepository: &mockrepository.OperationTypeInterface{},
				AccountRepository:       &mockrepository.AccountInterface{},
			},
			args: args{
				entity.Transaction{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          123.45,
				},
			},
			mockBehavior: func(f fields, a args) {
				f.AccountRepository.(*mockrepository.AccountInterface).On("Find", a.transaction.AccountID).Return(account, nil)
				f.OperationTypeRepository.(*mockrepository.OperationTypeInterface).On("Find", a.transaction.OperationTypeID).Return(operationType, nil)
				f.TransactionRepository.(*mockrepository.TransactionInterface).On("Insert", a.transaction).Return(response, nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.AccountRepository.(*mockrepository.AccountInterface).AssertExpectations(t)
				f.OperationTypeRepository.(*mockrepository.OperationTypeInterface).AssertExpectations(t)
				f.TransactionRepository.(*mockrepository.TransactionInterface).AssertExpectations(t)
			},
			want:    response,
			wantErr: false,
		},
		{
			name: "Insert transaction user not found",
			fields: fields{
				TransactionRepository:   &mockrepository.TransactionInterface{},
				OperationTypeRepository: &mockrepository.OperationTypeInterface{},
				AccountRepository:       &mockrepository.AccountInterface{},
			},
			args: args{
				entity.Transaction{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          123.45,
				},
			},
			mockBehavior: func(f fields, a args) {
				f.AccountRepository.(*mockrepository.AccountInterface).On("Find", a.transaction.AccountID).Return(entity.Account{}, errors.New("error"))
				f.OperationTypeRepository.(*mockrepository.OperationTypeInterface).On("Find", a.transaction.OperationTypeID).Return(operationType, nil)
				f.TransactionRepository.(*mockrepository.TransactionInterface).On("Insert", a.transaction).Return(response, nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.AccountRepository.(*mockrepository.AccountInterface).AssertExpectations(t)
				f.OperationTypeRepository.(*mockrepository.OperationTypeInterface).AssertExpectations(t)
				f.TransactionRepository.(*mockrepository.TransactionInterface).AssertExpectations(t)
			},
			want:    entity.Transaction{},
			wantErr: true,
		},
		{
			name: "Insert transaction operation type not found",
			fields: fields{
				TransactionRepository:   &mockrepository.TransactionInterface{},
				OperationTypeRepository: &mockrepository.OperationTypeInterface{},
				AccountRepository:       &mockrepository.AccountInterface{},
			},
			args: args{
				entity.Transaction{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          123.45,
				},
			},
			mockBehavior: func(f fields, a args) {
				f.AccountRepository.(*mockrepository.AccountInterface).On("Find", a.transaction.AccountID).Return(account, nil)
				f.OperationTypeRepository.(*mockrepository.OperationTypeInterface).On("Find", a.transaction.OperationTypeID).Return(entity.OperationType{}, errors.New("error"))
				f.TransactionRepository.(*mockrepository.TransactionInterface).On("Insert", a.transaction).Return(response, nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.AccountRepository.(*mockrepository.AccountInterface).AssertExpectations(t)
				f.OperationTypeRepository.(*mockrepository.OperationTypeInterface).AssertExpectations(t)
				f.TransactionRepository.(*mockrepository.TransactionInterface).AssertExpectations(t)
			},
			want:    entity.Transaction{},
			wantErr: true,
		},
		{
			name: "Fail at insert transaction",
			fields: fields{
				TransactionRepository:   &mockrepository.TransactionInterface{},
				OperationTypeRepository: &mockrepository.OperationTypeInterface{},
				AccountRepository:       &mockrepository.AccountInterface{},
			},
			args: args{
				entity.Transaction{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          123.45,
				},
			},
			mockBehavior: func(f fields, a args) {
				f.AccountRepository.(*mockrepository.AccountInterface).On("Find", a.transaction.AccountID).Return(account, nil)
				f.OperationTypeRepository.(*mockrepository.OperationTypeInterface).On("Find", a.transaction.OperationTypeID).Return(operationType, nil)
				f.TransactionRepository.(*mockrepository.TransactionInterface).On("Insert", a.transaction).Return(entity.Transaction{}, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.AccountRepository.(*mockrepository.AccountInterface).AssertExpectations(t)
				f.OperationTypeRepository.(*mockrepository.OperationTypeInterface).AssertExpectations(t)
				f.TransactionRepository.(*mockrepository.TransactionInterface).AssertExpectations(t)
			},
			want:    entity.Transaction{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			service := service.NewTransactionService(tt.fields.TransactionRepository, tt.fields.OperationTypeRepository, tt.fields.AccountRepository)
			got, err := service.Insert(tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactiontService.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactiontService.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
