package battleship2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func CreateUser(c echo.Context) error {
	in, err := NewCreateUserInput(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	usecase := NewUserUsecase()
	out, err := usecase.CreateUser(c.Request().Context(), in)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	return c.JSON(http.StatusCreated, out)
}

func DescribeUser(c echo.Context) error {
	in, err := NewDescribeUserInput(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	usecase := NewUserUsecase()
	out, err := usecase.DescribeUser(c.Request().Context(), in)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	return c.JSON(http.StatusOK, out)
}

func InQueue(c echo.Context) error {
	in, err := NewInQueueInput(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	usecase := NewUserUsecase()
	out, err := usecase.InQueue(c.Request().Context(), in)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	return c.JSON(http.StatusCreated, out)
}

func MatchMaking(l echo.Logger) {
	usecase := NewRoomUsecase()
	for {
		err := usecase.MatchMaking(context.Background(), nil)
		if err != nil {
			l.Error(err)
		}
	}
}

func IsMatched(c echo.Context) error {
	in, err := NewIsMatchedInput(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	usecase := NewRoomUsecase()
	out, err := usecase.IsMatched(c.Request().Context(), in)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	return c.JSON(http.StatusCreated, out)
}

func ConnectRoom(c echo.Context) error {
	in, err := NewConnectRoomInput(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	usecase := NewRoomUsecase()
	out, err := usecase.ConnectRoom(c.Request().Context(), in)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	_ = out

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		// 初回のメッセージを送信
		err := websocket.Message.Send(ws, "Server: Hello, Client!")
		if err != nil {
			c.Logger().Error(err)
		}

		for {
			// Client からのメッセージを読み込む
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}

			// Client からのメッセージを元に返すメッセージを作成し送信する
			err := websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
