package models

import "time"

// Device struct for one device
type Device struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	RoomID       string    `json:"roomID" bson:"room_id,omitempty"`
	Name         string    `json:"name" bson:"name,omitempty"`
	Model        string    `json:"model" bson:"model,omitempty"`
	MacAddress   string    `json:"macAddress" bson:"mac_address,omitempty"`
	Memo         string    `json:"memo" bson:"memo,omitempty"`
	SerialNumber string    `json:"serialNumber" bson:"serial_number,omitempty"`
	Status       int       `json:"status" bson:"status,omitempty"`
	Type         int       `json:"type" bson:"type,omitempty"`
	CreatedDate  time.Time `json:"createdDate" bson:"created_date,omitempty"`
	LastModified time.Time `json:"lastModified" bson:"last_modified,omitempty"`
}

// Update convert DeviceUpdate to struct
func (d *Device) Update(input DeviceUpdate) {
	if input.Name != nil {
		d.Name = *input.Name
	}
	if input.Model != nil {
		d.Model = *input.Model
	}
	if input.MacAddress != nil {
		d.MacAddress = *input.MacAddress
	}
	if input.Memo != nil {
		d.Memo = *input.Memo
	}
	if input.SerialNumber != nil {
		d.SerialNumber = *input.SerialNumber
	}
	if input.Status != nil {
		d.Status = *input.Status
	}
	if input.Type != nil {
		d.Type = *input.Type
	}
}
