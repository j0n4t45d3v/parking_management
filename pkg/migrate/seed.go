package migrate

const SeedDir string = "./database/seed"

func Seed() {
	createSeedsTable()
	runningQuerysInDirectory(SeedDir, "seed")
}

func createSeedsTable() {
	createTable := "CREATE TABLE IF NOT EXISTS seeds (" +
		"id SERIAL NOT NULL PRIMARY KEY , " +
		"seed VARCHAR(100) NOT NULL UNIQUE " +
		")"

	execQuery(createTable)
}

func insertSeed(nameFile string) {
	query := "INSERT INTO seeds (seed) VALUES ($1)"
	execQuery(query, nameFile)
}

func deleteSeed(nameFile string) {
	query := "DELETE FROM seeds WHERE seed = $1"
	execQuery(query, nameFile)
}

func seedExistsByFilename(fileName string) bool {
	query := "SELECT EXISTS(SELECT 1 FROM seeds WHERE seed = $1)"

	con := getConnection()
	var exists bool
	err := con.QueryRow(query, fileName).Scan(&exists)
	cleanError(err)
	return exists
}

