#include "common.h"
#include "bpf_tracing.h"
#include "bpf_helpers.h"

char __license[] SEC("license") = "Dual MIT/GPL";

SEC("lsm/bpf")
int BPF_PROG(lsm_bpf, int cmd)
{
    bpf_printk("lsm_bpf: cmd=%d", cmd);
    return 0;
}

SEC("lsm/locked_down")
int BPF_PROG(locked_down, enum lockdown_reason what, int ret)
{
    bpf_printk("locked_down: what=%d ret=%d\n", what, ret);
    return 0;
}