package controller

import (
	"fmt"
	"gindemo/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
)

type Trainer struct {
	Name string
	Age  int
	City string
}


func AddConfigTest(c *gin.Context) {
	client := database.DB.Mongo
	collection := client.Database(database.DBName).Collection(database.Collection)
	ash := Trainer{"Ash", 10, "Pallet Town"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}


