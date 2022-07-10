package entities

import (
	components "gogame/components"
)

type Player struct {
	components.InputMovement
	components.Position
	components.Sprite
	components.CameraTarget
	components.SpriteAnimation
	components.CollisionShape
}
