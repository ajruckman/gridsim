package abelian

import (
    "fmt"
    "github.com/ajruckman/lib/err"
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel/pixelgl"
    "golang.org/x/image/colornames"
    "math/rand"
    "sync"
    "time"
)

var (
    cfg = pixelgl.WindowConfig{
        Title:  "Abelian",
        Bounds: pixel.R(0, 0, 1024, 1024),
        //VSync:  true,
    }
    cellSize float64 = 1
    wg       sync.WaitGroup
    ops      int
)

func RunInt() {
    win, err := pixelgl.NewWindow(cfg)
    liberr.Err(err)

    win.SetPos(pixel.V(0, 26))

    var (
        draw1 = imdraw.New(nil)
        //g     = grid()
        cell = square()
    )

    draw1.Color = colornames.White

    for !win.Closed() {
        win.Clear(colornames.Black)
        cell.Clear()

        for i := 0; i < 1000; i++ {
            x := float64(rand.Intn(int(cfg.Bounds.Max.X / cellSize))) * cellSize
            y := float64(rand.Intn(int(cfg.Bounds.Max.Y / cellSize))) * cellSize

            cell.Push(offset(x, y)...)

            cell.Polygon(0)
        }

        cell.Draw(win)

        //g.Draw(win)
        win.Update()
        ops++
    }

    wg.Done()
}

func tick() {
    var (
        before = time.Second * 10
        dur    = time.Second * 60
        after  = time.Millisecond * 250
    )

    time.Sleep(before)
    ops1 := ops
    time.Sleep(dur - (2 * before))
    ops2 := ops
    time.Sleep(after)

    d := ops2 - ops1
    n := dur - before - after

    fmt.Println(d, n, (float64(d)/float64(n.Nanoseconds()))*float64(time.Second))

    wg.Done()
}

func Run() {
    wg.Add(1)
    go RunInt()
    go tick()
    wg.Wait()
    fmt.Println(ops)
}

func grid() (draw *imdraw.IMDraw) {
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

/* To use matrix instead of offset
var cellVector = []pixel.Vec{
    pixel.V(0, 0),
    pixel.V(cellSize, 0),
    pixel.V(cellSize, cellSize),
    pixel.V(0, cellSize),
}

cell.Push(cellVector...)
cell.SetMatrix(pixel.IM.Moved(pixel.V(x*cellSize, y*cellSize)))
*/

func square() (draw *imdraw.IMDraw) {
    draw = imdraw.New(nil)
    draw.Color = colornames.Red

    return
}

func offset(x, y float64) []pixel.Vec {
   return []pixel.Vec{
       pixel.V(x, y),
       pixel.V(x+cellSize, y),
       pixel.V(x+cellSize, y+cellSize),
       pixel.V(x, y+cellSize),
   }
}
