// Package handlers is a collection of handlers used by the HLGateway services.
//
// @title HLGateway API
// @version 1
// @description This is the HLGateway server API.
// @BasePath /api/v1
package handlers

import (
	"github.com/IAmFutureHokage/HLGateway/internal/dto"
	pb "github.com/IAmFutureHokage/HLGateway/proto/posts_service"
	"github.com/gofiber/fiber/v2"
)

type PostsHandler struct {
	grpcClient pb.PostsServiceClient
}

func NewPostsHandler(client pb.PostsServiceClient) *PostsHandler {
	return &PostsHandler{
		grpcClient: client,
	}
}

// @Summary Add post
// @Description Add a new post
// @Tags PostsService
// @Accept json
// @Produce json
// @Param request body dto.PostRequest true "Add Post Request"
// @Success 200 {object} dto.PostResponse
// @Router /api/add-post [post]
func (h *PostsHandler) AddPostHandler(c *fiber.Ctx) error {

	var request dto.PostRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.CreateRequest{
		Post: &pb.Post{
			Id:    request.Post.ID,
			Code:  request.Post.Code,
			Name:  request.Post.Name,
			River: request.Post.River,
		},
	}

	grpcResponse, err := h.grpcClient.Create(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	post := dto.Post{
		ID:    grpcResponse.Post.Id,
		Code:  grpcResponse.Post.Code,
		Name:  grpcResponse.Post.Name,
		River: grpcResponse.Post.River,
	}

	return c.Status(fiber.StatusOK).JSON(dto.PostResponse{Post: post})
}

// @Summary Delete the post
// @Description Delete the post by id
// @Tags PostsService
// @Accept json
// @Produce json
// @Param request body dto.DeletePostRequest true "Delete Post Request"
// @Success 200 {object} dto.PostResponse
// @Router /api/delete-post [delete]
func (h *PostsHandler) DeletePostHandler(c *fiber.Ctx) error {

	var request dto.DeletePostRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.DeleteRequest{Id: request.ID}

	grpcResponse, err := h.grpcClient.Delete(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	post := dto.Post{
		ID:    grpcResponse.Post.Id,
		Code:  grpcResponse.Post.Code,
		Name:  grpcResponse.Post.Name,
		River: grpcResponse.Post.River,
	}

	return c.Status(fiber.StatusOK).JSON(dto.PostResponse{Post: post})
}

// @Summary Update the post
// @Description Update the post
// @Tags PostsService
// @Accept json
// @Produce json
// @Param request body dto.PostRequest true "Update Post Request"
// @Success 200 {object} dto.PostResponse
// @Router /api/update-post [put]
func (h *PostsHandler) UpdatePostHandler(c *fiber.Ctx) error {

	var request dto.PostRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.UpdateRequest{
		Post: &pb.Post{
			Id:    request.Post.ID,
			Code:  request.Post.Code,
			Name:  request.Post.Name,
			River: request.Post.River,
		},
	}

	grpcResponse, err := h.grpcClient.Update(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	post := dto.Post{
		ID:    grpcResponse.Post.Id,
		Code:  grpcResponse.Post.Code,
		Name:  grpcResponse.Post.Name,
		River: grpcResponse.Post.River,
	}

	return c.Status(fiber.StatusOK).JSON(dto.PostResponse{Post: post})
}

// @Summary Gets posts
// @Description Get posts by page
// @Tags PostsService
// @Accept json
// @Produce json
// @Param page_number query int true "Page number"
// @Success 200 {object} dto.GetPostsPageResponse
// @Router /api/get-posts [get]
func (h *PostsHandler) GetPostsHandler(c *fiber.Ctx) error {

	pageNumber := c.QueryInt("page_number", 1)

	grpcRequest := pb.GetPageRequest{
		PageNumber: uint32(pageNumber),
	}

	grpcResponse, err := h.grpcClient.GetPage(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	posts := make([]dto.Post, len(grpcResponse.Posts))
	for i := 0; i < len(posts); i++ {
		posts[i] = dto.Post{
			ID:    grpcResponse.Posts[i].Id,
			Code:  grpcResponse.Posts[i].Code,
			Name:  grpcResponse.Posts[i].Name,
			River: grpcResponse.Posts[i].River,
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetPostsPageResponse{
		Posts:         posts,
		PageNumber:    grpcResponse.PageNumber,
		MaxPageNumber: grpcResponse.MaxPageNumber,
	})
}

// @Summary Get the post
// @Description Get the post by id
// @Tags PostsService
// @Accept json
// @Produce json
// @Param id query string true "Post ID"
// @Success 200 {object} dto.PostResponse
// @Router /api/get-post [get]
func (h *PostsHandler) GetPostHandler(c *fiber.Ctx) error {

	id := c.Query("id")

	grpcRequest := pb.GetRequest{Id: id}

	grpcResponse, err := h.grpcClient.Get(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	post := dto.Post{
		ID:    grpcResponse.Post.Id,
		Code:  grpcResponse.Post.Code,
		Name:  grpcResponse.Post.Name,
		River: grpcResponse.Post.River,
	}

	return c.Status(fiber.StatusOK).JSON(dto.PostResponse{Post: post})
}

// @Summary Find posts
// @Description Find posts by serchstring
// @Tags PostsService
// @Accept json
// @Produce json
// @Param substring query string true "Search Substring"
// @Success 200 {object} dto.PostsResponse
// @Router /api/find-posts [get]
func (h *PostsHandler) FindPostsHandler(c *fiber.Ctx) error {

	substring := c.Query("substring")

	grpcRequest := pb.FindRequest{
		Substring: substring,
	}

	grpcResponse, err := h.grpcClient.Find(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	posts := make([]dto.Post, len(grpcResponse.Posts))
	for i := 0; i < len(posts); i++ {
		posts[i] = dto.Post{
			ID:    grpcResponse.Posts[i].Id,
			Code:  grpcResponse.Posts[i].Code,
			Name:  grpcResponse.Posts[i].Name,
			River: grpcResponse.Posts[i].River,
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.PostsResponse{Posts: posts})
}

// @Summary Get all posts
// @Description Get all posts
// @Tags PostsService
// @Accept json
// @Produce json
// @Success 200 {object} dto.PostsResponse
// @Router /api/get-all-posts [get]
func (h *PostsHandler) GetAllPostsHandler(c *fiber.Ctx) error {

	grpcResponse, err := h.grpcClient.GetAll(c.Context(), &pb.GetAllRequest{})
	if err != nil {
		return handleGRPCError(c, err)
	}

	posts := make([]dto.Post, len(grpcResponse.Posts))
	for i := 0; i < len(posts); i++ {
		posts[i] = dto.Post{
			ID:    grpcResponse.Posts[i].Id,
			Code:  grpcResponse.Posts[i].Code,
			Name:  grpcResponse.Posts[i].Name,
			River: grpcResponse.Posts[i].River,
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.PostsResponse{Posts: posts})
}
