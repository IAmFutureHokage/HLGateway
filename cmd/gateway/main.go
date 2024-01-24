package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/IAmFutureHokage/HLGateway/docs"
	"github.com/IAmFutureHokage/HLGateway/internal/handlers"
	pb "github.com/IAmFutureHokage/HLGateway/proto"
)

func init() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	viper.SetConfigName(env)
	viper.AddConfigPath("../../config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	conn, err := grpc.Dial(viper.GetString("services.buffer-service"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC сервису: %v", err)
	}

	grpcClient := pb.NewHydrologyBufferServiceClient(conn)
	bufferHandler := handlers.NewHandler(grpcClient)

	go monitorServiceAvailability(viper.GetString("services.buffer-service"))

	// Глобальные middleware
	//app.Use(middleware.Logger())
	//app.Use(middleware.Recover())

	setupRoutes(app, bufferHandler)

	port := viper.GetInt("server.port")
	if port == 0 {
		log.Fatal("Server port is not set in the config file")
	}

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func setupRoutes(app *fiber.App, handler *handlers.Handler) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Post("/api/v1/api/add-telegram", handler.AddTelegramHandler)
}

func monitorServiceAvailability(serviceAddress string) {
	for {
		time.Sleep(5 * time.Minute)

		conn, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.FailOnNonTempDialError(true))
		if err != nil {
			log.Printf("Сервис недоступен: %v", err)
		} else {
			conn.Close()
		}
	}
}
