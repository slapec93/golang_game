package entities

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

// MouseTracker ...
type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}
