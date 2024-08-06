package migrate

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/j0n4t45d3v/parking_management/database"
)

func RollbackQueriesInDirectory(dir string, typeCmd string) {

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
  dropSeedsTable()
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
