package gridlib

import (
    "fmt"
    "strings"
)

// Check returns true if the vector passed is contained by the lattice.
// Generally, this function shouldn't be neccessary; vectors should be wrapped
// which guarantees they are contained by the lattice.
func (l Lattice) Check(v Vec) bool {
    if 0 <= v.X && v.X < l.Size.X {
        if 0 <= v.Y && v.Y < l.Size.Y {
            return true
        }
    }
    return false
}

// Wrap returns a vector with coordinates that are outside of the lattice
// wrapped to the other side of the lattice.
func (l Lattice) Wrap(v Vec) Vec {
    return Vec{
        X: Mod(v.X, l.Size.X),
        Y: Mod(v.Y, l.Size.Y),
    }
}

// Moore returns positions of cells in the Moore neighborhood around the vector
// passed.
func (l Lattice) Moore(v Vec, distance int) (coords VecSet) {
    set := func(px, py, d int) {
        coords = append(coords, l.Wrap(Vec{px, py}))
    }

    for d := 1; d < distance+1; d++ {
        px, py := v.X-d, v.Y-d

        for i := 0; i < d*2; i++ {
            set(px, py, d)
            px++
        }

        for i := 0; i < d*2; i++ {
            set(px, py, d+20)
            py++
        }

        for i := 0; i < d*2; i++ {
            set(px, py, d+30)
            px--
        }

        for i := 0; i < d*2; i++ {
            set(px, py, d+40)
            py--
        }
    }

    return
}

// VonNeumann returns positions of cells in the Von Neumann neighborhood around
// the vector passed.
func (l *Lattice) VonNeumann(v Vec, distance int) (coords VecSet) {
    set := func(px, py, d int) {
        coords = append(coords, l.Wrap(Vec{px, py}))
    }

    for d := 0; d < distance+1; d++ {
        px, py := v.X-d, v.Y

        for i := 0; i < d; i++ {
            set(px, py, d+10)
            px++
            py++
        }

        for i := 0; i < d; i++ {
            set(px, py, d+20)
            px++
            py--
        }

        for i := 0; i < d; i++ {
            set(px, py, d+30)
            px--
            py--
        }

        for i := 0; i < d; i++ {
            set(px, py, d+40)
            px--
            py++
        }

    }

    return
}

// Index wraps the vector passed and returns the cell located at the wrapped
// vector.
func (l Lattice) Index(v Vec) Cell {
    t := l.Wrap(v)

    return l.Grid[t.Y][t.X]
}

// Print prints a prettified lattice to the console.
func (l Lattice) Print() {
    fmt.Printf("┌%s┐\n", strings.Repeat("─", l.Size.X*4+2))
    for _, y := range l.Grid {
        fmt.Print("│ ")
        for _, x := range y {
            s := x.String()
            if x == nil {
                fmt.Printf("·   ")
            } else {
                fmt.Printf("%-4s", s)
            }
        }
        fmt.Println(" │")
    }
    fmt.Printf("└%s┘\n", strings.Repeat("─", l.Size.X*4+2))
}
