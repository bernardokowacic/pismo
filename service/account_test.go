package service_test

import (
	"errors"
	"pismo/entity"
	mockrepository "pismo/mocks/repository"
	"pismo/repository"
	"pismo/service"
	"reflect"
	"testing"
)

func TestAccountService_Insert(t *testing.T) {
	response := entity.Account{
		ID:             1,
		DocumentNumber: "12345678900",
	}
	type fields struct {
		AccountRepository repository.AccountInterface
	}
	type args struct {
		account entity.Account
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		want           entity.Account
		wantErr        bool
	}{
		{
			name: "Insert account",
			fields: fields{
				AccountRepository: &mockrepository.AccountInterface{},
			},
			args: args{
				entity.Account{
					DocumentNumber: "12345679800",
				},
			},
			mockBehavior: func(f fields, a args) {
				f.AccountRepository.(*mockrepository.AccountInterface).On("Insert", a.account).Return(response, nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.AccountRepository.(*mockrepository.AccountInterface).AssertExpectations(t)
			},
			want:    response,
			wantErr: false,
		},
		{
			name: "Fail at insert account",
			fields: fields{
				AccountRepository: &mockrepository.AccountInterface{},
			},
			args: args{
				entity.Account{
					DocumentNumber: "12345679800",
				},
			},
			mockBehavior: func(f fields, a args) {
				f.AccountRepository.(*mockrepository.AccountInterface).On("Insert", a.account).Return(entity.Account{}, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.AccountRepository.(*mockrepository.AccountInterface).AssertExpectations(t)
			},
			want:    entity.Account{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			service := service.NewAccountService(tt.fields.AccountRepository)

			got, err := service.Insert(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountService.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountService.Insert() = %v, want %v", got, tt.want)
			}

			if tt.assertBehavior != nil {
				tt.assertBehavior(t, tt.fields)
			}
		})
	}
}

func TestAccountService_Get(t *testing.T) {
	response := entity.Account{
		ID:             1,
		DocumentNumber: "12345678900",
	}

	type fields struct {
		AccountRepository repository.AccountInterface
	}
	type args struct {
		accountID uint64
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		want           entity.Account
		wantErr        bool
	}{
		{
			name: "Fail at find account",
			fields: fields{
				AccountRepository: &mockrepository.AccountInterface{},
			},
			args: args{
				accountID: 1,
			},
			mockBehavior: func(f fields, a args) {
				f.AccountRepository.(*mockrepository.AccountInterface).On("Find", a.accountID).Return(entity.Account{}, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.AccountRepository.(*mockrepository.AccountInterface).AssertExpectations(t)
			},
			want:    entity.Account{},
			wantErr: true,
		},
		{
			name: "Find account",
			fields: fields{
				AccountRepository: &mockrepository.AccountInterface{},
			},
			args: args{
				accountID: 1,
			},
			mockBehavior: func(f fields, a args) {
				f.AccountRepository.(*mockrepository.AccountInterface).On("Find", a.accountID).Return(response, nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.AccountRepository.(*mockrepository.AccountInterface).AssertExpectations(t)
			},
			want:    response,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			service := service.NewAccountService(tt.fields.AccountRepository)

			got, err := service.Get(tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountService.Get() = %v, want %v", got, tt.want)
			}

			if tt.assertBehavior != nil {
				tt.assertBehavior(t, tt.fields)
			}
		})
	}
}
