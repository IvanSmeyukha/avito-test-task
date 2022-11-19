package models

type Transaction struct {
	Id        int     `json:"id"`
	UserId    int     `json:"userId"`
	ServiceId int     `json:"serviceId"`
	OrderId   int     `json:"orderId"`
	ReserveId int     `json:"reserveId"`
	Value     float32 `json:"value"`
	Date      string  `json:"date"`
	Type      string  `json:"type"`
}
