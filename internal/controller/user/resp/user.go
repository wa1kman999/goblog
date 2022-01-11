package resp

// UserResp 用户信息
type UserResp struct {
	Id       uint   `json:"id"`
	UserName string `json:"userName"`
	Role     int    `json:"role"`
}
