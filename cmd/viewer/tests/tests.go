// Copyright (c) 2022 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tests serves a list of tests for tailscale.com/cmd/viewer.
package tests

import (
	"fmt"

	"inet.af/netaddr"
)

//go:generate go run tailscale.com/cmd/viewer --type=Value,Struct,Map,Slices

type Value struct {
	Int int
	Pfx netaddr.IPPrefix
}

type Map struct {
	M map[string]int
}

type Struct struct {
	Value *Value
	Int   *int
}

func (v *Struct) String() string { return fmt.Sprintf("%v", v.Int) }

func (v *Struct) Equal(v2 *Struct) bool {
	return v.Value == v2.Value
}

type Slices struct {
	Values         []Value
	ValuePointers  []*Value
	StructPointers []*Struct
	Structs        []Struct
	Ints           []*int

	Slice    []string
	Prefixes []netaddr.IPPrefix
	Data     []byte
}
