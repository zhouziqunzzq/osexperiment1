package main

import (
	"fmt"
	. "github.com/zhouziqunzzq/osexperiment1/PCB"
	"github.com/pkg/errors"
	"container/heap"
)

func PrintInputText(s string) {
	fmt.Println(s)
	fmt.Print("> ")
	return
}

func AssignNewPid() (newPid int) {
	newPid = nextPid
	nextPid++
	if nextPid > MaxPid {
		// Now just assign min pid to next pid without check
		nextPid = MinPid
	}
	return
}

func RearrangeRunningProcess() {
	if runningProcess == nil && !readyQueue.IsEmpty() {
		fmt.Println("No running process, dispatch one from ready queue")
		runningProcess = readyQueue.PriPop().(*PCB)
	} else if runningProcess != nil && readyQueue.First().(*PCB).Priority > runningProcess.Priority {
		fmt.Println("New process has higher priority, dispatch it from ready queue")
		tp := readyQueue.PriPop().(*PCB)
		readyQueue.PriPush(runningProcess)
		runningProcess = tp
	}
	return
}

func StartNewProcess() (err error) {
	newPCB := &PCB{}
	Clear()
	// Input process basic info
	fmt.Println("============= Start a New Process ===============")
	PrintInputText("Input process name:")
	fmt.Scanf("%v", &newPCB.Name)
	PrintInputText("Input process priority:")
	fmt.Scanf("%v", &newPCB.Priority)
	PrintInputText("Input process memory:")
	fmt.Scanf("%v", &newPCB.Memory)
	// Check memory
	if newPCB.Memory <= 0 || FreeMemory-newPCB.Memory < 0 {
		err = errors.New("Failed to start a new process: Out of memory")
		return
	}
	// Assign pid, allocate memory and insert into queue
	newPCB.Pid = AssignNewPid()
	FreeMemory -= newPCB.Memory
	processList = append(processList, newPCB)
	readyQueue.PriPush(newPCB)
	fmt.Printf("Successfully created a process(pid=%v) and inserted into ready queue\n", newPCB.Pid)
	RearrangeRunningProcess()
	return
}

func TimeoutRunningProcess() (err error) {
	Clear()
	if runningProcess == nil {
		err = errors.New("No running process")
		return
	}
	if !readyQueue.IsEmpty() {
		tp := readyQueue.PriPop().(*PCB)
		readyQueue.PriPush(runningProcess)
		runningProcess = tp
	} else {
		err = errors.New("Ready queue empty, no need to timeout")
	}
	return
}

func EventWaitRunningProcess() (err error) {
	Clear()
	if runningProcess == nil {
		err = errors.New("No running process")
		return
	}
	blockedQueue.PriPush(runningProcess)
	runningProcess = nil
	RearrangeRunningProcess()
	return
}

func EventOccursProcess() (err error) {
	Clear()
	fmt.Println("============= Event occurs a Process ===============")
	PrintInputText("Input pid:")
	var pid int
	fmt.Scanf("%v", &pid)
	// Check pid
	if pid > len(processList) {
		err = errors.New("Invalid pid")
		return
	}
	found := false
	for i := 0; i < len(blockedQueue); i++ {
		if blockedQueue[i].Pid == pid {
			found = true
			break
		}
	}
	if !found {
		err = errors.New("Invalid pid")
		return
	}
	// Put the process to ready queue
	heap.Remove(&blockedQueue, processList[pid-1].Index)
	readyQueue.PriPush(processList[pid-1])
	RearrangeRunningProcess()
	return
}
