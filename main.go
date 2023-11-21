package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/domingo1021/go-base-model/util"
	"log"

	"github.com/domingo1021/go-base-model/db"
	_ "github.com/lib/pq"
)

func main() {

	var err error
	var dbConn *sql.DB
	util.InitConfig("./config", "app")

	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		util.GetConfigSingleton().PGUser,
		util.GetConfigSingleton().PGPassword,
		util.GetConfigSingleton().PGHost,
		util.GetConfigSingleton().PGPort,
		util.GetConfigSingleton().PGDb,
	)

	dbConn, err = sql.Open(util.GetConfigSingleton().DBDriver, dbSource)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer dbConn.Close()

	userModel := &db.User{BaseModel: db.BaseModel{DB: dbConn}}
	ctx := context.Background()

	newUser := db.User{Name: "John Doe", Email: "johndoe@example.com"}
	logMessage := "Created new user John Doe"

	if err := userModel.TransactionalCreateUserAndLog(ctx, newUser, logMessage); err != nil {
		log.Fatalf("Failed to create user and log event: %v", err)
	} else {
		log.Println("User created and event logged successfully")
	}
}
