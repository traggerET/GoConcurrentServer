package main

import "flag"

var params execParams

func init() {
	params.host = *flag.String("h", "localhost", "TCP host")
	params.port = *flag.String("p", "8081", "TCP port")
}

func main() {
	flag.Parse()
	Run()
}
