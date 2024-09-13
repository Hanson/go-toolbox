package win

import (
	"fmt"
	"github.com/AllenDang/w32"
	"syscall"
	"testing"
	"unsafe"
)

func TestGetExeVersion(t *testing.T) {
	//log.Println(GetVersion("D:\\software\\WeChat\\[3.6.0.18]\\WeChatWin.dll"))
	//process, err := os.FindProcess(123)
	//if err != nil {
	//	log.Printf("err: %+v", err)
	//	return
	//}
	main1(0x22300E0)
	main1(0x223D90C)
	//main1(0x22300E0)
	//main1(0x22300E0)
	//for _, pid := range ListProcesses() {
	//	if pid == 30476 {
	//		log.Println(GetProcessFilePath(pid))
	//	}
	//}
}

func ListProcesses() []uint32 {
	sz := uint32(10000)
	procs := make([]uint32, sz)
	var bytesReturned uint32
	if w32.EnumProcesses(procs, sz, &bytesReturned) {
		return procs[:int(bytesReturned)/4]
	}
	return []uint32{}
}

func GetProcessFilePath1(id uint32) string {
	snapshot := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPMODULE, id)
	if snapshot == w32.ERROR_INVALID_HANDLE {
		return "<UNKNOWN>"
	}
	defer w32.CloseHandle(snapshot)

	var me w32.MODULEENTRY32
	me.Size = uint32(unsafe.Sizeof(me))
	me.SzExePath = [w32.MAX_PATH]uint16{}
	if w32.Module32First(snapshot, &me) {
		return w32.UTF16PtrToString(&me.SzModule[0])
	}

	return "<UNKNOWN>"
}

func GetProcessFilePath(id uint32) string {
	snapshot := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPMODULE, id)
	if snapshot == w32.ERROR_INVALID_HANDLE {
		return "<UNKNOWN>"
	}
	defer w32.CloseHandle(snapshot)

	var me w32.MODULEENTRY32
	me.Size = uint32(unsafe.Sizeof(me))
	me.SzExePath = [w32.MAX_PATH]uint16{}
	if w32.Module32First(snapshot, &me) {
		return w32.UTF16PtrToString(&me.SzExePath[0])
	}

	return "<UNKNOWN>"
}

func main1(offset int) {
	// 定义常量
	const (
		PROCESS_ALL_ACCESS     = 0x1F0FFF // 所有进程访问权限
		PAGE_EXECUTE_READWRITE = 0x40     // 读写和执行权限
	)

	// 假设你的WeChatWin.dll已经加载到了目标进程中
	dllName := "WeChatWin.dll"
	//dllName := "微信.exe"
	dll, err := syscall.LoadLibrary(dllName)
	if err != nil {
		fmt.Printf("加载DLL失败: %v\n", err)
		return
	}
	defer syscall.FreeLibrary(dll)

	// 获取模块基址
	moduleBase := uintptr(dll)

	// 完整地址
	targetAddr := moduleBase + uintptr(offset)

	// 获取当前进程句柄
	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		fmt.Printf("获取当前进程句柄失败: %v\n", err)
		return
	}

	// 更改目标内存保护为可写
	var oldProtect uint32
	ret, _, err := syscall.Syscall6(
		syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect").Addr(),
		4,
		targetAddr,
		uintptr(len([]byte{0x63, 0x09, 0x00, 0xB1})),
		PAGE_EXECUTE_READWRITE,
		uintptr(unsafe.Pointer(&oldProtect)),
		0, 0,
	)
	if ret == 0 {
		fmt.Printf("修改内存保护失败: %v\n", err)
		return
	}

	// 写入新的值
	newValue := []byte{0x63, 0x09, 0x00, 0xB1}
	_, _, err = syscall.Syscall6(
		syscall.NewLazyDLL("kernel32.dll").NewProc("WriteProcessMemory").Addr(),
		5,
		uintptr(handle),
		targetAddr,
		uintptr(unsafe.Pointer(&newValue[0])),
		uintptr(len(newValue)),
		0,
		0,
	)
	if err != nil && err.Error() != "The operation completed successfully." {
		fmt.Printf("写入内存失败: %v\n", err)
		return
	}

	// 恢复内存保护
	_, _, err = syscall.Syscall6(
		syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect").Addr(),
		4,
		targetAddr,
		uintptr(len(newValue)),
		uintptr(oldProtect),
		uintptr(unsafe.Pointer(&oldProtect)),
		0, 0,
	)
	if err != nil && err.Error() != "The operation completed successfully." {
		fmt.Printf("恢复内存保护失败: %v\n", err)
		return
	}

	fmt.Println("内存修改成功！")
}

const (
	PROCESS_ALL_ACCESS     = 0x1F0FFF // 所有进程访问权限
	PAGE_EXECUTE_READWRITE = 0x40     // 读写和执行权限
)
