// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Role            string `json:"role"`
	ProfileImageUrl string `json:"profile_image_url"`
	Token           string `json:"token"`
}

type RegisterRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type RegisterResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UserInfoResponse struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	ProfileImageUrl string `json:"profile_image_url"`
}
