package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type User struct {
	// bson is official annotation for mongodb
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Dob string `json:"dob,omitempty" bson:"dob,omitempty"`
	Address string `json:"address,omitempty" bson:"address,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Time string `json:"time,omitempty" bson:"time,omitempty"`
}


var client *mongo.Client

// for creating a user
func createUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var user User
	user.Time = time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("test").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

// for retreiving list of all users
func getUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var users []User
	collection := client.Database("test").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err!= nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err!= nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(users)
}

// for retreiving a user
func getUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	collection := client.Database("test").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

// for deleting a user
func deleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	collection := client.Database("test").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOneAndDelete(ctx, User{ID: id}).Decode(&user)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	response.Write([]byte("User deleted"))
}

// for updating a user
func updateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)

	collection := client.Database("test").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}


	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// used dotenv
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	// database connection
	clientOptions := options.Client().ApplyURI(uri)
	client, _ = mongo.Connect(ctx, clientOptions)

	router := mux.NewRouter()
	router.HandleFunc("/create", createUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/delete/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateUser).Methods("PATCH")
	http.ListenAndServe(":8080", router)
	
}