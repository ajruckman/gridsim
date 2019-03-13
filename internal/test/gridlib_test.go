package test

import (
    "fmt"
    "github.com/ajruckman/gridsim/internal/gridlib"
    "github.com/ajruckman/gridsim/internal/simulators/life"
    "github.com/stretchr/testify/assert"
    "hash/fnv"
    "testing"
)

var (
    size    = gridlib.Vec{8, 25}
    lifesim = life.Sim{Lattice: gridlib.Lattice{Size: size}}
)

func init() {
    lifesim.Lattice.Grid = NewGrid(lifesim.Lattice.Size)
}

func NewGrid(size gridlib.Vec) (grid [][]gridlib.Cell) {
    grid = make([][]gridlib.Cell, size.Y)

    cur := 0
    for iy := 0; iy < size.Y; iy++ {
        grid[iy] = make([]gridlib.Cell, size.X)

        for ix := 0; ix < size.X; ix++ {
            grid[iy][ix] = &life.Life{Value: cur}

            cur += 1
        }
    }
    return
}

func TestLattice_Print(t *testing.T) {
    h := fnv.New32a()

    _, err := h.Write([]byte(lifesim.Lattice.String()))
    assert.Equal(t, nil, err)
    assert.Equal(t, uint32(3127130728), h.Sum32())

    fmt.Println(lifesim.Lattice.String())
}

func TestLattice_Moore(t *testing.T) {
    var (
        coords gridlib.VecSet
        s      int
    )

    coords = lifesim.Lattice.Moore(lifesim.Lattice.Wrap(gridlib.Vec{100, 3}), 10, false)
    s = lifesim.Sum(coords)
    assert.Equal(t, 8824, s)

    coords = lifesim.Lattice.Moore(lifesim.Lattice.Wrap(gridlib.Vec{100, 3}), 10, true)
    s = lifesim.Sum(coords)
    assert.Equal(t, 41552, s)
}

func TestLattice_Corners(t *testing.T) {
    var (
        coords gridlib.VecSet
        s      int
    )

    coords = lifesim.Lattice.Corners(lifesim.Lattice.Wrap(gridlib.Vec{100, 3}), 10, false)
    s = lifesim.Sum(coords)
    assert.Equal(t, 512, s)

    coords = lifesim.Lattice.Corners(lifesim.Lattice.Wrap(gridlib.Vec{100, 3}), 10, true)
    s = lifesim.Sum(coords)
    assert.Equal(t, 3904, s)
}

func TestLattice_VonNeumann(t *testing.T) {
    var (
        coords gridlib.VecSet
        s      int
    )

    coords = lifesim.Lattice.VonNeumann(lifesim.Lattice.Wrap(gridlib.Vec{100, 3}), 10, false)
    s = lifesim.Sum(coords)
    assert.Equal(t, 3704, s)

    coords = lifesim.Lattice.VonNeumann(lifesim.Lattice.Wrap(gridlib.Vec{100, 3}), 10, true)
    s = lifesim.Sum(coords)
    assert.Equal(t, 15856, s)
}

func TestLattice_Check(t *testing.T) {
    for x := -10; x < lifesim.Lattice.Size.X+10; x++ {
        for y := -10; y < lifesim.Lattice.Size.Y+10; y++ {
            if (x < 0 || x >= lifesim.Lattice.Size.X) || (y < 0 || y >= lifesim.Lattice.Size.Y) {
                assert.False(t, lifesim.Lattice.Check(gridlib.Vec{x, y}),
                    fmt.Sprintf("%5d %d", x, y))
            } else {
                assert.True(t, lifesim.Lattice.Check(gridlib.Vec{x, y}),
                    fmt.Sprintf("%5d %d", x, y))
            }
        }
    }
}
