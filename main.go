package main

import (
	"fmt"
	"sort"
	"time"
)

func main()  {
	//fmt.Println("1")
	//result:=0
	//nums := [4]int{2,3,-2,4}

	nums:=[...]int{2,2,4,-9,6,1,100,200,500,800,3,2,5,7,8,3,2,3,4,2,4,3,2,3}

	result:=BiggestProduct1(nums[:])
	fmt.Println(result)

}
func BiggestProduct1(nums []int) int {
	var result int=0
	var numsResult =make([]int,0)
	if len(nums)>1{
		for i:=0;i<len(nums);i++{
			for j:=i+1;j<len(nums);j++ {
				if nums[i]<0 && nums[j]<0{
					continue
				}else {
					temp:=nums[i]*nums[j]
					numsResult = append(numsResult, temp)
				}
				i=j
			}
		}
	}
	if len(numsResult)>0{
		//冒泡排序
		cur1:=time.Now()
		for i:=0;i<len(numsResult);i++ {
			for j:=i+1;j<len(numsResult);j++{
				if numsResult[i] < numsResult[j]{
					temp := numsResult[i]
					numsResult[i] = numsResult[j]
					numsResult[j] = temp
				}
			}
		}
		result=numsResult[0]
		nowTime1 :=time.Now()
		fmt.Println(cur1)
		fmt.Println(nowTime1)

		cur := time.Now()
		var sortList []int
		for _,v := range numsResult {
			sortList = append(sortList, v)
		}
		sort.Ints(sortList)
		result=sortList[len(sortList)-1]
		nowTime :=time.Now()
		fmt.Println(cur)
		fmt.Println(nowTime)
	}
	return result
}

func BiggestProduct(nums []int) int {
	var result int=0
	var numsResult =make(map[int][2]int,0)
	if len(nums)>0{
		for i:=0;i<len(nums);i++{
			for j:=i+1;j<len(nums);j++ {
				if nums[i]<0 && nums[j]<0{
					continue
				}else {
					temp:=nums[i]*nums[j]
					valTemp:=[2]int{nums[i],nums[j]}
					numsResult[temp]=valTemp
				}
				i=j
			}
		}
	}
	if len(numsResult)>0{
		var keys []int
		for k := range numsResult {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		result=keys[len(keys)-1]
		fmt.Printf("子数组 [%d,%d] 有最大乘积 %d \n",numsResult[result][0],numsResult[result][1],result)
	}
	return result
}
