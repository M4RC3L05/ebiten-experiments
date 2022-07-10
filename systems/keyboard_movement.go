package systems

import (
	components "gogame/components"
	coreMath "gogame/core/math"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sedyh/mizu/pkg/engine"
)

type KeyboardMovement struct {
	*components.InputMovement
	*components.Position
	*components.SpriteAnimation
}

func (km *KeyboardMovement) Update(w engine.World) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		km.InputMovement.Add(coreMath.VECTOR_RIGHT)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		km.InputMovement.Remove(coreMath.VECTOR_RIGHT)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		km.InputMovement.Add(coreMath.VECTOR_LEFT)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) {
		km.InputMovement.Remove(coreMath.VECTOR_LEFT)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		km.InputMovement.Add(coreMath.VECTOR_DOWN)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
		km.InputMovement.Remove(coreMath.VECTOR_DOWN)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		km.InputMovement.Add(coreMath.VECTOR_UP)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
		km.InputMovement.Remove(coreMath.VECTOR_UP)
	}

	lastInput := km.InputMovement.Last()

	if lastInput.Equals(coreMath.VECTOR_LEFT) {
		km.SpriteAnimation.SetAnimation("WALK_LEFT")
	}

	if lastInput.Equals(coreMath.VECTOR_RIGHT) {
		km.SpriteAnimation.SetAnimation("WALK_RIGHT")
	}

	if lastInput.Equals(coreMath.VECTOR_UP) {
		km.SpriteAnimation.SetAnimation("WALK_UP")
	}

	if lastInput.Equals(coreMath.VECTOR_DOWN) {
		km.SpriteAnimation.SetAnimation("WALK_DOWN")
	}

	if lastInput.Equals(coreMath.VECTOR_ZERO) {
		currAnimName := strings.Split(km.SpriteAnimation.CurrentAnimation, "_")
		km.SpriteAnimation.SetAnimation("IDLE_" + currAnimName[1])
	}
}
