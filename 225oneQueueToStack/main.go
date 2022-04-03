package main

import "fmt"

//用队列实现栈：使用1个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
type MyStack struct {
	queue []int
}

func Constructor() (s MyStack) {
	return s

}

func (s *MyStack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	for n > 0 {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
		n--
	}

}

func (s *MyStack) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v

}

func (s *MyStack) Top() int {
	return s.queue[0]

}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
	fmt.Println(obj.Empty())

}
