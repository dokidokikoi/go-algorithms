package sorting

/*
 * Bucket sort - https://en.wikipedia.org/wiki/Bucket_sort
 */

var BucketSize = 5

func BucketSort(arr []int) {
	minValue, maxValue := arr[0], arr[0]
	for _, v := range arr {
		if v < minValue {
			minValue = v
		} else if v > maxValue {
			maxValue = v
		}
	}

	buckets := make([][]int, (maxValue-minValue)/BucketSize+1)
	for _, v := range arr {
		index := int((v - minValue) / BucketSize)
		buckets[index] = append(buckets[index], v)
	}

	arr = arr[:0]
	for _, bucket := range buckets {
		InsertionSort(bucket)
		arr = append(arr, bucket...)
	}
}
