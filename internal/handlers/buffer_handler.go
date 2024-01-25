// Package handlers is a collection of handlers used by the HLGateway services.
//
// @title HLGateway API
// @version 1
// @description This is the HLGateway server API.
// @BasePath /api/v1
package handlers

import (
	"time"

	"github.com/IAmFutureHokage/HLGateway/internal/dto"
	pb "github.com/IAmFutureHokage/HLGateway/proto/buffer_service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BufferHandler struct {
	grpcClient pb.HydrologyBufferServiceClient
}

func NewBufferHandler(client pb.HydrologyBufferServiceClient) *BufferHandler {
	return &BufferHandler{
		grpcClient: client,
	}
}

// @Summary Add telegram
// @Description Add a new telegram
// @Tags Buffer
// @Accept json
// @Produce json
// @Param request body dto.AddTelegramRequest true "Add Telegram Request"
// @Success 200 {object} dto.AddTelegramResponse
// @Router /api/add-telegram [post]
func (h *BufferHandler) AddTelegramHandler(c *fiber.Ctx) error {

	var request dto.AddTelegramRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.AddTelegramRequest{Code: request.Code}
	grpcResponse, err := h.grpcClient.AddTelegram(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	telegrams := buildTelegramsFromGRPCResponse(grpcResponse)

	return c.Status(fiber.StatusOK).JSON(dto.AddTelegramResponse{Telegrams: telegrams})
}

// // @Summary Remove telegram
// // @Description Remove existing telegram
// // @Tags Telegram
// // @Accept json
// // @Produce json
// // @Param request body dto.RemoveTelegramRequest true "Remove Telegram Request"
// // @Success 200 {object} dto.RemoveTelegramResponse
// // @Router /api/remove-telegram [post]
// func (h *BufferHandler) RemoveTelegramHandler(c *fiber.Ctx) error {

// 	var request dto.RemoveTelegramsRequest
// 	if err := c.BodyParser(&request); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
// 	}

// 	grpcRequest := pb.RemoveTelegramsRequest{Id: request.ID}
// 	grpcResponse, err := h.grpcClient.RemoveTelegrams(c.Context(), &grpcRequest)
// 	if err != nil {
// 		return handleGRPCError(c, err)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(dto.RemoveTelegramsResponse{Success: telegrams})
// }

// // @Summary Update Telegram By Info
// // @Description Update info about telegram By Info
// // @Tags Telegram
// // @Accept json
// // @Produce json
// // @Param request body dto.UpdateTelegramByInfoRequest true "Update Telegram By Info Request"
// // @Success 200 {object} dto.UpdateTelegramByInfoResponse
// // @Router /api/update-telegram-by-info [post]
// func (h *BufferHandler) UpdateTelegramByInfoHandler(c *fiber.Ctx) error {

// 	var request dto.UpdateTelegramByInfoRequest
// 	if err := c.BodyParser(&request); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
// 	}

// 	grpcRequest := pb.UpdateTelegramByInfoRequest{Telegram: request.Telegram}
// 	grpcResponse, err := h.grpcClient.UpdateTelegramByInfo(c.Context(), &grpcRequest)
// 	if err != nil {
// 		return handleGRPCError(c, err)
// 	}

// 	telegrams := buildTelegramsFromGRPCResponse(grpcResponse)

// 	return c.Status(fiber.StatusOK).JSON(dto.UpdateTelegramByInfoResponse{Telegrams: telegrams})
// }

func handleGRPCError(c *fiber.Ctx, err error) error {
	grpcErr, _ := status.FromError(err)

	switch grpcErr.Code() {
	case codes.Unknown:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": grpcErr.Message()})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": grpcErr.Message()})
	}
}

func buildTelegramsFromGRPCResponse(response *pb.AddTelegramResponse) []dto.Telegram {
	var telegrams []dto.Telegram

	for _, t := range response.Telegrams {
		telegram := buildTelegramFromGRPCMessage(t)
		telegrams = append(telegrams, telegram)
	}

	return telegrams
}

