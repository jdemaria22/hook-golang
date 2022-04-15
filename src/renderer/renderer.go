package renderer

import (
	Hook "framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"sync"
)

type Renderer struct {
	Width          int32
	Height         int32
	ViewMatrix     []float32
	ProjMatrix     []float32
	ViewProjMatrix [16]float32
}

const (
	INITIAL_SUM_MULTIPLU_MATIX float32 = 0
	INITIAL_VIEW_PROJ          int     = 0
	LAST_VALUE_VIEW_PROJ       int     = 4
)

var (
	hook         Hook.ProcessHook = Hook.HOOK
	RENDERER     Renderer
	wg           sync.WaitGroup
	rendererBase = 0
)

func Update() error {

	if rendererBase == 0 {
		rendererBaseValue, err := memory.ReadInt(hook.Process, hook.ModuleBaseAddr+offset.RENDERER)
		if err != nil {
			return err
		}
		rendererBase = rendererBaseValue
		RENDERER.Width = 1920
		RENDERER.Width = 1080
	}

	// rendererBaseBuff, err := memory.Read(hook.Process, rendererBase, 128)
	// if err != nil {
	// 	return err
	// }

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	// RENDERER.Width = memory.CopyInt(rendererBaseBuff, offset.RENDERERWIDTH)
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	RENDERER.Height = memory.CopyInt(rendererBaseBuff, offset.RENDERERHEIGHT)
	// }()

	viewProjMatrix := hook.ModuleBaseAddr + offset.VIEWPROJMATRICES
	viewProjMatricesBuff, err := memory.Read(hook.Process, viewProjMatrix, 128)
	if err != nil {
		return err
	}

	viewMatrix := make([]float32, 16)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			viewMatrix[i] = memory.Float32frombytes(viewProjMatricesBuff[i*4:])
		}(i)
	}

	projMatrix := make([]float32, 16)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			projMatrix[i] = memory.Float32frombytes(viewProjMatricesBuff[0x40+(i*4):])
		}(i)
	}
	wg.Wait()
	RENDERER.ProjMatrix = projMatrix
	RENDERER.ViewMatrix = viewMatrix

	var mMatrix [16]float32

	mMatrix, _ = multiplyMatrices(RENDERER)

	wg.Wait()
	RENDERER.ViewProjMatrix = mMatrix
	return nil
}

func multiplyMatrices(renderer Renderer) ([16]float32, error) {
	var mMatrix [16]float32
	for i := INITIAL_VIEW_PROJ; i < LAST_VALUE_VIEW_PROJ; i++ {
		for j := INITIAL_VIEW_PROJ; j < LAST_VALUE_VIEW_PROJ; j++ {
			var sum float32 = INITIAL_SUM_MULTIPLU_MATIX
			for k := INITIAL_VIEW_PROJ; k < LAST_VALUE_VIEW_PROJ; k++ {
				sum += renderer.ViewMatrix[i*LAST_VALUE_VIEW_PROJ+k] * renderer.ProjMatrix[k*LAST_VALUE_VIEW_PROJ+j]
			}
			mMatrix[i*LAST_VALUE_VIEW_PROJ+j] = sum
		}
	}
	return mMatrix, nil
}
