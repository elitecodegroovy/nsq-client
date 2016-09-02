package main

import (
	"fmt"
	"nsq"
)


//nsq consumer task entrance
func main() {
	fmt.Println("--start to process topic message...")
	nsq.ReadNsqMessage("007john", func(c *nsq.Config) {
		c.Snappy = true
	})

}
