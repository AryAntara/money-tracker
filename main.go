package main

import (
	"money-tracker/command"
	"money-tracker/config"
	"os"
)

func main() {

	var cmd string
	var args []string

	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	}

	if len(os.Args) >= 3 {
		args = os.Args[2:]
	}

	insertCommand := command.NewInsertExecutor()
	seeCommand := command.NewSeeExecutor()
	resetCommand := command.NewResetExecutor()
	updateCommand := command.NewUpdateExecutor()
	deleteCommand := command.NewDeleteExecutor()
	analyticCommand := command.NewAnalyticExecutor()

	db := config.NewDatabase()
	commands := []*config.Executor{
		insertCommand,
		seeCommand,
		resetCommand,
		updateCommand,
		deleteCommand,
		analyticCommand,
	}

	db.Migrate()

	flag := config.NewCommandFlag(args)

	for _, command := range commands {
		if cmd == command.Cmd {
			command.Handler(db.Sqlx, flag)
		}
	}

}
