package components

type Animation struct {
	Frames       []int
	TimeStep     float64
	CurrentFrame int
	InitialFrame int
}

type SpriteAnimation struct {
	Animations       map[string]*Animation
	CurrentAnimation string
	LastTime         float64
}

func (sa *SpriteAnimation) SetAnimation(animationName string) {
	if sa.CurrentAnimation == animationName {
		return
	}

	currentAnim := sa.Animations[sa.CurrentAnimation]
	currentAnim.CurrentFrame = currentAnim.InitialFrame

	sa.CurrentAnimation = animationName
	sa.LastTime = 0

	newCurrentAnimation := sa.Animations[sa.CurrentAnimation]
	newCurrentAnimation.CurrentFrame = newCurrentAnimation.InitialFrame
}
