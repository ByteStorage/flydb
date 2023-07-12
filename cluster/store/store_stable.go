package store

import (
	"github.com/ByteStorage/FlyDB/config"
	"github.com/hashicorp/raft"
)

func newStableLog() (raft.StableStore, error) {
	_ = Init()

	// Get the "memory" DataStoreFactory from the map
	return getDataStore("memory", config.Options{})
}
