package repository

type LoginRequest struct {
	UserName   string `json:"userName" from:"userName" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	DeviceType string `json:"deviceType" form:"deviceType" `
	DeviceId   string `json:"deviceId" form:"deviceId"`
}

type DefaultResponse struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

type LoginResponse struct {
	DefaultResponse                   // 这是继承
	Result          LoginResponseData `json:"result"` // 这是拥有
}

type LoginResponseData struct {
	UserId     string `json:"userId"`
	IsRealName string `json:"isRealname"`
	Token      string `json:"token"`
	IdCard     string `json:"idCard"`
	RealName   string `json:"realName"`
	MedicareNo string `json:"medicareNo"`
	Photo      string `json:"photo"`
}
