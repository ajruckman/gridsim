package particle

import (
    "fmt"
    "golang.org/x/image/colornames"
    "image/color"
    "math/rand"

    "github.com/ajruckman/gridsim/internal/gridlib"
)

type Particle struct {
    Value    int
    Static   bool
    HasMoved bool
    Color    *color.RGBA
}

func (p Particle) String() *string {
    if p.Value == 0 {
        return nil
    } else {
        s := fmt.Sprintf("%d", p.Value)
        return &s
    }
}

func (p Particle) Display() bool {
    return p.Value != 0
}

func (p Particle) GetColor() color.Color {
    if p.Color == nil {
        return colornames.Green
    } else {
        return *p.Color
    }
}

func (p *Particle) Set(v int) {
    p.Value = v
}

/////

type Sim struct {
    Lattice gridlib.Lattice
}

func (p *Sim) Init(size gridlib.Vec) {
    p.Lattice.Size = size
    p.Lattice.Grid = NewGrid(size, true, 2, 4)

    v := gridlib.Vec{X: size.X / 2, Y: size.Y / 16}

    p.IndexGrid(v).Static = true
    p.IndexGrid(v).Color = &colornames.Blue
    p.IndexGrid(v).Value = 1
}

/////

// We define IndexGrid on the implementor so we can cast it to the right type
func (p Sim) IndexGrid(v gridlib.Vec) *Particle {
    return p.Lattice.Grid[v.Y][v.X].(*Particle)
}

func (p Sim) IndexOverlay(v gridlib.Vec) *Particle {
    return p.Lattice.Overlay[v.Y][v.X].(*Particle)
}

func (p Sim) Sum(v gridlib.VecSet) (sum int) {
    for _, v := range v {
        sum += p.IndexGrid(v).Value
    }
    return
}

/////

func NewGrid(size gridlib.Vec, randomize bool, val, rarity int) (grid [][]gridlib.Cell) {
    grid = make([][]gridlib.Cell, size.Y)

    for iy := 0; iy < size.Y; iy++ {
        grid[iy] = make([]gridlib.Cell, size.X)

        for ix := 0; ix < size.X; ix++ {
            if randomize {
                var v int
                v = rand.Intn(val)

                for i := 0; i < rarity; i++ {
                    v &= rand.Intn(val)
                }

                grid[iy][ix] = &Particle{Value: v}
            } else {
                grid[iy][ix] = &Particle{Value: val}
            }
        }
    }
    return
}
