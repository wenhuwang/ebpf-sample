// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || loong64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadTcPing returns the embedded CollectionSpec for tcPing.
func loadTcPing() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_TcPingBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load tcPing: %w", err)
	}

	return spec, err
}

// loadTcPingObjects loads tcPing and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*tcPingObjects
//	*tcPingPrograms
//	*tcPingMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadTcPingObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadTcPing()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// tcPingSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tcPingSpecs struct {
	tcPingProgramSpecs
	tcPingMapSpecs
}

// tcPingSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tcPingProgramSpecs struct {
	ClsMain  *ebpf.ProgramSpec `ebpf:"cls_main"`
	Pingpong *ebpf.ProgramSpec `ebpf:"pingpong"`
}

// tcPingMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tcPingMapSpecs struct {
}

// tcPingObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadTcPingObjects or ebpf.CollectionSpec.LoadAndAssign.
type tcPingObjects struct {
	tcPingPrograms
	tcPingMaps
}

func (o *tcPingObjects) Close() error {
	return _TcPingClose(
		&o.tcPingPrograms,
		&o.tcPingMaps,
	)
}

// tcPingMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadTcPingObjects or ebpf.CollectionSpec.LoadAndAssign.
type tcPingMaps struct {
}

func (m *tcPingMaps) Close() error {
	return _TcPingClose()
}

// tcPingPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadTcPingObjects or ebpf.CollectionSpec.LoadAndAssign.
type tcPingPrograms struct {
	ClsMain  *ebpf.Program `ebpf:"cls_main"`
	Pingpong *ebpf.Program `ebpf:"pingpong"`
}

func (p *tcPingPrograms) Close() error {
	return _TcPingClose(
		p.ClsMain,
		p.Pingpong,
	)
}

func _TcPingClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed tcping_bpfel.o
var _TcPingBytes []byte
