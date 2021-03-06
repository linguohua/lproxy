package main

import (
	"flag"
	"fmt"
	"lproxy/server"
	"lproxy/servercfg"
	"os"
	"runtime"
	"runtime/pprof"

	log "github.com/sirupsen/logrus"

	_ "lproxy/handlers/auth"
	_ "lproxy/xport"
	_ "lproxy/handlers/sayhello"
	_ "lproxy/handlers/dv"
)

var (
	cfgFilepath    = ""
	serverUUID     = ""
)

func init() {
	flag.StringVar(&cfgFilepath, "c", "", "specify the config file path name")
	flag.StringVar(&serverUUID, "u", "", "specify the server UUID")
}

func main() {
	// only one thread
	runtime.GOMAXPROCS(1)

	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Printf("%s\n", server.GetVersion())
		os.Exit(0)
	}

	if serverUUID != "" {
		servercfg.ServerID = serverUUID
	}

	if cfgFilepath == "" {
		// 如果没有配置json文件，则必须提供uuid以及redis地址
		if serverUUID == "" {
			log.Fatal("must provide redis and uuid when json config file is omit")
		}
	}

	if cfgFilepath != "" {
		r := servercfg.ParseConfigFile(cfgFilepath)
		if r != true {
			log.Fatal("can't parse configure file:", cfgFilepath)
		}
	} else {
		log.Fatal("please specify a valid config file path")
	}

	log.Println("try to start  lproxy server, version:", server.GetVersion())

	server.OnCfgLoaded()

	// start http server
	server.CreateHTTPServer()
	log.Println("start lproxy server ok!")

	if servercfg.Daemon == "yes" {
		waitForSignal()
	} else {
		waitInput()
	}
	return
}

func waitInput() {
	var cmd string
	for {
		_, err := fmt.Scanf("%s\n", &cmd)
		if err != nil {
			//log.Println("Scanf err:", err)
			continue
		}

		switch cmd {
		case "exit", "quit":
			log.Println("exit by user")
			return
		case "gr":
			log.Println("current goroutine count:", runtime.NumGoroutine())
			break
		case "gd":
			pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
			break
		default:
			break
		}
	}
}

func dumpGoRoutinesInfo() {
	log.Println("current goroutine count:", runtime.NumGoroutine())
	// use DEBUG=2, to dump stack like golang dying due to an unrecovered panic.
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)
}
