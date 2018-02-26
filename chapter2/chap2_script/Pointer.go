func pointer_test(){
	//空指针，输出值为nil
	var P *int
	fmt.Printf("p:%v\n",p)
	*p = 8 
	fmt.Printf("p: %v,%v\n", p, *p)
	//数组的初始化及输出
	m := [3]int{3, 4, 5}
	fmt.Printf("m:%v--%v,%v,%v\n", m, m[0], m[1], m[2])
	//指针数组的初始化及输出
	x := [3]*int{&m[0], &m[1], &m[2]}
	fmt.Printf("x:%v,%v,%v\n", x[0], x[1], x[2])
	fmt.Printf("*x:%v,%v,%v\n", *x[0], *x[1], *x[2])
	}