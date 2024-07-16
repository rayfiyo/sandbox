package entity

type Money struct {
	amount int
}

func NewMoney(amount int) *Money {
	return &Money{amount: amount}
}

func (m *Money) Amount() int {
	return m.amount
}
