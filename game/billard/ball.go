package billard

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BallType int

const (
	FULL BallType = iota
	STRIPED
	BLACK
	WHITE
)

type Ball struct {
	Model    rl.Model
	Position rl.Vector3
	Rotation rl.Vector3
	Type     BallType
}

func NewBall(model string, ballType BallType) *Ball {

	var loadModel rl.Model = rl.LoadModel(model)
	loadModel.Transform = rl.MatrixRotateXYZ(rl.NewVector3(0, 0, 0))

	return &Ball{
		Model:    loadModel,
		Position: rl.NewVector3(0, 3.2, 0),
		Rotation: rl.NewVector3(0, 0, 0),
		Type:     ballType,
	}
}

func (b *Ball) Draw() {
	rl.DrawModel(b.Model, b.Position, 1.0, rl.White)
}

func (b *Ball) Unload() {
	rl.UnloadModel(b.Model)
}
