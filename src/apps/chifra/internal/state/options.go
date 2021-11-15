package statePkg

/*-------------------------------------------------------------------------------------------
 * qblocks - fast, easily-accessible, fully-decentralized data from blockchains
 * copyright (c) 2016, 2021 TrueBlocks, LLC (http://trueblocks.io)
 *
 * This program is free software: you may redistribute it and/or modify it under the terms
 * of the GNU General Public License as published by the Free Software Foundation, either
 * version 3 of the License, or (at your option) any later version. This program is
 * distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details. You should have received a copy of the GNU General
 * Public License along with this program. If not, see http://www.gnu.org/licenses/.
 *-------------------------------------------------------------------------------------------*/
/*
 * The file was auto generated with makeClass --gocmds. DO NOT EDIT.
 */

import (
	"net/http"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/cmd/globals"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type StateOptions struct {
	Addrs    []string
	Blocks   []string
	Parts    []string
	Changes  bool
	NoZero   bool
	Call     string
	ProxyFor string
	Globals  globals.GlobalOptionsType
}

func (opts *StateOptions) TestLog() {
	logger.TestLog(len(opts.Addrs) > 0, "Addrs: ", opts.Addrs)
	logger.TestLog(len(opts.Blocks) > 0, "Blocks: ", opts.Blocks)
	logger.TestLog(len(opts.Parts) > 0, "Parts: ", opts.Parts)
	logger.TestLog(opts.Changes, "Changes: ", opts.Changes)
	logger.TestLog(opts.NoZero, "NoZero: ", opts.NoZero)
	logger.TestLog(len(opts.Call) > 0, "Call: ", opts.Call)
	logger.TestLog(len(opts.ProxyFor) > 0, "ProxyFor: ", opts.ProxyFor)
	opts.Globals.TestLog()
}

func (opts *StateOptions) ToDashStr() string {
	options := ""
	for _, t := range opts.Parts {
		options += " --parts " + t
	}
	if opts.Changes {
		options += " --changes"
	}
	if opts.NoZero {
		options += " --no_zero"
	}
	if len(opts.Call) > 0 {
		options += " --call " + opts.Call
	}
	if len(opts.ProxyFor) > 0 {
		options += " --proxy_for " + opts.ProxyFor
	}
	options += " " + strings.Join(opts.Addrs, " ")
	return options
}

func FromRequest(w http.ResponseWriter, r *http.Request) *StateOptions {
	opts := &StateOptions{}
	for key, value := range r.URL.Query() {
		switch key {
		case "addrs":
			opts.Addrs = append(opts.Addrs, value...)
		case "blocks":
			opts.Blocks = append(opts.Blocks, value...)
		case "parts":
			opts.Parts = append(opts.Parts, value...)
		case "changes":
			opts.Changes = true
		case "no_zero":
			opts.NoZero = true
		case "call":
			opts.Call = value[0]
		case "proxy_for":
			opts.ProxyFor = value[0]
		}
	}
	opts.Globals = *globals.FromRequest(w, r)

	return opts
}