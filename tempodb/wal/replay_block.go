package wal

import (
	"os"

	"github.com/grafana/tempo/tempodb/encoding/common"
	v0 "github.com/grafana/tempo/tempodb/encoding/v0"
)

type ReplayBlock struct {
	block
}

func (r *ReplayBlock) Iterator() (common.Iterator, error) {
	name := r.fullFilename()
	f, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return v0.NewIterator(f), nil
}

func (r *ReplayBlock) TenantID() string {
	return r.meta.TenantID
}

func (r *ReplayBlock) Clear() error {
	if r.readFile != nil {
		_ = r.readFile.Close()
	}

	name := r.fullFilename()
	return os.Remove(name)
}
