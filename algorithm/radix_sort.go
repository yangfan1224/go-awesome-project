package algorithm



// 获取数组中的最大值
func getMax(intArray []int) (int){
	max := 0
	for _, value := range  intArray{
		if value > max{
			max = value
		}
	}
	return max
}

func getRadix(num int, exp int) (int){
	return (num/exp) % 10
}

func radixSort(sortArray []int, exp int ) {
	outputArray := make([]int, len(sortArray))
	bucket := make([]int, 10)

	//将基数出现的次数存储到bucket中
	for _, value := range sortArray{
		bucket[getRadix(value,exp)] ++
	}

	//将bucket中存储的值与output的位置对应
	for i:= 1; i < 10; i++ {
		bucket[i] += bucket[i-1]
	}

	//将数据存储到临时数组中
	for i := len(sortArray) -1; i >=0; i-- {
		value := sortArray[i]
		outputArray[bucket[getRadix(value,exp)] - 1] = value
		bucket[getRadix(value,exp)] --
	}

	//排序好的数组拷贝到sortArray
	copy(sortArray, outputArray)
}

func main(){
	arrayInt := [] int{53, 3, 542, 748, 14, 214, 154, 63, 616}
	max := getMax(arrayInt)
	for exp:=1; max/exp > 0; exp *= 10 {
		radixSort(arrayInt, exp)
	}
	for _,value := range arrayInt{
		println(value)
	}
}

