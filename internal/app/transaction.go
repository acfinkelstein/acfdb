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

// scan transactions for the queried name. Return the stored value and whether it hit.
func getCurrentTransactionValue(name string) (string, bool) {

	var currentValue string

	found := false
	if currentTransaction != -1 {
		for i := currentTransaction; i > -1; i-- {
			currentValue, found = databaseTransactions[i].currentNames[name]
			if found {
				break
			}
		}
	}

	// if not found in any transaction, return the committed data (if available)
	if !found {
		currentValue, found = databaseNames[name]
	}

	return currentValue, found
}

// set the current value based on whether a transaction is active
func setTransactionValue(name, value string) {
	if currentTransaction != -1 {
		databaseTransactions[currentTransaction].currentNames[name] = value
	} else {
		databaseNames[name] = value
	}
}

// scan transactions for the queried count. Return the stored count and whether it hit.
func getCurrentTransactionCount(value string) (int, bool) {

	var currentCount int

	found := false
	if currentTransaction != -1 {
		for i := currentTransaction; i > -1; i-- {
			currentCount, found = databaseTransactions[i].currentValues[value]
			if found {
				break
			}
		}
	}

	// if not found in any transaction, return the committed data (if available)
	if !found {
		currentCount, found = databaseValues[value]
	}

	return currentCount, found
}

// set the current count for values based on whether a transaction is active
func setTransactionCount(value string, count int) {
	if currentTransaction != -1 {
		databaseTransactions[currentTransaction].currentValues[value] = count
	} else {
		databaseValues[value] = count
	}
}
