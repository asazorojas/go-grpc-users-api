package data

import (
	userpb "github.com/asazorojas/go-grpc-users-api/common/protos/v1/user"
)

var Users = []*userpb.UserMessage{
	{
		UserId:      "1",
		Name:        "Alejandra",
		PhoneNumber: "01012341234",
		Age:         22,
	},
	{
		UserId:      "2",
		Name:        "Elisa",
		PhoneNumber: "01012341234",
		Age:         22,
	},
	{
		UserId:      "3",
		Name:        "Rodrigo",
		PhoneNumber: "01012341234",
		Age:         22,
	},
	{
		UserId:      "4",
		Name:        "Sebastián",
		PhoneNumber: "01012341234",
		Age:         22,
	},
	{
		UserId:      "5",
		Name:        "Charlie",
		PhoneNumber: "01012341234",
		Age:         22,
	},
}
