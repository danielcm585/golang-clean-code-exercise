package models

type transactions struct {
	data map[int]int
}

func NewTransactions(d []int) Transactions {
	data := make(map[int]int)

	for i, v := range d {
		data[i] = v;
	}

	return &transactions{data: data}
}

func (t *transactions) Get(idx int) int {
	return t.data[idx]
}

func (t *transactions) GetTotal() int {
	ret := 0
	for _, v := range t.data {
		ret += v
	}
	return ret
}

func (t *transactions) GetTotalWithinRange(i, j int) int {
	ret := 0
	for ; i <= min(j,len(t.data)); i++ {
		ret += t.data[i]
	}
	return ret
}