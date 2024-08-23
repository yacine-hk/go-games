package main

import (
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

type termios struct {
	Iflag, Oflag, Cflag, Lflag uint32
	Cc                         [20]byte
	Ispeed, Ospeed             uint32
}

func getTermios(fd uintptr) (*termios, error) {
	var t termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&t)), 0, 0, 0)
	if err != 0 {
		return nil, err
	}
	return &t, nil
}

func setTermios(fd uintptr, t *termios) error {
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(t)), 0, 0, 0)
	if err != 0 {
		return err
	}
	return nil
}

// Enable raw mode for real-time input
func enableRawMode() (*termios, error) {
	fd := os.Stdin.Fd()
	oldState, err := getTermios(fd)
	if err != nil {
		return nil, err
	}

	raw := *oldState
	raw.Lflag &^= syscall.ICANON | syscall.ECHO // Disable canonical mode and echo

	if err := setTermios(fd, &raw); err != nil {
		return nil, err
	}

	// Set non-blocking mode for stdin
	/*
		if err := syscall.SetNonblock(int(fd), true); err != nil {
			return nil, err
		}
	*/

	return oldState, nil
}

// Disable raw mode and restore terminal state
func disableRawMode(oldState *termios) error {
	fd := os.Stdin.Fd()
	return setTermios(fd, oldState)
}

// Clear the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func handleInput() string {
	var direction string
	var buf [3]byte
	n, err := syscall.Read(int(os.Stdin.Fd()), buf[:])
	if err == nil && n > 0 {
		// Handle WASD keys
		if n == 1 {
			switch buf[0] {
			case 'k':
				direction = "UP"
			case 'j':
				direction = "DOWN"
			case 'h':
				direction = "LEFT"
			case 'l':
				direction = "RIGHT"
			case 10:
				direction = "CR"
			}
		}

		// Handle arrow keys (multi-byte sequences)
		if n == 3 && buf[0] == 0x1b && buf[1] == '[' {
			switch buf[2] {
			case 'A':
				direction = "UP"
			case 'B':
				direction = "DOWN"
			case 'C':
				direction = "RIGHT"
			case 'D':
				direction = "LEFT"
			}
		}
	}
	return direction
}
