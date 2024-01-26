// Package handlers is a collection of handlers used by the HLGateway services.
//
// @title HLGateway API
// @version 1
// @description This is the HLGateway server API.
// @BasePath /api/v1

package handlers

import (
	"github.com/IAmFutureHokage/HLGateway/internal/dto"
	pb "github.com/IAmFutureHokage/HLGateway/proto/user_service"
	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	grpcClient pb.UsersServiceClient
}

func NewUsersHandler(client pb.UsersServiceClient) *UsersHandler {
	return &UsersHandler{
		grpcClient: client,
	}
}

// @Summary Add user
// @Description Add a new user
// @Tags UsersService
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "Add User Request"
// @Success 200 {object} dto.UserResponse
// @Router /api/add-user [post]
func (h *UsersHandler) AddUserHandler(c *fiber.Ctx) error {

	var request dto.UserRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.CreateUserRequest{
		User: &pb.UserDTO{
			Id:         request.User.ID,
			Role:       request.User.Role,
			FirstName:  request.User.FirstName,
			MiddleName: request.User.MiddleName,
			LastName:   request.User.LastName,
			Phone:      request.User.Phone,
			Login:      request.User.Login,
			Password:   request.User.Password,
		},
	}

	grpcResponse, err := h.grpcClient.CreateUser(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	user := dto.User{
		ID:         grpcResponse.User.Id,
		Role:       grpcResponse.User.Role,
		PostCode:   grpcResponse.User.PostCode,
		FirstName:  grpcResponse.User.FirstName,
		MiddleName: grpcResponse.User.MiddleName,
		LastName:   grpcResponse.User.LastName,
		Phone:      grpcResponse.User.Phone,
		Login:      grpcResponse.User.Login,
		Password:   grpcResponse.User.Password,
	}

	return c.Status(fiber.StatusOK).JSON(dto.UserResponse{User: user})
}

// @Summary Delete the user
// @Description Delete the user by id
// @Tags UsersService
// @Accept json
// @Produce json
// @Param request body dto.UserIDRequest true "Delete User Request"
// @Success 200 {object} dto.UserResponse
// @Router /api/delete-user [delete]
func (h *UsersHandler) DeleteUserHandler(c *fiber.Ctx) error {

	var request dto.UserIDRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.DeleteUserRequest{
		Id: request.ID,
	}

	grpcResponse, err := h.grpcClient.DeleteUser(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	user := dto.User{
		ID:         grpcResponse.User.Id,
		Role:       grpcResponse.User.Role,
		PostCode:   grpcResponse.User.PostCode,
		FirstName:  grpcResponse.User.FirstName,
		MiddleName: grpcResponse.User.MiddleName,
		LastName:   grpcResponse.User.LastName,
		Phone:      grpcResponse.User.Phone,
		Login:      grpcResponse.User.Login,
		Password:   grpcResponse.User.Password,
	}

	return c.Status(fiber.StatusOK).JSON(dto.UserResponse{User: user})
}

// @Summary Update user
// @Description Update user
// @Tags UsersService
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "Update User Request"
// @Success 200 {object} dto.UserResponse
// @Router /api/update-user [put]
func (h *UsersHandler) UpdateUserHandler(c *fiber.Ctx) error {

	var request dto.UserRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.UpdateUserRequest{
		User: &pb.UserDTO{
			Id:         request.User.ID,
			Role:       request.User.Role,
			FirstName:  request.User.FirstName,
			MiddleName: request.User.MiddleName,
			LastName:   request.User.LastName,
			Phone:      request.User.Phone,
			Login:      request.User.Login,
			Password:   request.User.Password,
		},
	}

	grpcResponse, err := h.grpcClient.UpdateUser(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	user := dto.User{
		ID:         grpcResponse.User.Id,
		Role:       grpcResponse.User.Role,
		PostCode:   grpcResponse.User.PostCode,
		FirstName:  grpcResponse.User.FirstName,
		MiddleName: grpcResponse.User.MiddleName,
		LastName:   grpcResponse.User.LastName,
		Phone:      grpcResponse.User.Phone,
		Login:      grpcResponse.User.Login,
		Password:   grpcResponse.User.Password,
	}

	return c.Status(fiber.StatusOK).JSON(dto.UserResponse{User: user})
}

// @Summary Get the user
// @Description Get the user by id
// @Tags UsersService
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} dto.UserResponse
// @Router /api/get-user [get]
func (h *UsersHandler) GetUserHandler(c *fiber.Ctx) error {

	id := c.Query("id")

	grpcRequest := pb.GetUserRequest{Id: id}

	grpcResponse, err := h.grpcClient.GetUser(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	user := dto.User{
		ID:         grpcResponse.User.Id,
		Role:       grpcResponse.User.Role,
		PostCode:   grpcResponse.User.PostCode,
		FirstName:  grpcResponse.User.FirstName,
		MiddleName: grpcResponse.User.MiddleName,
		LastName:   grpcResponse.User.LastName,
		Phone:      grpcResponse.User.Phone,
		Login:      grpcResponse.User.Login,
		Password:   grpcResponse.User.Password,
	}

	return c.Status(fiber.StatusOK).JSON(dto.UserResponse{User: user})
}

// @Summary Get all users
// @Description Get all users
// @Tags UsersService
// @Accept json
// @Produce json
// @Success 200 {object} dto.UsersResponse
// @Router /api/get-users [get]
func (h *UsersHandler) GetUsersHandler(c *fiber.Ctx) error {

	grpcRequest := pb.GetAllUsersRequest{}

	grpcResponse, err := h.grpcClient.GetAllUsers(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	users := make([]dto.User, len(grpcResponse.Users))

	for i := 0; i < len(users); i++ {
		users[i] = dto.User{
			ID:         grpcResponse.Users[i].Id,
			Role:       grpcResponse.Users[i].Role,
			PostCode:   grpcResponse.Users[i].PostCode,
			FirstName:  grpcResponse.Users[i].FirstName,
			MiddleName: grpcResponse.Users[i].MiddleName,
			LastName:   grpcResponse.Users[i].LastName,
			Phone:      grpcResponse.Users[i].Phone,
			Login:      grpcResponse.Users[i].Login,
			Password:   grpcResponse.Users[i].Password,
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.UsersResponse{Users: users})
}
