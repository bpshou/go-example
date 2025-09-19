package types

type ValidatorReq struct {
	Name string `json:"name" binding:"required,min=3,max=20"`
	Age  int    `json:"age" binding:"required,min=18,max=100"`
}

type ValidatorResp struct {
	Data string `json:"data"`
}
