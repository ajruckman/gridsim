package abelian

import (
    "github.com/ajruckman/abelian/internal/gridlib"
    "github.com/ajruckman/abelian/internal/schema"
    "github.com/ajruckman/lib/err"
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel/pixelgl"
    "golang.org/x/image/colornames"
    "math/rand"
    "time"
)

var (
    size             = gridlib.Vec{X: 128, Y: 128}
    cellSize float64 = 5
    ops      int

    cfg = pixelgl.WindowConfig{
        Title:  "Abelian",
        Bounds: pixel.R(0, 0, float64(size.X)*cellSize, float64(size.Y)*cellSize),
        VSync:  true,
    }
)

func runInt() {
    win, err := pixelgl.NewWindow(cfg)
    liberr.Err(err)

    win.SetPos(pixel.V(0, 26))

    var (
        grid  = genGrid()
        cells = genSquare()
        p     = schema.ParticleSim{}
    )
    _ = grid

    p.Init(size)
    //p.Lattice.Print()

    for range time.Tick(time.Second / 30) {
        if win.Closed() {
            break
        }

        win.Clear(colornames.Black)
        cells.Clear()

        p.Tick()

        drawLattice(p.Lattice, cells)

        cells.Draw(win)
        //grid.Draw(win)
        win.Update()
    }
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func Run() {
    runInt()
}

func drawLattice(l gridlib.Lattice, cells *imdraw.IMDraw) {
    for yi, y := range l.Grid {
        for xi, x := range y {
            if x.Val() != 0 {
                cells.Push(cellVec(float64(xi)*cellSize, float64(yi)*cellSize)...)
                cells.Polygon(0)
            }
        }
    }
}

func genGrid() (draw *imdraw.IMDraw) {
    draw = imdraw.New(nil)
    draw.Color = colornames.White

    for i := float64(0); i < cfg.Bounds.Max.X; i += cellSize {
        draw.Push(pixel.V(i, 0))
        draw.Push(pixel.V(i, cfg.Bounds.Max.Y))
        draw.Line(1)
    }

    for i := float64(0); i < cfg.Bounds.Max.Y; i += cellSize {
        draw.Push(pixel.V(0, i))
        draw.Push(pixel.V(cfg.Bounds.Max.X, i))
        draw.Line(1)
    }

    return
}

/* To use matrix instead of cellVec
var cellVector = []pixel.Vec{
    pixel.V(0, 0),
    pixel.V(cellSize, 0),
    pixel.V(cellSize, cellSize),
    pixel.V(0, cellSize),
}

cell.Push(cellVector...)
cell.SetMatrix(pixel.IM.Moved(pixel.V(x*cellSize, y*cellSize)))
*/

func genSquare() (draw *imdraw.IMDraw) {
    draw = imdraw.New(nil)
    draw.Color = colornames.Red

    return
}

func cellVec(x, y float64) []pixel.Vec {
    return []pixel.Vec{
        pixel.V(x, y),
        pixel.V(x+cellSize, y),
        pixel.V(x+cellSize, y+cellSize),
        pixel.V(x, y+cellSize),
    }
}
