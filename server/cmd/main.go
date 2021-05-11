package main

import (
<<<<<<< HEAD
=======
	"context"
	"fmt"
	"io/ioutil"
	"net"
>>>>>>> upstream/master
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"server"
	"server/log"
	"server/settings"
	"server/torr"
	"server/version"
	"server/web/api/utils"
)

type args struct {
<<<<<<< HEAD
	Addr     string `arg:"-p" help:"web server host:port"`
	Path     string `arg:"-d" help:"database path"`
	LogPath  string `arg:"-l" help:"log path"`
	RDB      bool   `arg:"-r" help:"start in read-only DB mode"`
	HttpAuth bool   `arg:"-a" help:"http auth on all requests"`
	DontKill bool   `arg:"-k" help:"dont kill server on signal"`
=======
	Port        string `arg:"-p" help:"web server port"`
	Path        string `arg:"-d" help:"database path"`
	LogPath     string `arg:"-l" help:"log path"`
	WebLogPath  string `arg:"-w" help:"web log path"`
	RDB         bool   `arg:"-r" help:"start in read-only DB mode"`
	HttpAuth    bool   `arg:"-a" help:"http auth on all requests"`
	DontKill    bool   `arg:"-k" help:"dont kill server on signal"`
	UI          bool   `arg:"-u" help:"run page torrserver in browser"`
	TorrentsDir string `arg:"-t" help:"autoload torrent from dir"`
>>>>>>> upstream/master
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
	log.Init(params.LogPath, params.WebLogPath)

	//Preconfig(params.DontKill)

<<<<<<< HEAD
	server.Start(params.Addr, params.RDB)
=======
	if params.TorrentsDir != "" {
		go watchTDir(params.TorrentsDir)
	}

	server.Start(params.Port, params.RDB)
>>>>>>> upstream/master
	log.TLogln(server.WaitServer())
	log.Close()
	time.Sleep(time.Second * 3)
	os.Exit(0)
}

<<<<<<< HEAD
=======
func dnsResolve() {
	addrs, err := net.LookupHost("www.google.com")
	if len(addrs) == 0 {
		fmt.Println("Check dns", addrs, err)

		fn := func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", "1.1.1.1:53")
		}

		net.DefaultResolver = &net.Resolver{
			Dial: fn,
		}

		addrs, err = net.LookupHost("www.themoviedb.org")
		fmt.Println("Check new dns", addrs, err)
	}
}

func watchTDir(dir string) {
	time.Sleep(5 * time.Second)
	path, err := filepath.Abs(dir)
	if err != nil {
		path = dir
	}
	for {
		files, err := ioutil.ReadDir(path)
		if err == nil {
			for _, file := range files {
				filename := filepath.Join(path, file.Name())
				if strings.ToLower(filepath.Ext(file.Name())) == ".torrent" {
					sp, err := utils.ParseLink("file://" + filename)
					if err == nil {
						tor, err := torr.AddTorrent(sp, "", "", "")
						if err == nil {
							if tor.GotInfo() {
								if tor.Title == "" {
									tor.Title = tor.Name()
								}
								torr.SaveTorrentToDB(tor)
								tor.Drop()
								os.Remove(filename)
								time.Sleep(time.Second)
							}
						}
					}
				}
			}
		}
		time.Sleep(time.Second)
	}
}
>>>>>>> upstream/master
