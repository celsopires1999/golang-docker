package faker

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
)

type CostFakeBuilder struct {
	CostID           string
	ProjectID        string
	CostType         domain.CostType
	Description      string
	Comment          string
	Amount           float64
	Currency         domain.Currency
	InstallmentProps []domain.NewInstallmentProps
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewCostFakeBuilder() *CostFakeBuilder {
	return &CostFakeBuilder{
		CostID:      uuid.New().String(),
		ProjectID:   uuid.New().String(),
		CostType:    domain.CostType(randomdata.StringSample("one_time", "running", "investment")),
		Description: randomdata.Paragraph(),
		Comment:     randomdata.SillyName(),
		Amount:      100.0,
		Currency:    domain.Currency(randomdata.StringSample("BRL", "USD", "EUR")),
		InstallmentProps: []domain.NewInstallmentProps{
			{Year: 2020, Month: time.January, Amount: 60.},
			{Year: 2020, Month: time.August, Amount: 40.},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (b *CostFakeBuilder) WithCostID(costID string) *CostFakeBuilder {
	b.CostID = costID
	return b
}

func (b *CostFakeBuilder) WithProjectID(projectID string) *CostFakeBuilder {
	b.ProjectID = projectID
	return b
}

func (b *CostFakeBuilder) WithCostType(costType domain.CostType) *CostFakeBuilder {
	b.CostType = costType
	return b
}

func (b *CostFakeBuilder) WithDescription(description string) *CostFakeBuilder {
	b.Description = description
	return b
}

func (b *CostFakeBuilder) WithComment(comment string) *CostFakeBuilder {
	b.Comment = comment
	return b
}

func (b *CostFakeBuilder) WithAmount(amount float64) *CostFakeBuilder {
	b.Amount = amount
	return b
}

func (b *CostFakeBuilder) WithCurrency(currency domain.Currency) *CostFakeBuilder {
	b.Currency = currency
	return b
}

func (b *CostFakeBuilder) WithInstallmentProps(installments []domain.NewInstallmentProps) *CostFakeBuilder {
	b.InstallmentProps = installments
	return b
}

func (b *CostFakeBuilder) WithCreatedAt(createdAt time.Time) *CostFakeBuilder {
	b.CreatedAt = createdAt
	return b
}

func (b *CostFakeBuilder) WithUpdatedAt(updatedAt time.Time) *CostFakeBuilder {
	b.UpdatedAt = updatedAt
	return b
}

func (b *CostFakeBuilder) Build() *domain.Cost {
	installments := make([]domain.Installment, len(b.InstallmentProps))

	for i, v := range b.InstallmentProps {
		installments[i] = newInstallment(v.Year, v.Month, v.Amount)
	}

	props := domain.RestoreCostProps{
		ProjectID:    b.ProjectID,
		CostType:     b.CostType,
		Description:  b.Description,
		Comment:      b.Comment,
		Amount:       b.Amount,
		Currency:     b.Currency,
		Installments: installments,
	}

	cost := domain.RestoreCost(props)
	err := cost.Validate()
	if err != nil {
		panic(err)
	}

	return cost
}

func newInstallment(year int, month time.Month, amount float64) domain.Installment {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return domain.Installment{
		PaymentDate: date,
		Amount:      amount,
	}
}
