package systems

import (
	"gogame/components"

	"github.com/sedyh/mizu/pkg/engine"
)

type Collision struct {
	*components.InputMovement
	*components.Position
	*components.CollisionShape
}

func (c *Collision) Update(w engine.World) {
	if c.CollisionShape.Static {
		return
	}

	to := c.InputMovement.Last().MultScalar(80.0).MultScalar(1.0 / 60.0)

	if collisionX := c.CollisionShape.Shape.Check(to.X, 0); collisionX != nil {
		to.X = collisionX.ContactWithObject(collisionX.Objects[0]).X()
	}

	if collisionY := c.CollisionShape.Shape.Check(0, to.Y); collisionY != nil {
		to.Y = collisionY.ContactWithObject(collisionY.Objects[0]).Y()
	}

	c.Position.Value = c.Position.Value.Add(to)
	c.CollisionShape.Shape.X = c.Position.Value.X - 8
	c.CollisionShape.Shape.Y = c.Position.Value.Y

	c.CollisionShape.Shape.Update()
}
