package battleship2

import (
	"container/list"
	"fmt"
	"sync"
)

var (
	_data = NewDataStore()
	dMu   sync.Mutex
)

func D() DataStore {
	dMu.Lock()
	d := _data
	dMu.Unlock()
	return d
}

type DataStore interface {
	Create(k, v any)
	Read(k any) (any, error)
	Update(k, v any)
	Delete(k any)
}

func NewDataStore() DataStore {
	return &dataStore{}
}

type dataStore struct {
	m sync.Map
}

func (d *dataStore) Create(k, v any) {
	d.m.Store(k, v)
}

func (d *dataStore) Read(k any) (any, error) {
	v, ok := d.m.Load(k)
	if !ok {
		return nil, fmt.Errorf("error not found by key=%v", k)
	}
	return v, nil
}

func (d *dataStore) Update(k, v any) {
	d.m.Delete(k)
	d.Create(k, v)
}

func (d *dataStore) Delete(k any) {
	d.m.Delete(k)
}

var (
	_queue = NewQueue[UserAndSession]()
	qMu    sync.Mutex
)

type DataStoreClient[T Model] interface {
	Create(T) (T, error)
	Get(id string) (T, error)
	Update(T) (T, error)
	Delete(id string) error
}

func NewDataStoreClient[T Model]() DataStoreClient[T] {
	return dataStoreClient[T]{}
}

type dataStoreClient[T Model] struct{}

type Model interface {
	ID() string
}

func (dataStoreClient[T]) Create(v T) (T, error) {
	d, _ := D().Read(v.ID())
	if d != nil {
		var v T
		return v, fmt.Errorf("create duplicate key error")
	}
	D().Create(v.ID(), v)
	return v, nil
}
func (dataStoreClient[T]) Get(id string) (T, error) {
	d, err := D().Read(id)
	if err != nil {
		var v T
		return v, err
	}

	v, ok := d.(T)
	if !ok {
		var v T
		return v, fmt.Errorf("type assertion error. do not %T to %T", d, v)
	}

	return v, nil
}

func (c dataStoreClient[T]) Update(v T) (T, error) {
	if err := c.Delete(v.ID()); err != nil {
		var v T
		return v, err
	}

	v, err := c.Create(v)
	if err != nil {
		var v T
		return v, err
	}

	return v, nil
}

func (dataStoreClient[V]) Delete(id string) error {
	_, err := D().Read(id)
	if err != nil {
		return err
	}

	D().Delete(id)
	return nil
}

func Q() Queue[UserAndSession] {
	qMu.Lock()
	q := _queue
	qMu.Unlock()
	return q
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{l: list.New()}
}

type Queue[T any] interface {
	Push(*T)
	Pop() *T
	Len() int
}

type queue[T any] struct {
	l *list.List
}

func (q *queue[T]) Push(v *T) {
	q.l.PushBack(v)
}

func (q *queue[T]) Pop() *T {
	v := q.l.Front()
	q.l.Remove(v)
	return v.Value.(*T)
}

func (q *queue[T]) Len() int {
	return q.l.Len()
}
