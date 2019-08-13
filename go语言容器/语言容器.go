package main

import "fmt"

/*import "fmt"

func main() {
	var a [3]int //三个整数的数组
	fmt.Print(a[0]) //打印第一个元素
	fmt.Println(a[len(a)-1]) //打印最后一个元素

	//打印索引和元素
	for i,v := range a{
		fmt.Printf("%d %d\n",i,v)
	}
	//仅打印元素
	for _, v := range a{
		fmt.Printf("%d\n",v)
	}
}
*/

//func main()  { //切片的实现 从数组生成切片
//	var a = [3]int{1,2,3}
//	fmt.Println(a,a[1:2]) //[1 2 3] [2]
//}
/*func main()  { 从指定范围中生成切片
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间的所有元素
	fmt.Println(highRiseBuilding[:2])
}*/
/*
//直接声明新的切片
func main()  {
	// 声明字符串切片
	var strList []string

	// 声明整型切片
	var numList []int

	// 声明一个空切片
	var numListEmpty = []int{}

	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)

	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))

	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
}*/
//使用 make() 函数构造切片
//如果需要动态地创建一个切片，可以使用 make() 内建函数，格式如下：
//make( []T, size, cap )
func main()  {
	a := make([]int ,2)
	b := make([]int,2,10)
	fmt.Println(a,b)
	fmt.Println(len(a),len(b))
	//[0 0] [0 0]
	//2 2
	//a 和 b 均是预分配 2 个元素的切片，只是 b 的内部存储空间已经分配了 10 个，但实际使用了 2 个元素。
	//
	//容量不会影响当前的元素个数，因此 a 和 b 取 len 都是 2。
}
















