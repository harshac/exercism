package listops

//IntList exposes a new type of array of integers with additional functions
type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(n int) int

//Foldr reduces the list into result starting from right
func (i IntList) Foldr(fn binFunc, result int) int {
	reverse := i.Reverse()
	for _, elem := range reverse {
		result = fn(elem, result)
	}
	return result
}

//Foldl reduces the list into result starting from left 
func (i IntList) Foldl(fn binFunc, result int) int {
	for _, elem := range i {
		result = fn(result, elem)
	}
	return result
}

//Reverse returns a reversed list
func (i IntList) Reverse() IntList {
	reversed := make(IntList, i.Length())
	for n, elem := range i {
		reversed[i.Length()-1-n] = elem
	}
	return reversed
}

//Filter returns a new list with elements matching the predicate
func (i IntList) Filter(fn predFunc) IntList {
	filtered := IntList{}
	for _, elem := range i {
		if fn(elem) {
			filtered = filtered.Append(IntList{elem})
		}
	}
	return filtered
}

//Map applies given function to every element of the list
func (i IntList) Map(fn unaryFunc) IntList {
	result := make(IntList, i.Length())
	for n, elem := range i {
		result[n] = fn(elem)
	}
	return result
}

//Appends adds all elements of second list to the original list
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

//Concat flattens and appends all elements of given lists
func (i IntList) Concat(slice []IntList) IntList {
	var result = i
	for _, list := range slice {
		result = result.Append(list)
	}
	return result
}

//Length calculate the number of elements in given list
func (i IntList) Length() int {
	return len(i)
}
