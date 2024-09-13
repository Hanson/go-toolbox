package win

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os/exec"
	"strconv"
	"strings"
	"unsafe"
)

type Process struct {
	Pid    int
	Handle windows.Handle
}

func NewProcessByName(name string) (p *Process, err error) {
	pid := GetPidByName(name)

	p, err = NewProcessByPid(pid)
	if err != nil {
		return
	}

	return
}

func NewProcessByPid(pid int) (p *Process, err error) {
	handle, err := windows.OpenProcess(0x000F0000|0x00100000|0xFFF, false, uint32(pid))
	if err != nil {
		return
	}
	p = &Process{
		Pid:    pid,
		Handle: handle,
	}
	return
}

func GetPidByName(name string) (pid int) {
	// 查找指定名称的进程ID
	cmd := exec.Command("tasklist", "/NH", "/FI", "IMAGENAME eq "+name)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("command error:", err)
		return
	}
	lines := strings.Split(string(output), "\r\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		pid, err = strconv.Atoi(fields[1])
		if err == nil {
			return
		}
	}
	return
}

func (p *Process) ReadProcessMemory(addr uintptr, size int, value *int) (err error) {
	var nread uintptr
	err = windows.ReadProcessMemory(p.Handle, addr, (*byte)(unsafe.Pointer(value)), uintptr(size), &nread)
	if err != nil {
		fmt.Println("read process memory error:", err)
		return
	}

	return
}

func (p *Process) WriteProcessMemory(addr uintptr, size int, value *int) (err error) {
	var nread uintptr
	err = windows.WriteProcessMemory(p.Handle, addr, (*byte)(unsafe.Pointer(value)), uintptr(size), &nread)
	if err != nil {
		fmt.Println("read process memory error:", err)
		return
	}

	return
}
