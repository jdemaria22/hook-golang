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
	hook     Hook.ProcessHook = Hook.HOOK
	RENDERER Renderer
)

func Update() error {
	var wg sync.WaitGroup

	rendererBase, err := memory.ReadInt(hook.Process, hook.ModuleBaseAddr+offset.RENDERER)
	if err != nil {
		return err
	}

	rendererBaseBuff, err := memory.Read(hook.Process, rendererBase, 128)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		RENDERER.Width = memory.CopyInt(rendererBaseBuff, offset.RENDERERWIDTH)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		RENDERER.Height = memory.CopyInt(rendererBaseBuff, offset.RENDERERHEIGHT)
	}()

	viewProjMatrix := hook.ModuleBaseAddr + offset.VIEWPROJMATRICES
	viewProjMatricesBuff, err := memory.Read(hook.Process, viewProjMatrix, 128)
	if err != nil {
		return err
	}

	viewMatrix := make([]float32, 16)
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[0] = memory.Float32frombytes(viewProjMatricesBuff[0x0 : 0x0+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[0] = memory.Float32frombytes(viewProjMatricesBuff[0x0 : 0x0+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[1] = memory.Float32frombytes(viewProjMatricesBuff[0x4 : 0x4+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[2] = memory.Float32frombytes(viewProjMatricesBuff[0x8 : 0x8+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[3] = memory.Float32frombytes(viewProjMatricesBuff[0xC : 0xC+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[4] = memory.Float32frombytes(viewProjMatricesBuff[0x10 : 0x10+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[5] = memory.Float32frombytes(viewProjMatricesBuff[0x14 : 0x14+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[6] = memory.Float32frombytes(viewProjMatricesBuff[0x18 : 0x18+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[7] = memory.Float32frombytes(viewProjMatricesBuff[0x1C : 0x1C+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[8] = memory.Float32frombytes(viewProjMatricesBuff[0x20 : 0x20+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[9] = memory.Float32frombytes(viewProjMatricesBuff[0x24 : 0x24+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[10] = memory.Float32frombytes(viewProjMatricesBuff[0x28 : 0x28+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[11] = memory.Float32frombytes(viewProjMatricesBuff[0x2c : 0x2c+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[12] = memory.Float32frombytes(viewProjMatricesBuff[0x30 : 0x30+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[13] = memory.Float32frombytes(viewProjMatricesBuff[0x34 : 0x34+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[14] = memory.Float32frombytes(viewProjMatricesBuff[0x38 : 0x38+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		viewMatrix[15] = memory.Float32frombytes(viewProjMatricesBuff[0x3c : 0x3c+4])
	}()
	wg.Wait()
	RENDERER.ViewMatrix = viewMatrix

	projMatrix := make([]float32, 16)
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[0] = memory.Float32frombytes(viewProjMatricesBuff[0x40 : 0x40+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[1] = memory.Float32frombytes(viewProjMatricesBuff[0x44 : 0x44+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[2] = memory.Float32frombytes(viewProjMatricesBuff[0x48 : 0x48+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[3] = memory.Float32frombytes(viewProjMatricesBuff[0x4c : 0x4c+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[4] = memory.Float32frombytes(viewProjMatricesBuff[0x50 : 0x50+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[5] = memory.Float32frombytes(viewProjMatricesBuff[0x54 : 0x54+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[6] = memory.Float32frombytes(viewProjMatricesBuff[0x58 : 0x58+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[7] = memory.Float32frombytes(viewProjMatricesBuff[0x5c : 0x5c+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[8] = memory.Float32frombytes(viewProjMatricesBuff[0x60 : 0x60+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[9] = memory.Float32frombytes(viewProjMatricesBuff[0x64 : 0x64+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[10] = memory.Float32frombytes(viewProjMatricesBuff[0x68 : 0x68+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[11] = memory.Float32frombytes(viewProjMatricesBuff[0x6c : 0x6c+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[12] = memory.Float32frombytes(viewProjMatricesBuff[0x70 : 0x70+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[13] = memory.Float32frombytes(viewProjMatricesBuff[0x74 : 0x74+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[14] = memory.Float32frombytes(viewProjMatricesBuff[0x78 : 0x78+4])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		projMatrix[15] = memory.Float32frombytes(viewProjMatricesBuff[0x7c : 0x7c+4])
	}()
	wg.Wait()
	RENDERER.ProjMatrix = projMatrix

	var mMatrix [16]float32
	wg.Add(1)
	go func() {
		defer wg.Done()
		mMatrix, _ = multiplyMatrices(RENDERER)
	}()

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
