package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var osType = runtime.GOOS
var hidden = flag.Bool("hidden", false, "hidden arg")

// var time = flag.Int("time", 1000, "time arg")

func main() {
	// os.Args[0] = "custom name"
	// SetProcessName("custom name sanduo")
	// time.Sleep(10 * time.Second)
	// getCurrentDirectory()
	// selfProecessPath := getCurrentDirectory()
	// RemoveSelf()
	flag.Parse()
	// 调用自身,隐藏自身
	var testProgressName = os.Args[0]
	// fmt.Println(os.Args[0])
	exec.Command(testProgressName, "-hidden=true").Start()
	// cmd := exec.Command(testProgressName, "-hidden=true")
	// stdout, _ := cmd.StdoutPipe()
	// cmd.Start()
	RemoveSelf()

	// for {
	// 	tmp := make([]byte, 1024)
	// 	_, err := stdout.Read(tmp)
	// 	fmt.Print(string(tmp))
	// 	if err != nil {
	// 		break
	// 	}
	// }
	// fmt.Println(*hidden)

	if *hidden {
		customName := RandStringRunes(16)
		fmt.Println("customName is ", customName)
		// 两个一起调用
		SetProcessName(customName)
		SetProcessName1(customName)
		println("sleeping")
		// sleepTime, _ := strconv.Atoi(*time)
		time.Sleep(1000 * time.Second)
		println("done")
	}
}

// htop ps 显示进程名已随机，pstree 未显示进程名随机
func SetProcessName(name string) error {
	argv0str := (*reflect.StringHeader)(unsafe.Pointer(&os.Args[0]))
	argv0 := (*[1 << 30]byte)(unsafe.Pointer(argv0str.Data))[:argv0str.Len]

	n := copy(argv0, name)
	if n < len(argv0) {
		argv0[n] = 0
	}

	return nil
}

// pstree 显示进程名已随机，ps，htop未显示随机
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
		fileNameSplit := strings.Split(os.Args[0], "\\")
		fileName := fileNameSplit[len(fileNameSplit)-1]
		dir += "\\" + fileName
	} else {
		fileNameSplit := strings.Split(os.Args[0], "/")
		fileName := fileNameSplit[len(fileNameSplit)-1]
		dir += "/" + fileName
	}
	// fmt.Println(dir)
	return dir
}
