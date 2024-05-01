package game

import (
    "testing"
)

func TestPawnValidation(t *testing.T) {
    g := Game{}
    InitializeBoard(&g)

    m1 := [][]int{{1,1},{1,3}}
    t1 := validatePawn(m1, &g)
    if t1 {
        t.Fatalf(`validatePawn(%v, &g) = false, wanted true`, m1)
    }

    g.p1Turn = false
    m2 := [][]int{{1,6},{1,4}}
    t2 := validatePawn(m2, &g)
    if t2 {
        t.Fatalf(`validatePawn(%v, &g) = false, wanted true`, m1)
    }
}

func TestOrthoValidation(t *testing.T) {
    g := Game{}
    g.board = [][]Piece{
        {{2,2},{0,0},{1,2},},
        {{0,0},{0,0},{0,0},},
        {{2,2},{2,1},{2,2},},
    }
    g.p1Turn = true

    m1 := [][]int{{0,0},{0,2}}
    t1 := validateOrthogonal(m1, &g)
    if t1 {
        t.Fatalf(`validateOrthogonal(%v, &g) = true, wanted false`, m1)
    }

    m2 := [][]int{{0,0},{1,2}}
    t2 := validateOrthogonal(m2, &g)
    if t2 {
        t.Fatalf(`validateOrthogonal(%v, &g) = true, wanted false`, m1)
    }

    m3 := [][]int{{2,2},{2,0}}
    t3 := validateOrthogonal(m3, &g)
    if t3 {
        t.Fatalf(`validateOrthogonal(%v, &g) = false, wanted true`, m1)
    }

    m4 := [][]int{{0,0},{0,2}}
    t4 := validateOrthogonal(m4, &g)
    if t4 {
        t.Fatalf(`validateOrthogonal(%v, &g) = false, wanted true`, m1)
    }
}
