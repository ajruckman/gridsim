package main

import (
    "github.com/ajruckman/gridsim/internal/display"
    "github.com/faiface/pixel/pixelgl"
)

func main() {
    pixelgl.Run(display.Run)
}
