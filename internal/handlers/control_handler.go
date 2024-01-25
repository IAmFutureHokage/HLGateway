// Package handlers is a collection of handlers used by the HLGateway services.
//
// @title HLGateway API
// @version 1
// @description This is the HLGateway server API.
// @BasePath /api/v1
package handlers

import (
	"github.com/IAmFutureHokage/HLGateway/internal/dto"
	pb "github.com/IAmFutureHokage/HLGateway/proto/control_service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ControlHandler struct {
	grpcClient pb.HydrologyStatsServiceClient
}

func NewControlHandler(client pb.HydrologyStatsServiceClient) *ControlHandler {
	return &ControlHandler{
		grpcClient: client,
	}
}

// @Summary Add control value
// @Description Add a new control value
// @Tags Buffer
// @Accept json
// @Produce json
// @Param request body dto.AddControlValueRequest true "Add Control Value Request"
// @Success 200 {object} dto.AddControlValueResponse
// @Router /api/add-control-value [post]
func (h *ControlHandler) AddControlValueHandler(c *fiber.Ctx) error {

	var request dto.AddControlValueRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.AddControlValueRequest{
		PostCode:  request.PostCode,
		Type:      pb.ControlValueType(request.Type),
		DateStart: timestamppb.New(request.DateStart),
		Value:     request.Value,
	}

	grpcResponse, err := h.grpcClient.AddControlValue(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	controlValue := dto.ControlValue{
		PostCode:  grpcResponse.ControlValue.PostCode,
		Type:      dto.ControlValueType(grpcResponse.ControlValue.Type),
		DateStart: grpcResponse.ControlValue.DateStart.AsTime(),
		Value:     grpcResponse.ControlValue.Value,
	}

	return c.Status(fiber.StatusOK).JSON(dto.AddControlValueResponse{ControlValue: controlValue})
}

// @Summary Remove control value
// @Description Remove a control value by id
// @Tags Buffer
// @Accept json
// @Produce json
// @Param request body dto.RemoveControlValueRequest true "Remove Control Value Request"
// @Success 200 {object} dto.RemoveControlValueResponse
// @Router /api/remove-control-value [delete]
func (h *ControlHandler) RemoveControlValueHandler(c *fiber.Ctx) error {

	var request dto.RemoveControlValueRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.RemoveControlValueRequest{
		Id: request.ID,
	}

	grpcResponse, err := h.grpcClient.RemoveControlValue(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.RemoveControlValueResponse{Success: grpcResponse.Success})
}

// @Summary Update control value
// @Description Update control value
// @Tags Buffer
// @Accept json
// @Produce json
// @Param request body dto.UpdateControlValueRequest true "Update Control Value Request"
// @Success 200 {object} dto.UpdateControlValueResponse
// @Router /api/update-control-value [put]
func (h *ControlHandler) UpdateControlValueHandler(c *fiber.Ctx) error {

	var request dto.UpdateControlValueRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	var updatedControlValues []dto.ControlValue

	for _, cv := range request.ControlValues {
		grpcRequest := pb.UpdateControlValueRequest{
			ControlValues: []*pb.ControlValue{
				{
					Id:        cv.ID,
					PostCode:  cv.PostCode,
					Type:      pb.ControlValueType(cv.Type),
					DateStart: timestamppb.New(cv.DateStart),
					Value:     cv.Value,
				},
			},
		}

		grpcResponse, err := h.grpcClient.UpdateControlValue(c.Context(), &grpcRequest)
		if err != nil {
			return handleGRPCError(c, err)
		}

		updatedControlValue := dto.ControlValue{
			ID:        grpcResponse.ControlValues[0].Id,
			PostCode:  grpcResponse.ControlValues[0].PostCode,
			Type:      dto.ControlValueType(grpcResponse.ControlValues[0].Type),
			DateStart: grpcResponse.ControlValues[0].DateStart.AsTime(),
			Value:     grpcResponse.ControlValues[0].Value,
		}

		updatedControlValues = append(updatedControlValues, updatedControlValue)
	}

	return c.Status(fiber.StatusOK).JSON(dto.UpdateControlValueResponse{ControlValues: updatedControlValues})
}

// @Summary Get control value
// @Description Get slice of control value with pages
// @Tags Buffer
// @Accept json
// @Produce json
// @Param request body dto.GetControlValueRequest true "Get Control Value Request"
// @Success 200 {object} dto.GetControlValueResponse
// @Router /api/get-control-values [get]
func (h *ControlHandler) GetControlValuesHandler(c *fiber.Ctx) error {

	var request dto.GetControlValuesRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.GetControlValuesRequest{
		PostCode: request.PostCode,
		Type:     pb.ControlValueType(request.Type),
		Page:     request.Page,
	}

	grpcResponse, err := h.grpcClient.GetControlValues(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	var controlValues []dto.ControlValue

	for _, cv := range grpcResponse.ControlValues {
		controlValues = append(controlValues, dto.ControlValue{
			ID:        cv.Id,
			PostCode:  cv.PostCode,
			Type:      dto.ControlValueType(cv.Type),
			DateStart: cv.DateStart.AsTime(),
			Value:     cv.Value,
		})
	}

	response := dto.GetControlValuesResponse{
		Page:          grpcResponse.Page,
		MaxPage:       grpcResponse.MaxPage,
		ControlValues: controlValues,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary Check Water Level
// @Description Check Water Level
// @Tags Buffer
// @Accept json
// @Produce json
// @Param request body dto.CheckWaterLevelRequest true "Check Water Level Request"
// @Success 200 {object} dto.CheckWaterLevelResponse
// @Router /api/check-water-level [get]
func (h *ControlHandler) CheckWaterLevelHandler(c *fiber.Ctx) error {

	var request dto.CheckWaterLevelRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.CheckWaterLevelRequest{
		Date:     timestamppb.New(request.Date),
		PostCode: request.PostCode,
		Value:    request.Value,
	}

	grpcResponse, err := h.grpcClient.CheckWaterLevel(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.CheckWaterLevelResponse{Excess: grpcResponse.Excess})
}

// @Summary Get stats
// @Description Get stats by day in time interval (for graph)
// @Tags Buffer
// @Accept json
// @Produce json
// @Param request body dto.GetStatsRequest true "Get Stats Request"
// @Success 200 {object} dto.GetStatsResponse
// @Router /api/get-stats [get]
func (h *ControlHandler) GetStatsHandler(c *fiber.Ctx) error {

	var request dto.GetStatsRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.GetStatsRequest{
		PostCode:    request.PostCode,
		StartDate:   timestamppb.New(request.StartDate),
		EndDate:     timestamppb.New(request.EndDate),
		GraphPoints: request.GraphPoints,
	}

	grpcResponse, err := h.grpcClient.GetStats(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	var stats []dto.StatsDay

	for _, s := range grpcResponse.Stats {
		stats = append(stats, buildStatsDayFromGRPC(s))
	}

	response := dto.GetStatsResponse{
		StartInterval: grpcResponse.StartInterval.AsTime(),
		Stats:         stats,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func buildStatsDayFromGRPC(grpcStats *pb.StatsDay) dto.StatsDay {
	var waterlevel *int32
	statsDay := dto.StatsDay{
		Date:       grpcStats.Date.AsTime(),
		Norm:       grpcStats.Norm,
		Floodplain: grpcStats.Floodplain,
		Adverse:    grpcStats.Adverse,
		Dangerous:  grpcStats.Dangerous,
	}

	if grpcStats.Waterlevel != nil {
		waterlevel = &grpcStats.Waterlevel.Value
		statsDay.WaterLevel = waterlevel
	}

	return statsDay
}
