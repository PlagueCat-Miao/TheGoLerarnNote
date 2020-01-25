package toolNya

func Exp2(x []int)(y []int , z *[]int ){ //形参
	z= &x;
	y =x; //x 与 y 同名 操作有效
	y[0]=-5;
	x[0]=-4;
	return  y,z;
}

func Exp3(x [5]int)(y [5]int , z *[5]int ){ //形参
	z= &x;
	y =x; //x 与 y 同名 操作有效
	y[0]=-5;
	x[0]=-4;
	return  y,z;
}