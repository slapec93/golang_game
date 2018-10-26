package entities

import (
	"image"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// HUD ...
type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Setup ...
func (h *HUD) Setup(size engo.Point) {
	h.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 200},
		Width:    size.X,
		Height:   size.Y,
	}
	hudImage := image.NewUniform(color.RGBA{205, 205, 205, 255})
	hudNRGBA := common.ImageToNRGBA(hudImage, 200, 200)
	hudImageObj := common.NewImageObject(hudNRGBA)
	hudTexture := common.NewTextureSingle(hudImageObj)

	h.RenderComponent = common.RenderComponent{
		Drawable: hudTexture,
		Scale:    engo.Point{X: 1, Y: 1},
	}
	h.RenderComponent.SetShader(common.HUDShader)
	h.RenderComponent.SetZIndex(1)
}
