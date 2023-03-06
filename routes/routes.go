package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"ecoadiner.com/v1/conection"
	"go.mongodb.org/mongo-driver/bson"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func V1Route(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to v1"))
}

func GetEstudiantes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client := conection.MongoConection(ctx)
	db := client.Database("ecoa_diner")
	estudiantes := db.Collection("estudiantes")
	cursor, err := estudiantes.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var results []bson.M
	err = cursor.All(ctx, &results); 
    if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}
