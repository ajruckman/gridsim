package main

import (
    "testing"

    "github.com/faiface/pixel/pixelgl"

    "github.com/ajruckman/abelian/internal/abelian"
)

func BenchmarkRun(b *testing.B) {
    pixelgl.Run(abelian.Run)

    b.ReportAllocs()
}
