package models

type Account struct {
	Id     int     `json:"id"`
	UserId int     `json:"userId"`
	Value  float32 `json:"value"`
}
