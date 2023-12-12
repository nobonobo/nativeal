package main

import (
	"log"

	"github.com/nobonobo/nativeal"
)

func main() {
	ctx := nativeal.NewAudioContext()
	defer ctx.Close()
	log.Println(ctx)
}
