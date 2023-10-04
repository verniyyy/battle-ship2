package battleship2

import "context"

type Adapter[T, Conn any] struct {
	service Service[T, Conn]
	conn    Conn
}

func NewAdapter[T, Conn any](c Conn, s Service[T, Conn]) *Adapter[T, Conn] {
	return &Adapter[T, Conn]{
		service: s,
		conn:    c,
	}
}

func (ad Adapter[T, Conn]) Get(ctx context.Context, id ID[T]) (*T, error) {
	return ad.service.Get(ctx, ad.conn, id)
}

func (ad Adapter[T, Conn]) Create(ctx context.Context, v *T) (*T, error) {
	return ad.service.Create(ctx, ad.conn, v)
}

func (ad Adapter[T, Conn]) Update(ctx context.Context, v *T) (*T, error) {
	return ad.service.Update(ctx, ad.conn, v)
}

func (ad Adapter[T, Conn]) Delete(ctx context.Context, id ID[T]) error {
	return ad.service.Delete(ctx, ad.conn, id)
}

func NewMatchingQueueAdapter() *MatchingQueueAdapter {
	return &MatchingQueueAdapter{}
}

type MatchingQueueAdapter struct {
	service MatchingQueueService
	conn    any
}

func (ad MatchingQueueAdapter) Push(ctx context.Context, userAndSession *UserAndSession) error {
	return ad.service.Push(ctx, ad.conn, userAndSession)
}

func (ad MatchingQueueAdapter) Pop(ctx context.Context) (*UserAndSession, error) {
	return ad.service.Pop(ctx, ad.conn)
}

func (ad MatchingQueueAdapter) Len() int {
	return ad.service.Len(ad.conn)
}
