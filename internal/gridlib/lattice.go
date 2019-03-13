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

func (l Lattice) Corners(v Vec, distance int, inclusive bool) (coords VecSet) {
    set := func(px, py, d int) {
        coords = append(coords, l.Wrap(Vec{px, py}))
    }

    var d int
    if inclusive {
        d = 1
    } else {
        d = distance
    }
    for d = d; d < distance+1; d++ {
        px, py := v.X-d, v.Y-d

        set(px, py, d)
        px += d * 2

        set(px, py, d)
        py += d * 2

        set(px, py, d)
        px -= d * 2

        set(px, py, d)
        px -= d * 2
    }

    return
}

// Moore returns positions of cells in the Moore neighborhood around the vector
// passed.
func (l Lattice) Moore(v Vec, distance int, inclusive bool) (coords VecSet) {
    set := func(px, py, d int) {
        coords = append(coords, l.Wrap(Vec{px, py}))
    }

    var d int
    if inclusive {
        d = 1
    } else {
        d = distance
    }
    for d = d; d < distance+1; d++ {
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
func (l *Lattice) VonNeumann(v Vec, distance int, inclusive bool) (coords VecSet) {
    set := func(px, py, d int) {
        coords = append(coords, l.Wrap(Vec{px, py}))
    }

    var d int
    if inclusive {
        d = 1
    } else {
        d = distance
    }
    for d = d; d < distance+1; d++ {
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

// IndexGrid wraps the vector passed and returns the cell located at the wrapped
// vector.
func (l Lattice) IndexGrid(v Vec) Cell {
    t := l.Wrap(v)

    return l.Grid[t.Y][t.X]
}

// IndexOverlay wraps the vector passed and returns the cell located at the
// wrapped vector.
func (l Lattice) IndexOverlay(v Vec) Cell {
    t := l.Wrap(v)

    return l.Overlay[t.Y][t.X]
}

// Print returns a prettified lattice string.
func (l Lattice) String() (res string) {

    res += fmt.Sprintf("┌%s┐\n", strings.Repeat("─", l.Size.X*4+2))
    for _, y := range l.Grid {
        res += "│ "
        for _, x := range y {
            s := x.String()
            if s == nil {
                res += "·   "
            } else {
                res += fmt.Sprintf("%-4s", *s)
            }
        }
        res += " │\n"
    }
    res += fmt.Sprintf("└%s┘\n", strings.Repeat("─", l.Size.X*4+2))

    return
}
