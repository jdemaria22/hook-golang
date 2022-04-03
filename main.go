package main

import (
	"fmt"
	"framework-memory-go/src/hook"
	"framework-memory-go/src/minimap"
	"framework-memory-go/src/renderer"
	"framework-memory-go/src/time"
	"framework-memory-go/src/unitmanager"
	"sync"
)

var (
	HOOK hook.ProcessHook
)

func init() {
	processHook, err := hook.GetHook()
	if err != nil {
		fmt.Println(err)
	}
	HOOK = processHook
}

func main() {
	var wg sync.WaitGroup
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = time.Update()
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = renderer.Update()
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = minimap.Update()
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = unitmanager.Update()
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
	fmt.Println("Time: ", time.TIME)
	fmt.Println("Renderer :", renderer.RENDERER)
	fmt.Println("Minimap :", minimap.MINIMAP)
	fmt.Println("Unit manager champ: .", unitmanager.UNITMANAGER.Champions)
	fmt.Println("Unit manager minios: .", len(unitmanager.UNITMANAGER.Minions))
}
