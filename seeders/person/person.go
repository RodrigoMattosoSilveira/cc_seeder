package person

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MyError struct {
	Code string
	Message string
}
func (m *MyError) Error() string {
	return fmt.Sprintf("Error: %s: %s", m.Code, m.Message)
}
func PersonSeeder (db *gorm.DB)  error {
	// Open the CSV file

	var count int64
	db.Model(&Person{}).Count(&count)
	if (count > 0) {
		log.Println("Database already seeded.")
		return errors.New("database already seeded")
	}
	file, err := os.Open("./data/people.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all rows from the CSV file
	people, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading CSV file:", err)
		return err
	}

	// Process the rows
	var person Person
	var persons []Person
	for _, row := range people {
		// fmt.Printf("Row %d: %v\n", i, row)
		person.Name = row[NAME]
		person.Email = row[EMAIL]
		person.Cell = row[CELL]
		hashedPassword, err := HashPassword(row[PASSWORD])
		if err != nil {
			return errors.New("unable to hash password")
		}
		err = CheckPassword(row[PASSWORD], hashedPassword)
			if err != nil {
			fmt.Println("Invalid password")
		}
		person.Password = hashedPassword
		person.Role = row[ROLE]
		persons = append(persons, person)
	}
	db.Create(persons)
	log.Println("populated people database")
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}