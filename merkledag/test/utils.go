package mdutils

import (
	bsrv "github.com/ipfs/go-ipfs/blockservice"
	"github.com/ipfs/go-ipfs/exchange/offline"
	dag "github.com/ipfs/go-ipfs/merkledag"

	ds "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	ipld "github.com/ipfs/go-ipld-format"
)

// Mock returns a new thread-safe, mock DAGService.
func Mock() ipld.DAGService {
	return dag.NewDAGService(Bserv())
}

// Bserv returns a new, thread-safe, mock BlockService.
func Bserv() bsrv.BlockService {
	bstore := blockstore.NewBlockstore(dssync.MutexWrap(ds.NewMapDatastore()))
	return bsrv.New(bstore, offline.Exchange(bstore))
}
