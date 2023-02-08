package app

type transaction struct {
	currentNames  map[string]string
	currentValues map[string]int
	deletedNames  map[string]bool
}

// maintain a transaction stack
var databaseTransactions map[int]transaction

// shortcut to checking whether we have transactions
var currentTransaction int

func initTransaction() transaction {
	newTransaction := transaction{
		currentNames:  make(map[string]string, 10),
		currentValues: make(map[string]int, 10),
		deletedNames:  make(map[string]bool, 10),
	}

	return newTransaction
}

// get the current transaction's value and count for a particular name
func getCurrentTransactionData(name string) (string, int, bool) {
	var currentValue string
	var currentCount int

	found := false
	if currentTransaction != -1 {
		for i := currentTransaction; i > -1; i-- {
			// If the value is deleted, return as if the key was not found
			_, hit := databaseTransactions[i].deletedNames[name]
			if hit {
				return "", 0, false
			}

			currentValue, found = databaseTransactions[i].currentNames[name]
			if found {
				currentCount = databaseTransactions[i].currentValues[currentValue]
				break
			}
		}
	}

	// if not found in any transaction, return the committed data (if available)
	if !found {
		currentValue, found = databaseNames[name]
		if found {
			currentCount = databaseValues[currentValue]
		}
	}

	return currentValue, currentCount, found
}

// set the current value based on whether a transaction is active
func setTransactionValue(name, value string) {
	if currentTransaction != -1 {
		databaseTransactions[currentTransaction].currentNames[name] = value

		// if the name was previously deleted this transaction, remove it from the list of deleted names
		delete(databaseTransactions[currentTransaction].deletedNames, name)
	} else {
		databaseNames[name] = value
	}
}

// delete the current value for the target name.
func deleteTransactionValue(name string) {
	value, count, ok := getCurrentTransactionData(name)

	if ok {
		if currentTransaction != -1 {
			// make note of the deleted data in the transaction as well as remove it from the records
			delete(databaseTransactions[currentTransaction].currentNames, name)
			databaseTransactions[currentTransaction].deletedNames[name] = true
		} else {
			delete(databaseNames, name)
		}

		setTransactionCount(value, count-1)
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
