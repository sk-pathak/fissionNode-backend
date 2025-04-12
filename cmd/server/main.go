package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	config "github.com/sk-pathak/fissionNode-backend/configs"
	"github.com/sk-pathak/fissionNode-backend/internal/app/auth"
	handler "github.com/sk-pathak/fissionNode-backend/internal/app/handler"
	repo "github.com/sk-pathak/fissionNode-backend/internal/app/repository"
	routes "github.com/sk-pathak/fissionNode-backend/internal/app/routes"
	service "github.com/sk-pathak/fissionNode-backend/internal/app/service"
	db "github.com/sk-pathak/fissionNode-backend/internal/db"
	middlewares "github.com/sk-pathak/fissionNode-backend/internal/middlewares"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	connStr := cfg.DBDriver + "://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName

	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer dbPool.Close()

	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	queries := db.New(dbPool)

	userRepo := repo.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	store := auth.NewAuth()
	authHandler := handler.NewAuthHandler(service.NewAuthService([]byte(cfg.SessionSecret)))

	r := gin.Default()
	r.Use(middlewares.SetupCORS())

	routes.RegisterAuthRoutes(r, authHandler)
	routes.RegisterUserRoutes(r, userHandler, store)

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
