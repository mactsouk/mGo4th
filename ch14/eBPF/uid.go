package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"os/signal"

	bpf "github.com/iovisor/gobpf/bcc"
)

import "C"

const source string = `
#include <uapi/linux/ptrace.h>
BPF_HASH(counts);

int count(struct pt_regs *ctx) {
	if (!PT_REGS_PARM1(ctx))
		return 0;

	u64 *pointer;
	u64 times = 0;
	u64 uid;

    uid = bpf_get_current_uid_gid() & 0xFFFFFFFF;
	pointer = counts.lookup(&uid);
    	if (pointer !=0)
        	times = *pointer;

	times++;
        counts.update(&uid, &times);

	return 0;
}
`

func main() {
	pid := flag.Int("pid", -1, "attach to pid, default is all processes")
	flag.Parse()
	m := bpf.NewModule(source, []string{})
	defer m.Close()

	Uprobe, err := m.LoadUprobe("count")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load uprobe count: %s\n", err)
		return
	}

	err = m.AttachUprobe("c", "getuid", Uprobe, *pid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to attach uprobe to getuid: %s\n", err)
		return
	}

	table := bpf.NewTable(m.TableId("counts"), m)
	fmt.Println("Tracing getuid()... Press Ctrl-C to end.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	fmt.Printf("%s\t%s\n", "User ID", "COUNT")
	for it := table.Iter(); it.Next(); {
		k := binary.LittleEndian.Uint64(it.Key())
		v := binary.LittleEndian.Uint64(it.Leaf())
		fmt.Printf("%d\t\t%d\n", k, v)
	}
}
