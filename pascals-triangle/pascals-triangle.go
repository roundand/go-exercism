// Package pascal provides a method for calculating Pascal's triangle.
package pascal

// Function Triangle calculates Pascal's Triangle to the specified depth
func Triangle(depth int) [][]int {

	// terminating case
	if depth == 1 {
		return [][]int{
			{1},
		}
	}

	// recursive case - add new layer to root triangle
	root := Triangle(depth - 1) // get the rest of the triangle
	base := root[depth-2]       // grab current base layer for convenience

	// create and populate the new layer
	var layer []int = make([]int, depth) // next layer will be as wide as triangle is deep
	layer[0], layer[depth-1] = 1, 1      // initialise first and last elements
	for i := 1; i < depth-1; i++ {       // calculate remaining elements
		layer[i] = base[i-1] + base[i]
	}

	// ... et voila
	root = append(root, layer)
	return root
}
