package main

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。
// (若队列中没有元素，deleteHead 操作返回 -1 )
//示例 1：
//输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]

//示例 2：
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]

//提示：
//1 <= values <= 10000
//最多会对 appendTail、deleteHead 进行 10000 次调用

type CQueue struct {
	In  *Stack
	Out *Stack
}

type Stack struct {
	Nums []int
}

func newStack() *Stack {
	return &Stack{
		Nums: []int{},
	}
}

//入栈
func (s *Stack) push(n int) {
	s.Nums = append(s.Nums, n)
}

//出栈
func (s *Stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	result := s.Nums[len(s.Nums)-1]
	s.Nums = s.Nums[:len(s.Nums)-1]
	return result
}

func (s *Stack) len() int {
	return len(s.Nums)
}

//判断栈是否为空
func (s *Stack) isEmpty() bool {
	if len(s.Nums) == 0 {
		return true
	}
	return false
}

func Constructor() CQueue {
	return CQueue{
		In:  newStack(),
		Out: newStack(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.In.push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.Out.isEmpty() {
		if this.In.isEmpty() {
			return -1
		} else {
			for this.In.len() > 0 {
				this.Out.push(this.In.pop())
			}
		}
	}
	return this.Out.pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
