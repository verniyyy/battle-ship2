package battleship2

import "context"

type Service[T, Conn any] interface {
	Get(context.Context, Conn, ID[T]) (*T, error)
	Create(context.Context, Conn, *T) (*T, error)
	Update(context.Context, Conn, *T) (*T, error)
	Delete(context.Context, Conn, ID[T]) error
}

type UserService struct{}

func (UserService) Get(ctx context.Context, conn DataStoreClient[User], id ID[User]) (*User, error) {
	user, err := conn.Get(id.String())
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (UserService) Create(ctx context.Context, conn DataStoreClient[User], user *User) (*User, error) {
	newUser, err := conn.Create(*user)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (UserService) Update(ctx context.Context, conn DataStoreClient[User], user *User) (*User, error) {
	newUser, err := conn.Update(*user)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (UserService) Delete(ctx context.Context, conn DataStoreClient[User], id ID[User]) error {
	return conn.Delete(id.String())
}

type UserAndSessionService struct{}

func (UserAndSessionService) Get(ctx context.Context, conn DataStoreClient[UserAndSession], id ID[UserAndSession]) (*UserAndSession, error) {
	userAndSession, err := conn.Get(id.String())
	if err != nil {
		return nil, err
	}
	return &userAndSession, nil
}

func (UserAndSessionService) Create(ctx context.Context, conn DataStoreClient[UserAndSession], userAndSession *UserAndSession) (*UserAndSession, error) {
	newUserAndSession, err := conn.Create(*userAndSession)
	if err != nil {
		return nil, err
	}

	return &newUserAndSession, nil
}

func (UserAndSessionService) Update(ctx context.Context, conn DataStoreClient[UserAndSession], userAndSession *UserAndSession) (*UserAndSession, error) {
	newUserAndSession, err := conn.Update(*userAndSession)
	if err != nil {
		return nil, err
	}

	return &newUserAndSession, nil
}

func (UserAndSessionService) Delete(ctx context.Context, conn DataStoreClient[UserAndSession], id ID[UserAndSession]) error {
	return conn.Delete(id.String())
}

type RoomService struct{}

func (RoomService) Get(ctx context.Context, conn DataStoreClient[Room], id ID[Room]) (*Room, error) {
	room, err := conn.Get(id.String())
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (RoomService) Create(ctx context.Context, conn DataStoreClient[Room], room *Room) (*Room, error) {
	newRoom, err := conn.Create(*room)
	if err != nil {
		return nil, err
	}

	return &newRoom, nil
}

func (RoomService) Update(ctx context.Context, conn DataStoreClient[Room], room *Room) (*Room, error) {
	newRoom, err := conn.Update(*room)
	if err != nil {
		return nil, err
	}

	return &newRoom, nil
}

func (RoomService) Delete(ctx context.Context, conn DataStoreClient[Room], id ID[Room]) error {
	return conn.Delete(id.String())
}

type MatchingQueueService struct{}

func (MatchingQueueService) Push(ctx context.Context, conn any, userAndSession *UserAndSession) error {
	Q().Push(userAndSession)
	return nil
}

func (MatchingQueueService) Pop(ctx context.Context, conn any) (*UserAndSession, error) {
	userAndSession := Q().Pop()
	return userAndSession, nil
}

func (MatchingQueueService) Len(conn any) int {
	return Q().Len()
}
