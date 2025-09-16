package main

import (
	status "github.com/keainya/status/kitex_gen/status/statusservice"
	"log"
)

func main() {
	svr := status.NewServer(new(StatusServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
