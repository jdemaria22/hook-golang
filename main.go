package main

import (
	"fmt"
	"framework-memory-go/src/hook"
	"framework-memory-go/src/minimap"
	"framework-memory-go/src/renderer"
	"framework-memory-go/src/time"
)

func main() {
	fmt.Println("inject initializing...")
	processHook, err := hook.Hook()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("updating time...")
	time, err := time.Update(processHook)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Time: ", time)

	fmt.Println("updating Renderer...")
	renderer, err := renderer.Update(processHook)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Renderer :", renderer)

	fmt.Println("updating Minimap...")
	minimap, err := minimap.Update(processHook)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Minimap :", minimap)

}
