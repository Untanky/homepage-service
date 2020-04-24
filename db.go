package main

import (
	"fmt"
	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Init() {
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		fmt.Println("Error connecting to database")
	} else {
		fmt.Println("Database: Connected")
	}
}