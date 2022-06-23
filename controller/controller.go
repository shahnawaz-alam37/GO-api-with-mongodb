package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shahnawaz-alam37/newrepo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

const connectionString = "mongodb+srv://shahnawaz:<password>@cluster0.fdzblgv.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"			//enter the valid password before testing
const colName = "watchlist"


//important
var collection *mongo.Collection

//connect with mongoDB

func init()  {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)	

	//connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb connection successfull")
	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("collection instance is ready")

}

//mongodb helpers
//insert 1 data
func insertOneMovie(movie models.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted 1 movie in db with id",inserted.InsertedID)
}

//update 1 data
func updateOneMovie(movieId string){
	id, _:=primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"watched":true}}

	result, err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("modified count: ",result.ModifiedCount)
}

//delete 1 movie
func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	deleteCount, err := collection.DeleteOne(context.Background(),filter)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println("movie got delete with delete count:",deleteCount)
}

//delete all the movies
func deleteAllmovie()  int64{
	deleteresult , err := collection.DeleteMany(context.Background(), bson.D{{}},nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of movies delete:",deleteresult.DeletedCount)
	return deleteresult.DeletedCount
}

//get all movies
func getAllmovies() []primitive.M {
	cur, err := collection.Find(context.Background(),bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M 
		err := cur.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies

}

//actual controller 
func Getmyallmovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/x-www-form-urlencode")
	allMovies := getAllmovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}	

func Markaswatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func Deletemymovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}


func Deletemyallmovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	count := deleteAllmovie()
	json.NewEncoder(w).Encode(count)
}