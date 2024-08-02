package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/j0n4t45d3v/parking_management/database"
)

const migrateDir string = "./database/migration"
const seedDir string = "./database/seed"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Operation not info..")
	}

	rollbackCmd := (len(args) > 1 && args[1] == "rollback")

	switch args[0] {
	case "migrate":
		if rollbackCmd {
			fmt.Println("Run Rollback Migrations...")
			rollbackQueriesInDirectory(migrateDir, "migrate")
			return
		}
		fmt.Println("Run Migrations...")
		migration()
		break
	case "seed":
		if rollbackCmd {
			fmt.Println("Run Rollback Seeds...")
			rollbackQueriesInDirectory(migrateDir, "seed")
			return
		}
		fmt.Println("Run Seeds...")
		seed()
		break
	default:
		fmt.Println("Operation not exists; use 'migrate' or 'seed'")
	}
}

func migration() {
	createMigrationsTable()
	runningQuerysInDirectory(migrateDir, "migrate")
}

func seed() {
	createSeedsTable()
	runningQuerysInDirectory(seedDir, "seed")
}

func rollbackQueriesInDirectory(dir string, typeCmd string) {

	delete := map[string]func(string){
		"migrate": deleteMigrate,
		"seed":    deleteSeed,
	}

	files := listDirectory(dir)
	con := getConnection()

	for _, file := range files {
		filename := file.Name()
		pathFile := fmt.Sprintf("%v/%v", dir, filename)
		content, err := os.ReadFile(pathFile)
		cleanError(err)
		contentToString := string(content)
		query := strings.Split(contentToString, "--ROLLBACK")
		if len(query) > 1 {
			_, errReadFile := con.Exec(query[1])
			cleanError(errReadFile)
			fmt.Println("Rollback query", filename)
			delete[typeCmd](filename)
		}
	}
}

func runningQuerysInDirectory(dir string, typeCmd string) {
	con := getConnection()
	files := listDirectory(dir)

	validate := map[string]func(string) bool{
		"migrate": migrateExistsByFilename,
		"seed":    seedExistsByFilename,
	}

	insert := map[string]func(string){
		"migrate": insertMigrate,
		"seed":    insertSeed,
	}

	for _, file := range files {
		fileName := file.Name()
		if validate[typeCmd](fileName) {
			continue
		}
		pathFile := fmt.Sprintf("%v/%v", dir, fileName)
		content, err := os.ReadFile(pathFile)
		cleanError(err)
		contentAsString := string(content)
    queries := strings.Split(contentAsString, "--ROLLBACK")
		_, errReadFile := con.Exec(queries[0])
		cleanError(errReadFile)
		fmt.Println("Success running query", fileName)
		insert[typeCmd](fileName)
	}
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

func execQuery(query string, params ...interface{}) {
	con := getConnection()
	var err error
	if len(params) > 0 {
		_, err = con.Exec(query, params...)
	} else {
		_, err = con.Exec(query)
	}
	cleanError(err)
}

func getConnection() *sql.DB {
	con, err := database.GetConnection()
	cleanError(err)
	return con
}

func listDirectory(dir string) []fs.DirEntry {
	files, err := os.ReadDir(dir)
	cleanError(err)
	return files
}

func cleanError(err error) {
	if err != nil {
		panic(err)
	}
}
