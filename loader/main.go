package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product represents a single product data
type Product struct {
	Name         string `json:"name"`
	ImageURL     string `json:"imageURL"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	TotalReviews int    `json:"totalReviews"`
}

// ProductDetails represents product document in mongodb
type ProductDetails struct {
	ID        string    `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product   *Product  `json:"product"`
}

func UploadProductDetails(ctx context.Context, product ProductDetails, db *mongo.Database) (string, error) {
	collection := db.Collection("products")
	res, err := collection.InsertOne(ctx, product)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func main() {
	URI := os.Getenv("MONGODB_CONNSTR")
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database("sellerapp")
	router := gin.Default()
	router.POST("/products", func(c *gin.Context) {
		product := ProductDetails{}
		c.ShouldBind(&product)
		ID, err := UploadProductDetails(context.TODO(), product, db)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		product.ID = ID
		c.AbortWithStatusJSON(http.StatusCreated, product)
		return
	})
	router.Run("0.0.0.0:3001")
}
