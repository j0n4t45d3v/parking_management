package main

import (
	"fmt"
	"os"

	"github.com/j0n4t45d3v/parking_management/pkg/migrate"
)

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
			migrate.RollbackQueriesInDirectory(migrate.MigrateDir, "migrate")
			return
		}
		fmt.Println("Run Migrations...")
		migrate.Migration()
		break
	case "seed":
		if rollbackCmd {
			fmt.Println("Run Rollback Seeds...")
			migrate.RollbackQueriesInDirectory(migrate.SeedDir, "seed")
			return
		}
		fmt.Println("Run Seeds...")
		migrate.Seed()
		break
	default:
		fmt.Println("Operation not exists; use 'migrate' or 'seed'")
	}
}
