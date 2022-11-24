package types

type Response struct {
	Status int         `json:"Status" redis:"Status"`
	Error  interface{} `json:"Error" redis:"Error"`
	Data   interface{} `json:"-"`
}
