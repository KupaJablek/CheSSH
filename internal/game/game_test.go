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
        t.Fatalf(`ValidatePawn(%v, &g) = false, wanted true`, m1)
    }

    g.p1Turn = false
    m2 := [][]int{{1,6},{1,4}}
    t2 := validatePawn(m2, &g)
    if t2 {
        t.Fatalf(`ValidatePawn(%v, &g) = false, wanted true`, m1)
    }
}
