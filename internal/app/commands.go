package app

// maintain a map of name/value pairs
var databaseNames map[string]string

// maintain a count of names with a value
var databaseValues map[string]int

func Init() {
	// start out with a database with 100 available entries and grow as needed.
	databaseNames = make(map[string]string, 100)
	databaseValues = make(map[string]int, 100)
}

func setValue(name, value string) {
	databaseNames[name] = value
	databaseValues[value] += 1
}

func getValue(name string) string {
	value, ok := databaseNames[name]

	if ok {
		return value
	}

	return "NULL"
}
