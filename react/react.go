// Package react implements types and methods for reactive cell-based programming.
package react

import (
//	"fmt"
)

const testVersion = 4

// reactor implements the Reactor interface.
type xReactor struct {
	nextCell int // id of the next cell to be created - simpler than pointers, since method signatures use value cell parameters.
}

// New returns an xReactor which implements Reactor.
func New() Reactor {
	return xReactor{}
}

// CreateInput is a Reactor method that returns an InputCell.
func (r xReactor) CreateInput(value int) InputCell {
	r.nextCell++
	return xInputCell{xCell{r.nextCell, value}}
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (r xReactor) CreateCompute1(other Cell, calc func(int) int) ComputeCell {
	r.nextCell++
	cell := xComputeCell1{}
  cell.id = r.nextCell
	cell.other = other.(xCell)
	cell.calc = calc
	return cell
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (r xReactor) CreateCompute2(other1, other2 Cell, calc func(int, int) int) ComputeCell {
  r.nextCell++
	cell := xComputeCell2{}
  cell.id = r.nextCell
	cell.other1 = other1.(xCell)
  cell.other2 = other2.(xCell)
	cell.calc = calc
	return cell
}


// xCell implements the Cell interface.
type xCell struct {
	id    int
	value int
}

// Value implements the Cell interface.
func (cell xCell) Value() int {
	return cell.value
}

// xInputCell implements InputCell.
type xInputCell struct {
	xCell
}

// SetValue sets the value of an xInputCell.
func (cell xInputCell) SetValue(value int) {
	cell.value = value
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
	other xCell
	calc  func(int) int
}

// xCompute2 specialises ComputeCell for a two-cell dependency.
type xComputeCell2 struct {
	xComputeCell
	other1, other2 xCell
	calc  func(int, int) int
}
