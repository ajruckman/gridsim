package life

import (
    "fmt"
    "github.com/ajruckman/gridsim/internal/gridlib"
    "golang.org/x/image/colornames"
    "image/color"
    "math/rand"
)

type Life struct {
    Value int
}

func (p Life) String() *string {
    if p.Value == 0 {
        return nil
    } else {
        s := fmt.Sprintf("%d", p.Value)
        return &s
    }
}

func (p Life) Display() bool {
    return p.Value != 0
}

func (p Life) GetColor() color.Color {
    return colornames.Red
}

func (p *Life) Set(v int) {
    p.Value = v
}

/////

type Sim struct {
    Lattice gridlib.Lattice
}

func (p *Sim) Init(size gridlib.Vec) {
    p.Lattice.Size = size
    p.Lattice.Grid = NewGrid(size, true, 2, 1)
}

/////

// We define IndexGrid on the implementor so we can cast it to the right type
func (p Sim) IndexGrid(v gridlib.Vec) *Life {
    return p.Lattice.IndexGrid(v).(*Life)
    //return p.Lattice.Grid[v.Y][v.X].(*Life)
}

func (p Sim) IndexOverlay(v gridlib.Vec) *Life {
    return p.Lattice.IndexOverlay(v).(*Life)
    //return p.Lattice.Overlay[v.Y][v.X].(*Life)
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

                grid[iy][ix] = &Life{Value: v}
            } else {
                grid[iy][ix] = &Life{Value: val}
            }
        }
    }
    return
}
