package main

import (
	"mini-project-3/entity"
	"mini-project-3/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Initialize database connection
	dsn := "postgres://cxlwwkljmgpncn:1a2581d43a1d0a9c81fec6091ac369809dd1d0055ae76b0f17f03b0064faf1ec@ec2-52-21-233-246.compute-1.amazonaws.com:5432/d7e6k2bou44qvu"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	// Auto Migrate the database
	db.AutoMigrate(&entity.User{})

	// Use middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWT([]byte("secret")))

	// Initialize Registration handler
	registrationHandler := handlers.NewRegistration(db)

	// Initialize Login handler
	loginHandler := handlers.NewLogin(db)

	// Initialize User service
	userService := handlers.NewUserService(db)

	// Routes
	e.POST("/register", registrationHandler.RegisterHandler)
	e.POST("/login", loginHandler.LoginHandler)
	e.POST("/topup", userService.TopUp)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
