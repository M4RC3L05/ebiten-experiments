package systems

import (
	"gogame/components"
	"image/color"
	_ "image/png"
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct{}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	cameraEntity, _ := w.View(components.Camera{}).Get()
	var cameraComponent *components.Camera
	cameraEntity.Get(&cameraComponent)

	minX := int(math.Max(0, float64(int(cameraComponent.Camera.X)-(cameraComponent.Camera.Width/2))))
	minY := int(math.Max(0, float64(int(cameraComponent.Camera.Y)-(cameraComponent.Camera.Height/2))))
	maxX := int(math.Max(0, float64(int(cameraComponent.Camera.X)+(cameraComponent.Camera.Width/2))))
	maxY := int(math.Max(0, float64(int(cameraComponent.Camera.Y)+(cameraComponent.Camera.Height/2))))

	renderables := w.View(components.Position{}, components.Sprite{}).Filter()
	sort.SliceStable(renderables, func(i, j int) bool {
		e1 := renderables[i]
		e2 := renderables[j]

		var e1Pos, e2Pos *components.Position
		var e1Sprite, e2Sprite *components.Sprite

		e1.Get(&e1Pos, &e1Sprite)
		e2.Get(&e2Pos, &e2Sprite)

		_, e1h := e1Sprite.GetFrameSize()
		_, e2h := e2Sprite.GetFrameSize()

		e1Pivot := e1Pos.Value.Y + float64(e1h)*e1Sprite.SortPivot.Y
		e2Pivot := e2Pos.Value.Y + float64(e2h)*e2Sprite.SortPivot.Y

		return e1Pivot < e2Pivot
	})

	cameraComponent.Camera.Surface.Clear()
	cameraComponent.Camera.Surface.Fill(color.RGBA{0, 0, 0, 255})

	for _, e := range renderables {
		var position *components.Position
		var sprite *components.Sprite
		var cameraTarget *components.CameraTarget

		e.Get(&position, &sprite, &cameraTarget)

		w, h := sprite.GetFrameSize()
		renderW := float64(w) * sprite.RenderPivot.X
		renderH := float64(h) * sprite.RenderPivot.Y

		if cameraTarget == nil && int(position.Value.X+float64(w)) < minX || int(position.Value.Y+float64(h)) < minY || int(position.Value.X-float64(w/2)) > maxX || int(position.Value.Y-float64(h/2)) > maxY {
			// println("fora")
			continue
		}

		op := cameraComponent.Camera.GetTranslation(0, 0)

		if cameraTarget != nil {
			op = cameraComponent.Camera.GetTranslation(position.Value.X, position.Value.Y)
		} else {
			op.GeoM.Translate(position.Value.X, position.Value.Y)
		}

		op.GeoM.Translate(-renderW, -renderH)

		spriteFrameBounds := sprite.GetFrameBounds()
		cameraComponent.Camera.Surface.DrawImage(sprite.Image.SubImage(spriteFrameBounds).(*ebiten.Image), op)
	}

	cameraComponent.Camera.Blit(screen)

}
