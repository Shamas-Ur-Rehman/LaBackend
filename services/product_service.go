package services

import (
	"Laorgaincs/config"
	"Laorgaincs/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(product models.Product) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product.ID = primitive.NewObjectID()
	collection := config.GetCollection("products")

	fmt.Printf("Inserting product: %+v\n", product) // ✅ log the product for debugging

	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		fmt.Printf("Insert error: %v\n", err) // ✅ log the actual error
		return primitive.NilObjectID, err
	}

	fmt.Println("Inserted product ID:", result.InsertedID)
	return product.ID, nil
}
func GetProducts() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetCollection("products")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	cursor.All(ctx, &products)
	return products, nil
}

func GetProductByID(id string) (*models.Product, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetCollection("products")
	var product models.Product
	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	return &product, err
}

func UpdateProduct(id string, product models.Product) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetCollection("products")

	update := bson.M{
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"inventory":   product.Inventory,
		"imageUrl":    product.ImageURL,
		"category":    product.Category,
		"thc":         product.THC,
		"cbd":         product.CBD,
		"strain":      product.Strain,
		"effects":     product.Effects,
		"badge":       product.Badge,
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	return err
}

func DeleteProduct(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetCollection("products")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
