package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state/snapshot"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/web3ext"
	"github.com/ethereum/go-ethereum/log"
	"math"
	"time"
)

const DebugJs = `
web3._extend({
	property: 'debug',
	methods: [
		new web3._extend.Method({
			name: 'accountRange',
			call: 'debug_accountRange',
			params: 6,
			inputFormatter: [web3._extend.formatters.inputDefaultBlockNumberFormatter, null, null, null, null, null],
		}),
		new web3._extend.Method({
			name: 'printBlock',
			call: 'debug_printBlock',
			params: 1,
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'getRawHeader',
			call: 'debug_getRawHeader',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getRawBlock',
			call: 'debug_getRawBlock',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getRawReceipts',
			call: 'debug_getRawReceipts',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getRawTransaction',
			call: 'debug_getRawTransaction',
			params: 1
		}),
		new web3._extend.Method({
			name: 'setHead',
			call: 'debug_setHead',
			params: 1
		}),
		new web3._extend.Method({
			name: 'seedHash',
			call: 'debug_seedHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'dumpBlock',
			call: 'debug_dumpBlock',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter]
		}),
		new web3._extend.Method({
			name: 'chaindbProperty',
			call: 'debug_chaindbProperty',
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'chaindbCompact',
			call: 'debug_chaindbCompact',
		}),
		new web3._extend.Method({
			name: 'verbosity',
			call: 'debug_verbosity',
			params: 1
		}),
		new web3._extend.Method({
			name: 'vmodule',
			call: 'debug_vmodule',
			params: 1
		}),
		new web3._extend.Method({
			name: 'backtraceAt',
			call: 'debug_backtraceAt',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'stacks',
			call: 'debug_stacks',
			params: 1,
			inputFormatter: [null],
			outputFormatter: console.log
		}),
		new web3._extend.Method({
			name: 'freeOSMemory',
			call: 'debug_freeOSMemory',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'setGCPercent',
			call: 'debug_setGCPercent',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'memStats',
			call: 'debug_memStats',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'gcStats',
			call: 'debug_gcStats',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'cpuProfile',
			call: 'debug_cpuProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startCPUProfile',
			call: 'debug_startCPUProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'stopCPUProfile',
			call: 'debug_stopCPUProfile',
			params: 0
		}),
		new web3._extend.Method({
			name: 'goTrace',
			call: 'debug_goTrace',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startGoTrace',
			call: 'debug_startGoTrace',
			params: 1
		}),
		new web3._extend.Method({
			name: 'stopGoTrace',
			call: 'debug_stopGoTrace',
			params: 0
		}),
		new web3._extend.Method({
			name: 'blockProfile',
			call: 'debug_blockProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'setBlockProfileRate',
			call: 'debug_setBlockProfileRate',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeBlockProfile',
			call: 'debug_writeBlockProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'mutexProfile',
			call: 'debug_mutexProfile',
			params: 2
		}),
		new web3._extend.Method({
			name: 'setMutexProfileFraction',
			call: 'debug_setMutexProfileFraction',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeMutexProfile',
			call: 'debug_writeMutexProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'writeMemProfile',
			call: 'debug_writeMemProfile',
			params: 1
		}),
		new web3._extend.Method({
			name: 'traceBlock',
			call: 'debug_traceBlock',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockFromFile',
			call: 'debug_traceBlockFromFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBadBlock',
			call: 'debug_traceBadBlock',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'standardTraceBadBlockToFile',
			call: 'debug_standardTraceBadBlockToFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'intermediateRoots',
			call: 'debug_intermediateRoots',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'standardTraceBlockToFile',
			call: 'debug_standardTraceBlockToFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockByNumber',
			call: 'debug_traceBlockByNumber',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter, null]
		}),
		new web3._extend.Method({
			name: 'traceBlockByHash',
			call: 'debug_traceBlockByHash',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceTransaction',
			call: 'debug_traceTransaction',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'traceCall',
			call: 'debug_traceCall',
			params: 3,
			inputFormatter: [null, null, null]
		}),
		new web3._extend.Method({
			name: 'preimage',
			call: 'debug_preimage',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getBadBlocks',
			call: 'debug_getBadBlocks',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'storageRangeAt',
			call: 'debug_storageRangeAt',
			params: 5,
		}),
		new web3._extend.Method({
			name: 'getModifiedAccountsByNumber',
			call: 'debug_getModifiedAccountsByNumber',
			params: 2,
			inputFormatter: [null, null],
		}),
		new web3._extend.Method({
			name: 'getModifiedAccountsByHash',
			call: 'debug_getModifiedAccountsByHash',
			params: 2,
			inputFormatter:[null, null],
		}),
		new web3._extend.Method({
			name: 'freezeClient',
			call: 'debug_freezeClient',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'getAccessibleState',
			call: 'debug_getAccessibleState',
			params: 2,
			inputFormatter:[web3._extend.formatters.inputBlockNumberFormatter, web3._extend.formatters.inputBlockNumberFormatter],
		}),
		new web3._extend.Method({
			name: 'dbGet',
			call: 'debug_dbGet',
			params: 1
		}),
		new web3._extend.Method({
			name: 'dbAncient',
			call: 'debug_dbAncient',
			params: 2
		}),
		new web3._extend.Method({
			name: 'dbAncients',
			call: 'debug_dbAncients',
			params: 0
		}),
		new web3._extend.Method({
			name: 'setTrieFlushInterval',
			call: 'debug_setTrieFlushInterval',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getTrieFlushInterval',
			call: 'debug_getTrieFlushInterval',
			params: 0
		}),

		new web3._extend.Method({
			name: 'snapshotStats',
			call: 'debug_snapshotStats',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getAggregatorMemoryLimit',
			call: 'debug_getAggregatorMemoryLimit',
			params: 0,
		}),
		new web3._extend.Method({
			name: 'setAggregatorMemoryLimit',
			call: 'debug_setAggregatorMemoryLimit',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'getPathDBLayers',
			call: 'debug_getPathDBLayers',
			params: 0,
		}),
	],
	properties: []
});
`

