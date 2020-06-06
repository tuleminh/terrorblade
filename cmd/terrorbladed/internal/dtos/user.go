package dtos

import _commonDTOs "terrorblade/dtos"

//
// CreateUser
//

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   string `json:"status"`
}

type CreateUserResponse struct {
	Metadata _commonDTOs.Metadata `json:"_metadata"`
	Data     CreateUserData       `json:"data"`
}

type CreateUserData struct {
	ID int64 `json:"id"`
}

//
// GetUser
//

type GetUserResponse struct {
	Metadata _commonDTOs.Metadata `json:"_metadata"`
	Data     GetUserData          `json:"data"`
}

type GetUserData struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   string `json:"status"`
}
