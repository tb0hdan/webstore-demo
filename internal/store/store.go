package store

import (
	"webstore-demo/internal/store/memory"
	"webstore-demo/pkg/types"
)

const (
	Memory = "memory"
)

func New(storeType string) types.Store {
	switch storeType {
	case Memory:
		return memory.NewStore()
	default:
		return nil
	}
}
