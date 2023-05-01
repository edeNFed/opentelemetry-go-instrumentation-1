// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package grpc

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfGrpcRequestT struct {
	StartTime uint64
	EndTime   uint64
	Method    [50]int8
	Target    [50]int8
	Sc        bpfSpanContext
	Psc       bpfSpanContext
	_         [4]byte
}

type bpfHeadersBuff struct{ Buff [500]uint8 }

type bpfSpanContext struct {
	TraceID [16]uint8
	SpanID  [8]uint8
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	UprobeClientConnInvoke              *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Invoke"`
	UprobeClientConnInvokeReturns       *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Invoke_Returns"`
	UprobeHttp2ClientCreateHeaderFields *ebpf.ProgramSpec `ebpf:"uprobe_Http2Client_CreateHeaderFields"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	AllocMap         *ebpf.MapSpec `ebpf:"alloc_map"`
	Events           *ebpf.MapSpec `ebpf:"events"`
	GrpcEvents       *ebpf.MapSpec `ebpf:"grpc_events"`
	HeadersBuffMap   *ebpf.MapSpec `ebpf:"headers_buff_map"`
	TrackedSpans     *ebpf.MapSpec `ebpf:"tracked_spans"`
	TrackedSpansBySc *ebpf.MapSpec `ebpf:"tracked_spans_by_sc"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	AllocMap         *ebpf.Map `ebpf:"alloc_map"`
	Events           *ebpf.Map `ebpf:"events"`
	GrpcEvents       *ebpf.Map `ebpf:"grpc_events"`
	HeadersBuffMap   *ebpf.Map `ebpf:"headers_buff_map"`
	TrackedSpans     *ebpf.Map `ebpf:"tracked_spans"`
	TrackedSpansBySc *ebpf.Map `ebpf:"tracked_spans_by_sc"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.AllocMap,
		m.Events,
		m.GrpcEvents,
		m.HeadersBuffMap,
		m.TrackedSpans,
		m.TrackedSpansBySc,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	UprobeClientConnInvoke              *ebpf.Program `ebpf:"uprobe_ClientConn_Invoke"`
	UprobeClientConnInvokeReturns       *ebpf.Program `ebpf:"uprobe_ClientConn_Invoke_Returns"`
	UprobeHttp2ClientCreateHeaderFields *ebpf.Program `ebpf:"uprobe_Http2Client_CreateHeaderFields"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.UprobeClientConnInvoke,
		p.UprobeClientConnInvokeReturns,
		p.UprobeHttp2ClientCreateHeaderFields,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel.o
var _BpfBytes []byte
