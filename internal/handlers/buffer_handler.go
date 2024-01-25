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
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

// @Summary Remove telegrams
// @Description Remove existing telegrams
// @Tags Telegram
// @Accept json
// @Produce json
// @Param request body dto.RemoveTelegramsRequest true "Remove Telegrams Request"
// @Success 200 {object} dto.RemoveTelegramsResponse
// @Router /api/remove-telegrams [delete]
func (h *BufferHandler) RemoveTelegramsHandler(c *fiber.Ctx) error {

	var request dto.RemoveTelegramsRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.RemoveTelegramsRequest{Id: request.ID}
	grpcResponse, err := h.grpcClient.RemoveTelegrams(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.RemoveTelegramsResponse{Success: grpcResponse.Success})
}

// @Summary Update Telegram By Info
// @Description Update info about telegram By Info
// @Tags Telegram
// @Accept json
// @Produce json
// @Param request body dto.UpdateTelegramByInfoRequest true "Update Telegram By Info Request"
// @Success 200 {object} dto.UpdateTelegramByInfoResponse
// @Router /api/update-telegram-by-info [put]
func (h *BufferHandler) UpdateTelegramByInfoHandler(c *fiber.Ctx) error {

	var request dto.UpdateTelegramByInfoRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.UpdateTelegramByInfoRequest{
		Telegram: buildGRPCMessageFromDTO(request.Telegram),
	}

	grpcResponse, err := h.grpcClient.UpdateTelegramByInfo(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	telegram := buildTelegramFromGRPCMessage(grpcResponse.Telegram)

	return c.Status(fiber.StatusOK).JSON(dto.UpdateTelegramByInfoResponse{Telegram: telegram})
}

// @Summary Update Telegram By Code
// @Description Update info about telegram By Code
// @Tags Telegram
// @Accept json
// @Produce json
// @Param request body dto.UpdateTelegramByCodeRequest true "Update Telegram By Code Request"
// @Success 200 {object} dto.UpdateTelegramByCodeResponse
// @Router /api/update-telegram-by-code [put]
func (h *BufferHandler) UpdateTelegramByCodeHandler(c *fiber.Ctx) error {

	var request dto.UpdateTelegramByCodeRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.UpdateTelegramByCodeRequest{
		Id:           request.ID,
		TelegramCode: request.TelegramCode,
	}

	grpcResponse, err := h.grpcClient.UpdateTelegramByCode(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	telegram := buildTelegramFromGRPCMessage(grpcResponse.Telegram)

	return c.Status(fiber.StatusOK).JSON(dto.UpdateTelegramByCodeResponse{Telegram: telegram})
}

// @Summary Get Telegram
// @Description Get Telegram by id
// @Tags Telegram
// @Accept json
// @Produce json
// @Param request body dto.GetTelegram true "Get Telegram Request"
// @Success 200 {object} dto.GetTelegram
// @Router /api/get-telegram [get]
func (h *BufferHandler) GetTelegramHandler(c *fiber.Ctx) error {

	var request dto.GetTelegramRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.GetTelegramRequest{
		Id: request.ID,
	}

	grpcResponse, err := h.grpcClient.GetTelegram(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	telegram := buildTelegramFromGRPCMessage(grpcResponse.Telegram)

	return c.Status(fiber.StatusOK).JSON(dto.GetTelegramResponse{Telegram: telegram})
}

// @Summary Get Telegrams
// @Description Get all Telegrams
// @Tags Telegram
// @Accept json
// @Produce json
// @Param request body dto.GetTelegrams true "Get Telegrams Request"
// @Success 200 {object} dto.GetTelegrams
// @Router /api/get-telegrams [get]
func (h *BufferHandler) GetTelegramsHandler(c *fiber.Ctx) error {

	var request dto.GetTelegramsRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.GetTelegramsRequest{}

	grpcResponse, err := h.grpcClient.GetTelegrams(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	telegrams := buildGRPCtoDTOTelegrams(grpcResponse.Telegrams)

	return c.Status(fiber.StatusOK).JSON(dto.GetTelegramsResponse{Telegrams: telegrams})
}

// @Summary Transfer To System
// @Description Transfer To System
// @Tags Telegram
// @Accept json
// @Produce json
// @Param request body dto.TransferToSystem true "Transfer To System Request"
// @Success 200 {object} dto.TransferToSystem
// @Router /api/get-telegrams [put]
func (h *BufferHandler) TransferToSystemHandler(c *fiber.Ctx) error {

	var request dto.TransferToSystemRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка декодирования запроса"})
	}

	grpcRequest := pb.TransferToSystemRequest{}

	grpcResponse, err := h.grpcClient.TransferToSystem(c.Context(), &grpcRequest)
	if err != nil {
		return handleGRPCError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.TransferToSystemResponse{Success: grpcResponse.Success})
}

func buildGRPCtoDTOTelegrams(grpcTelegrams []*pb.Telegram) []dto.Telegram {
	telegrams := make([]dto.Telegram, 0, len(grpcTelegrams))

	for _, grpcTelegram := range grpcTelegrams {
		telegram := buildTelegramFromGRPCMessage(grpcTelegram)
		telegrams = append(telegrams, telegram)
	}

	return telegrams
}

func handleGRPCError(c *fiber.Ctx, err error) error {
	grpcErr, _ := status.FromError(err)

	switch grpcErr.Code() {
	case codes.Unknown:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": grpcErr.Message()})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Внутренняя ошибка сервера"})
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

func buildGRPCMessageFromDTO(t dto.Telegram) *pb.Telegram {
	var waterlevelOnTime, deltaWaterlevel, waterLevelOn20h, airTemperature, icePhenomeniaState, iceHeight,
		snowHeight, precipitationDuration, headwaterLevel, averageReservoirLevel, downstreamLevel *wrapperspb.Int32Value
	var waterTemperature, waterFlow, precipitationValue, reservoirVolume, inflow, reset *wrapperspb.DoubleValue
	var icePhenomenias []*pb.IcePhenomenia
	var reservoirDate, reservoirWaterInflowDate *timestamppb.Timestamp

	if t.WaterLevelOnTime != nil {
		waterlevelOnTime = &wrapperspb.Int32Value{Value: *t.WaterLevelOnTime}
	}

	if t.DeltaWaterLevel != nil {
		deltaWaterlevel = &wrapperspb.Int32Value{Value: *t.DeltaWaterLevel}
	}

	if t.WaterLevelOn20h != nil {
		waterLevelOn20h = &wrapperspb.Int32Value{Value: *t.WaterLevelOn20h}
	}

	if t.WaterTemperature != nil {
		waterTemperature = &wrapperspb.DoubleValue{Value: *t.WaterTemperature}
	}

	if t.AirTemperature != nil {
		airTemperature = &wrapperspb.Int32Value{Value: *t.AirTemperature}
	}

	if t.IcePhenomeniaState != nil {
		icePhenomeniaState = &wrapperspb.Int32Value{Value: *t.IcePhenomeniaState}
	}

	if t.IceHeight != nil {
		iceHeight = &wrapperspb.Int32Value{Value: *t.IceHeight}
	}

	if t.SnowHeight != nil {
		snowHeight = &wrapperspb.Int32Value{Value: *t.SnowHeight}
	}

	if t.WaterFlow != nil {
		waterFlow = &wrapperspb.DoubleValue{Value: *t.WaterFlow}
	}

	if t.PrecipitationValue != nil {
		precipitationValue = &wrapperspb.DoubleValue{Value: *t.PrecipitationValue}
	}

	if t.PrecipitationDuration != nil {
		precipitationDuration = &wrapperspb.Int32Value{Value: *t.PrecipitationDuration}
	}

	if t.ReservoirDate != nil {
		reservoirDate = &timestamppb.Timestamp{Seconds: t.ReservoirDate.Unix()}
	}

	if t.HeadwaterLevel != nil {
		headwaterLevel = &wrapperspb.Int32Value{Value: *t.HeadwaterLevel}
	}

	if t.AverageReservoirLevel != nil {
		averageReservoirLevel = &wrapperspb.Int32Value{Value: *t.AverageReservoirLevel}
	}

	if t.DownstreamLevel != nil {
		downstreamLevel = &wrapperspb.Int32Value{Value: *t.DownstreamLevel}
	}

	if t.ReservoirVolume != nil {
		reservoirVolume = &wrapperspb.DoubleValue{Value: *t.ReservoirVolume}
	}

	if t.ReservoirWaterInflowDate != nil {
		reservoirWaterInflowDate = &timestamppb.Timestamp{Seconds: t.ReservoirWaterInflowDate.Unix()}
	}

	if t.Inflow != nil {
		inflow = &wrapperspb.DoubleValue{Value: *t.Inflow}
	}

	if t.Reset != nil {
		reset = &wrapperspb.DoubleValue{Value: *t.Reset}
	}

	if len(t.IcePhenomenias) != 0 {
		icePhenomenias = make([]*pb.IcePhenomenia, 0, len(t.IcePhenomenias))

		for _, ip := range t.IcePhenomenias {
			intensity := &wrapperspb.Int32Value{Value: *ip.Intensity}
			icePhenomenias = append(icePhenomenias, &pb.IcePhenomenia{
				Phenomen:  ip.Phenomen,
				Intensity: intensity,
			})
		}
	}

	grpcMessage := &pb.Telegram{
		Id:                       t.ID,
		GroupId:                  t.GroupID,
		TelegramCode:             t.TelegramCode,
		PostCode:                 t.PostCode,
		Datetime:                 &timestamppb.Timestamp{Seconds: t.Datetime.Unix()},
		IsDangerous:              t.IsDangerous,
		WaterLevelOnTime:         waterlevelOnTime,
		DeltaWaterLevel:          deltaWaterlevel,
		WaterLevelOn20H:          waterLevelOn20h,
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
		Reset_:                   reset,
	}

	return grpcMessage
}
