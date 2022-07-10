package scenes

import (
	"gogame/components"
	coreMath "gogame/core/math"
	"gogame/entities"
	"gogame/systems"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/solarlune/resolv"
)

type MainScene struct{}

func (s *MainScene) Setup(w engine.World) {
	tileMap := loadMap()
	space := resolv.NewSpace(tileMap.Width*tileMap.TileWidth, tileMap.Height*tileMap.TileHeight, tileMap.TileWidth, tileMap.TileHeight)
	playerPos := coreMath.MakeVector2(16+8, 8)
	playerShape := resolv.NewObject(playerPos.X, playerPos.Y, 16, 8)
	space.Add(playerShape)

	for _, og := range tileMap.ObjectGroups {
		for _, ob := range og.Objects {
			if ob.Type == "start" {
				playerPos.X = ob.X
				playerPos.Y = ob.Y
			}
		}
	}

	w.AddComponents(
		components.CameraTarget{},
		components.InputMovement{},
		components.Position{},
		components.Sprite{},
		components.Camera{},
		components.SpriteAnimation{},
		components.MapTile{},
		components.CollisionShape{},
		components.CollisionSpace{},
	)

	makeMapTileEntity(w, tileMap, space)
	w.AddEntities(
		&entities.CollisionSpace{
			CollisionSpace: components.CollisionSpace{
				Space: space,
			},
		},
		// &entities.MapTile{
		// 	MapTile:  components.MapTile{},
		// 	Position: components.Position{Value: playerPos},
		// 	Sprite: components.Sprite{
		// 		Image:       loadTileSetHouseSprite(),
		// 		Frame:       0,
		// 		HFrames:     7,
		// 		VFrames:     7,
		// 		RenderPivot: coreMath.VECTOR_ZERO,
		// 		SortPivot:   coreMath.VECTOR_ONE.MultScalar(0.5),
		// 	},
		// },
		&entities.Player{
			CollisionShape: components.CollisionShape{Shape: playerShape, Static: false},
			InputMovement:  components.InputMovement{Stack: []coreMath.Vector2{coreMath.VECTOR_ZERO}},
			Position:       components.Position{Value: playerPos},
			Sprite: components.Sprite{
				Image:       loadPlayerSprite(),
				Frame:       0,
				HFrames:     4,
				VFrames:     7,
				RenderPivot: coreMath.VECTOR_ONE.MultScalar(0.5),
				SortPivot:   coreMath.MakeVector2(0.5, 0.5),
			},
			CameraTarget: components.CameraTarget{},
			SpriteAnimation: components.SpriteAnimation{
				CurrentAnimation: "IDLE_DOWN",
				LastTime:         0.0,
				Animations: map[string]*components.Animation{
					"IDLE_DOWN": {
						Frames:       []int{0},
						TimeStep:     10,
						CurrentFrame: 0,
						InitialFrame: 0,
					},
					"IDLE_UP": {
						Frames:       []int{1},
						TimeStep:     10,
						CurrentFrame: 0,
						InitialFrame: 0,
					},
					"IDLE_LEFT": {
						Frames:       []int{2},
						TimeStep:     10,
						CurrentFrame: 0,
						InitialFrame: 0,
					},
					"IDLE_RIGHT": {
						Frames:       []int{3},
						TimeStep:     10,
						CurrentFrame: 0,
						InitialFrame: 0,
					},
					"WALK_DOWN": {
						Frames:       []int{0, 4, 8, 12},
						TimeStep:     .1,
						CurrentFrame: 0,
						InitialFrame: 1,
					},
					"WALK_UP": {
						Frames:       []int{1, 5, 9, 13},
						TimeStep:     .1,
						CurrentFrame: 0,
						InitialFrame: 1,
					},
					"WALK_LEFT": {
						Frames:       []int{2, 6, 10, 14},
						TimeStep:     .1,
						CurrentFrame: 0,
						InitialFrame: 1,
					},
					"WALK_RIGHT": {
						Frames:       []int{3, 7, 11, 15},
						TimeStep:     .1,
						CurrentFrame: 0,
						InitialFrame: 1,
					},
				},
			},
		},
		&entities.Camera{Camera: components.Camera{Camera: camera.NewCamera(1920/6, 1080/6, 0, 0, 0, 1)}},
	)
	w.AddSystems(
		&systems.KeyboardMovement{},
		&systems.Movement{},
		&systems.Collision{},
		&systems.SpriteAnimation{},
		&systems.Render{},
	)
}

