package store

import (
	"webstore-demo/internal/store/memory"
	"webstore-demo/pkg/types"
)

const (
	StoreTypeMemory = "memory"
)

func New(storeType string) types.Store {
	switch storeType {
	case StoreTypeMemory:
		return memory.NewStore()
	default:
		return nil
	}
}
