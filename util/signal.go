package util

import (
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

var (
	gSigReloadFns   = make(map[uintptr]SigFunc, 0)
	gSigShutdownFns = make(map[uintptr]SigFunc, 0)
)

type SigFunc func()

func TrapSignals() {
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGUSR1)

		for sig := range sigchan {
			switch sig {
			case syscall.SIGTERM,
				syscall.SIGQUIT,
				syscall.SIGINT:
				if gSigShutdownFns != nil {
					for _, fn := range gSigShutdownFns {
						fn()
					}
				}
				os.Exit(0)

			case syscall.SIGUSR1:
				// reloading
				if gSigReloadFns != nil {
					for _, fn := range gSigReloadFns {
						fn()
					}
				}

			}
		}
	}()
}

func OnShutdownEvent(fn SigFunc) {

	v := reflect.ValueOf(fn)
	addr := v.Pointer()
	if addr > 0 {
		if _, ok := gSigShutdownFns[addr]; ok {
			return
		}
		gSigShutdownFns[addr] = fn
	}
}

func OnReloadEvent(fn SigFunc) {
	v := reflect.ValueOf(fn)
	addr := v.Pointer()

	if addr > 0 {
		if _, ok := gSigReloadFns[addr]; ok {
			return
		}
		gSigReloadFns[addr] = fn
	}
}

func FireReload() {
	syscall.Kill(os.Getpid(), syscall.SIGUSR1)
}

func FireTerminate() {
	syscall.Kill(os.Getpid(), syscall.SIGINT)
}
