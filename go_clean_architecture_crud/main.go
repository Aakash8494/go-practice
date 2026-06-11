package main

import (
	"fmt"
	"myapp/domain"
	"myapp/repository"
	"myapp/service"
	// "myapp/domain"
	// "myapp/repository"
	// "myapp/service"
)

func main() {
	fmt.Println("Clean Architecture CRUD App Initialized!")
	fmt.Println("Review the 'main.go' file to see how to inject Mongo or SQL repositories.")

	// --- SCENARIO A: MongoDB ---
	mongoCollection := connectToMongo() // You would initialize your real DB connection here
	mongoRepo := repository.NewMongoUserRepository(mongoCollection)
	userService := service.NewUserService(mongoRepo)

	// --- SCENARIO B: The Boss says switch to SQL! ---
	// sqlDB := connectToSQL() // Initialize real SQL DB connection here
	// sqlRepo := repository.NewSQLUserRepository(sqlDB)
	// userService := service.NewUserService(sqlRepo) // The service DOES NOT change!

	// --- Example Usage ---
	newUser := &domain.User{ID: "1", Name: "Aakash", Email: "aakash@example.com"}
	err := userService.RegisterUser(newUser)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User registered successfully!")
	}
}
