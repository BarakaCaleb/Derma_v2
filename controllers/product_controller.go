package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "dermadelight/models"
    "dermadelight/utils"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = utils.ConnectDB()

func GetProducts(w http.ResponseWriter, r *http.Request) {
    var products []models.Product
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var product models.Product
        cursor.Decode(&product)
        products = append(products, product)
    }
    json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    json.NewDecoder(r.Body).Decode(&product)
    product.ID = primitive.NewObjectID()
    _, err := collection.InsertOne(context.Background(), product)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(product)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var product models.Product
    err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var product models.Product
    json.NewDecoder(r.Body).Decode(&product)
    _, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": product})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    _, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode("Product deleted")
}
