package models

type UserIdWrapper[T any] struct {
	UserId string `json:"userId"`
	Data   T      `json:"data"`
}
