package gridlib

import (
    "image/color"
)

// A vector is an (X,Y) coordinate on a lattice.
type Vec struct {
    X, Y int
}

// A vector set is a slice of vectors. Vector sets are returned by neighborhood
// functions to indicate the positions of all cells in the neighborhood around a
// cell.
type VecSet []Vec

// A cell is a value on a lattice.
type Cell interface {
    // String returns the string representation of the cell. This could be its
    // height, value, or anything else. nil is returned if the cell should not
    // be printed by a lattice's Print() method.
    String() *string

    // Display returns a boolean indicating whether the cell should be displayed
    // on the screen.
    Display() bool

    // GetColor returns the color a cell should be displayed in.
    GetColor() color.Color
}

// The lattice is a 2-dimensional slice of Cells bounded by Size.
type Lattice struct {
    // Size holds the dimensions of the lattice's grid and overlay.
    Size Vec

    // Grid holds the 2-dimensional orthogonal grid of cells on the lattice.
    Grid [][]Cell

    // Overlay can be used for simulations that need to finish reading the
    // lattice before updating it, like in Conway's Game of Life; if changes to
    // the lattice were written to the lattice immediately, cells later in the
    // update process would not be handled correctly.
    Overlay [][]Cell
}

// Simulation is used to define different types of simulations to run on a
// lattice.
type Simulation interface {
    // Init initializes a simulation with a lattice the size of the vector
    // passed.
    Init(Vec)

    // BeginTick prepares the simulation for a tick - for example, by clearing
    // the lattice's overlay.
    BeginTick()

    // Tick moves the lattice one step forward in time. This is when cells are
    // updated according to the simulation's rules.
    Tick()

    // EndTick completes a simulation tick - for example, by overwriting the
    // lattice's grid with the lattice's overlay. After this step, the lattice
    // is ready to be displayed.
    EndTick()
}
