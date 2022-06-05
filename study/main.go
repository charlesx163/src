package main

import "fmt"

type user struct {
	name  string
	email string
}
type notifier interface {
	notify()
}

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration:%d", *d)
}
func main() {

	// var arr1 [2][2]int
	// var arr2 [2][2]int
	// arr1 = [2][2]int{{10, 20}, {30, 40}}
	// arr1[0][0] = 10
	// arr1[0][1] = 20
	// arr1[1][0] = 30
	// arr1[1][1] = 40
	// arr2 = arr1

	// fmt.Println(arr2[0][1])

	// var arr3 [2]int = arr1[1]
	// for _, ele := range arr3 {
	// 	fmt.Println(ele)
	// }
	// var arr4 [2]int = arr1[0]
	// for _, ele := range arr4 {

	// 	fmt.Println(ele)
	// }

	// slice := make([]string, 5)
	// slice = []string{"red", "blue", "green", "yellow", "pink"}
	// newSlice := slice[1:3]

	// for index, item := range newSlice {
	// 	fmt.Println(index, item)
	// }
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))
	//修改切片的元素底层数组也会被改变
	// slice := []int{10, 20, 0, 40, 50}
	// newSlice := slice[1:3]
	// newSlice[1] = 30
	// fmt.Println(slice[2])
	// newSlice = append(newSlice, 60)
	// fmt.Println(newSlice[2])
	// fmt.Println(slice[3])

	//如果切片的底层数组没有足够的可用容量， append 函数会创建一个新的底层数组，将被引
	//用的现有的值复制到新数组里，再追加新的值，
	// slice := []int{10, 20, 30, 40, 50}
	// newSlice := append(slice, 60)
	// for index, ele := range newSlice {
	// 	fmt.Println(index, ele)
	// }
	// for index, ele := range slice {
	// 	fmt.Println(index, ele)
	// }
	// fmt.Println(newSlice[5])

	//3个索引创建切片
	//使用第三个索引：容量限制索引的好处是：当使用append的时候可以限制用量达到不修改底层的效果
	//如果不加第三个索引，由于剩余的所有容量都属于 slice，向 slice 追加 Kiwi 会改变
	//原有底层数组索引为 3 的元素的值 Banana。不过在代码清单 4-36 中我们限制了 slice 的容
	//量为 1。当我们第一次对 slice 调用 append 的时候，会创建一个新的底层数组，这个数组包
	//括 2 个元素，并将水果 Plum 复制进来，再追加新水果 Kiwi，并返回一个引用了这个底层数组
	//的新切片，如图 4-18 所示。
	//因为新的切片 slice 拥有了自己的底层数组，所以杜绝了可能发生的问题。我们可以继续
	//向新切片里追加水果，而不用担心会不小心修改了其他切片里的水果。同时，也保持了为切片申请新的底层数组的简洁
	// source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// slice := source[2:3:3]
	// for index, ele := range slice {
	// 	fmt.Println(index, ele)
	// }
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))
	// fmt.Println("append")
	// slice = append(slice, "Kiwi")
	// for index, ele := range slice {
	// 	fmt.Println(index, ele)
	// }
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// fmt.Println("source")
	// fmt.Println(source[2])
	// fmt.Println(source[3])

	//切片叠加
	// s1 := []int{1, 2}
	// s2 := []int{3, 4}
	// fmt.Println("%v\n", append(s1, s2...))

	//map
	// dict := map[string]string{"Red": "#da337", "Orange": "#e952aa"}

	// value, exists := dict["Red"]
	// fmt.Println(value)
	// fmt.Println(exists)

	//值类型和指针类型
	// bill := user{"Bill", "bill@email.com"}
	// bill.notify()
	// bill1 := &user{"Bill", "bill@email.com"}
	// bill1.changeName("charles")
	// bill.changeName("1111111111")
	// (&bill).changeName("22222")

	//接口
	u := user{"Frank", "bill@email.com"}
	sendNotification(&u)
}

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// func (u *user) changeName(name string) {
// 	u.name = name
// 	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
// }

func sendNotification(n notifier) {
	n.notify()
}
