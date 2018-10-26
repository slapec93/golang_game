package entities

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// City ...
type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// CreateNewCity ...
func CreateNewCity(position *engo.Point, size *engo.Point) City {
	city := City{BasicEntity: ecs.NewBasic()}
	city.SpaceComponent = common.SpaceComponent{
		Position: *position,
		Width:    size.X,
		Height:   size.Y,
	}
	texture, err := common.LoadedSprite("textures/city.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	city.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: .1, Y: .1},
	}
	return city
}
