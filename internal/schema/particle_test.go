package schema

import (
    "github.com/ajruckman/abelian/internal/gridlib"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestParticle_VonNeumann(t *testing.T) {
    var l = gridlib.Lattice{
        Lattice: [][]int{
            {0, 0, 0, 0, 0, 1, 1},
            {0, 1, 0, 0, 0, 1, 0},
            {0, 1, 0, 0, 1, 0, 0},
            {0, 0, 1, 0, 1, 0, 0},
            {0, 0, 0, 0, 0, 0, 0},
            {0, 0, 1, 1, 0, 0, 0},
            {0, 0, 0, 0, 1, 1, 1},
        },
    }

    l.SetBounds()

    var p Particle
    _ = p

    for y := l.Min.Y - 10; y <= l.Max.Y+10; y++ {
        for x := l.Min.X - 10; x <= l.Max.X+10; x++ {
            if y < l.Min.Y || y > l.Max.Y {
                assert.False(t, l.Contains(gridlib.Vector{x, y}))
            } else if x < l.Min.X || x > l.Max.X {
                assert.False(t, l.Contains(gridlib.Vector{x, y}))
            } else {
                assert.True(t, l.Contains(gridlib.Vector{x, y}))
            }
        }
    }

    p.CellPosition = gridlib.Vector{0, 0}
    for i := 0; i < 7; i++ {
        v := l.VonNeumann(p, i)
        assert.Equal(t, (i+1)*(i+1)-1, len(v))
    }

    p.CellPosition = gridlib.Vector{3, 3}
    for i := 0; i < 10; i++ {
        e := (i*2+1)*(i*2+1) - 1
        if e > l.Size.X * l.Size.Y - 1 {
            e = l.Size.X * l.Size.Y - 1
        }
        v := l.VonNeumann(p, i)
        assert.Equal(t, e, len(v))
    }

    p.CellPosition = gridlib.Vector{6, 6}
    for i := 0; i < 10; i++ {
        e := (i+1)*(i+1) - 1
        if e > l.Size.X * l.Size.Y - 1 {
            e = l.Size.X * l.Size.Y - 1
        }
        v := l.VonNeumann(p, i)
        assert.Equal(t, e, len(v))
    }

    p.CellPosition = gridlib.Vector{3, 3}
    for i := 0; i < 10; i++ {
        e := (i*2+1)*(i*2+1) - 1
        if e > l.Size.X * l.Size.Y - 1 {
           e = l.Size.X * l.Size.Y - 1
        }
        v := l.VonNeumann(p, i)
        assert.Equal(t, e, len(v))
    }
}
