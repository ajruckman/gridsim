package schema

import (
    "github.com/ajruckman/abelian/internal/gridlib"
)

type Particle struct {
    CellPosition gridlib.Vector
    CellVal      gridlib.CellVal
}

func (p Particle) Pos() gridlib.Vector {
    return p.CellPosition
}
func (p Particle) Val() gridlib.CellVal {
    return p.CellVal
}


