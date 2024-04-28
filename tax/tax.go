package tax

type Tax struct {
	Tax       float64    `json:"tax"`
}

type Allowance struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount" validate:"gte=0"`
}

type Person struct {
	TotalIncome float64     `json:"totalIncome" validate:"required,gt=0"`
	WHT         float64     `json:"wht" validate:"gte=0"`
	Allowances  []Allowance `json:"allowances"`
}