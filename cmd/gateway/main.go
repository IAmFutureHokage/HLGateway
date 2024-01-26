package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/IAmFutureHokage/HLGateway/docs"
	"github.com/IAmFutureHokage/HLGateway/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	buffer_pb "github.com/IAmFutureHokage/HLGateway/proto/buffer_service"
	control_pb "github.com/IAmFutureHokage/HLGateway/proto/control_service"
	posts_pb "github.com/IAmFutureHokage/HLGateway/proto/posts_service"
	users_pb "github.com/IAmFutureHokage/HLGateway/proto/user_service"
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
		AllowHeaders: "Origin, Content-Type, Accept, ngrok-skip-browser-warning",
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

	//controlService
	connControl, err := grpc.Dial(viper.GetString("services.control-service"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC сервису: %v", err)
	}

	grpcControlClient := control_pb.NewHydrologyStatsServiceClient(connControl)
	controlHandler := handlers.NewControlHandler(grpcControlClient)

	//usersService
	connUsers, err := grpc.Dial(viper.GetString("services.users-service"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC сервису: %v", err)
	}

	grpcUsersClient := users_pb.NewUsersServiceClient(connUsers)
	usersHandler := handlers.NewUsersHandler(grpcUsersClient)

	//мониторинг сервисов
	go monitorServiceAvailability(viper.GetString("services.buffer-service"))
	go monitorServiceAvailability(viper.GetString("services.posts-service"))
	go monitorServiceAvailability(viper.GetString("services.control-service"))
	go monitorServiceAvailability(viper.GetString("services.users-service"))

	// Глобальные middleware
	// app.Use(middleware.Logger())
	// app.Use(middleware.Recover())

	setupBufferRoutes(app, bufferHandler)
	setupPostsRoutes(app, postsHandler)
	setupControlRoutes(app, controlHandler)
	setupUsersRoutes(app, usersHandler)

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
	app.Delete("/api/v1/api/remove-telegrams", handler.RemoveTelegramsHandler)
	app.Put("/api/v1/api/update-telegram-by-info", handler.UpdateTelegramByInfoHandler)
	app.Put("/api/v1/api/update-telegram-by-code", handler.UpdateTelegramByCodeHandler)
	app.Get("/api/v1/api/get-telegram", handler.GetTelegramHandler)
	app.Get("/api/v1/api/get-telegrams", handler.GetTelegramsHandler)
	app.Get("/api/v1/api/transfer-to-system", handler.TransferToSystemHandler)
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

func setupControlRoutes(app *fiber.App, handler *handlers.ControlHandler) {
	app.Post("/api/v1/api/add-control-value", handler.AddControlValueHandler)
	app.Delete("/api/v1/api/remove-control-value", handler.RemoveControlValueHandler)
	app.Put("/api/v1/api/update-control-value", handler.UpdateControlValueHandler)
	app.Get("/api/v1/api/get-control-values", handler.GetControlValuesHandler)
	app.Get("/api/v1/api/check-water-level", handler.CheckWaterLevelHandler)
	app.Get("/api/v1/api/get-stats", handler.GetStatsHandler)
}

func setupUsersRoutes(app *fiber.App, handler *handlers.UsersHandler) {
	app.Post("/api/v1/api/add-user", handler.AddUserHandler)
	app.Delete("/api/v1/api/delete-user", handler.DeleteUserHandler)
	app.Put("/api/v1/api/update-user", handler.UpdateUserHandler)
	app.Get("/api/v1/api/get-user", handler.GetUserHandler)
	app.Get("/api/v1/api/get-users", handler.GetUsersHandler)
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
