package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

type Customer struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func main() {
	// Define command-line flags
	username := flag.String("username", "user", "Database username")
	password := flag.String("password", "password", "Database password")
	dbname := flag.String("dbname", "yourdb", "Database name")
	flag.Parse()

	// Construct dbURL
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", *username, *password, *dbname)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	customers, err := getCustomers(conn)
	if err != nil {
		log.Fatal("Error retrieving customers:", err)
	}

	fmt.Println("Sample Customer Data:")
	for _, customer := range customers {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Created At: %s\n",
			customer.ID, customer.Name, customer.Email, customer.CreatedAt)
	}
}

func getCustomers(conn *pgx.Conn) ([]Customer, error) {
	rows, err := conn.Query(context.Background(),
		"SELECT customer_id, name, email, created_at FROM Customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
