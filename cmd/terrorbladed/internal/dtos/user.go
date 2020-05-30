package dtos

import "terrorblade/dtos"

type CreateUserRequest struct {
	FullName string
	Username string
	Password string
	Status   string
}

type CreateUserResponse struct {
	Metadata dtos.Metadata
	Data     CreateUserData
}

type CreateUserData struct {
	ID int64
}

//
// GetUser
//

type GetUserResponse struct {
	Metadata dtos.Metadata
	Data     GetUserData
}

type GetUserData struct {
	ID       int64
	FullName string
	Username string
	Password string
	Status   string
}
