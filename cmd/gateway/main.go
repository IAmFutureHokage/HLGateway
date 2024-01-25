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
	buffer_pb "github.com/IAmFutureHokage/HLGateway/proto/buffer_service"
	posts_pb "github.com/IAmFutureHokage/HLGateway/proto/posts_service"
)

func init() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	viper.SetConfigName(env)
	viper.AddConfigPath("./config")
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

	//bufferService
	connBuffer, err := grpc.Dial(viper.GetString("services.buffer-service"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC сервису: %v", err)
	}

	grpcbufferClient := buffer_pb.NewHydrologyBufferServiceClient(connBuffer)
	bufferHandler := handlers.NewBufferHandler(grpcbufferClient)

	//postsService
	connPosts, err := grpc.Dial(viper.GetString("services.posts-service"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC сервису: %v", err)
	}

	grpcPostsClient := posts_pb.NewPostsServiceClient(connPosts)
	postsHandler := handlers.NewPostsHandler(grpcPostsClient)

	// мониторинг сервисов
	go monitorServiceAvailability(viper.GetString("services.buffer-service"))
	go monitorServiceAvailability(viper.GetString("services.posts-service"))

	// Глобальные middleware
	//app.Use(middleware.Logger())
	//app.Use(middleware.Recover())

	setupBufferRoutes(app, bufferHandler)
	setupPostsRoutes(app, postsHandler)

	port := viper.GetInt("server.port")
	if port == 0 {
		log.Fatal("Server port is not set in the config file")
	}

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func setupBufferRoutes(app *fiber.App, handler *handlers.BufferHandler) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Post("/api/v1/api/add-telegram", handler.AddTelegramHandler)
}

func setupPostsRoutes(app *fiber.App, handler *handlers.PostsHandler) {
	app.Post("/api/v1/api/add-post", handler.AddPostHandler)
	app.Delete("/api/v1/api/delete-post", handler.DeletePostHandler)
	app.Put("/api/v1/api/update-post", handler.UpdatePostHandler)
	app.Get("/api/v1/api/get-posts", handler.GetPostsHandler)
	app.Get("/api/v1/api/get-post", handler.GetPostHandler)
	app.Get("/api/v1/api/find-posts", handler.FindPostsHandler)
	app.Get("/api/v1/api/get-all-posts", handler.GetAllPostsHandler)
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
