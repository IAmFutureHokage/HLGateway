package dto

import "time"

type Telegram struct {
	ID                       string          `json:"id"`
	GroupID                  string          `json:"group_id"`
	TelegramCode             string          `json:"telegram_code"`
	PostCode                 string          `json:"post_code"`
	Datetime                 time.Time       `json:"datetime"`
	IsDangerous              bool            `json:"is_dangerous"`
	WaterLevelOnTime         *int32          `json:"water_level_on_time,omitempty"`
	DeltaWaterLevel          *int32          `json:"delta_water_level,omitempty"`
	WaterLevelOn20h          *int32          `json:"water_level_on20h,omitempty"`
	WaterTemperature         *float64        `json:"water_temperature,omitempty"`
	AirTemperature           *int32          `json:"air_temperature,omitempty"`
	IcePhenomeniaState       *int32          `json:"ice_phenomenia_state,omitempty"`
	IcePhenomenias           []IcePhenomenia `json:"ice_phenomenias"`
	IceHeight                *int32          `json:"ice_height,omitempty"`
	SnowHeight               *int32          `json:"snow_height,omitempty"`
	WaterFlow                *float64        `json:"water_flow,omitempty"`
	PrecipitationValue       *float64        `json:"precipitation_value,omitempty"`
	PrecipitationDuration    *int32          `json:"precipitation_duration,omitempty"`
	ReservoirDate            *time.Time      `json:"reservoir_date,omitempty"`
	HeadwaterLevel           *int32          `json:"headwater_level,omitempty"`
	AverageReservoirLevel    *int32          `json:"average_reservoir_level,omitempty"`
	DownstreamLevel          *int32          `json:"downstream_level,omitempty"`
	ReservoirVolume          *float64        `json:"reservoir_volume,omitempty"`
	ReservoirWaterInflowDate *time.Time      `json:"reservoir_water_inflow_date,omitempty"`
	Inflow                   *float64        `json:"inflow,omitempty"`
	Reset                    *float64        `json:"reset,omitempty"`
}

type IcePhenomenia struct {
	Phenomen  int32  `json:"phenomen"`
	Intensity *int32 `json:"intensity,omitempty"`
}

type AddTelegramRequest struct {
	Code string `json:"code"`
}

type AddTelegramResponse struct {
	Telegrams []Telegram `json:"telegrams"`
}

type RemoveTelegramsRequest struct {
	ID []string `json:"id"`
}

type RemoveTelegramsResponse struct {
	Success bool `json:"success"`
}

type UpdateTelegramByInfoRequest struct {
	Telegram Telegram `json:"telegram"`
}

type UpdateTelegramByInfoResponse struct {
	Telegram Telegram `json:"telegram"`
}

type UpdateTelegramByCodeRequest struct {
	ID           string `json:"id"`
	TelegramCode string `json:"telegram_code"`
}

type UpdateTelegramByCodeResponse struct {
	Telegram Telegram `json:"telegram"`
}

type GetTelegramResponse struct {
	Telegram Telegram `json:"telegram"`
}

type GetTelegramsResponse struct {
	Telegrams []Telegram `json:"telegrams"`
}

type TransferToSystemResponse struct {
	Success bool `json:"success"`
}
