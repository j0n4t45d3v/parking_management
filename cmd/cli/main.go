package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"

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
	runningQuerysInDirectory(migrateDir)
}

func seed() {
	const seedDir string = "./database/seed"
	runningQuerysInDirectory(seedDir)
}

func runningQuerysInDirectory(dir string) {

	con := getConnection()
	files := listDirectory(dir)
	for _, file := range files {
		fileName := file.Name()
		pathFile := fmt.Sprintf("%v/%v", dir, fileName)
		content, err := os.ReadFile(pathFile)
		cleanError(err)
		query := string(content)
		_, errReadFile := con.Exec(query)
		cleanError(errReadFile)
		fmt.Println("Success running query")
	}
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
