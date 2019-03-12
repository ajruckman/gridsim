package gridlib

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestLattice_Print(t *testing.T) {
    var l = Lattice{Size: Vec{8, 25}}
    l.Init(true, 10, 3)

    l.Print()
}

func TestLattice_Set(t *testing.T) {
    var l = Lattice{Size: Vec{8, 25}}
    l.Init(false, 0, -1)

    l.BeginTick()
    l.Set(Vec{100, 3}, -50)
    l.Set(Vec{4, 5}, 1)
    l.Set(Vec{7, 23}, 2)
    l.Set(Vec{0, 24}, 4)
    l.Set(Vec{0, 0}, 8)
    l.EndTickSet()

    l.Print()

    coords := l.Moore(Vec{5, 5}, 35)
    s := l.Sum(coords)

    assert.Equal(t, 15, s)
}

func TestLattice_Check(t *testing.T) {
    var l = Lattice{Size: Vec{8, 25}}
    l.Init(false, 0, -1)

    for x := -10; x < l.Size.X+10; x++ {
        for y := -10; y < l.Size.Y+10; y++ {
            if (x < 0 || x >= l.Size.X) || (y < 0 || y >= l.Size.Y) {
                assert.False(t, l.Check(Vec{x, y}),
                    fmt.Sprintf("%5d %d", x, y))
            } else {
                assert.True(t, l.Check(Vec{x, y}),
                    fmt.Sprintf("%5d %d", x, y))
            }
        }
    }
}

func checkCoords(t *testing.T, l Lattice, coords VecSet) {
    for i, v := range coords {
        i++
        l.Grid[v.Y][v.X] = i
    }

    l.Print()

    n := len(coords)
    e := (n*n + n) / 2 // nth triangle number
    s := l.Sum(coords)

    assert.Equal(t, e, s)

    fmt.Println("Count:    ", n)
    fmt.Println("Expected: ", e)
    fmt.Println("Sum:      ", s)
}

func TestLattice_Moore(t *testing.T) {
    var l = Lattice{Size: Vec{8, 25}}
    l.Init(false, 0, -1)

    l.Grid[5][5] = -50

    coords := l.Moore(Vec{5, 5}, 5)

    checkCoords(t, l, coords)
}

func TestLattice_VonNeumann(t *testing.T) {
    var l = Lattice{Size: Vec{8, 25}}
    l.Init(false, 0, -1)

    l.Grid[5][5] = -50

    coords := l.VonNeumann(Vec{5, 5}, 5)

    checkCoords(t, l, coords)
}
