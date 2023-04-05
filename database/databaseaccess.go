package database

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func Connectdatabase() {
	// Define connection parameters
	server := "localhost"
	port := 1433
	user := "him"
	password := "Dhiman@123"
	database := "test"

	// Create connection string
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)

	// Open database connection
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		fmt.Println("Error connecting to database:", err.Error())
		return
	}
	defer db.Close()

	// Execute SELECT statement
	rows, err := db.Query("SELECT * FROM dbo.Numbers")
	if err != nil {
		fmt.Println("Error querying database:", err.Error())
		return
	}
	defer rows.Close()

	// Iterate over rows and print results
	for rows.Next() {
		//var id int
		var numb string
		//var email string
		//err = rows.Scan(&id, &name, &email)
		err = rows.Scan(&numb)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return
		}
		fmt.Printf("Nunmber: %s,\n", numb)
		//fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

}
