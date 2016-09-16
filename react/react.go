// Package react implements types and methods for reactive cell-based programming.
package react

import "fmt"

const testVersion = 4

// we'll be using a lot of these
type cellID int

// reactor implements the Reactor interface.
type xReactor struct{}

var state struct { // state of the reactor universe
	nextCell cellID // id of the next cell to be created - simpler than pointers, since method signatures use value cell parameters.

	// state for all cells
	value     map[cellID]int      // current value for each cell, by id.
	comp1call map[cellID][]cellID // which compute1 cells does each cell update?
	comp2call map[cellID][]cellID // which compute2 cells does each cell update?

	// additional state for compute cells
	calc map[cellID]func() int // function for each compute cell
}

func nextCell() cellID {
	state.nextCell++
	return state.nextCell
}

// New returns an xReactor which implements Reactor.
func New() Reactor {
	state.nextCell = 42
	state.value = make(map[cellID]int)
	state.calc = make(map[cellID]func() int)
	state.comp1call = make(map[cellID][]cellID)
	state.comp2call = make(map[cellID][]cellID)
	return xReactor{}
}

// CreateInput is a Reactor method that returns an InputCell.
func (r xReactor) CreateInput(value int) InputCell {
	cell := xInputCell{}
	cell.id = nextCell()
	cell.SetValue(value)
	return cell
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (r xReactor) CreateCompute1(other Cell, calc func(int) int) ComputeCell {
	cell := xComputeCell1{}
	cell.id = nextCell()
	oc := other.(cellIDer).cellID()
	state.calc[cell.id] = pack1(calc, oc)
	state.comp1call[oc] = append(state.comp1call[oc], cell.id)
	state.value[cell.id] = state.calc[cell.id]()
	return cell
}

// pack1 uses closure to pack a single parameter function and its parameter into a zero parameter function
func pack1(calc func(int) int, other cellID) func() int {
	return func() int {
		return calc(state.value[other])
	}
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (r xReactor) CreateCompute2(other1, other2 Cell, calc func(int, int) int) ComputeCell {
	cell := xComputeCell2{}
	cell.id = nextCell()

	oc1 := other1.(cellIDer).cellID()
	state.comp2call[oc1] = append(state.comp2call[oc1], cell.id)
	oc2 := other2.(cellIDer).cellID()
	state.comp2call[oc2] = append(state.comp2call[oc2], cell.id)
	state.calc[cell.id] = pack2(calc, oc1, oc2)
	state.value[cell.id] = state.calc[cell.id]()
	return cell
}

// pack1 uses closure to pack a single parameter function and its parameter into a zero parameter function
func pack2(calc func(int, int) int, other1 cellID, other2 cellID) func() int {
	return func() int {
		v1, v2 := state.value[other1], state.value[other2]
		fmt.Printf("pack2() v1: %d, v2: %d.\n", v1, v2)
		return calc(v1, v2)
	}
}

// xCell implements the Cell interface.
type xCell struct {
	id cellID
}

// Value implements the Cell interface.
func (cell xCell) Value() int {
	return state.value[cell.id]
}

func (cell xCell) cellID() cellID {
	return cell.id
}

// Value implements the Cell interface.
type cellIDer interface {
	cellID() cellID
}

// xInputCell implements InputCell.
type xInputCell struct {
	xCell
}

// SetValue sets the value of an xInputCell and updates all direct and indirect dependencies.
func (cell xInputCell) SetValue(value int) {

	if state.value[cell.id] == value {
		return // no action to take if value unchanged.
	}
	state.value[cell.id] = value
	fmt.Printf("SetValue(): state.value: %v\n", state.value)

	// create todo1 and todo2 maps so we can follow minimal re-work strategies
	todo1 := make(map[cellID]bool)
	todo2 := make(map[cellID]bool)

	// seed the todos
	updateTodos(&todo1, &todo2, cell.id)

	// keep going while there are any todos
	for len(todo1)+len(todo2) > 0 {
		switch {
		case len(todo1) > 0:
			ci := pop(&todo1)
			recalc(&todo1, &todo2, ci)
		case len(todo2) > 0:
			ci := pop(&todo2)
			recalc(&todo1, &todo2, ci)
		}
	}
}

func recalc(todo1, todo2 *map[cellID]bool, ci cellID) {
	val := state.calc[ci]()
	if val != state.value[ci] {
		state.value[ci] = val
		updateTodos(todo1, todo2, ci)
	}
}

// update todo1 and todo2 maps with dependencies of a given cell
func updateTodos(todo1, todo2 *map[cellID]bool, cell cellID) {
	for _, dep := range state.comp1call[cell] {
		(*todo1)[dep] = true
	}
	for _, dep := range state.comp2call[cell] {
		(*todo2)[dep] = true
	}
}

// removes one value from a map and returns it
func pop(todo *map[cellID]bool) (ci cellID) {
	for ci := range *todo {
		delete(*todo, ci)
		return ci
	}
	panic("empty todo passed to pop")
}

// xCompute implements ComputeCell callbacks.
type xComputeCell struct {
	xCell
	nextCbh   int
	callbacks map[int]func(int)
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (cell xComputeCell) AddCallback(cb func(int)) CallbackHandle {
	cell.callbacks[cell.nextCbh] = cb
	cell.nextCbh++
	return &cb
}

// RemoveCallback removes a previously added callback, if it exists.
func (cell xComputeCell) RemoveCallback(h CallbackHandle) {
	hint := h.(int) // panic if it's not actually an int, because we know it is.
	delete(cell.callbacks, hint)
}

// xCompute1 specialises ComputeCell for a single cell dependency.
type xComputeCell1 struct {
	xComputeCell
}

// xCompute2 specialises ComputeCell for a two-cell dependency.
type xComputeCell2 struct {
	xComputeCell
}
