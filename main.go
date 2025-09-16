package main

import (
	"log"

	status "github.com/keainya/status/kitex_gen/status/statusservice"
)

func main() {
	svr := status.NewServer(new(StatusServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
