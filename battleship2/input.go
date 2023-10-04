package battleship2

import (
	"fmt"

	"github.com/labstack/echo"
)

func NewCreateUserInput(c echo.Context) (*CreateUserInput, error) {
	in := new(CreateUserInput)
	if err := c.Bind(in); err != nil {
		return nil, err
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}

	return in, nil
}

type CreateUserInput struct {
	*User
}

func (i CreateUserInput) Validate() error {
	if i.Name == "" || i.Password == "" || i.Email == "" {
		return fmt.Errorf("invalid name or password or email")
	}
	return nil
}

func NewDescribeUserInput(c echo.Context) (*DescribeUserInput, error) {
	in := new(DescribeUserInput)
	if err := c.Bind(in); err != nil {
		return nil, err
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}

	return in, nil
}

type DescribeUserInput struct {
	UserID string `json:"user_id" xml:"user_id" form:"user_id" query:"user_id"`
}

func (i DescribeUserInput) Validate() error {
	if i.UserID == "" {
		return fmt.Errorf("invalid user_id")
	}
	return nil
}

func NewInQueueInput(c echo.Context) (*InQueueInput, error) {
	in := new(InQueueInput)
	if err := c.Bind(in); err != nil {
		return nil, err
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}

	return in, nil
}

type InQueueInput struct {
	UserID string `json:"user_id" xml:"user_id" form:"user_id" query:"user_id"`
}

func (i InQueueInput) Validate() error {
	if i.UserID == "" {
		return fmt.Errorf("invalid user_id")
	}
	return nil
}

type MatchMakingInput struct{}

func NewIsMatchedInput(c echo.Context) (*IsMatchedInput, error) {
	in := new(IsMatchedInput)
	if err := c.Bind(in); err != nil {
		return nil, err
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}

	return in, nil
}

type IsMatchedInput struct {
	SessionID string `json:"session_id" xml:"session_id" form:"session_id" query:"session_id"`
}

func (i IsMatchedInput) Validate() error {
	if i.SessionID == "" {
		return fmt.Errorf("invalid session_id")
	}
	return nil
}

func NewConnectRoomInput(c echo.Context) (*ConnectRoomInput, error) {
	in := new(ConnectRoomInput)
	if err := c.Bind(in); err != nil {
		return nil, err
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}

	return in, nil
}

type ConnectRoomInput struct {
	RoomID    string `json:"room_id" xml:"room_id" form:"room_id" query:"room_id"`
	SessionID string `json:"session_id" xml:"session_id" form:"session_id" query:"session_id"`
}

func (i ConnectRoomInput) Validate() error {
	if i.RoomID == "" || i.SessionID == "" {
		return fmt.Errorf("invalid room_id or session_id")
	}
	return nil
}
