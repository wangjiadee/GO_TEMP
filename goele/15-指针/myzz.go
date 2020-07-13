package main

//func main()  {
//	var num int  = 10
//	Test(num)
//	fmt.Println(num)
//}
//func Test(n int) {
//	n = 30
//	fmt.Println(n)
//}

//func main()  {
//	var a int  =10
//	var p *int
//	p = &a
//	fmt.Printf("%p\n", &a)
//	fmt.Printf("%p\n", p)
//	fmt.Println(*p)
//	*p = 222
//	fmt.Println(a) //222
//}
//
//func main()  {
//	var num int = 10
//	Update(&num)
//	fmt.Println(num)
//}
//func Update(p * int)  {
//	*p =60
//}
//
//数组指针：
//	func main() {
//		nums:=[10]int{1,2,3,4,5,6,7,8,9,10}
//		var p *[10]int
//		p=&nums
//		Update(p)
//		fmt.Println(nums)
//	}
//	func Update(p*[10]int) {
//		p[0]=100
//		(*p)[2]=200
//	}

//func main() {
//	var p [2]*int
//	var i int = 10
//	var j int = 20
//	p[0] = &i
//	p[1] = &j
//	fmt.Println(p) // 打印的结构都是地址
//	fmt.Println(*p[0]) // 不要加括号
//	for i := 0; i < len(p); i++ {
//		fmt.Println(*p[i])
//	}
//	for k, v := range p {
//		fmt.Println("k-", k)
//		fmt.Println("v-", v) //打印的是内存地址
//	}
//}

//func main() {
//	s := []int{1, 2, 3, 4, 54, 6, 7, 8}
//	var p *[]int
//	p = &s
//	fmt.Println((*p)[0])
//	(*p)[0] = 200
//	fmt.Println(s)
//
//	for i := 0; i < len(*p); i++ {
//		fmt.Println((*p)[i])
//	}
//
//	for k,v:=range *p{
//		fmt.Println("k=",k)
//		fmt.Println("value=",v)
//	}
//}

//type Mjie struct {
//	id   int
//	name string
//	age  int
//}
//
//	func main() {
//		s := Mjie{1, "zzz", 12}
//		var p *Mjie
//		p=&s
//		UpdateStruct(p)
//		fmt.Println(s)
//
//	}
//
//	func UpdateStruct(s *Mjie) {
//		s.age =22
//	}

//func main() {
//	var a int =10
//	var p *int
//	p =&a
//	var pp **int
//	pp =&p
//	**pp = 2000
//	fmt.Println(a)
//}