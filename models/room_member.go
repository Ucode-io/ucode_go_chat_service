package models

type RoomMember struct {
	Id        string `json:"id"`
	RoomId    string `json:"room_id"`
	RowId     string `json:"row_id"`
	ToName    string `json:"to_name"`
	ToRowId   any    `json:"to_row_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateRoomMember struct {
	RoomId  string `json:"room_id"`
	RowId   string `json:"row_id"`
	ToName  string `json:"to_name"`
	ToRowId any    `json:"to_row_id"`
}

type UpdateLastReadAtReq struct {
	RoomId string `json:"room_id"`
	RowId  string `json:"row_id"`
}
