package models

type Reserve struct {
	Id      int  `json:"id"`
	UserId  int  `json:"userId"`
	Value   int  `json:"value"`
	Deleted bool `json:"deleted"`
}
