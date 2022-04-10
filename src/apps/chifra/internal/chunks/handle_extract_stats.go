// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package chunksPkg

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/cache"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/index"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type ChunkStats struct {
	Start         uint64 `json:"start"`
	End           uint64 `json:"end"`
	NAddrs        uint64 `json:"nAddrs"`
	NApps         uint64 `json:"nApps"`
	NBlocks       uint64 `json:"nBlocks"`
	NBlooms       uint64 `json:"nBlooms"`
	AddrsPerBlock string `json:"addrsPerBlock"`
	AppsPerBlock  string `json:"appsPerBlock"`
	AppsPerAddr   string `json:"appsPerAddr"`
	RecWid        uint64 `json:"recWid"`
	BloomSz       int64  `json:"bloomSz"`
	ChunkSz       int64  `json:"chunkSz"`
	Ratio         string `json:"ratio"`
}

func NewChunkStats(path string) ChunkStats {
	chunk, err := index.NewChunk(path)
	if err != nil {
		fmt.Println(err)
	}
	var ret ChunkStats
	ret.Start = chunk.Range.First
	ret.End = chunk.Range.Last
	ret.NAddrs = uint64(chunk.Data.Header.AddressCount)
	ret.NApps = uint64(chunk.Data.Header.AppearanceCount)
	ret.NBlocks = ret.End - ret.Start + 1
	ret.NBlooms = uint64(chunk.Bloom.Count)
	if ret.NBlocks > 0 {
		ret.AddrsPerBlock = fmt.Sprintf("%.3f", float64(ret.NAddrs)/float64(ret.NBlocks))
		ret.AppsPerBlock = fmt.Sprintf("%.3f", float64(ret.NApps)/float64(ret.NBlocks))
	}
	if ret.NAddrs > 0 {
		ret.AppsPerAddr = fmt.Sprintf("%.3f", float64(ret.NApps)/float64(ret.NAddrs))
	}
	ret.RecWid = 4 + index.BLOOM_WIDTH_IN_BYTES
	ret.BloomSz = file.FileSize(path)
	p := strings.Replace(strings.Replace(path, ".bloom", ".bin", -1), "blooms", "finalized", -1)
	ret.ChunkSz = file.FileSize(p)
	ret.Ratio = fmt.Sprintf("%.3f", float64(ret.ChunkSz)/float64(ret.BloomSz))
	return ret
}

func (opts *ChunksOptions) showStats(path string, first bool) {
	if opts.Globals.TestMode {
		r, _ := cache.RangeFromFilename(path)
		if r.First > 2000000 && r.First < 3000000 {
			return
		}
		if r.First > 4000000 {
			return
		}
	}

	var results []ChunkStats
	results = append(results, NewChunkStats(path))

	// TODO: Fix export without arrays
	if opts.Globals.ApiMode {
		opts.Globals.Respond2(opts.Globals.Writer, http.StatusOK, results, !first)

	} else {
		err := opts.Globals.Output2(os.Stdout, opts.Globals.Format, results, !first)
		if err != nil {
			logger.Log(logger.Error, err)
		}
	}
}