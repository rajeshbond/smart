package dto

type OEE struct {
	Availability float64 `json:"availability"`
	Performance  float64 `json:"performance"`
	Quality      float64 `json:"quality"`
	OEE          float64 `json:"oee"`
}
