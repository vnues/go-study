package main

import (
	"fmt"
	"os"
)

//拿到文件数据
//显然写在外面是不正确的编程习惯，函数式编程
//file,err := os.open('./maze.in')
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	var row, col int

	//%d %d 这样的参数变量的意义
	//传指针过去因为需要修改值
	//注意写成%d,%d是错误的不要加逗号
	fmt.Fscanf(file, "%d %d", &row, &col)
	//把读到的flie写入变量
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

//点（位置）的结构
type point struct {
	//跟x,y坐标不一样
	i, j int
}

//定义搜索的四个方向
//数组第一行第二列a[0][1]往上走是多少跟我们x,y轴不一样
var dirs = [4]point{
	//	向右
	//	{0,1},
	//向上
	{-1, 0},
	//向左
	{0, -1},
	//向下
	{1, 0},
	//向右
	{0, 1},
}

//当前走了一步的左标
//结构体方法
func (p point) Add(r point) point {
	//注意结构体写法
	return point{p.i + r.i, p.j + r.j}

}

//判断是否越界 返回这个二维数组对应的元素值
//这里拿到的值已经是走一步的值了
func (p point) At(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {

		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[0]) {
		//越界了就返回value 0给他
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	//复制一份迷宫长度相同的路径，就数我们走了多少步，这步是路径
	//二维数组的创建
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))

	}
	//定义一个队列，如果探索到了，可以走通就放进去，再继续拿出来探索
	//point类型
	//首先我们是从起点探索来的
	Q := []point{start}
	//我们是通过这个队列是否有探索的点（位置）
	for len(Q) > 0 {
		//当前探索的点
		cur := Q[0]
		//如果到了终点就终止
		if cur == end {
			break
		}
		//取出队列，我们去除去它
		Q = Q[1:] //slice方便操作
		//那我们就走这四个位置看看可以通过吗
		for _, dir := range dirs {
			//拿到当前前进一步的坐标
			//nexr是坐标
			next := cur.Add(dir)
			//判断越界没有
			val, ok := next.At(maze) //OK 意思是没有越界
			//判断遇到墙的情况，1代表有墙
			if !ok || val == 1 {
				//跳出循环
				continue
			}
			//判断是否走过这个路线
			//为什么要判断路线呢？
			val, ok = next.At(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			cursteps, _ := cur.At(steps)
			steps[next.i][next.j] = cursteps + 1
			Q = append(Q, next)

		}
	}

	return steps
}

func main() {
	maze := readMaze("./maze/maze.in")
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%d ", maze[i][j])
		}
		fmt.Println()
	}

	fmt.Println("--------------------路线-------------------------")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	//打印路线
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}

//由于我们采用的是上左下右的判断方式 越是到最后的坐标 越是最节省时间的 所以关键点在于 如果上坐标可以移动
//下坐标也可以移动 下坐标会覆盖上坐标 因为我们是求最短路线 关键点之
/*
 Fprintf 将参数列表 a 填写到格式字符串 format 的占位符中
 并将填写后的结果写入 w 中，返回写入的字节数（返回写入的）
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
*/

//continue 语句 跳过当前循环的剩余语句，然后继续进行下一轮循环。

//广度迷宫算法就是往四个方向走，探索到这四个方向如果不是墙或者越界就存起来继续探索，但是遇到的问题就是可能会走以前走过的路线，所以得判断

//我们得知道位置
