package entities

import "gogame/components"

type MapTile struct {
	components.MapTile
	components.Position
	components.Sprite
}
