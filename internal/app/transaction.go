package app

type transaction struct {
	currentNames  map[string]string
	currentValues map[string]int
}

// maintain a transaction stack
var databaseTransactions map[int]transaction

// shortcut to checking whether we have transactions
var currentTransaction int

func initTransaction() transaction {
	newTransaction := transaction{
		currentNames:  make(map[string]string, 10),
		currentValues: make(map[string]int, 10),
	}

	return newTransaction
}
