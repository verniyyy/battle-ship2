package battleship2

import "context"

type Repository[T any] interface {
	Get(context.Context, ID[T]) (*T, error)
	Create(context.Context, *T) (*T, error)
	Update(context.Context, *T) (*T, error)
	Delete(context.Context, ID[T]) error
}

type MatchingQueue interface {
	Push(context.Context, *UserAndSession) error
	Pop(context.Context) (*UserAndSession, error)
	Len() int
}
