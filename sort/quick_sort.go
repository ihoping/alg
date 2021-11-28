package sort

//快速排序
func QuickSort(arr []int) []int {
	if arr == nil {
		return nil
	}

	quickSortMerge(arr, 0, len(arr)-1)
	return arr
}

func quickSortMerge(arr []int, l, r int) {
	if l >= r {
		return
	}

	pivot := arr[l]

	i := l + 1
	j := l + 1

	for i <= r {
		if arr[i] < pivot {
			swap(arr, i, j)
			j++
		}
		i++
	}

	swap(arr, l, j-1)

	quickSortMerge(arr, l, j-2)
	quickSortMerge(arr, j, r)
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
