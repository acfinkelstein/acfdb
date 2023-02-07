package app

type transaction struct {
	currentNames  map[string]string
	currentValues map[string]int
}

func initTransaction() transaction {
	newTransaction := transaction{
		currentNames:  make(map[string]string, 10),
		currentValues: make(map[string]int, 10),
	}

	return newTransaction
}
