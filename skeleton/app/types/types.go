package types

type GenerateJwtReq struct {
	Uid int64 `json:"uid" binding:"required,min=1"`
}

type GenerateJwtResp struct {
	Token string `json:"token"`
}

type ValidatorReq struct {
	Name string `json:"name" binding:"required,min=3,max=20"`
	Age  int    `json:"age" binding:"required,min=18,max=100"`
}

type ValidatorResp struct {
	Data string `json:"data"`
}
