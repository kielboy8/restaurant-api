package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Request struct {
	Action string `json:"action"`
}

func handleRequest(ctx context.Context, args string) (string, error) {
	dbPass := os.Getenv("IAM_DB_PASS")
	dbHost := os.Getenv("IAM_DB_HOST")
	dbUser := os.Getenv("IAM_DB_USER")
	dbPort := os.Getenv("IAM_DB_PORT")
	dbName := os.Getenv("IAM_DB_NAME")
	migrationsPath := os.Getenv("MIGRATIONS_PATH")

	log.Printf("Connecting to database %s:%s/%s as user %s", dbHost, dbPort, dbName, dbUser)
	dbUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, url.QueryEscape(dbPass), dbHost, dbPort, dbName)
	m, err := migrate.New(migrationsPath, dbUri)
	if err != nil {
		log.Println("Connection error: ", err)
		return err.Error(), nil
	}

	a := strings.Split(args, ",")
	log.Println("args:", a)
	switch a[0] {
	case "up":
		if len(a) > 1 {
			forceVersion, _ := strconv.Atoi(a[1])
			if err := m.Force(forceVersion); err != nil {
				if err == migrate.ErrNilVersion || err == migrate.ErrNoChange {
					log.Println("No change..")
					return err.Error(), nil
				}

				log.Println("Migration force error: ", err)
				return err.Error(), nil
			}
		}

		log.Println("Running migrations up...")
		if err := m.Up(); err != nil {
			if err == migrate.ErrNilVersion || err == migrate.ErrNoChange {
				log.Println("No change..")
				return err.Error(), nil
			}

			log.Println("Migration up error: ", err)
			return err.Error(), nil
		}
	case "down":
		log.Println("Running migration down...")
		if err := m.Down(); err != nil {
			log.Println("Migration down error: ", err)
			return err.Error(), nil
		}
	case "drop":
		log.Println("Running migration drop...")
		if err := m.Drop(); err != nil {
			log.Println("Migration drop err: ", err)
			return err.Error(), nil
		}
	default:
		return "Unknown command", nil
	}

	return "success", nil
}

func main() {
	lambda.Start(handleRequest)
}
