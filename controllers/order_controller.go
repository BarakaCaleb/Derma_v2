package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "dermadelight/models"
    "dermadelight/utils"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var orderCollection = utils.ConnectDB().Collection("orders")

func CreateOrder(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    json.NewDecoder(r.Body).Decode(&order)
    order.ID = primitive.NewObjectID()
    order.CreatedAt = time.Now()

    _, err := orderCollection.InsertOne(context.Background(), order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(order)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var order models.Order
    err := orderCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var order models.Order
    json.NewDecoder(r.Body).Decode(&order)
    _, err := orderCollection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": order})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    _, err := orderCollection.DeleteOne(context.Background(), bson.M{"_id": id})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode("Order deleted")
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
    var orders []models.Order
    cursor, err := orderCollection.Find(context.Background(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var order models.Order
        cursor.Decode(&order)
        orders = append(orders, order)
    }
    json.NewEncoder(w).Encode(orders)
}
