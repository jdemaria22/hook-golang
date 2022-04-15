package module

import (
	"fmt"
	"framework-memory-go/src/minimap"
	"framework-memory-go/src/renderer"
	"framework-memory-go/src/time"
	"framework-memory-go/src/unitmanager"
	"sync"
)

func Update() {
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
	// for i := 0; i < len(unitmanager.UNITMANAGER.Champions); i++ {
	// 	fmt.Println(unitmanager.UNITMANAGER.Champions[i].BasicAtkWindup)
	// }
	wg.Wait()
}
