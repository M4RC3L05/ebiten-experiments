package components

import (
	coreMath "gogame/core/math"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image       *ebiten.Image
	Frame       int
	HFrames     int
	VFrames     int
	RenderPivot coreMath.Vector2
	SortPivot   coreMath.Vector2
}

func (s *Sprite) GetFrameBounds() image.Rectangle {
	frameX := s.Frame % s.HFrames
	frameY := s.Frame / s.HFrames

	frameSizeX, frameSizeY := s.GetFrameSize()

	return image.Rect(frameX*frameSizeX, frameY*frameSizeY, (frameX*frameSizeX)+frameSizeX, (frameY*frameSizeY)+frameSizeY)
}

func (s *Sprite) GetFrameSize() (int, int) {
	imgSize := s.Image.Bounds()

	frameSizeX := imgSize.Max.X / s.HFrames
	frameSizeY := imgSize.Max.Y / s.VFrames

	return frameSizeX, frameSizeY
}
