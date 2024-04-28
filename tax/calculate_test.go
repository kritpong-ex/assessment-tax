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

func TestCalculateAllowances(t *testing.T) {
	tests := []struct {
		name       string
		person     Person
		wantDonate float64
		wantKRec   float64
	}{
		{name: "No allowances", person: Person{Allowances: []Allowance{}}, wantDonate: 0, wantKRec: 0},
		{name: "Donation less than limit", person: Person{Allowances: []Allowance{{AllowanceType: "donation", Amount: 80000}}}, wantDonate: 80000, wantKRec: 0},
		{name: "Donation exceeds limit", person: Person{Allowances: []Allowance{{AllowanceType: "donation", Amount: 120000}}}, wantDonate: 100000, wantKRec: 0},
		{name: "K-receipt less than limit", person: Person{Allowances: []Allowance{{AllowanceType: "k-receipt", Amount: 40000}}}, wantDonate: 0, wantKRec: 40000},
		{name: "K-receipt exceeds limit", person: Person{Allowances: []Allowance{{AllowanceType: "k-receipt", Amount: 60000}}}, wantDonate: 0, wantKRec: kReceiptAmount},
		{name: "Multiple allowances", person: Person{Allowances: []Allowance{{AllowanceType: "donation", Amount: 90000}, {AllowanceType: "k-receipt", Amount: 45000}}}, wantDonate: 90000, wantKRec: 45000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDonate, gotKRec := calculateAllowances(&tt.person)
			assert.Equal(t, tt.wantDonate, gotDonate)
			assert.Equal(t, tt.wantKRec, gotKRec)
		})
	}
}

func TestCalculateNetIncome(t *testing.T) {
	tests := []struct {
		name string
		p    Person
		want float64
	}{
		{"Income less than 0 without donation or k-receipt", Person{TotalIncome: 0, WHT: 0, Allowances: []Allowance{{AllowanceType: "donation", Amount: 0}, {AllowanceType: "k-receipt", Amount: 0}}}, 0},
		{"Income less than 500000 with donation", Person{TotalIncome: 500000.0, WHT: 0, Allowances: []Allowance{{AllowanceType: "donation", Amount: 100000.0}, {AllowanceType: "k-receipt", Amount: 0}}}, 340000.0},
		{"Income less than 500000 with k-receipt", Person{TotalIncome: 500000.0, WHT: 0, Allowances: []Allowance{{AllowanceType: "donation", Amount: 0}, {AllowanceType: "k-receipt", Amount: 200000.0}}}, 390000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateNetIncome(&tt.p)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculateWht(t *testing.T) {
    tests := []struct {
        name     string
        person   Person
        tax      float64
        wantTax  float64
    }{
        {name: "No WHT", person: Person{WHT: 0}, tax: 100000, wantTax: 100000},
        {name: "WHT less than tax", person: Person{WHT: 20000}, tax: 100000, wantTax: 80000},
        {name: "WHT greater than tax", person: Person{WHT: 200000}, tax: 100000, wantTax: 0},
        {name: "Tax is negative", person: Person{WHT: 50000}, tax: -10000, wantTax: 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotTax := calculateWht(&tt.person, tt.tax)
            assert.Equal(t, tt.wantTax, gotTax)
        })
    }
}
