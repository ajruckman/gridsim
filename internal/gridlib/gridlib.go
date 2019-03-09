package gridlib

import (
    "fmt"
    "strings"
)

type Cell interface {
    Pos() Vector
    Val() CellVal
}

type CellVal int

type Lattice struct {
    Lattice [][]int
    Size    Vector
    //Center   Vector
    Min, Max Vector // Extrapolated from Size and Center
}

func (l *Lattice) SetBounds() {
    l.Size.X = len(l.Lattice[0])
    l.Size.Y = len(l.Lattice)

    l.Min = Vector{0, 0}
    l.Max = Vector{l.Size.X - 1, l.Size.Y - 1}

    //l.Max = Vector{l.Size.X - l.Center.X - 1, l.Size.Y - l.Center.Y - 1}
    //l.Min = Vector{l.Max.X - l.Size.X + 1, l.Max.Y - l.Size.Y + 1}
}

type Vector struct {
    X, Y int
}

func (l Lattice) Contains(pos Vector) bool {
    if l.Min.X <= pos.X && pos.X <= l.Max.X && l.Min.Y <= pos.Y && pos.Y <= l.Max.Y {
        return true
    }
    return false
}

func (l Lattice) Print() {
    fmt.Println("+-" + strings.Repeat("--", l.Size.Y) + "+")
    for _, y := range l.Lattice {
        fmt.Print("| ")

        for _, x := range y {
            if x == 0 {
                fmt.Print("  ")
            } else {
                fmt.Print(x, " ")
            }
        }
        fmt.Println("|")
    }
    fmt.Println("+-" + strings.Repeat("--", l.Size.Y) + "+")
}

func (l Lattice) VonNeumann(cell Cell, r int) (positions []Vector) {
    //fmt.Printf("(%-5d,%-5d) ... (%-5d,%-5d)\n",
    //    cell.Pos().X-r, cell.Pos().Y-r,
    //    cell.Pos().X+r, cell.Pos().Y+r,
    //)

    for y := cell.Pos().Y - r; y <= cell.Pos().Y+r; y++ {
        for x := cell.Pos().X - r; x <= cell.Pos().X+r; x++ {
            if x == cell.Pos().X && y == cell.Pos().Y {
                continue
            }
            if l.Contains(Vector{x, y}) {
                positions = append(positions, Vector{x, y})
            }
        }
    }

    //r := 3
    //for x := -r; x <= r; x++ {
    //   r_x := r - abs(x)
    //   for y := -r_x; y <= r_x; y++ {
    //       r_y := r_x - abs(y)
    //
    //       fmt.Println(r_x, r_y)
    //   }
    //}

    return
}

func (l Lattice) Index(vector Vector) int {
    if l.Contains(vector) {
        return l.Lattice[vector.Y][vector.X]
    } else {
        return 0
    }
}

func (l Lattice) SumVectors(vectors []Vector) (sum int) {
    for _, v := range vectors {
        sum += l.Index(v)
    }
    return
}

//func (l Lattice) VonNeumannOld(cell Cell, distance int) (vec []Vector) {
//    //var m = [14][14]int{}
//
//    add := func(px, py int) {
//        if l.Contains(px, py) {
//            fmt.Println(fmt.Sprintf(" + %5d %5d %5d", py, px, len(vec)))
//            vec = append(vec, Vector{px, py})
//
//            //m[py][px] = 1
//
//            //    vec = append(vec, Vector{px, py})
//        } else {
//            fmt.Println(fmt.Sprintf("-  %5d %5d %5d", py, px, len(vec)))
//        }
//        //fmt.Println(l.Contains(px, py))
//    }
//
//    _ = add
//
//    //for d := distance; d > 0; d-- {
//
//    //px := cell.Pos().X - distance
//    py := cell.Pos().Y - distance
//
//    for d := 0; d < distance*2; d++ {
//
//        for px := cell.Pos().X - distance; px < cell.Pos().X+distance; px++ {
//            add(px, py)
//        }
//
//        py++
//
//        //px := cell.Pos().X - d
//        //py := cell.Pos().Y - d
//        //for c := 0; c < cell.Pos().X; c++ {
//        //    add(px, py)
//        //}
//        //py--
//
//        //px := cell.Pos().X - d
//        //py := cell.Pos().Y - d
//
//        //for c := 0; c < d * 2; c++ {
//        //    add(px, py)
//        //    px++
//        //}
//        //for c := 0; c < d * 2; c++ {
//        //    add(px, py)
//        //    py--
//        //}
//        //for c := 0; c < d * 2; c++ {
//        //    add(px, py)
//        //    px--
//        //}
//        //for c := 0; c < d * 2; c++ {
//        //    add(px, py)
//        //    py++
//        //}
//
//        //var px = cell.Pos().X - d
//        //var py = cell.Pos().Y - d
//        //
//        //for c := px; c <= px + d; c++ {
//        //    add(px, py)
//        //    px++
//        //}
//        //for c := py; c <= py - d; c++ {
//        //    add(px, py)
//        //    py--
//        //}
//        //for c := px; c >= px - d; c-- {
//        //    add(px, py)
//        //    px--
//        //}
//        //for c := py; c >= py - d; c-- {
//        //    add(px, py)
//        //    py++
//        //}
//
//        //for c := 0; c < d; c++ {
//        //   add(px, py)
//        //   px++
//        //   py--
//        //}
//        //for c := 0; c < d; c++ {
//        //   add(px, py)
//        //   px++
//        //   py++
//        //}
//        //for c := 0; c < d; c++ {
//        //   add(px, py)
//        //   px--
//        //   py--
//        //}
//        //for c := 0; c < d; c++ {
//        //   add(px, py)
//        //   px--
//        //   py--
//        //}
//
//        //fmt.Println("------")
//        //fmt.Println(" x", len(vec), d)
//        //fmt.Println("------")
//        //fmt.Println()
//        //fmt.Println()
//        //
//        //m[8][10] = 2
//
//        //for _, y := range m {
//        //   for _, x := range y {
//        //       switch x {
//        //       case 0:
//        //           fmt.Print(".")
//        //       case 1:
//        //           fmt.Print("x")
//        //       case 2:
//        //           fmt.Print("+")
//        //       }
//        //   }
//        //   fmt.Println()
//        //}
//        //fmt.Println()
//
//        //m = [14][14]int{}
//    }
//
//    //for d := distance; d > 0; d-- {
//    //    var px = cell.Pos().X - d
//    //    var py = cell.Pos().Y
//    //
//    //    for c := 0; c < (d * 2); c++ {
//    //        add(px, py)
//    //        px++
//    //    }
//    //    for c := 0; c < (d * 2); c++ {
//    //        add(px, py)
//    //        px++
//    //
//    //    }
//    //    for c := 0; c < (d * 2); c++ {
//    //        add(px, py)
//    //        px++
//    //
//    //    }
//    //    for c := 0; c < (d * 2); c++ {
//    //        add(px, py)
//    //        px++
//    //
//    //    }
//    //
//    //    fmt.Println("------")
//    //    fmt.Println(" x", len(vec), d)
//    //    fmt.Println("------")
//    //    fmt.Println()
//    //    fmt.Println()
//    //}
//
//    fmt.Println("------")
//    fmt.Println(" *", len(vec))
//    fmt.Println("------")
//    fmt.Println()
//    fmt.Println()
//
//    return
//}

func abs(i int) int {
    n := int64(i)
    y := n >> 63
    return int((n ^ y) - y)
}
