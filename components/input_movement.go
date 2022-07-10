package components

import (
	coreMath "gogame/core/math"
)

type InputMovement struct {
	Stack []coreMath.Vector2
}

func (im *InputMovement) Add(input coreMath.Vector2) {
	im.Stack = append(im.Stack, input)
}

func (im *InputMovement) Remove(input coreMath.Vector2) {
	for i := len(im.Stack) - 1; i >= 0; i -= 1 {
		if im.Stack[i].Equals(input) {
			im.Stack = append(im.Stack[:i], im.Stack[i+1:]...)
		}
	}
}

func (im *InputMovement) Last() coreMath.Vector2 {
	if len(im.Stack) <= 0 {
		im.Add(coreMath.VECTOR_ZERO)
	}

	return im.Stack[len(im.Stack)-1]
}
