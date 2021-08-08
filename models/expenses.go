package models

type expenses struct {
        data map[int]int
}

func NewExpenses(d []int) Transactions {
        return NewTransactions(d)
}