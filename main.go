//
// main.go
// Copyright (C) 2021 rmelo <Ricardo Melo <rmelo@ludia.com>>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"log"
	"os"

	"github.com/christopher-dG/go-obs-websocket"
)

func main() {

	// load arguments
	if len(os.Args[1:]) != 2 {
		log.Fatalln("Programs waiting two arguments!")
	}
	sceneList := os.Args[1:]

	// Connect a client.
	c := obsws.Client{Host: "localhost", Port: 4444}
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer c.Disconnect()

	scene, err := obsws.NewGetSceneListRequest().SendReceive(c)
	if err != nil {
		log.Fatal(err)
	}
	if scene.CurrentScene == sceneList[0] {
		_, err := obsws.NewSetCurrentSceneRequest(sceneList[1]).SendReceive(c)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := obsws.NewSetCurrentSceneRequest(sceneList[0]).SendReceive(c)
		if err != nil {
			log.Fatal(err)
		}
	}
}
