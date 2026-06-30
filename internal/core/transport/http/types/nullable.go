package core_http_types

import (
	"encoding/json"

	domain "github.com/QBL25079/teams/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarashalJSON(b []byte) error {
	n.Set = true

	if string(b) == "null" {
		n.Value = nil
		return nil
	}

	var value T

	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.Value = &value

	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{
		Value: n.Value,
		Set: n.Set,
	}
}