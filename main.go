package main

import (
	"freshers-bootcamp-last/handlers"
	"freshers-bootcamp-last/middleware"
	"freshers-bootcamp-last/models"
	"freshers-bootcamp-last/repositories"
	"freshers-bootcamp-last/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Connect to MySQL
	dsn := "root:root3004@tcp(localhost:3306)/retailer_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate models
	db.AutoMigrate(&models.Product{}, &models.Order{}, &models.User{}, &models.Transaction{})

	// Initialize repositories
	productRepo := repositories.NewProductRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo, productRepo)
	authService := services.NewAuthService(userRepo)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)
	orderHandler := handlers.NewOrderHandler(orderService)
	authHandler := handlers.NewAuthHandler(authService)

	// Setup Gin router
	r := gin.Default()

	// Public routes (no authentication required)
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.GET("/products", productHandler.GetAllProducts)
	r.GET("/products/:id", productHandler.GetProductById)
	r.POST("/admin/register", authHandler.RegisterAdmin)
	// Protected routes (require authentication)
	adminGroup := r.Group("/")
	adminGroup.Use(middleware.AdminMiddleware())
	{
		adminGroup.POST("/products", productHandler.CreateProduct)
		adminGroup.PUT("/products/:id", productHandler.UpdateProduct)
		adminGroup.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware())
	{

		// Order routes
		authGroup.POST("/orders", orderHandler.PlaceOrder)
		authGroup.GET("/orders", orderHandler.GetOrderHistory)
		authGroup.GET("/orders/:id", orderHandler.GetOrderDetails)
	}

	// Start server
	r.Run(":8080")
}
