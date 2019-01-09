package main

import "fmt"
func BunnleSort(arr *[10]int){
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-1-i ;j++  {
			if arr[j]<arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
}
func main(){
	arr:=[10]int{123,421,36,21,63,71,634,124,731,1}
	BunnleSort(&arr)
	fmt.Println(arr)

}
