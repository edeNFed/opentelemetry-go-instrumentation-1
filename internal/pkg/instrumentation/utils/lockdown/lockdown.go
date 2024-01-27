package lockdown

import (
	"github.com/cilium/ebpf/link"
	"github.com/go-logr/logr"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target amd64,arm64 -cc clang -cflags $CFLAGS bpf ./bpf/probe.bpf.c

type Lockdown struct {
	lsmLockDown link.Link
	lsmEbpf     link.Link
}

func Load(log logr.Logger) *Lockdown {
	log.Info("Loading lockdown")
	obj := &bpfObjects{}
	err := loadBpfObjects(obj, nil)
	if err != nil {
		log.Error(err, "Failed to load lockdown")
		return nil
	}

	log.Info("Loaded lockdown")

	ld, err := link.AttachLSM(link.LSMOptions{
		Program: obj.LockedDown,
	})

	if err != nil {
		log.Error(err, "Failed to attach lockdown")
		return nil
	}

	lsmEbpf, err := link.AttachLSM(link.LSMOptions{
		Program: obj.LsmBpf,
	})

	if err != nil {
		log.Error(err, "Failed to attach lockdown")
		return nil
	}

	return &Lockdown{
		lsmLockDown: ld,
		lsmEbpf:     lsmEbpf,
	}
}

func (l *Lockdown) Close() {
	if l.lsmLockDown != nil {
		l.lsmLockDown.Close()
	}

	if l.lsmEbpf != nil {
		l.lsmEbpf.Close()
	}
}
