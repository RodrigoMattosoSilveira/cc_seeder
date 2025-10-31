package main

import (
	"github.com/RodrigoMattosoSilveira/cc_seeder/seeders"
	"github.com/RodrigoMattosoSilveira/cc_seeder/database"
)

func main() {
	DB := database.ConnectDB()
	seeders.PersonSeeder(DB)
}