func loadPlayerSprite() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/player.png")

	if err != nil {
		log.Fatal(err)
	}

	return img
}

// func loadTileSetHouseSprite() *ebiten.Image {
// 	img, _, err := ebitenutil.NewImageFromFile("assets/TilesetHouse.png")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return img
// }

func loadMap() *tiled.Map {
	gameMap, _ := tiled.LoadFile("assets/sampleMap.tmx")

	return gameMap
}

func makeMapTileEntity(w engine.World, tileMap *tiled.Map, space *resolv.Space) {
	layerImage, _, _ := ebitenutil.NewImageFromFile("assets/tilemap.png")

	for _, og := range tileMap.ObjectGroups {
		if og.Name != "Collisions" {
			continue
		}

		for _, o := range og.Objects {
			shape := resolv.NewObject(o.X, o.Y, o.Width, o.Height)
			space.Add(shape)
			w.AddEntities(
				&entities.Collider{
					CollisionShape: components.CollisionShape{Shape: shape, Static: true},
				},
			)
		}
	}

	for _, layer := range tileMap.Layers {
		if !layer.Visible {
			continue
		}

		for y := 0; y < tileMap.Height; y++ {
			for x := 0; x < tileMap.Width; x++ {

				tile := layer.Tiles[y*tileMap.Width+x]

				if tile.IsNil() {
					continue
				}

				// if tile.ID == 148 || tile.ID == 150 {
				// 	mergedRect := image.Rectangle{Min: image.Pt(0, 0), Max: image.Pt(16*3, 16)}
				// 	merged := image.NewRGBA(mergedRect)

				// 	if tile.ID == 148 {
				// 		for i, x := range []uint32{148, 157, 166} {
				// 			spriteRect := tile.Tileset.GetTileRect(x)
				// 			tileImage := ebiten.NewImageFromImage(layerImage.SubImage(spriteRect))
				// 			draw.Draw(merged, image.Rectangle{Min: image.Pt(i*16, 0), Max: image.Pt((i*16)+16, 16)}, tileImage, image.Pt(0, 0), draw.Src)
				// 		}
				// 	}

				// 	if tile.ID == 150 {
				// 		for i, x := range []uint32{150, 159, 168, 177} {
				// 			spriteRect := tile.Tileset.GetTileRect(x)
				// 			tileImage := ebiten.NewImageFromImage(layerImage.SubImage(spriteRect))
				// 			draw.Draw(merged, image.Rectangle{Min: image.Pt(i*16, 0), Max: image.Pt((i*16)+16, 16)}, tileImage, image.Pt(0, 0), draw.Src)
				// 		}
				// 	}

				// 	w.AddEntities(
				// 		&entities.AnimatedMapTile{
				// 			MapTile:  components.MapTile{},
				// 			Position: components.Position{Value: coreMath.MakeVector2(float64(x*tileMap.TileWidth), float64(y*tileMap.TileHeight))},
				// 			Sprite: components.Sprite{
				// 				Image:       ebiten.NewImageFromImage(merged),
				// 				Frame:       0,
				// 				HFrames:     3,
				// 				VFrames:     1,
				// 				RenderPivot: coreMath.VECTOR_ZERO,
				// 				SortPivot:   coreMath.VECTOR_ZERO,
				// 			},
				// 			SpriteAnimation: components.SpriteAnimation{
				// 				Animations: map[string]*components.Animation{
				// 					"DEFAULT": {
				// 						Frames:       []int{0, 1, 2},
				// 						TimeStep:     0.250,
				// 						CurrentFrame: 0,
				// 						InitialFrame: 0,
				// 					},
				// 				},
				// 				CurrentAnimation: "DEFAULT",
				// 				LastTime:         0,
				// 			},
				// 		},
				// 	)
				// } else {
				spriteRect := tile.Tileset.GetTileRect(tile.ID)
				tileImage := ebiten.NewImageFromImage(layerImage.SubImage(spriteRect))

				w.AddEntities(
					&entities.MapTile{
						MapTile:  components.MapTile{},
						Position: components.Position{Value: coreMath.MakeVector2(float64(x*tileMap.TileWidth), float64(y*tileMap.TileHeight))},
						Sprite: components.Sprite{
							Image:       tileImage,
							Frame:       0,
							HFrames:     1,
							VFrames:     1,
							RenderPivot: coreMath.VECTOR_ZERO,
							SortPivot:   coreMath.VECTOR_ZERO,
						},
					},
				)
				// }
			}
		}
	}
}
