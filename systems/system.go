package systems

import "engo.io/ecs"

// System is an interface which implements an ECS-System. A System
// should iterate over its Entities on `Update`, in any way
// suitable for the current implementation.
type System interface {
	// Update is ran every frame, with `dt` being the time
	// in seconds since the last frame
	Update(dt float32)

	// Remove removes an Entity from the System
	Remove(ecs.BasicEntity)
}
