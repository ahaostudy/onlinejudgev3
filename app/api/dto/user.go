package dto

type User struct {
	Id        int64  `json:"id"`
	Nickname  string `json:"nickname"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
	Role      int32  `json:"role"`
}

const (
	RoleUser  int32 = 0
	RoleAdmin int32 = 1
)

type LoginReq struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Captcha  *string `json:"captcha"`
}

type LoginResp struct {
	BaseResp
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

type RegisterReq struct {
	Email    string `json:"email,required"`
	Captcha  string `json:"captcha"`
	Password string `json:"password"`
}

type RegisterResp struct {
	BaseResp
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
}

type GetUserByIdReq struct {
	Id int64 `path:"id,require"`
}

type GetUserByUsernameReq struct {
	Username string `path:"username,require"`
}

type GetUserResp struct {
	BaseResp
	User *User `json:"user"`
}

type CreateUserReq struct {
	User
}

type CreateUserResp struct {
	BaseResp
	UserId int64 `json:"user_id"`
}

type UpdateUserReq struct {
	Id        int64   `path:"id,required"`
	Nickname  *string `json:"nickname"`
	Username  *string `json:"username"`
	Password  *string `json:"password"`
	Email     *string `json:"email"`
	Avatar    *string `json:"avatar"`
	Signature *string `json:"signature"`
	Role      *int32  `json:"role"`
}

type UpdateUserResp struct {
	BaseResp
}

type GetCaptchaReq struct {
	Email string `json:"email,required"`
}

type GetCaptchaResp struct {
	BaseResp
}

type UpdateAvatarResp struct {
	BaseResp
}

type DeleteAvatarResp struct {
	BaseResp
}
