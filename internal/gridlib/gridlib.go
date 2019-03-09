package gridlib

import (
    "fmt"
    "math/rand"
    "strings"
)

type Simulation interface {
    Init(Vec)
    Tick()
}

type Cell interface {
    Val() int
    Set(int)
}

type Vec struct {
    X, Y int
}

type Lattice struct {
    Grid [][]Cell
    Size Vec

    overlay [][]Cell
}

type Coords []Vec

func (l *Lattice) Init(randomize bool, val, rarity int) {
    for y := 0; y < l.Size.Y; y++ {
        l.Grid = append(l.Grid, []Cell{})
        for x := 0; x < l.Size.X; x++ {
            var v Cell

            if randomize {
                v.Set(rand.Intn(val))
                for i := 0; i < rarity; i++ {
                    v.Set(v.Val() & rand.Intn(val))
                }
            } else {
                v.Set(val)
            }

            l.Grid[y] = append(l.Grid[y], v)
        }
    }
}

// Clear overlay
func (l *Lattice) BeginTransaction() {
    l.overlay = make([][]Cell, l.Size.Y)
    for yi := 0; yi < l.Size.Y; yi++ {
        l.overlay[yi] = make([]Cell, l.Size.X)
    }
}

func (l *Lattice) Set(v Vec, n int) {
    //if l.Check(v) {
    //    l.overlay[v.Y][v.X] = n
    //}
    t := l.Wrap(v)
    l.overlay[t.Y][t.X].Set(n)
}

func (l *Lattice) Add(v Vec, n int) {
    //if l.Check(v) {
    //    l.overlay[v.Y][v.X] += n
    //}
    t := l.Wrap(v)
    l.overlay[t.Y][t.X].Set(l.overlay[t.Y][t.X].Val() + n)
}

// Replace grid with overlay
func (l *Lattice) CommitSet() {
    l.Grid = l.overlay
}

// Add overlay cells to grid cells
func (l *Lattice) CommitAdd() {
    for yi := 0; yi < l.Size.Y; yi++ {
        for xi := 0; xi < l.Size.X; xi++ {
            l.Grid[yi][xi].Set(l.Grid[yi][xi].Val() + l.overlay[yi][xi].Val())
        }
    }
}

// Check if coordinates are inside grid
func (l Lattice) Check(v Vec) bool {
    if 0 <= v.X && v.X < l.Size.X {
        if 0 <= v.Y && v.Y < l.Size.Y {
            return true
        }
    }
    return false
}

func (l Lattice) Wrap(v Vec) Vec {
    return Vec{
        X: mod(v.X, l.Size.X),
        Y: mod(v.Y, l.Size.Y),
    }
}

func (l Lattice) Print() {
    fmt.Printf("┌%s┐\n", strings.Repeat("─", l.Size.X*4+2))
    for _, y := range l.Grid {
        fmt.Print("│ ")
        for _, x := range y {
            if x.Val() != 0 {
                fmt.Printf("%-4d", x)
            } else {
                fmt.Printf("·   ")
            }
        }
        fmt.Println(" │")
    }
    fmt.Printf("└%s┘\n", strings.Repeat("─", l.Size.X*4+2))
}

func (l Lattice) Moore(v Vec, distance int) (coords Coords) {
    set := func(px, py, d int) {
        //if l.Check(Vec{px, py}) {
        //    coords = append(coords, Vec{px, py})
        //}
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

func (l *Lattice) VonNeumann(v Vec, distance int) (coords Coords) {
    set := func(px, py, d int) {
        //if l.Check(Vec{px, py}) {
        //    coords = append(coords, Vec{px, py})
        //}
        coords = append(coords, l.Wrap(Vec{px, py}))
    }

    for d := 0; d < distance; d++ {
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

func (l Lattice) Index(v Vec) (val Cell) {
    //if l.Check(v) {
    //    return l.Grid[v.Y][v.X]
    //}
    //return

    t := l.Wrap(v)
    return l.Grid[t.Y][t.X]
}

func (l Lattice) Sum(coords Coords) (sum int) {
    for _, v := range coords {
        sum += l.Index(v).Val()
    }
    return
}

func mod(d, m int) (rem int) {
    rem = d % m
    if rem < 0 {
        rem += m
    }
    return
}
