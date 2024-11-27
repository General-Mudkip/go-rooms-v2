package models

type Room struct {
	RoomID      int    `json:"room_id" gorm:"primaryKey"`
	RoomName    string `json:"room_name"`
	RoomMembers string `json:"room_members"` // String delimited by commas
	RoomPin     string `json:"room_pin"`
}
