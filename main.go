package main

// Recipes API
//
// This is a sample recipes API. You can find out more about the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
//	Schemes: http
//  Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//	Contact: Mohamed Labouardy <mohamed@labouardy.com> https://labouardy.com
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta

import (
	"context"
	"distributego/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler

// var recipes []models.Recipe
// var Client *mongo.Client = database.DBSet()

// func bulkload() {
// 	recipes = make([]models.Recipe, 0)
// 	file, _ := ioutil.ReadFile("recipes.json")
// 	_ = json.Unmarshal([]byte(file), recipes)
// 	//json file loaded into recipes

// 	var listOfRecipes []interface{} //db range loaded here
// 	for _, recipe := range recipes {
// 		listOfRecipes = append(listOfRecipes, recipe)
// 	}
// 	collection := Client.Database("shadman").Collection("recipes")
// 	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	intestManyResult, err := collection.InsertMany(ctx, listOfRecipes)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("Inserted recipes: ", len(intestManyResult.InsertedIDs))

// }

func init() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/shadman?directConnection=true"))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection := client.Database(os.Getenv("shadman")).Collection("recipes")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping()
	log.Println(status)

	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)

}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/recipes", recipesHandler.NewRecipeHandler)
		api.GET("/recipes", recipesHandler.ListRecipesHandler)
		api.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
		api.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
		api.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
		//api.POST("/user/register", controller.RegisterUser)
		//secured := api.Group("/secured").Use(middleware.Auth())
		// {
		// 	secured.GET("/ping", controller.Ping)
		// 	secured.GET("/allproducts", controller.GetProducts) //ger prodcuts
		// 	secured.GET("/getproduct", controller.GetProduct)
		// 	secured.GET("/addproduct", controller.AddProduct)
		// 	secured.GET("/updateproduct", controller.UpdateProduct)
		// 	secured.GET("/deleteproduct", controller.DeleteProduct)
		// }
	}
	return router
}

func main() {

	//LoadAppConfig()

	router := initRouter()
	//router.GET("/:name", IndexHandler)
	// router.POST("/recipes", NewRecipeHandler)
	// router.GET("/recipes", ListRecipesHandler)
	// router.PUT("/recipes/:id", UpdateRecipeHandler)
	// router.DELETE("/recipes/:id", DeleteRecipeHandler)
	// router.GET("/recipes/search", SearchRecipesHandler)
	router.Run(":9090")

}
