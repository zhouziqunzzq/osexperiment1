package main

import (
	"container/heap"
	"fmt"
	. "github.com/zhouziqunzzq/osexperiment1/PCB"
	. "github.com/zhouziqunzzq/osexperiment1/priorityqueue"
)

const (
	DefaultQueueLength = 0
	ClearScreenCMD     = "\033[2J\033[1;1H"
	MemoryCapacity     = 1024
	MemoryBarLength    = 30
	MaxPid             = 65535
	MinPid             = 2
)

var (
	processList    = make([]*PCB, DefaultQueueLength)
	readyQueue     = make(PriorityQueue, DefaultQueueLength)
	blockedQueue   = make(PriorityQueue, DefaultQueueLength)
	runningProcess *PCB
	FreeMemory     = MemoryCapacity
	nextPid        = 2
)

func Clear() {
	fmt.Print(ClearScreenCMD)
	return
}

func PrintMemoryBar() {
	t := int(float64(FreeMemory) / float64(MemoryCapacity) * MemoryBarLength)
	fmt.Print("Memory: [")
	for i := 0; i < MemoryBarLength-t; i++ {
		fmt.Print("*")
	}
	for i := 0; i < t; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("]\t%v/%v\n", MemoryCapacity-FreeMemory, MemoryCapacity)
	return
}

func PrintMenu() {
	fmt.Println("============= Process Control Panel ===============")
	fmt.Println("Please choose an operation:")
	fmt.Println("1. Start a new process")
	fmt.Println("2. Timeout running process")
	fmt.Println("3. Event wait running process")
	fmt.Println("4. Event occurs a process")
	fmt.Println("5. Print status")
	fmt.Println("6. Shutdown")
	fmt.Print("> ")
	return
}

func PrintStatus() {
	fmt.Println("============= Process Status ===============")
	PrintMemoryBar()
	if runningProcess != nil {
		fmt.Println("### Running Process ###")
		fmt.Printf("PID\tName\tPrority\n")
		runningProcess.PrintPCB()
		fmt.Println()
	}

	fmt.Println("### Ready queue ###")
	readyQueue.PrintQueue()
	fmt.Println()

	fmt.Println("### Blocked queue ###")
	blockedQueue.PrintQueue()
	fmt.Println()
}

func main() {
	Clear()
	fmt.Println("Starting BSOS v0.1...")
	heap.Init(&readyQueue)
	heap.Init(&blockedQueue)
	//readyQueue.PriPush(NewPCB(2, "test1", 3, 10))
	//readyQueue.PriPush(NewPCB(3, "test2", 2, 10))
	//readyQueue.PriPush(NewPCB(4, "test3", 1, 10))
	//readyQueue.PrintQueue()
	//fmt.Println(readyQueue.PriPop().(*PCB))

	fmt.Println("Initiating initial process(pid = 1)...")
	runningProcess = NewPCB(1, "System", 1, 128)
	FreeMemory -= 128
	processList = append(processList, runningProcess)

	fmt.Println("Init done. Welcome to BSOS v0.1")

	choice := 0
	for choice != 6 {
		PrintMenu()
		fmt.Scanf("%v", &choice)
		switch choice {
		case 1:
			if err := StartNewProcess(); err != nil {
				fmt.Println(err)
			}
		case 2:
			if err := TimeoutRunningProcess(); err != nil {
				fmt.Println(err)
			}
		case 3:
			if err := EventWaitRunningProcess(); err != nil {
				fmt.Println(err)
			}
		case 4:
			if err := EventOccursProcess(); err != nil {
				fmt.Println(err)
			}
		case 5:
			Clear()
			PrintStatus()
		case 6:
			break
		}
	}
	fmt.Println("Shuting down BSOS v0.1...")
	fmt.Println("BYE")
	return
}
