package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name string
		p    Person
		want float64
	}{
		{"Income less than 0", Person{TotalIncome: 0}, 0},
		{"Income less than 150,000", Person{TotalIncome: 150000}, 0},
		{"Income less than 150,001", Person{TotalIncome: 150001}, 0.1},
		{"Income less than 500,000", Person{TotalIncome: 500000}, 35000},
		{"Income less than 500,001", Person{TotalIncome: 500001}, 35000.15},
		{"Income less than 1,000,000", Person{TotalIncome: 1000000}, 110000},
		{"Income less than 1,000,001", Person{TotalIncome: 1000001}, 135000.2},
		{"Income less than 2,000,000", Person{TotalIncome: 2000000}, 335000},
		{"Income less than 2,000,001", Person{TotalIncome: 2000001}, 335000.35},
		{"Income less than 3,000,000", Person{TotalIncome: 3000000}, 685000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := calculateTax(tt.p.TotalIncome)
			assert.Equal(t, tt.want, got)
		})
	}
}
