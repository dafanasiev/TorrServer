package main

import (
	"os"
	"time"

	"github.com/alexflint/go-arg"
	"server"
	"server/log"
	"server/settings"
	"server/version"
)

type args struct {
	Addr     string `arg:"-p" help:"web server host:port"`
	Path     string `arg:"-d" help:"database path"`
	LogPath  string `arg:"-l" help:"log path"`
	RDB      bool   `arg:"-r" help:"start in read-only DB mode"`
	HttpAuth bool   `arg:"-a" help:"http auth on all requests"`
	DontKill bool   `arg:"-k" help:"dont kill server on signal"`
}

func (args) Version() string {
	return "TorrServer " + version.Version
}

var params args

func main() {
	arg.MustParse(&params)

	if params.Path == "" {
		params.Path, _ = os.Getwd()
	}

	if params.Addr == "" {
		params.Addr = "127.0.0.1:8090"
	}

	settings.Path = params.Path
	settings.HttpAuth = params.HttpAuth
	log.Init(params.LogPath)

	//Preconfig(params.DontKill)

	server.Start(params.Addr, params.RDB)
	log.TLogln(server.WaitServer())
	time.Sleep(time.Second * 3)
	os.Exit(0)
}

