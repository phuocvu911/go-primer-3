package mysort

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Sort data using quicksort algo
func Sort(data Interface) {
	quickSort(data, 0, data.Len()-1)
}

// quickSort implement quick sort algo.
func quickSort(data Interface, low, high int) {
	if low < high {
		p := partition(data, low, high)
		quickSort(data, low, p-1)
		quickSort(data, p+1, high)
	}
}

// helper func of quicksort.
func partition(data Interface, low, high int) int {
	i := low - 1
	for j := low; j < high; j++ {
		if data.Less(j, high) { // element[j] < pivot, pivot here is last elem
			i++
			data.Swap(i, j)
		}
	}
	data.Swap(i+1, high) // drop the pivot into its sorted slot
	return i + 1         //return the pos of that pivot elem for the next call.
}

type IntSlice []int

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Float64Slice []float64

func (x Float64Slice) Len() int { return len(x) }

// have to manage the NaN cases
func (x Float64Slice) Less(i, j int) bool {
	a, b := x[i], x[j]
	if a != a { // a is NaN
		return b == b // NaN goes first, unless b is NaN too
	}
	if b != b { // b is NaN, a is not
		return false
	}
	return a < b
}
func (x Float64Slice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func Ints(x []int) {
	Sort(IntSlice(x))
}
func Float64s(x []float64) {
	Sort(Float64Slice(x))
}
