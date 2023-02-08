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
	oldValue, update := databaseNames[name]
	databaseNames[name] = value

	if update {
		oldCount, ok := databaseValues[oldValue]
		if ok && oldCount > 0 {
			oldCount -= 1
			databaseValues[oldValue] = oldCount
		}
	}

	databaseValues[value] += 1
}

func getValue(name string) string {
	value, ok := databaseNames[name]

	if ok {
		return value
	}

	return "NULL"
}

func countValues(value string) int {
	count, ok := databaseValues[value]

	if ok {
		return count
	}

	return 0
}

func deleteValue(name string) {
	value, ok := databaseNames[name]

	if ok {
		delete(databaseNames, name)
		count := databaseValues[value]

		if count > 0 {
			count -= 1
			databaseValues[value] = count
		}
	}
}

func beginTransaction() {
	currentTransaction += 1
	databaseTransactions[currentTransaction] = initTransaction()
}

func rollbackTransaction() string {
	if currentTransaction == -1 {
		return "TRANSACTION NOT FOUND"
	}

	delete(databaseTransactions, currentTransaction)
	currentTransaction -= 1

	return ""
}
