package migrate

const MigrateDir string = "./database/migration"

func Migration() {
	createMigrationsTable()
	runningQuerysInDirectory(MigrateDir, "migrate")
}

func createMigrationsTable() {
	createTable := "CREATE TABLE IF NOT EXISTS migrations (" +
		"id SERIAL NOT NULL PRIMARY KEY , " +
		"migrate VARCHAR(100) NOT NULL UNIQUE " +
		")"

	execQuery(createTable)
}

func insertMigrate(nameFile string) {
	query := "INSERT INTO migrations (migrate) VALUES ($1)"
	execQuery(query, nameFile)
}

func deleteMigrate(nameFile string) {
	query := "DELETE FROM migrations WHERE migrate = $1"
	execQuery(query, nameFile)
}

func migrateExistsByFilename(fileName string) bool {
	query := "SELECT EXISTS(SELECT 1 FROM migrations WHERE migrate = $1)"

	con := getConnection()
	var exists bool
	err := con.QueryRow(query, fileName).Scan(&exists)
	cleanError(err)
	return exists
}
