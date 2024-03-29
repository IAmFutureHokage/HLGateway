package dto

import "time"

type ControlValue struct {
	ID        string           `json:"id"`
	PostCode  string           `json:"post_code"`
	Type      ControlValueType `json:"type"`
	DateStart time.Time        `json:"date_start"`
	Value     uint32           `json:"value"`
}

type StatsDay struct {
	Date       time.Time `json:"date"`
	Norm       uint32    `json:"norm"`
	Floodplain uint32    `json:"floodplain"`
	Adverse    uint32    `json:"adverse"`
	Dangerous  uint32    `json:"dangerous"`
	WaterLevel *int32    `json:"waterLevel,omitempty"`
}

type ControlValueType int

const (
	None ControlValueType = iota
	Norm
	Floodplain
	Adverse
	Dangerous
)

type AddControlValueRequest struct {
	PostCode  string           `json:"post_code"`
	Type      ControlValueType `json:"type"`
	DateStart time.Time        `json:"date_start"`
	Value     uint32           `json:"value"`
}

type AddControlValueResponse struct {
	ControlValue ControlValue `json:"control_value"`
}

type RemoveControlValueRequest struct {
	ID string `json:"id"`
}

type RemoveControlValueResponse struct {
	Success bool `json:"success"`
}

type UpdateControlValueRequest struct {
	ControlValues []ControlValue `json:"control_values"`
}

type UpdateControlValueResponse struct {
	ControlValues []ControlValue `json:"control_values"`
}

type GetControlValuesResponse struct {
	Page          uint32         `json:"page"`
	MaxPage       uint32         `json:"max_page"`
	ControlValues []ControlValue `json:"control_values"`
}

type CheckWaterLevelResponse struct {
	Excess uint32 `json:"excess"`
}

type GetStatsResponse struct {
	StartInterval time.Time  `json:"start_interval"`
	Stats         []StatsDay `json:"stats"`
}
