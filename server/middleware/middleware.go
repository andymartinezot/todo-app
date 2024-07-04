package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andymartinezot/todo-app/server/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURL(connectionString)

	client, err := mongo.connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to the MongoDB")

	collection = client.Database(dbName).Collection(collName)
	fmt.Println("Collection instance created")

}

func init() {
	loadTheEnv()
	createDBInstance()
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)
}

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDoList
	json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteOneTask(["id"])

}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count := deleteAllTasks()
	json.NewDecoder(w).Encode(count)

}

//---------- Core funcs ----------

func getAllTasks() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()){
		var result bson.M
		e := cur.Decode(&result)
		if e != nil{
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil{
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

func insertOneTask(task models.ToDoList){
	insertResult, err := collection.insertOne(context.Background(), task)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted a single record", insertResult.InsertedID)
}

func taskComplete(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"status":true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result)

}

func undoTask(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"status":false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("modified count:", result)
}

func deleteOneTask(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id":id}
	delete, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted document:", delete)

}

func deleteAllTasks() int64{
	delete, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted document:", delete)
	return delete.DeletedCount
}