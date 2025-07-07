package usecase_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/i0li/super_shiharai_kun/internal/domain"
	mockrepo "github.com/i0li/super_shiharai_kun/internal/repository/mocks"
	"github.com/i0li/super_shiharai_kun/internal/usecase"
	"github.com/i0li/super_shiharai_kun/testdata"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestInvoiceUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		userID         int64
		paymentAmount  decimal.Decimal
		paymentDueDate time.Time
	}
	tests := []struct {
		name      string
		args      args
		setupMock func(m *mockrepo.MockInvoiceRepository)
		want      int64
		wantErr   error
	}{
		{
			name: "create invoice success",
			args: args{
				userID:         int64(1),
				paymentAmount:  testdata.StandardInvoice().PaymentAmount,
				paymentDueDate: testdata.StandardInvoice().PaymentDueDate,
			},
			setupMock: func(m *mockrepo.MockInvoiceRepository) {
				m.EXPECT().Create(gomock.Any()).DoAndReturn(func(inv *domain.Invoice) error {
					inv.ID = 1
					inv.CreatedAt = time.Now()
					inv.UpdatedAt = time.Now()
					return nil
				})
			},
			want:    int64(1),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockrepo.NewMockInvoiceRepository(ctrl)
			tt.setupMock(mockRepo)

			uc := usecase.NewInvoiceUsecase(mockRepo)
			invoiceID, err := uc.Create(
				tt.args.userID,
				tt.args.paymentAmount,
				tt.args.paymentDueDate,
			)

			require.Equal(t, invoiceID, tt.want)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestInvoiceUsecase_FindPayableInvoicesInPeriod(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		userID    int64
		startDate time.Time
		endDate   time.Time
		limit     int
		offset    int
	}
	type want struct {
		invoices   []*usecase.InvoiceOutput
		pagination usecase.Pagination
		err        error
	}
	tests := []struct {
		name      string
		args      args
		setupMock func(m *mockrepo.MockInvoiceRepository)
		want      want
		wantErr   error
	}{
		{
			name: "find payable invoice success",
			args: args{
				userID:    int64(1),
				startDate: time.Date(2025, 7, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2025, 7, 15, 0, 0, 0, 0, time.Local),
				limit:     10,
				offset:    0,
			},
			setupMock: func(m *mockrepo.MockInvoiceRepository) {
				userID := int64(1)
				start := time.Date(2025, 7, 1, 0, 0, 0, 0, time.Local)
				end := time.Date(2025, 7, 15, 0, 0, 0, 0, time.Local)
				limit := 10
				offset := 0

				m.EXPECT().FindByPaymentDueDatePeriod(
					gomock.Eq(userID),
					gomock.Eq(start),
					gomock.Eq(end),
					gomock.Eq(limit),
					gomock.Eq(offset),
				).
					Return([]*domain.Invoice{
						{
							ID:             testdata.StandardInvoice().ID,
							UserID:         testdata.StandardInvoice().UserID,
							IssueDate:      testdata.StandardInvoice().IssueDate,
							PaymentAmount:  testdata.StandardInvoice().PaymentAmount,
							Fee:            testdata.StandardInvoice().Fee,
							FeeRate:        testdata.StandardInvoice().FeeRate,
							TaxAmount:      testdata.StandardInvoice().TaxAmount,
							TaxRate:        testdata.StandardInvoice().TaxRate,
							TotalAmount:    testdata.StandardInvoice().TotalAmount,
							PaymentDueDate: testdata.StandardInvoice().PaymentDueDate,
							CreatedAt:      testdata.StandardInvoice().CreatedAt,
							UpdatedAt:      testdata.StandardInvoice().UpdatedAt,
						},
					}, nil)

				m.EXPECT().CountByPaymentDueDatePeriod(
					gomock.Eq(userID),
					gomock.Eq(start),
					gomock.Eq(end),
				).Return(int64(1), nil)
			},
			want: want{
				invoices: []*usecase.InvoiceOutput{
					{
						ID:             testdata.StandardInvoice().ID,
						IssueDate:      testdata.StandardInvoice().IssueDate,
						PaymentAmount:  testdata.StandardInvoice().PaymentAmount,
						Fee:            testdata.StandardInvoice().Fee,
						FeeRate:        testdata.StandardInvoice().FeeRate,
						TaxAmount:      testdata.StandardInvoice().TaxAmount,
						TaxRate:        testdata.StandardInvoice().TaxRate,
						TotalAmount:    testdata.StandardInvoice().TotalAmount,
						PaymentDueDate: testdata.StandardInvoice().PaymentDueDate,
					},
				},
				pagination: usecase.Pagination{
					Limit:  10,
					Offset: 0,
					Total:  1,
				},
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockrepo.NewMockInvoiceRepository(ctrl)
			tt.setupMock(mockRepo)

			uc := usecase.NewInvoiceUsecase(mockRepo)
			invoices, pagination, err := uc.FindPayableInvoicesInPeriod(
				tt.args.userID,
				tt.args.startDate,
				tt.args.endDate,
				tt.args.limit,
				tt.args.offset,
			)

			require.Equal(t, invoices, tt.want.invoices)
			require.Equal(t, pagination, tt.want.pagination)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
