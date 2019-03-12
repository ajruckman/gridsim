package particle

import (
    "math/rand"

    "golang.org/x/image/colornames"

    "github.com/ajruckman/gridsim/internal/gridlib"
)

func (p *Sim) BeginTick() {
    //p.Lattice.Overlay = NewGrid(p.Lattice.Size, false, 0, 0)
}

func r(rarity int) (v bool) {
    v = rand.Intn(2) == 1
    for i := 0; i < rarity; i++ {
        if v {
            v = rand.Intn(2) == 1
        } else {
            return
        }
    }
    return
}

func move(vec gridlib.Vec) (res gridlib.Vec) {
    if r(0) {
        if r(0) {
            res.X = vec.X + 1
        } else {
            res.X = vec.X - 1
        }
        res.Y = vec.Y
    } else {
        if r(2) {
            res.Y = vec.Y + 1
        } else {
            res.Y = vec.Y - 1
        }
        res.X = vec.X
    }

    return
}

func (p *Sim) Tick() {
    for iy := 0; iy < p.Lattice.Size.Y; iy++ {
        for ix := 0; ix < p.Lattice.Size.X; ix++ {

            var (
                this = gridlib.Vec{ix, iy}
                cur  = p.IndexGrid(this).Value
            )

            if cur == 0 {
                continue
            } else if p.IndexGrid(this).Static {
                continue
            } else if p.IndexGrid(this).HasMoved {
                continue
            }

            var (
                neighbors = p.Lattice.VonNeumann(this, 2)
                num       = p.Sum(neighbors)
                n         = this
                _         = num
            )

            if num > 0 {
                for _, v := range neighbors {
                    if p.IndexGrid(v).Static {
                        p.IndexGrid(this).Value = 1
                        p.IndexGrid(this).Static = true
                        p.IndexGrid(this).Color = &colornames.Red
                        goto next
                    }
                }
            }

            n = move(this)
            n = p.Lattice.Wrap(n)

            if p.IndexGrid(n).Value == 0 {
                p.IndexGrid(this).Value = 0
                p.IndexGrid(n).Value = 1
                p.IndexGrid(n).HasMoved = true
            }

            goto next
        next:
        }
    }

    for iy := 0; iy < p.Lattice.Size.Y; iy++ {
        for ix := 0; ix < p.Lattice.Size.X; ix++ {
            p.IndexGrid(gridlib.Vec{ix, iy}).HasMoved = false
        }
    }
}

func (p *Sim) EndTick() {
    //p.Lattice.Grid = p.Lattice.Overlay
}
