package main

import "fmt"
//思想它的基本思想是：通过一趟排序将要排序的数据分割成独立的两部分
// 其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，
// 整个排序过程可以递归进行，以此达到整个数据变成有序序列。
func QuickSort(left int ,right int ,arr *[10]int){
	l:=left
	r:=right
	//pivot是中轴数  一组数中的中间数
	pivot:=arr[(left+right)/2]
	//for循环的目的主要是找出中间e数两边的数字
	for ; l<r;  {
		for ;arr[l]<pivot ;  {
			l++
		}
		for ;arr[r]>pivot ;  {
			r--
		}
		if l>=r{
			break
		}
		arr[l],arr[r]=arr[r],arr[l]
		if arr[l]==pivot{
			l++
		}
		if arr[r]==pivot{
			r--
		}
		if l==r{
			l++
			r--
		}
		if left<r{
			QuickSort(left,r,arr)
		}
		if right>l{
			QuickSort(l,right,arr)
		}
	}


}
func main (){
	arr:=[10]int{123,45,24,56,236,1,62,623,12,21312}
	QuickSort(0,len(arr)-1,&arr)
	fmt.Println(arr)
}
