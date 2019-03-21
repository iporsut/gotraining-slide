package todo

import "context"

type Service interface {
	Add(context.Context, *TodoList) (*TodoList, error)
}
