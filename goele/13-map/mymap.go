package main

//func main() {
	////key是唯一的
	//var m map[int]string = map[int]string{1: "张三", 2: "李四", 3: "王五", 4: "itcast"}
	//ma1 := map[int]string{1: "张三", 2: "李四", 3: "王五", 4: "itcast"}
	//ma2 := make(map[string]int,10)
	//fmt.Println("m=",m)
	//fmt.Println(ma1)
	//ma2["宝宝"]=20
	//ma2["宝宝"]=190
	//fmt.Println(ma2)
	//fmt.Println(len(ma2))




	//var m map[int]string = map[int]string{1:"宝宝",2:"丑东西"}
	//fmt.Println(m[2])
	//value, ok := m[6]
	//if ok {
	//	fmt.Println(value)
	//}else {
	//	fmt.Println("ERROR")
	//}

	//for k,v := range m{
	//	fmt.Println(k)
	//	fmt.Println(v)
	//}

	//delete(m,3)
	//fmt.Println(m)

//}

//func main()  {
//	var map1 map[int]string = map[int]string{1:"宝宝",2:"我",3:"爱",4:"你"}
//	DeleteMap(map1)
//	PrintMap(map1)
//	fmt.Println(map1)
//}
//
//func PrintMap(map1 map[int]string)  {
//	for k,v:=range map1{
//		fmt.Println("key=",k)
//		fmt.Println("value=",v)
//	}
//}
//
//func DeleteMap(map1 map[int]string)  {
//	delete(map1,1)
//}


//func main() {
//// 有一個字符串 统计每个字母出现的次数
//	// 定义一个变量
//	var str string = "xiejierwoainiya"
//	//循环定义的变量 取出每个字母
//	// 在存入map里面 注意key 里面的是是字符类型
//	m:= make(map[byte]int)
//	for i:=0;i<len(str) ;i++  {
//		Single := str[i]
//		m[Single] = m[Single] + 1
//	}
//	//循环map格式化输出
//	for k,v :=range m{
//		fmt.Printf("%c:%d\n",k,v)
//	}
//}



