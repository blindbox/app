package utils

import (
	"github.com/srs-project/app/config"
	"github.com/srs-project/app/migrations"
	"fmt"
)

// SetupTesting the testing environment
func SetupTesting() {
	config.SetEnviroment(config.Environments["test"])

	teardown()

	err := migrations.Create()
	if err != nil {
		fmt.Printf("%s\n\n", err.Error())
	}

	migrations.Migrate()
}

func teardown() {
	err := migrations.Destroy()
	if err != nil {
		fmt.Printf("%s\n\n", err.Error())
	}
}
