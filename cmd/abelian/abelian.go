package main

import (
    "github.com/ajruckman/abelian/internal/abelian"
    "github.com/faiface/pixel/pixelgl"
)

func main() {
    pixelgl.Run(abelian.Run)
}
