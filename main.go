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
	v1 := echo.New()

	// Initialize database connection
	dsn := "postgres://dyxksqpzvdmatc:4c8ff428fdee5f98bac542757bdeeda4fc99fe27140ce72b1b1f92f78dee7a3c@ec2-52-54-200-216.compute-1.amazonaws.com:5432/d3ld79k0m2u1ft"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		v1.Logger.Fatal(err.Error())
	}

	// Auto Migrate the database
	db.AutoMigrate(&entity.User{}, &entity.EquipmentType{}, &entity.Transaction{}, &entity.EquipmentRentalHistory{})

	// Use middleware
	v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())

	// Initialize handlers
	registrationHandler := handlers.NewRegistration(db)
	loginHandler := handlers.NewLogin(db)
	bookingService := handlers.NewBookingService(db)
	topUpService := handlers.NewUserService(db)
	rentProductHandler := handlers.RentProductService(db)
	equipmentHandler := handlers.NewEquipmentService(db)
	userServiceHandler := handlers.NewUserService(db)

	// Routes
	v1.POST("/register", registrationHandler.RegisterHandler) //tested { "email": "example@gmail.com", "password": "test", "deposit": 10 }
	v1.POST("/login", loginHandler.LoginHandler)              //tested {"email": "example@gmail.com", "password": "test"}
	v1.POST("/book", bookingService.BookEquipment)
	v1.POST("/topup", topUpService.TopUp)
	v1.POST("/rent", rentProductHandler.RentEquipmentHandler)
	v1.POST("/equipment", equipmentHandler.PostEquipmentHandler)

	v1.GET("/equipment", equipmentHandler.GetAllEquipmentHandler)
	v1.GET("/users/:id", userServiceHandler.FindUserByID) //tested localhost:8080/users/18

	// Start the server
	v1.Logger.Fatal(v1.Start(":8080"))
}
