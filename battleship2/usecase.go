package battleship2

import (
	"context"

	"github.com/verniyyy/battle-ship2/lib"
)

// EmptyOutput ...
type EmptyOutput struct{}

// UserUsecase ...
type UserUsecase struct {
	user          Repository[User]
	queue         MatchingQueue
	randGenerator lib.RandGenerator
}

// NewUserUsecase ...
func NewUserUsecase() UserUsecase {
	return UserUsecase{
		user:          NewAdapter(NewDataStoreClient[User](), UserService{}),
		queue:         NewMatchingQueueAdapter(),
		randGenerator: lib.NewRandGenerator(),
	}
}

// CreateUser ...
func (u UserUsecase) CreateUser(ctx context.Context, in *CreateUserInput) (*User, error) {
	user := &User{
		UserID:   ID[User](u.randGenerator.ULID()),
		Name:     in.Name,
		Password: lib.Hash(in.Password),
		Email:    in.Email,
	}
	return u.user.Create(ctx, user)
}

// DescribeUser ...
func (u UserUsecase) DescribeUser(ctx context.Context, in *DescribeUserInput) (*User, error) {
	return u.user.Get(ctx, ID[User](in.UserID))
}

// InQueue ...
func (u UserUsecase) InQueue(ctx context.Context, in *InQueueInput) (*UserAndSession, error) {
	user, err := u.user.Get(ctx, ID[User](in.UserID))
	if err != nil {
		return nil, err
	}

	userAndSession := &UserAndSession{
		User:      user,
		SessionID: ID[UserAndSession](u.randGenerator.ULID()),
	}
	err = u.queue.Push(ctx, userAndSession)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// RoomUsecase ...
type RoomUsecase struct {
	queue          MatchingQueue
	userAndSession Repository[UserAndSession]
	room           Repository[Room]
	randGenerator  lib.RandGenerator
}

// NewRoomUsecase ...
func NewRoomUsecase() RoomUsecase {
	return RoomUsecase{
		queue:          NewMatchingQueueAdapter(),
		userAndSession: NewAdapter(NewDataStoreClient[UserAndSession](), UserAndSessionService{}),
		room:           NewAdapter(NewDataStoreClient[Room](), RoomService{}),
		randGenerator:  lib.NewRandGenerator(),
	}
}

func (u RoomUsecase) MatchMaking(ctx context.Context, in *MatchMakingInput) error {
	if u.queue.Len() < 2 {
		return nil
	}

	playerX, err := u.queue.Pop(ctx)
	if err != nil {
		return err
	}

	playerY, err := u.queue.Pop(ctx)
	if err != nil {
		return err
	}

	room := &Room{
		RoomID:  ID[Room](u.randGenerator.ULID()),
		PlayerX: playerX,
		PlayerY: playerY,
	}
	room, err = u.room.Create(ctx, room)
	if err != nil {
		return err
	}

	playerX.RoomID = room.RoomID
	u.userAndSession.Create(ctx, playerX)
	if err != nil {
		return err
	}

	playerY.RoomID = room.RoomID
	u.userAndSession.Create(ctx, playerY)
	if err != nil {
		return err
	}

	return nil
}

// IsMatched ...
func (u RoomUsecase) IsMatched(ctx context.Context, in *IsMatchedInput) (*Room, error) {
	userAndSession, err := u.userAndSession.Get(ctx, ID[UserAndSession](in.SessionID))
	if err != nil {
		return nil, err
	}

	room, err := u.room.Get(ctx, userAndSession.RoomID)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// ConnectRoom  ...
func (u RoomUsecase) ConnectRoom(ctx context.Context, in *ConnectRoomInput) (*EmptyOutput, error) {
	room, err := u.room.Get(ctx, ID[Room](in.RoomID))
	if err != nil {
		return nil, err
	}

	room.StandBy(ID[UserAndSession](in.SessionID))
	_, err = u.room.Update(ctx, room)
	if err != nil {
		return nil, err
	}

	return &EmptyOutput{}, nil
}