func init() {
	web3ext.Modules["debug"] = DebugJs
}

type SnapshotRet struct {
	Index     int         `json:"index"`
	Number    uint64      `json:"number"`
	Root      common.Hash `json:"root"`
	BlockHash common.Hash `json:"blockhash"`
}

type SnapshotStatsRet struct {
	Total     int           `json:"total"`
	Header    common.Hash   `json:"header"`
	DiskRoot  common.Hash   `json:"diskroot"`
	Snapshots []SnapshotRet `json:"snapshots"`
}

func (api *DebugAPI) SnapshotStats(r *common.Hash) (SnapshotStatsRet, error) {
	snaptree := api.eth.blockchain.Snapshots()
	header := api.eth.APIBackend.CurrentBlock()
	var root = header.Root
	if r != nil {
		root = *r
	}
	snaps := snaptree.Snapshots(root, math.MaxInt, false)
	log.Info("Snapshots info", "blockRoot", header.Root, "total", len(snaps), "diskRoot", snaptree.DiskRoot().String())
	head := header
	getHead := func(snap snapshot.Snapshot) *types.Header {
		timeout := time.After(time.Minute)
		for head != nil {
			select {
			case <-timeout:
				return nil
			default:
			}
			if snap.Root() == head.Root {
				return head
			}
			if head.Number.Uint64() <= 0 {
				return nil
			}
			head = api.eth.blockchain.GetHeader(head.ParentHash, head.Number.Uint64()-1)
		}
		return nil
	}
	sr := SnapshotStatsRet{
		Total:     len(snaps),
		Header:    header.Root,
		DiskRoot:  snaptree.DiskRoot(),
		Snapshots: []SnapshotRet{},
	}

	for i, snap := range snaps {
		shead := getHead(snap)
		if shead != nil {
			sr.Snapshots = append(sr.Snapshots, SnapshotRet{
				Index:     i,
				Number:    shead.Number.Uint64(),
				Root:      snap.Root(),
				BlockHash: shead.Hash(),
			})
			log.Info("snapshot", "index", i, "root", snap.Root().String(), "blockHash", shead.Hash().String(), "number", shead.Number.String())
		} else {
			sr.Snapshots = append(sr.Snapshots, SnapshotRet{
				Index: i,
				Root:  snap.Root(),
			})
			log.Info("snapshot", "index", i, "root", snap.Root().String())
		}
	}
	return sr, nil
}

func (api *DebugAPI) GetAggregatorMemoryLimit() (uint64, error) {
	return snapshot.GetAggregatorMemoryLimit(), nil
}

func (api *DebugAPI) SetAggregatorMemoryLimit(limit uint64) error {
	snapshot.SetAggregatorMemoryLimit(limit)
	return nil
}

func (api *DebugAPI) GetPathDBLayers() (interface{}, error) {
	trdb := api.eth.blockchain.TrieDB()
	header := api.eth.APIBackend.CurrentBlock()
	var root = header.Root
	return trdb.Debug(root), nil
}
