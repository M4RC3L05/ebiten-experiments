package systems

import (
	"gogame/components"

	"github.com/sedyh/mizu/pkg/engine"
)

type Movement struct {
	*components.InputMovement
	*components.Position
}

func (km *Movement) Update(w engine.World) {
	cameraEntity, _ := w.View(components.Camera{}).Get()
	var cameraComponent *components.Camera
	cameraEntity.Get(&cameraComponent)

	cameraComponent.Camera.SetPosition(km.Position.Value.X, km.Position.Value.Y)
}
