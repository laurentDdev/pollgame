package billard

import rl "github.com/gen2brain/raylib-go/raylib"

type Billard struct {
	Model    rl.Model
	Position rl.Vector3
}

func NewBillard(model string) *Billard {

	var loadModel rl.Model = rl.LoadModel(model)
	loadModel.Transform = rl.MatrixRotateXYZ(rl.NewVector3(0, 0, 0))

	return &Billard{
		Model:    loadModel,
		Position: rl.NewVector3(0, 2.5, 0),
	}
}

func (b *Billard) Unload() {
	rl.UnloadModel(b.Model)
}
