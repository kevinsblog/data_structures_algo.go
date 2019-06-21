package msort

import (
	"fmt"
	"math"
)

const version = "0.1"

func Version() {
	fmt.Println("msort ver:", version)
}

//insert sort
func InsertSort(nums []int) {
	var j int
	for p := 1; p < len(nums); p++ {
		tmp := nums[p]
		// mind the j scope
		for j = p; j > 0 && tmp < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = tmp
	}
}

//bubble sort
func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		var doSwap bool //if no swap done in this loop, the whole array must be in-order
		for j := len(nums) - 1; j > 0; j-- {
			if nums[j-1] > nums[j] {
				doSwap = true
				nums[j], nums[j-1] = nums[j-1], nums[j] //bubble up
			}
		}

		if !doSwap {
			break
		}
	}
}

//selection sort
func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		var minIdx int
		minKey := math.MaxInt32
		//select the min element in 'unorder' region
		for j := i; j < len(nums); j++ {
			if nums[j] < minKey {
				minIdx = j
				minKey = nums[j]
			}
		}
		//put min element in the back of 'ordered region'
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

//shell sort
func ShellSort(nums []int) {
	//compare ans swap element with distance gap
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		//offset i on current gap
		for i := gap; i < len(nums); i++ {
			var j int
			tmp := nums[i] //to put nums[i] at its right place
			//insert sort, move all nums[j] at gap 'gap' that is greater
			//	than tmp
			for j = i; j >= gap && nums[j] < nums[j-gap]; j -= gap {
				nums[j] = nums[j-gap]
			}
			nums[j] = tmp //found the right place for tmp
		}
	}
}

//quick sort
func QuickSort(nums []int) {
	if len(nums) > 1 {
		var smaller, equal, larger []int
		pivot := nums[len(nums)/2]
		//split input into three disjoint subset
		for _, n := range nums {
			if n < pivot {
				smaller = append(smaller, n)
			} else if n > pivot {
				larger = append(larger, n)
			} else {
				equal = append(equal, n)
			}
		}

		//sort the two subset
		QuickSort(smaller)
		QuickSort(larger)
		//combine the subset
		copy(nums[:], smaller)
		copy(nums[len(smaller):], equal)
		copy(nums[len(smaller)+len(equal):], larger)
	}
}

//in-place quick sort
func partition(nums []int, low, high int) int {
	pivot := nums[high] //select the last element as pivot
	i := low            //set 'smaller' boundary at low

	//traversal over all elements
	for j := low; j < high; j++ {
		if pivot > nums[j] {
			//if element is smaller than pivot, put it
			//	at low boundary and adjust boundary
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	//put pivot element at the low boundary
	nums[i], nums[high] = nums[high], nums[i]
	return i
}

func quickSortHelper(nums []int, low, high int) {
	if low < high {
		pi := partition(nums, low, high)
		quickSortHelper(nums, low, pi-1)
		quickSortHelper(nums, pi+1, high)
	}
}

func QuickSort2(nums []int) {
	quickSortHelper(nums, 0, len(nums)-1)
}

//heap sort
func leftChild(i int) int {
	return 2*i + 1
}

func precDown(heap []int, i, hsize int) {
	var child int
	var tmp int

	for tmp = heap[i]; leftChild(i) < hsize; i = child {
		child = leftChild(i)
		if child < hsize-1 && heap[child] < heap[child+1] {
			child++
		}
		//exchange left child
		if tmp < heap[child] {
			heap[i] = heap[child]
		} else {
			break
		}
	}
	heap[i] = tmp
}

func HeapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		precDown(nums, i, len(nums))
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		precDown(nums, 0, i)
	}
}

//merge sort
func merge(nums []int, low, mid, high int) {
	var L, R []int

	//left must include mid, since low could be equal with mid when
	//	high - low < 2, otherwise L would be empty and sort would fail
	L = append(L, nums[low:mid+1]...)
	R = append(R, nums[mid+1:high+1]...)

	var left, right, idx int
	llen, rlen := len(L), len(R) //mid-low+1, high-mid

	idx = low
	for left < llen && right < rlen {
		//fmt.Println(len(L), len(R), left, right)
		if L[left] < R[right] {
			nums[idx] = L[left]
			left++
		} else {
			nums[idx] = R[right]
			right++
		}
		idx++
	}

	if left < llen {
		copy(nums[idx:], L[left:])
	} else if right < rlen {
		copy(nums[idx:], R[right:])
	}
}

func mergeSortHelper(nums []int, low, high int) {
	if low < high {
		mid := low + (high-low)/2 //[0, 0, 1], low = mid
		mergeSortHelper(nums, low, mid)
		mergeSortHelper(nums, mid+1, high)
		merge(nums, low, mid, high)
	}
}

func MergeSort(nums []int) {
	mergeSortHelper(nums, 0, len(nums)-1)
}

//radix sort
func RadixSort(strs []string) {
	if len(strs) < 2 {
		return
	}
	const BUCKET = 256
	strLen := len(strs[0])
	buckets := make([][]string, BUCKET) //create buckets

	for i := strLen - 1; i >= 0; i-- {
		//put strings into buckets
		for _, s := range strs {
			buckets[s[i]] = append(buckets[s[i]], s)
		}

		var idx int
		for _, b := range buckets {
			//build string array with buckets
			for _, s := range b {
				strs[idx] = s
				idx++
			}

			buckets = make([][]string, BUCKET) //reset buckets
		}
	}
}

func CountingRadixSort(strs []string) {
	if len(strs) < 2 {
		return
	}
	const BUCKET = 256
	strLen := len(strs[0])
	arrLen := len(strs)

	in := strs //reference to input array
	out := make([]string, len(strs))

	//bucket sort on the i-th char of all strings
	for pos := strLen - 1; pos >= 0; pos-- {
		count := make([]int, BUCKET+1)

		//iterate over all strings
		for j := 0; j < arrLen; j++ {
			count[in[j][pos]+1]++ //the i-th char of the j-th string
		}

		//accumlate counting of buckets
		for b := 1; b <= BUCKET; b++ {
			count[b] += count[b-1]
		}

		//put back string at the right place
		for j := 0; j < arrLen; j++ {
			out[count[in[j][pos]]] = in[j]
			count[in[j][pos]]++
		}

		//exchange bucket
		in, out = out, in
	}

	//if odd num of passes, in is buffer while out is arr, copy buffer
	//	to arr
	if strLen%2 == 1 {
		copy(out[:], in[:])
	}

	//copy to origin array
	copy(strs[:], out[:])
}
