package graphql

import "context"

type StreamField[T any] interface {
	GetContext() context.Context
	GetField() *T
}
