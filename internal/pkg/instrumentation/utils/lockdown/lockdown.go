package lockdown

import "github.com/go-logr/logr"

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target amd64,arm64 -cc clang -cflags $CFLAGS bpf ./bpf/probe.bpf.c

func Load(log logr.Logger) {
	log.Info("Loading lockdown")
	err := loadBpfObjects(&bpfObjects{}, nil)
	if err != nil {
		log.Error(err, "Failed to load lockdown")
	}
}
