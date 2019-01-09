package main

import "fmt"
//选择出当前范围内的最小值；将该最小值放入当前范围的第一个位置；缩小排序的范围，直到只有范围中只有一个元素为止。
//复杂度
//
//时间复杂度：O(n*n) 每趟排序比较的次数是一个等差数列，是一种不稳定的排序算法
//空间复杂度： O(1)
func SelectSort(arr *[10]int){
	for i:=0;i<len(arr)-1 ;i++  {
		MAX:=arr[i]
		MaxInster:=i
		for j:=i+1;j<len(arr);j++  {
			if MAX<arr[j]{
				MAX=arr[j]
				MaxInster=j
			}

			}
		if MaxInster!=i{
			arr[i],arr[MaxInster]=arr[MaxInster],arr[i]
		}
	}
}
func main(){
	arr:=[10]int{2,3,64,1,546,124476,21,786,231,54}
	SelectSort(&arr)
	fmt.Println(arr)
}
