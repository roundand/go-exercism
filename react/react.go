// Package react implements types and methods for reactive cell-based programming.
package react

const testVersion = 4

type todoList map[*xComputeCell]bool

type compcaller interface {
	addComp1(*xComputeCell)
	addComp2(*xComputeCell)
}

////////////////////////////////////////////////
// xReactor implements the Reactor interface. //
////////////////////////////////////////////////
type xReactor struct{}

// New returns an xReactor, which implements Reactor.
func New() Reactor {
	return xReactor{}
}

// CreateInput is a Reactor method that returns an InputCell.
func (r xReactor) CreateInput(value int) InputCell {
	cell := xInputCell{}
	cell.SetValue(value)
	return &cell
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (r xReactor) CreateCompute1(other Cell, calc func(int) int) ComputeCell {
	cell := xComputeCell{callbacks: map[*func(int)]func(int){}}
	cell.calc = pack1(calc, other)
	cell.value = cell.calc()
	other.(compcaller).addComp1(&cell)
	return &cell
}

// pack1 uses closure to pack a single parameter function and its parameter into a zero parameter function
func pack1(calc func(int) int, oc Cell) func() int {
	return func() int {
		return calc(oc.Value())
	}
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (r xReactor) CreateCompute2(other1, other2 Cell, calc func(int, int) int) ComputeCell {
	cell := xComputeCell{callbacks: map[*func(int)]func(int){}}
	cell.calc = pack2(calc, other1, other2)
	cell.value = cell.calc()
	other1.(compcaller).addComp2(&cell)
	other2.(compcaller).addComp2(&cell)
	return &cell
}

// pack1 uses closure to pack a single parameter function and its parameter into a zero parameter function
func pack2(calc func(int, int) int, oc1, oc2 Cell) func() int {
	return func() int {
		return calc(oc1.Value(), oc2.Value())
	}
}

///////////////////////////////////////////
// *xCell implements the Cell interface. //
///////////////////////////////////////////
type xCell struct {
	value     int
	comp1call []*xComputeCell
	comp2call []*xComputeCell
}

// Value implements Cell.Value.
func (cell *xCell) Value() int {
	return (*cell).value
}

///////////////////////////////////////
// *xInputCell implements InputCell. //
///////////////////////////////////////
type xInputCell struct {
	xCell
}

// also implements addcomper interface
func (cell *xInputCell) addComp1(dep *xComputeCell) {
	cell.comp1call = append(cell.comp1call, dep)
}
func (cell *xInputCell) addComp2(dep *xComputeCell) {
	cell.comp2call = append(cell.comp2call, dep)
}

// SetValue sets the value of an xInputCell and updates all direct and indirect dependencies.
func (cell *xInputCell) SetValue(value int) {

	if (*cell).value == value {
		return // no action to take if value unchanged.
	}
	(*cell).value = value

	// create todo1 and todo2 maps so we can follow minimal re-work strategies
	todo1 := make(todoList)
	todo2 := make(todoList)

	// seed the todos
	cell.updateTodos(&todo1, &todo2)

	// keep going while there are any todos
	for len(todo1)+len(todo2) > 0 {
		switch {
		case len(todo1) > 0:
			ci := todo1.pop()
			ci.recalc(&todo1, &todo2)
		case len(todo2) > 0:
			ci := (&todo2).pop()
			ci.recalc(&todo1, &todo2)
		}
	}
}

// update xCell todo1 and todo2 maps with dependencies of a given cell
func (cell *xCell) updateTodos(todo1, todo2 *todoList) {
	for _, dep := range cell.comp1call {
		(*todo1)[dep] = true
	}
	for _, dep := range cell.comp2call {
		(*todo2)[dep] = true
	}
}

// removes one value from a map and returns it
func (todo *todoList) pop() (cell *xComputeCell) {
	for ci := range *todo {
		delete(*todo, ci)
		return ci
	}
	panic("empty todo passed to pop")
}

/////////////////////////////////////////////////////
// *xComputeCell implements ComputeCell callbacks. //
/////////////////////////////////////////////////////
type xComputeCell struct {
	xCell
	calc      func() int
	callbacks map[*func(int)]func(int)
}

// also implements addcomper interface
func (cell *xComputeCell) addComp1(dep *xComputeCell) {
	cell.comp1call = append(cell.comp1call, dep)
}
func (cell *xComputeCell) addComp2(dep *xComputeCell) {
	cell.comp2call = append(cell.comp2call, dep)
}

// should be able to recalculate itself too.
func (cell *xComputeCell) recalc(todo1, todo2 *todoList) {
	val := cell.calc()
	if val != cell.value {
		cell.value = val
		cell.updateTodos(todo1, todo2)
		for calc := range cell.callbacks {
			(*calc)(cell.Value())
		}
	}
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (cell *xComputeCell) AddCallback(cb func(int)) CallbackHandle {
	(*cell).callbacks[&cb] = cb
	return &cb
}

// RemoveCallback removes a previously added callback, if it exists.
func (cell *xComputeCell) RemoveCallback(h CallbackHandle) {
	delete((*cell).callbacks, h.(*func(int)))
}
