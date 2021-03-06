package main

import (
	"github.com/prest/adapters/postgres"
	"github.com/prest/cmd"
	"github.com/himanshumps/config"
)

func main() {
	config.Load()
	postgres.Load()
	cmd.Execute()
}
