package main

import "sync"

var (
	Sync	sync.Mutex
	Datas	map[string]string
)

func DataInit() {
	Sync.Lock()
	Datas = make(map[string]string)
	Sync.Unlock()
}

func DataCreate(key string, val string) {
	Sync.Lock()
	Datas[key] = val
	Sync.Unlock()
}

func DataGet(key string) string {
	Sync.Lock()
	val,_ := Datas[key]
	Sync.Unlock()
	return val
}

func DataUpdate(key string, val string) {
	Sync.Lock()
	if _,ok := Datas[key]; ok {
		Datas[key] = val
	}
	Sync.Unlock()
}

func DataDelete(key string) {
	Sync.Lock()
	if _,ok := Datas[key]; ok {
		delete(Datas, key)
	}
	Sync.Unlock()
}

func DataCount() int {
	Sync.Lock()
	count := len(Datas)
	Sync.Unlock()
	return count
}
