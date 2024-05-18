package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/gurdeep/crud/configs"
	"github.com/gurdeep/crud/models"
	"github.com/gurdeep/crud/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Operations
//1. Create a movie - DOne
//2. Read using id - Done
//3. Update watched using id
//4. Delete a movie - DONE
//5. List all movies - Done

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection = configs.GetCollection(dbName, colName)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		sendErrorResponse(&w, err, http.StatusBadRequest)
		return
	}
	if err := validator.New().Struct(movie); err != nil {
		sendErrorResponse(&w, err, http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	movie.Id = primitive.NewObjectID()
	result, err := collection.InsertOne(ctx, &movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := responses.MovieResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := responses.MovieResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(w).Encode(response)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	objectid, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var movie models.Movie
	err := collection.FindOne(ctx, bson.M{"_id": objectid}).Decode(&movie)
	if err != nil {
		sendErrorResponse(&w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := responses.MovieResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": movie}}
	json.NewEncoder(w).Encode(response)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var movies []models.Movie
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		sendErrorResponse(&w, err, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var current models.Movie
		if err := cursor.Decode(&current); err != nil {
			sendErrorResponse(&w, err, http.StatusInternalServerError)
			return
		}
		movies = append(movies, current)
	}
	w.WriteHeader(http.StatusOK)
	response := responses.MovieResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": movies}}
	json.NewEncoder(w).Encode(response)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idObject, _ := primitive.ObjectIDFromHex(id)
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()
	filter := bson.M{"_id": idObject}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		sendErrorResponse(&w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := responses.MovieResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(w).Encode(response)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	ctx, close := context.WithTimeout(context.Background(), 10*time.Minute)
	defer close()
	params := mux.Vars(r)
	objectid, _ := primitive.ObjectIDFromHex(params["id"])

	//validate body
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		sendErrorResponse(&w, err, http.StatusBadRequest)
		return
	}
	//validate for required fields
	if err := validator.New().Struct(movie); err != nil {
		sendErrorResponse(&w, err, http.StatusBadRequest)
		return
	}
	//update
	update := bson.M{"movie_name": movie.Name, "watched": movie.Watched}
	if _, err := collection.UpdateOne(ctx, bson.M{"_id": objectid}, bson.M{"$set": update}); err != nil {
		sendErrorResponse(&w, err, http.StatusInternalServerError)
		return
	}
	//get updated record and set in reponse
	var updatedMovie models.Movie
	if err := collection.FindOne(ctx, bson.M{"_id": objectid}).Decode(&updatedMovie); err != nil {
		sendErrorResponse(&w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := responses.MovieResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedMovie}}
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w *http.ResponseWriter, err error, statusCode int) {
	(*w).WriteHeader(statusCode)
	response := responses.MovieResponse{Status: statusCode, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
	json.NewEncoder(*w).Encode(response)
}
