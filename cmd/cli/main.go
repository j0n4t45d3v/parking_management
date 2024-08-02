package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/j0n4t45d3v/parking_management/database"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Operation not info..")
	}

	switch args[0] {
	case "migrate":
		migration()
		break
	case "seed":
		seed()
		break
	default:
		fmt.Println("Operation not exists; use 'migrate' or 'seed'")
	}
}

func migration() {
	const migrateDir string = "./database/migration"
  createMigrationsTable()
	runningQuerysInDirectory(migrateDir, "migrate")
}

func seed() {
	const seedDir string = "./database/seed"
  createSeedsTable()
	runningQuerysInDirectory(seedDir, "seed")
}

func runningQuerysInDirectory(dir string, typeCmd string) {
	con := getConnection()
	files := listDirectory(dir)

  validate := map[string]func(string) bool {
    "migrate": migrateExistsByFilename,
    "seed": seedExistsByFilename,
  }

  insert := map[string]func(string) {
    "migrate": insertMigrate,
    "seed": insertSeed,
  }


	for _, file := range files {
		fileName := file.Name()
    if validate[typeCmd](fileName) || strings.Contains(fileName, "rollback") {
      continue
    }
		pathFile := fmt.Sprintf("%v/%v", dir, fileName)
		content, err := os.ReadFile(pathFile)
		cleanError(err)
		query := string(content)
		_, errReadFile := con.Exec(query)
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

func seedExistsByFilename(fileName string) bool {
  query := "SELECT EXISTS(SELECT 1 FROM seeds WHERE seed = $1)"

  con := getConnection()
  var exists bool
  err := con.QueryRow(query, fileName).Scan(&exists)
  cleanError(err)
  return exists
}

func execQuery(query string, params ...interface{} ) {
  con := getConnection()
  var err error
  if len(params) > 0 {
    _, err = con.Exec(query, params...)
  }else{
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
