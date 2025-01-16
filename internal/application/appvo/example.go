package appvo

type ExampleReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type ExampleResp struct {
	Token string `json:"token"`
}
