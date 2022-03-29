package renderer

import (
	"framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"unsafe"
)

type Renderer struct {
	Width          uint32
	Height         uint32
	ViewMatrix     []float32
	ProjMatrix     []float32
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

	rendererBaseBuff, err := memory.Read(hook, rendererBase, 128)
	if err != nil {
		return renderer, err
	}

	destint := make([]uint32, 1)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rendererBaseBuff)), rendererBaseBuff[offset.RENDERERWIDTH:])
	renderer.Width = destint[0]
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rendererBaseBuff)), rendererBaseBuff[offset.RENDERERHEIGHT:])
	renderer.Height = destint[0]

	viewProjMatrix := hook.ModuleBaseAddr + offset.VIEWPROJMATRICES
	viewProjMatricesBuff, err := memory.Read(hook, viewProjMatrix, 128)
	if err != nil {
		return renderer, err
	}

	viewMatrix := make([]float32, 16)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[0])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x0:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[1])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x4:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[2])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x8:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[3])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0xC:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[4])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x10:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[5])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x14:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[6])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x18:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[7])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x1c:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[8])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x20:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[9])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x24:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[10])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x28:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[11])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x2c:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[12])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x30:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[13])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x34:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[14])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x38:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&viewMatrix[15])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x3c:])
	renderer.ViewMatrix = viewMatrix

	projMatrix := make([]float32, 16)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[0])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x40:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[1])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x44:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[2])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x48:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[3])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x4c:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[4])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x50:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[5])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x54:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[6])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x58:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[7])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x5c:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[8])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x60:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[9])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x64:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[10])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x68:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[11])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x6c:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[12])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x70:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[13])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x74:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[14])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x78:])
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&projMatrix[15])), unsafe.Sizeof(viewProjMatricesBuff)), viewProjMatricesBuff[0x7c:])
	renderer.ProjMatrix = projMatrix

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
