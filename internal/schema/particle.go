package schema

import (
    "github.com/ajruckman/abelian/internal/gridlib"
)

type ParticleSim struct {
    Lattice gridlib.Lattice
}

type Particle struct {
    Value  int
    Static bool
}

func (p *Particle) Val() int {
    return p.Value
}

func (p *Particle) Set(v int) {
    p.Value = v
}

func (p *ParticleSim) Init(size gridlib.Vec) {
    p.Lattice.Size = size
    p.Lattice.Init(true, 2, 1)
}

func (p *ParticleSim) Tick() {
    p.Lattice.BeginTransaction()

    for iy := 0; iy < p.Lattice.Size.Y; iy++ {
        for ix := 0; ix < p.Lattice.Size.X; ix++ {
            var (
                this   = gridlib.Vec{ix, iy}
                cur    = p.Lattice.Index(this)
                coords = p.Lattice.Moore(this, 1)
                s      = p.Lattice.Sum(coords)
            )

            if cur.Val() == 0 {
                if s == 3 {
                    p.Lattice.Set(this, 1)
                }
            } else {
                switch {
                case s < 2:
                    p.Lattice.Set(this, 0)
                case s < 4:
                    p.Lattice.Set(this, 1)
                default:
                    p.Lattice.Set(this, 0)
                }
            }

        }
    }

    p.Lattice.CommitSet()
}
