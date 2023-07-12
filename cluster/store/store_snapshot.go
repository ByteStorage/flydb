package store

import (
	"errors"
	"github.com/ByteStorage/FlyDB/config"
	"github.com/hashicorp/raft"
	"os"
	"path"
)

func newSnapshotStore() (raft.SnapshotStore, error) {
	conf := config.DefaultOptions
	snStore, err := getSnapShotStore("memory", conf)
	if err != nil {
		return nil, err
	}
	return snStore, nil
}

func getSnapShotStore(snapshotStore string, conf config.Options) (raft.SnapshotStore, error) {
	snapshotStoreDir := path.Join(conf.DirPath, "snapshot")

	switch snapshotStore {
	case "memory":
		return raft.NewInmemSnapshotStore(), nil
	case "discard":
		return raft.NewDiscardSnapshotStore(), nil
	case "file":
		return raft.NewFileSnapshotStore(snapshotStoreDir, 5, os.Stderr) //todo get the retain param from config
	default:
		return nil, errors.New("the datastore does not exist")
	}
}
