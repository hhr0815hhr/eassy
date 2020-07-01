package util

import "log"

func RecoverPanic(f func()) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("回收错误:%s", p)
		}

	}()
	f()
}
