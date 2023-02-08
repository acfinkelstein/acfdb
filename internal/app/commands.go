package app

// maintain a map of name/value pairs
var databaseNames map[string]string

// maintain a count of names with a value
var databaseValues map[string]int

func Init() {
	// start out with a database with 100 available entries and grow as needed.
	databaseNames = make(map[string]string, 100)
	databaseValues = make(map[string]int, 100)
	databaseTransactions = make(map[int]transaction, 5)
	currentTransaction = -1
}

func setValue(name, value string) {
	currentValue, currentCount, update := getCurrentTransactionData(name)
	setTransactionValue(name, value)

	if update && currentCount > 0 {
		setTransactionCount(currentValue, currentCount-1)
	}

	newCount, ok := getCurrentTransactionCount(value)
	if ok {
		setTransactionCount(value, newCount+1)
	} else {
		setTransactionCount(value, 1)
	}
}

func getValue(name string) string {
	value, _, ok := getCurrentTransactionData(name)

	if ok {
		return value
	}

	return "NULL"
}

func countValues(value string) int {
	count, ok := getCurrentTransactionCount(value)

	if ok {
		return count
	}

	return 0
}

func deleteValue(name string) {
	deleteTransactionValue(name)
}

func beginTransaction() {
	currentTransaction += 1
	databaseTransactions[currentTransaction] = initTransaction()
}

func rollbackTransaction() bool {
	// Ensure there is a transaction to rollback
	if currentTransaction == -1 {
		return false
	}

	// remove the most recent transaction
	delete(databaseTransactions, currentTransaction)
	currentTransaction -= 1

	return true
}

func commitTransaction() {
	if currentTransaction != -1 {
		for i := 0; i < currentTransaction+1; i++ {
			transaction := databaseTransactions[i]
			for name, value := range transaction.currentNames {
				databaseNames[name] = value
			}
			for value, count := range transaction.currentValues {
				databaseValues[value] = count
			}
			for name, deleted := range transaction.deletedNames {
				if deleted {
					delete(databaseNames, name)
				}
			}

			// cleanup
			delete(databaseTransactions, i)
		}
	}

	// no more transactions
	currentTransaction = -1
}
