package main

import "fmt"
//最优：时间效率O（n） 已经有序，只需要比较n-1
//最差：时间复杂度O(n^2)
//数组全部为逆序，一共需要比较（n-1)+(n-2)+…+1,是一个等差数列
//空间复杂度：O(1),它是一种稳定的排序算法。
//它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，
// 找到相应位置并插入。插入排序在实现上，通常采用in-place排序（即只需用到的额外空间的排序），
// 因而在从后向前扫描过程中，需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。
func InsterSort(arr *[5]int){
	for i := 1; i<len(arr);i++  {
	insterVal:=arr[i]
	insterIndex:=i-1
		for insterIndex>=0&&arr[insterIndex]<insterVal{
			arr[insterIndex+1]=arr[insterIndex]
		}
		for insterIndex +1!=i{
			arr[insterIndex+1]=insterVal
		}
	}
}
func main(){
	arr:=[5]int{3,721,65,21,7}
	InsterSort(&arr)
	fmt.Println(arr)
}
