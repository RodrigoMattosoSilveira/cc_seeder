package main

import (
	"github.com/RodrigoMattosoSilveira/cc_seeder/seeders/person"
	"github.com/RodrigoMattosoSilveira/cc_seeder/database"
)

func main() {
	DB := database.ConnectDB()
	person.PersonSeeder(DB)
}
