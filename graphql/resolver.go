//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"github.com/scorpionknifes/gqlopenhab/mongodb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver struct
type Resolver struct {
	DeviceRepo mongodb.DeviceRepo
	RoomRepo   mongodb.RoomRepo
	UserRepo   mongodb.UserRepo
}
