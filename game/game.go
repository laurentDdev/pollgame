package game

import (
	"billarggame/game/billard"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Game struct {
	Billard *billard.Billard
	Balls   []*billard.Ball
}

func NewGame() *Game {
	return &Game{
		Billard: billard.NewBillard("assets/models/billard/billard.gltf"),
		Balls: []*billard.Ball{
			billard.NewBall("assets/models/balls/fullyellow/fullyellow.gltf", billard.FULL),
			billard.NewBall("assets/models/balls/fullblue/fullblue.gltf", billard.FULL),
			billard.NewBall("assets/models/balls/fullred/fullred.gltf", billard.FULL),
			billard.NewBall("assets/models/balls/fullviolet/fullviolet.gltf", billard.FULL),
			billard.NewBall("assets/models/balls/fullorange/fullorange.gltf", billard.FULL),
			billard.NewBall("assets/models/balls/fullgreen/fullgreen.gltf", billard.FULL),
			billard.NewBall("assets/models/balls/fullbrown/fullbrown.gltf", billard.FULL),
			billard.NewBall("assets/models/balls/fullblack/fullblack.gltf", billard.BLACK),
			billard.NewBall("assets/models/balls/stripedyellow/stripedyellow.gltf", billard.STRIPED),
			billard.NewBall("assets/models/balls/stripedblue/stripedblue.gltf", billard.STRIPED),
			billard.NewBall("assets/models/balls/stripedred/stripedred.gltf", billard.STRIPED),
			billard.NewBall("assets/models/balls/stripedviolet/stripedviolet.gltf", billard.STRIPED),
			billard.NewBall("assets/models/balls/stripedorange/stripedorange.gltf", billard.STRIPED),
			billard.NewBall("assets/models/balls/stripedgreen/stripedgreen_.gltf", billard.STRIPED),
			billard.NewBall("assets/models/balls/stripedbrown/stripedbrown.gltf", billard.STRIPED),
			billard.NewBall("assets/models/balls/fullwhite/fullwhite.gltf", billard.WHITE),
		},
	}
}

func rotatePosition(x, z, angle float32) (float32, float32) {
	// Convertir l'angle en radians
	rad := angle * (math.Pi / 180)
	// Appliquer la rotation 2D autour de l'axe Y
	rotatedX := x*float32(math.Cos(float64(rad))) - z*float32(math.Sin(float64(rad)))
	rotatedZ := x*float32(math.Sin(float64(rad))) + z*float32(math.Cos(float64(rad)))
	return rotatedX, rotatedZ
}

func placeBallsWithRotation(balls []*billard.Ball) {
	var startX float32 = 0
	var startZ float32 = 3
	var radius float32 = 0.2         // Rayon de la boule
	var spacing float32 = radius * 2 // Espacement entre les boules

	// Initialiser le compteur de boules
	count := 0
	for row := 0; row < 5; row++ {
		for col := 0; col <= row; col++ {
			if count >= len(balls)-1 { // Exclure la boule blanche
				break
			}

			// Calculer la position de chaque boule dans le triangle
			x := startX + float32(col)*spacing - float32(row)*spacing/2
			z := startZ + float32(row)*spacing
			rotatedX, rotatedZ := rotatePosition(x, z, 90)

			// Positionner la boule
			balls[count].Position = rl.NewVector3(rotatedX, 3.2, rotatedZ)
			count++
		}
	}
}

func (g *Game) Init() {
	// Initialize ball positions

	placeBallsWithRotation(g.Balls)
	g.Balls[len(g.Balls)-1].Position = rl.NewVector3(3, 3.2, 0)
}

func (g *Game) Update() {
	rl.DrawModel(g.Billard.Model, g.Billard.Position, 1.0, rl.White)

	for _, ball := range g.Balls {
		ball.Draw()
	}

}

func (g *Game) Unload() {
	g.Billard.Unload()

	for _, ball := range g.Balls {
		ball.Unload()
	}

}
