package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

var osType = runtime.GOOS

func main() {
	// os.Args[0] = "custom name"
	// SetProcessName("custom name sanduo")
	// time.Sleep(10 * time.Second)
	// getCurrentDirectory()
	// selfProecessPath := getCurrentDirectory()
	// var testProgressName = "./test"
	// exec.Command(testProgressName).Start()
	// exec.Command(testProgressName).Start()
	RemoveSelf()
	customName := RandStringRunes(16)
	SetProcessName(customName)
	println("sleeping")
	time.Sleep(10000 * time.Second)
	println("done")
}

func SetProcessName(name string) error {
	argv0str := (*reflect.StringHeader)(unsafe.Pointer(&os.Args[0]))
	argv0 := (*[1 << 30]byte)(unsafe.Pointer(argv0str.Data))[:argv0str.Len]

	n := copy(argv0, name)
	if n < len(argv0) {
		argv0[n] = 0
	}

	return nil
}

func SetProcessName1(name string) error {
	bytes := append([]byte(name), 0)
	ptr := unsafe.Pointer(&bytes[0])
	if _, _, errno := syscall.RawSyscall6(syscall.SYS_PRCTL, syscall.PR_SET_NAME, uintptr(ptr), 0, 0, 0, 0); errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RemoveSelf() {
	err := os.Remove(getCurrentDirectory())
	if err != nil {
		fmt.Println(err)
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir((os.Args[0])))
	if err != nil {
		fmt.Println(err)
	}
	if osType == "windows" {
		dir += "\\" + os.Args[0]
	} else {
		dir += "/" + os.Args[0]
	}
	fmt.Println(dir)
	return dir
}
