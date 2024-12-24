package main

import (
	"database/sql"
	"fmt"
	pb "github.com/amyrd/gomicro/proto/"
	"google.golang.org/grpc"
	"log"
)

const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "Admin123"
	dbName     = "user"
)

// pointer to an object that allows you to interact with db
var db *sql.DB

func main() {
	var err error

	// Open db connection
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPassword, dbName)
	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal()
	}
	// apparently runs before it closes to makoe sure everything closes properly
	defer func() {
		if db.Close(); err != nil {
			log.Printf("error closing db: %s", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// grpc servcer setup
	grpc := grpc.NewServer()

}
