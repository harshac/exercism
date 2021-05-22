package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(n int) int

func (i IntList) Foldr(fn binFunc, result int) int {
	reverse := i.Reverse()
	for _, elem := range reverse {
		result = fn(elem, result)
	}
	return result
}

func (i IntList) Foldl(fn binFunc, result int) int {
	for _, elem := range i {
		result = fn(result, elem)
	}
	return result
}

func (i IntList) Reverse() IntList {
	reversed := make(IntList, i.Length())
	for n, elem := range i {
		reversed[i.Length()-1-n] = elem
	}
	return reversed
}

func (i IntList) Filter(fn predFunc) IntList {
	filtered := IntList{}
	for _, elem := range i {
		if fn(elem) {
			filtered = filtered.Append(IntList{elem})
		}
	}
	return filtered
}

func (i IntList) Map(fn unaryFunc) IntList {
	result := make(IntList, i.Length())
	for n, elem := range i {
		result[n] = fn(elem)
	}
	return result
}

func (i IntList) Append(slice IntList) IntList {
	length := i.Length()
	result := make(IntList, length+slice.Length())
	for n, elem := range i {
		result[n] = elem
	}
	for n, elem := range slice {
		result[length+n] = elem
	}
	return result
}

func (i IntList) Concat(slice []IntList) IntList {
	var result = i
	for _, list := range slice {
		result = result.Append(list)
	}
	return result
}

func (i IntList) Length() int {
	return len(i)
}
