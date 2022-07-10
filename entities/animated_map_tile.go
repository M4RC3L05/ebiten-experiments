package entities

import "gogame/components"

type AnimatedMapTile struct {
	components.MapTile
	components.Position
	components.Sprite
	components.SpriteAnimation
}
