package generics2



func Sum[A int](numbers []A) A {
	add := func(a, b A) A {return a + b}
	return Reduce(numbers, add, 0)
}

func SumAllTails(numsToSum ...[]int) []int {
	addTails := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}
		return append(acc, Sum(x[1:]))
	}
	return Reduce(numsToSum, addTails, []int{})
}

func Reduce[A, B any](arr []A, combiner func(B, A) B, initial B) B {
	var result = initial
	for _, x := range arr {
		result = combiner(result, x)
	}
	return result
}


type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceForR(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}

func BalanceFor(transactions []Transaction, name string) float64 {
	changeBalance := func(balance float64, t Transaction) float64 {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
		return balance
	}
	return Reduce(transactions, changeBalance, 0)
}