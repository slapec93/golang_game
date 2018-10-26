package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/slapec93/golang_game/entities"
	"github.com/slapec93/golang_game/systems"
)

const (
	windowWidth  = 800
	windowHeight = 800
	cameraSpeed  = 400
)

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load("textures/city.png", "tilemap/TrafficMap.tmx")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	engo.Input.RegisterButton("AddCity", engo.KeyF1)
	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	kbs := common.NewKeyboardScroller(
		cameraSpeed,
		engo.DefaultHorizontalAxis,
		engo.DefaultVerticalAxis,
	)
	world.AddSystem(kbs)
	world.AddSystem(&common.MouseZoomer{0.125})
	// world.AddSystem(&common.EdgeScroller{
	// 	ScrollSpeed: cameraSpeed,
	// 	EdgeMargin:  20,
	// })

	hud := entities.HUD{BasicEntity: ecs.NewBasic()}
	hud.Setup(engo.Point{X: 200, Y: 200})

	tiles := entities.CreateNewTiles()

	// And finally add it to the world:
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
			for _, v := range tiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}

	world.AddSystem(&systems.CityBuildingSystem{})
}

func main() {
	opts := engo.RunOptions{
		Title:          "Traffic Manager",
		Width:          windowWidth,
		Height:         windowHeight,
		StandardInputs: true,
	}
	engo.Run(opts, &myScene{})
}
