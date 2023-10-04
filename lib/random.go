package lib

import ulidv2 "github.com/oklog/ulid/v2"

type RandGenerator interface {
	ULID() string
}

func NewRandGenerator() RandGenerator {
	return new(randGenerator)
}

type randGenerator struct{}

func (randGenerator) ULID() string {
	return ulidv2.Make().String()
}

func Hash(s string) string {
	// TODO
	return s
}
