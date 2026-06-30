package models

type ResponseData struct {
	Quantity    int    `json:"quantity"`
	PackDetails string `json:"pack_details"`
}

type S3PackSizesConfig struct {
	PackSizes []int `json:"pack-sizes"`
}
