package request

type RegisterReq struct {
	UserName        string `binding:"required,min=8,max=20" json:"user_name,omitempty"`
	Password        string `binding:"required,min=8,max=20" json:"password,omitempty"`
	ConfirmPassword string `binding:"required,eqfield=Password" json:"confirm_password,omitempty"`
	Email           string `binding:"required,email" json:"email,omitempty"`
	Gender          string `binding:"oneof=male female" json:"gender,omitempty"`
}

type LoginReq struct {
	Body struct {
		UserName string `json:"user_name" binding:"required,e164|email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	Header struct {
		Platform string `json:"platform" binding:"required,oneof=H5 APP"`
	}
}
