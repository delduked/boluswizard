package types

type Response struct {
	Status int `json:"Status" redis:"Status"`
	Error  any `json:"Error" redis:"Error"`
	Data   any `json:"-"`
}
