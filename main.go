package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	PORT          string `mapstructure:"PORT"`
	DB_CONNECTION string `mapstructure:"DB_CONNECTION"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}
	}

	config := Config{
		PORT:          viper.GetString("PORT"),
		DB_CONNECTION: viper.GetString("DB_CONNECTION"),
	}

	db, err := database.ConnectDB(config.DB_CONNECTION)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Dependency Injection for Category
	CategoryRepository := repositories.NewCategoryRepository(db)
	CategoryService := services.NewCategoryService(CategoryRepository)
	CategoryHandler := handlers.NewCategoryHandler(CategoryService)

	// Setup category routes
	http.HandleFunc("/api/categories", CategoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", CategoryHandler.HandleCategoryByID)

	// Dependency Injection for Product
	ProductRepository := repositories.NewProductRepository(db)
	ProductService := services.NewProductService(ProductRepository)
	ProductHandler := handlers.NewProductHandler(ProductService)

	// Setup product routes
	http.HandleFunc("/api/products", ProductHandler.HandleProducts)
	http.HandleFunc("/api/products/", ProductHandler.HandleProductByID)

	// Dependency Injection for Transaction
	TransactionRepository := repositories.NewTransactionRepository(db)
	TransactionService := services.NewTransactionService(TransactionRepository)
	TransactionHandler := handlers.NewTransactionHandler(TransactionService)

	// Setup transaction routes
	http.HandleFunc("/api/checkout", TransactionHandler.HandleCheckout)

	// Dependency Injection for Report
	ReportRepository := repositories.NewReportRepository(db)
	ReportService := services.NewReportService(ReportRepository)
	ReportHandler := handlers.NewReportHandler(ReportService)

	// Setup report routes
	http.HandleFunc("/api/report", ReportHandler.HandleReport)
	http.HandleFunc("/api/report/today", ReportHandler.HandleReportToday)

	// API Endpoint: Root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Welcome to the Kasir API using Golang!",
		})
	})

	// API Endpoint: Server Health Check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Server is healthy",
		})
	})

	addr := "0.0.0.0:" + config.PORT
	fmt.Println("Server running locally on " + addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
