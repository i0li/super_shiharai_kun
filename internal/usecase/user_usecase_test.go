package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	mockrepo "github.com/i0li/super_shiharai_kun/internal/repository/mocks"
	"github.com/i0li/super_shiharai_kun/internal/testdata"
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
				corpName: testdata.Alice().CompanyName,
				userName: testdata.Alice().Name,
				email:    testdata.Alice().Email,
				password: testdata.Alice().Password,
			},
			setupMock: func(m *mockrepo.MockUserRepository) {
				m.EXPECT().FindByEmail(testdata.Alice().Email).Return(nil, gorm.ErrRecordNotFound)
				m.EXPECT().Create(gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "deplicate email",
			args: args{
				corpName: testdata.Tom().CompanyName,
				userName: testdata.Tom().Name,
				email:    testdata.Tom().Email,
				password: testdata.Tom().Password,
			},
			setupMock: func(m *mockrepo.MockUserRepository) {
				m.EXPECT().FindByEmail(testdata.Tom().Email).Return(testdata.Tom(), nil)
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