func buildTelegramFromGRPCMessage(t *pb.Telegram) dto.Telegram {
	location, _ := time.LoadLocation("Local")

	var waterlevelOnTime, deltaWaterlevel, waterLevelOn20h, airTemperature, icePhenomeniaState, iceHeight,
		snowHeight, precipitationDuration, headwaterLevel, averageReservoirLevel, downstreamLevel, intensity *int32
	var waterTemperature, waterFlow, precipitationValue, reservoirVolume, inflow, reset *float64
	var icePhenomenias []dto.IcePhenomenia
	var reservoirDate, reservoirWaterInflowDate *time.Time

	if t.WaterLevelOnTime != nil {
		waterlevelOnTime = &t.WaterLevelOnTime.Value
	}

	if t.DeltaWaterLevel != nil {
		deltaWaterlevel = &t.DeltaWaterLevel.Value
	}

	if t.WaterLevelOn20H != nil {
		waterLevelOn20h = &t.WaterLevelOn20H.Value
	}

	if t.WaterTemperature != nil {
		waterTemperature = &t.WaterTemperature.Value
	}

	if t.AirTemperature != nil {
		airTemperature = &t.AirTemperature.Value
	}

	if t.IcePhenomeniaState != nil {
		icePhenomeniaState = &t.IcePhenomeniaState.Value
	}

	if t.IceHeight != nil {
		iceHeight = &t.IceHeight.Value
	}

	if t.SnowHeight != nil {
		snowHeight = &t.SnowHeight.Value
	}

	if t.WaterFlow != nil {
		waterFlow = &t.WaterFlow.Value
	}

	if t.PrecipitationValue != nil {
		precipitationValue = &t.PrecipitationValue.Value
	}

	if t.PrecipitationDuration != nil {
		precipitationDuration = &t.PrecipitationDuration.Value
	}

	if t.ReservoirDate != nil {
		buffer := t.ReservoirDate.AsTime().In(location)
		reservoirDate = &buffer
	}

	if t.HeadwaterLevel != nil {
		headwaterLevel = &t.HeadwaterLevel.Value
	}

	if t.AverageReservoirLevel != nil {
		averageReservoirLevel = &t.AverageReservoirLevel.Value
	}

	if t.DownstreamLevel != nil {
		downstreamLevel = &t.DownstreamLevel.Value
	}

	if t.ReservoirVolume != nil {
		reservoirVolume = &t.ReservoirVolume.Value
	}

	if t.ReservoirWaterInflowDate != nil {
		buffer := t.ReservoirWaterInflowDate.AsTime().In(location)
		reservoirWaterInflowDate = &buffer
	}

	if t.Inflow != nil {
		inflow = &t.Inflow.Value
	}

	if t.Reset_ != nil {
		reset = &t.Reset_.Value
	}

	if len(t.IcePhenomenias) != 0 {
		icePhenomenias = make([]dto.IcePhenomenia, 0, len(t.IcePhenomenias))

		for _, ip := range t.IcePhenomenias {
			if ip.Intensity != nil {
				intensity = &ip.Intensity.Value
			}

			icePhenomenias = append(icePhenomenias, dto.IcePhenomenia{
				Phenomen:  ip.Phenomen,
				Intensity: intensity,
			})
		}
	}

	telegram := dto.Telegram{
		ID:                       t.Id,
		GroupID:                  t.GroupId,
		TelegramCode:             t.TelegramCode,
		PostCode:                 t.PostCode,
		Datetime:                 t.Datetime.AsTime().In(location),
		IsDangerous:              t.IsDangerous,
		WaterLevelOnTime:         waterlevelOnTime,
		DeltaWaterLevel:          deltaWaterlevel,
		WaterLevelOn20h:          waterLevelOn20h,
		WaterTemperature:         waterTemperature,
		AirTemperature:           airTemperature,
		IcePhenomeniaState:       icePhenomeniaState,
		IcePhenomenias:           icePhenomenias,
		IceHeight:                iceHeight,
		SnowHeight:               snowHeight,
		WaterFlow:                waterFlow,
		PrecipitationValue:       precipitationValue,
		PrecipitationDuration:    precipitationDuration,
		ReservoirDate:            reservoirDate,
		HeadwaterLevel:           headwaterLevel,
		AverageReservoirLevel:    averageReservoirLevel,
		DownstreamLevel:          downstreamLevel,
		ReservoirVolume:          reservoirVolume,
		ReservoirWaterInflowDate: reservoirWaterInflowDate,
		Inflow:                   inflow,
		Reset:                    reset,
	}

	return telegram
}
