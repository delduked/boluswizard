package types

type Users struct {
	Uid      string `json:"Uid" redis:"Uid"`
	Username string `json:"Username" redis:"Username"`
	Password string `json:"Password" redis:"Password"`
}
type SigninData struct {
	Username string `json:"Username" redis:"Username"`
	Token    string `json:"Token" redis:"Token"`
}
