package life

import (
    "github.com/ajruckman/gridsim/internal/gridlib"
)

func (p *Sim) BeginTick() {
    p.Lattice.Overlay = NewGrid(p.Lattice.Size, false, 0, 0)
}

func (p *Sim) Tick() {
    for iy := 0; iy < p.Lattice.Size.Y; iy++ {
        for ix := 0; ix < p.Lattice.Size.X; ix++ {
            var (
                this      = gridlib.Vec{ix, iy}
                cur       = p.IndexGrid(this).Value
                neighbors = p.Lattice.Moore(this, 1, true)
                num       = p.Sum(neighbors)
            )

            if cur == 0 {
                if num == 3 {
                    p.IndexOverlay(this).Value = 1
                }
            } else {
                switch {
                case num < 2:
                    p.IndexOverlay(this).Value = 0
                case num < 4:
                    p.IndexOverlay(this).Value = 1
                default:
                    p.IndexOverlay(this).Value = 0
                }
            }
        }
    }
}

func (p *Sim) EndTick() {
    p.Lattice.Grid = p.Lattice.Overlay
}
