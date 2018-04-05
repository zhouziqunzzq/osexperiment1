package PCB

import "fmt"

// An Item is something we manage in a priority queue.
type PCB struct {
	Pid      int	// Process ID range from 1 to 65535
	Name     string // Process readable name
	Priority int    // Process priority
	Memory int // Process memory usage
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

func NewPCB(pid int, name string, priority int, memory int) (*PCB) {
	return &PCB{
		pid,
		name,
		priority,
		memory,
		-1,
	}
}

func (pcb PCB) PrintPCB() {
	fmt.Printf("%v\t%v\t%v\n", pcb.Pid, pcb.Name, pcb.Priority)
}