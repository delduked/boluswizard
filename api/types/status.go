package types

type Response[T any] struct {
	Status int `json:"Status" redis:"Status"`
	Error  any `json:"Error" redis:"Error"`
	Data   T
}
