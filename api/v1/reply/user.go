package reply

type RegisterResp struct {
	UserID uint `json:"user_id"`
}

type LoginResp struct {
	Token     string `json:"token"`
	Duration  int64  `json:"duration"`
	CreatedAt string `json:"created_time"`
}
