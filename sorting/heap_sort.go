package sorting

/*
 * Heap sort - https://en.wikipedia.org/wiki/Heapsort
 */

func sift(arr []int, i int, arrLen int) []int {
	done := false
	maxChild := 0

	// 寻找当前非叶子节点的最大子节点
	for (i*2+1 < arrLen) && (!done) {
		if i*2+1 == arrLen-1 {
			maxChild = i*2 + 1
		} else if arr[i*2+1] > arr[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		// 当前非叶子节点小于最大子节点，交换节点
		if arr[i] < arr[maxChild] {
			arr[i], arr[maxChild] = arr[maxChild], arr[i]
			i = maxChild
		} else {
			done = true
		}
	}

	return arr
}

// 构建大顶堆
// https://www.cnblogs.com/sunshineliulu/p/12995910.html
func buildMaxHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		arr = sift(arr, i, len(arr))
	}
}

func HeapSort(arr []int) {
	buildMaxHeap(arr)

	for i := len(arr) - 1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		arr = sift(arr, 0, i)
	}
}
