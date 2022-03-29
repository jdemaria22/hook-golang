package renderer

import (
	"fmt"
	"framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"framework-memory-go/src/size"
	"unsafe"
)

type Renderer struct {
	Width          int
	Height         int
	ViewMatrix     [16]float32
	ProjMatrix     [16]float32
	ViewProjMatrix [16]float32
}

const (
	INITIAL_SUM_MULTIPLU_MATIX float32 = 0
	INITIAL_VIEW_PROJ          int     = 0
	LAST_VALUE_VIEW_PROJ       int     = 4
)

func Update(hook hook.ProcessHook) (Renderer, error) {
	var renderer Renderer
	rendererBase, err := memory.ReadInt(hook, hook.ModuleBaseAddr+offset.RENDERER)
	if err != nil {
		return renderer, err
	}

	width, err := memory.ReadInt(hook, rendererBase+offset.RENDERERWIDTH)
	if err != nil {
		return renderer, err
	}
	renderer.Width = width

	height, err := memory.ReadInt(hook, rendererBase+offset.RENDERERHEIGHT)
	if err != nil {
		return renderer, err
	}
	renderer.Height = height

	rendererBase1, err := memory.Read(hook, rendererBase, 128)
	if err != nil {
		return renderer, err
	}

	dest := make([]uint16, 1)

	copy(unsafe.Slice((*byte)(unsafe.Pointer(&dest[0])), unsafe.Sizeof(rendererBase1)), rendererBase1[0xC:])
	fmt.Println("memcpy:", dest)
	// fmt.Println("memcpy:", int(dest[offset.RENDERERWIDTH]))

	viewProjMatrices := hook.ModuleBaseAddr + offset.VIEWPROJMATRICES
	if err != nil {
		return renderer, err
	}

	for i := 0; i < len(renderer.ViewMatrix); i++ {
		viewMatrixVal, err := memory.ReadFloat(hook, viewProjMatrices+(i*int(size.Float)))
		if err != nil {
			return renderer, err
		}
		renderer.ViewMatrix[i] = viewMatrixVal
	}

	for i := 0; i < len(renderer.ProjMatrix); i++ {
		viewMatrixVal, err := memory.ReadFloat(hook, viewProjMatrices+64+(i*int(size.Float)))
		if err != nil {
			return renderer, err
		}
		renderer.ProjMatrix[i] = viewMatrixVal
	}

	mMatrix, _ := multiplyMatrices(renderer)
	renderer.ViewProjMatrix = mMatrix
	return renderer, nil
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
