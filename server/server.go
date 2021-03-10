package server

import (
	"server/settings"
	"server/web"
)

func Start(addr string, roSets bool) {
	settings.InitSets(roSets)
	if addr == "" {
		panic("addr not set")
	}
	web.Start(addr)
}

func WaitServer() string {
	err := web.Wait()
	if err != nil {
		return err.Error()
	}
	return ""
}

func Stop() {
	web.Stop()
	settings.CloseDB()
}
