package systems

import (
	"gogame/components"

	"github.com/sedyh/mizu/pkg/engine"
)

type SpriteAnimation struct {
	*components.Sprite
	*components.SpriteAnimation
}

func (sa *SpriteAnimation) Update(w engine.World) {
	currentAnimation := sa.SpriteAnimation.Animations[sa.SpriteAnimation.CurrentAnimation]

	sa.SpriteAnimation.LastTime += 1.0 / 60.0

	if sa.SpriteAnimation.LastTime > currentAnimation.TimeStep {
		sa.SpriteAnimation.LastTime = 0
		currentAnimation.CurrentFrame += 1
		currentAnimation.CurrentFrame %= len(currentAnimation.Frames)

	}

	sa.Sprite.Frame = currentAnimation.Frames[currentAnimation.CurrentFrame]
}
