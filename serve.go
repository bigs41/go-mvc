package main

import (
	"flag"
	"mvcProjectV1/app"
	"mvcProjectV1/config"
)

func main() {
	serverApp := app.CreateApp() // Create Application fiber

	_port := config.Config("PORT")
	_host := config.Config("HOST")
	host := flag.String("host", _host+":"+_port, "listen default 127.0.0.1:"+_port)
	flag.Parse()

	serverApp.Listen(*host)
}
