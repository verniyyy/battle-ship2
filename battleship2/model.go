package battleship2

type ID[T any] string

func (t ID[T]) String() string {
	return string(t)
}

type User struct {
	UserID   ID[User] `json:"user_id" xml:"user_id" form:"user_id" query:"user_id"`
	Name     string   `json:"name" xml:"name" form:"name" query:"name"`
	Password string   `json:"password" xml:"password" form:"password" query:"password"`
	Email    string   `json:"email" xml:"email" form:"email" query:"email"`
}

func (t User) ID() string {
	return t.UserID.String()
}

type UserAndSession struct {
	*User
	SessionID ID[UserAndSession] `json:"session_id" xml:"session_id" form:"session_id" query:"session_id"`
	RoomID    ID[Room]           `json:"room_id" xml:"room_id" form:"room_id" query:"room_id"`
}

func (t UserAndSession) ID() string {
	return t.SessionID.String()
}

type Room struct {
	RoomID           ID[Room]        `json:"room_id" xml:"room_id" form:"room_id" query:"room_id"`
	PlayerX          *UserAndSession `json:"player_x" xml:"player_x" form:"player_x" query:"player_x"`
	PlayerY          *UserAndSession `json:"player_y" xml:"player_y" form:"player_y" query:"player_y"`
	IsPlayerXStandBy bool
	IsPlayerYStandBy bool
	Turn             int8
	Event            *Event
}

func (t Room) ID() string {
	return t.RoomID.String()
}

func (r *Room) StandBy(sessionID ID[UserAndSession]) {
	if r.PlayerX.SessionID == sessionID {
		r.IsPlayerXStandBy = true
	}
	r.IsPlayerYStandBy = true
}

type Event struct {
	ActionBy *UserAndSession
	Action   int
	Detail   *Detail
}

type Detail struct{}
