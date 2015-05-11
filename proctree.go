package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/mgutz/ansi"
)

const (
	PID = 0
	NAME = 1
	PPID = 3
	UTIME = 13
	STIME = 14
	VSIZE = 22
	RSIZE = 23

	StatLength = 44 // fields
	BaseDir = "/proc"
	PageSize = 4 * 1024
)

var processes map[string]*Process

type Process struct {
	// Store as strings, lazily parse the ones we care about
	Name string
	Pid string
	PPid string
	Stime string
	Utime string
	Rsize string
	Vsize string
	Children []*Process
}

func NewProcess(pid string) *Process {
	return &Process{
		Pid: pid,
		PPid: "0",
		Stime: "0",
		Utime: "0",
		Rsize: "0",
		Vsize: "0",
		Children: make([]*Process, 0, 5),
	}
}

func IsPid(filename string) bool {
	_, err := strconv.ParseInt(filename, 10, 0)
	return err == nil
}

func exitWithError(err error, str string) {
	warnWithError(err, str)
	if err != nil {
		os.Exit(1)
	}
}

func warnWithError(err error, str string) {
	if err != nil {
		fmt.Printf("%s: %v", str, err)
	}
}

func processFromFields(fields []string) *Process {
	if len(fields) < StatLength {
		return nil
	}

	return &Process{
		Name: fields[NAME],
		Pid: fields[PID],
		PPid: fields[PPID],
		Stime: fields[STIME],
		Utime: fields[UTIME],
		Rsize: fields[RSIZE],
		Vsize: fields[VSIZE],
		Children: make([]*Process, 0, 5),
	}
}

func insertEntry(process *Process) {
	var entry *Process
	var ok bool

	// Parent
	if entry, ok = processes[process.PPid]; !ok {
		entry = NewProcess(process.PPid)
		processes[process.PPid] = entry
	}
	entry.Children = append(entry.Children, process)

	// Us
	if entry, ok = processes[process.Pid]; !ok {
		processes[process.Pid] = process
	} else {
		// Keep the existing children, then swap the record
		process.Children = entry.Children
		processes[process.Pid] = process
	}

}

func printChildren(process *Process, depth int, siblings []bool) {
	var str string
	for i:= 0; i < depth; i++ {
		if siblings[i] {
			str += " | "
		} else {
			str += "   "
		}
	}

	var accum Accumulator
	recursiveAccumulate(process.Pid, &accum)

	str += " \\_ " + ansi.Color(process.Pid, "blue") + " " + ansi.Color(process.Name, "green")
	fmt.Printf(
		"%-60s  %5s  %5s   %5d  %5d  %-10d  %-10s   %-10d  %10d\n",
		str,
		process.Utime,
		process.Stime,
		accum.Utime,
		accum.Stime,
		getInt(process.Rsize)*PageSize,
		process.Vsize,
		accum.Rsize*PageSize,
		accum.Vsize,
	)

	if len(process.Children) > 0 {
		for i, child := range process.Children {
			printChildren(child, depth+1, append(siblings, i < len(process.Children)-1))
		}
	}
}

type Accumulator struct {
	Stime int64
	Utime int64
	Rsize int64
	Vsize int64
}

func (a *Accumulator) Accumulate(process *Process) {
	a.Stime += getInt(process.Stime)
	a.Utime += getInt(process.Utime)
	a.Rsize += getInt(process.Rsize)
	a.Vsize += getInt(process.Vsize)
}

func getInt(str string) int64 {
	val, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0
	}

	return val
}

func recursiveAccumulate(pid string, accum *Accumulator) {
	process := processes[pid]
	if process == nil {
		return
	}

	accum.Accumulate(process)

	if len(process.Children) > 0 {
		for _, child := range process.Children {
			if child == nil {
				continue
			}
			recursiveAccumulate(child.Pid, accum)
		}
	}
}

func main() {
	matchPid := os.Args[1]
	processes = make(map[string]*Process, 200)

	startTime := time.Now()
	fileList, err := ioutil.ReadDir(BaseDir)
	exitWithError(err, "Can't read " + BaseDir)

	for _, entry := range fileList {
		pid := path.Base(entry.Name())
		if !IsPid(pid) {
			continue
		}

		filename := BaseDir + "/" + pid + "/stat"
		data, err := ioutil.ReadFile(filename)
		exitWithError(err, "Can't read " + filename)

		fields := strings.Split(string(data), " ")
		process := processFromFields(fields)

		insertEntry(process)
	}


	var accum Accumulator
	recursiveAccumulate(matchPid, &accum)
	endTime := time.Now()

	fmt.Printf(ansi.Color("%-42s  %5s  %5s   %5s  %5s  %-10s  %-10s  %-10s  %-10s", "yellow") + "\n",
		"Process", "UTime", "STime", "Ttl S", "Ttl U", "Rsize", "Vsize",
		"Ttl Rs", "Ttl Vs",
	)
	printChildren(processes[matchPid], 0, []bool{false})
	fmt.Printf("\n\n\nTime: %s\n", endTime.Sub(startTime))

	fmt.Printf("%#v\n", accum)
}
