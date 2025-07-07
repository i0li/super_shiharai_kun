package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/domain"
	mockrepo "github.com/i0li/super_shiharai_kun/internal/repository/mocks"
	"github.com/i0li/super_shiharai_kun/internal/usecase"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		corpName string
		userName string
		email    string
		password string
	}
	tests := []struct {
		name      string
		args      args
		setupMock func(m *mockrepo.MockUserRepository)
		wantErr   error
	}{
		{
			name: "create user success",
			args: args{
				corpName: "new corp",
				userName: "new user",
				email:    "new@example.com",
				password: "newUserPassword",
			},
			setupMock: func(m *mockrepo.MockUserRepository) {
				m.EXPECT().FindByEmail("new@example.com").Return(nil, gorm.ErrRecordNotFound)
				m.EXPECT().Create(gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "deplicate email",
			args: args{
				corpName: "new corp",
				userName: "new user",
				email:    "existing@example.com",
				password: "newUserPassword",
			},
			setupMock: func(m *mockrepo.MockUserRepository) {
				m.EXPECT().FindByEmail("existing@example.com").Return(&domain.User{
					ID:          1,
					CompanyName: "existing corp",
					Name:        "existing user",
					Email:       "existing@example.com",
					Password:    "xxxxx",
				}, nil)
			},
			wantErr: apperr.ErrEmailAlreadyExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockrepo.NewMockUserRepository(ctrl)
			tt.setupMock(mockRepo)

			uc := usecase.NewUserUsecase(mockRepo)
			err := uc.RegisterUser(
				tt.args.corpName,
				tt.args.userName,
				tt.args.email,
				tt.args.password,
			)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
