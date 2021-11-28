package sort

//冒泡
func BubbleSort(arr []int) []int {
	if arr == nil {
		return arr
	}

	length := len(arr)

	for i := 0; i < length; i++ { //i = 已冒泡的个数
		flag := false
		for j := 0; j < length-i-1; j++ {
			if arr[j] < arr[j+1] {
				swap(arr, j, j+1)
				flag = true
			}
		}

		if !flag {
			break
		}

	}
	return arr
}

//选择，不稳定
func SelectSort(arr []int) []int {

	if arr == nil {
		return arr
	}

	length := len(arr)

	for i := 0; i < length; i++ { //已排序的最大值
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		swap(arr, i, min)

	}
	return arr
}

//插入,相比冒泡、选择，插入排序最优
func InsertSort(arr []int) []int {
	if arr == nil {
		return arr
	}

	length := len(arr)

	for i := 1; i < length; i++ { //待插入的值
		value := arr[i]

		j := i - 1
		for ; j >= 0; j-- { //已有序的值
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}

		arr[j+1] = value

	}

	return arr
}

func HeapSort(arr []int) []int {
	if arr == nil {
		return nil
	}

	buildHeap(arr)

	j := len(arr) - 1

	for 0 < j {
		swap(arr, 0, j)
		heapCompareAndSwap(arr, 0, j)
		j--
	}
	return arr
}

func buildHeap(arr []int) {
	length := len(arr)
	lastNodeIndex := length/2 - 1
	for lastNodeIndex >= 0 {
		heapCompareAndSwap(arr, lastNodeIndex, length)
		lastNodeIndex--
	}
}

func heapCompareAndSwap(arr []int, node, max int) {
	if node > max/2-1 { //超过最后一个非叶子节点
		return
	}

	if 2*node+1 < max && arr[2*node+1] > arr[node] {
		swap(arr, 2*node+1, node)
		heapCompareAndSwap(arr, 2*node+1, max)
	}
	if 2*node+2 < max && arr[2*node+2] > arr[node] {
		swap(arr, 2*node+2, node)
		heapCompareAndSwap(arr, 2*node+1, max)
	}
}
