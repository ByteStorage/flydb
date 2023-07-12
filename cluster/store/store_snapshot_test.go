package store

import (
	"github.com/ByteStorage/FlyDB/config"
	"github.com/hashicorp/raft"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRegisterSnapshotStore(t *testing.T) {
	type test struct {
		name        string
		datastore   string
		want        raft.SnapshotStore
		expectError bool
	}
	tests := []test{
		{
			name:        "memory",
			datastore:   "memory",
			want:        &raft.InmemSnapshotStore{},
			expectError: false,
		},
		{
			name:        "correctly get file",
			datastore:   "file",
			want:        &raft.FileSnapshotStore{},
			expectError: false,
		},
		{
			name:        "correctly get discard",
			datastore:   "discard",
			want:        &raft.DiscardSnapshotStore{},
			expectError: false,
		},
		{
			name:        "a wrong datastore",
			datastore:   "wrong",
			want:        nil,
			expectError: true,
		},
	}
	for _, tc := range tests {
		conf := config.DefaultOptions
		snst, err := getSnapShotStore(tc.datastore, conf)
		if tc.expectError {
			assert.NotNil(t, err)
		} else {
			assert.IsType(t, tc.want, snst)
		}
	}
}

func TestStoreSnapshot_newSnapshot(t *testing.T) {
	sn, err := newSnapshotStore()
	assert.NoError(t, err)

	assert.True(t, reflect.TypeOf(sn).Implements(reflect.TypeOf((*raft.SnapshotStore)(nil)).Elem()))
}
