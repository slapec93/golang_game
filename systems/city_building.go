package systems

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/slapec93/golang_game/entities"
)

// CityBuildingSystem ...
type CityBuildingSystem struct {
	world        *ecs.World
	mouseTracker entities.MouseTracker
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (*CityBuildingSystem) Remove(ecs.BasicEntity) {}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (cb *CityBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddCity").JustPressed() {
		fmt.Println("The gamer pressed F1")

		city := entities.CreateNewCity(&engo.Point{
			X: cb.mouseTracker.MouseX,
			Y: cb.mouseTracker.MouseY,
		}, &engo.Point{
			X: 30,
			Y: 64,
		})
		for _, system := range cb.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
			}
		}
	}
}

// New is the initialisation of the System
func (cb *CityBuildingSystem) New(w *ecs.World) {
	fmt.Println("CityBuildingSystem was added to the Scene")

	cb.world = w
	cb.mouseTracker.BasicEntity = ecs.NewBasic()
	cb.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&cb.mouseTracker.BasicEntity, &cb.mouseTracker.MouseComponent, nil, nil)
		}
	}
}
