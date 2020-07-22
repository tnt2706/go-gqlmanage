package graphql

import (
	"context"

	"github.com/scorpionknifes/gqlopenhab/models"
)

// Room returns RoomResolver implementaion.
func (r *Resolver) Room() RoomResolver { return &roomResolver{r} }

type roomResolver struct{ *Resolver }

func (r *roomResolver) Devices(ctx context.Context, obj *models.Room) ([]*models.Device, error) {
	return r.DeviceRepo.GetDevicesByDeviceID(obj.ID)
}