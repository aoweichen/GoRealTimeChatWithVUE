package AuthHandlerModels

type LoginResponse struct {
	ID          int64  `json:"id"`
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	ExpireTime  int64  `json:"expire_time"`
	TokenToLive int64  `json:"token_duration"`
}
