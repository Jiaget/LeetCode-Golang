package LeetCode_Golang

type Object interface {}

type Node struct {
	Data Object // 数据域
	Next *Node	// 地址域
}

type List struct {
	headNode *Node //头节点
}

func (this *List) IsEmpty() bool  {
	if this.headNode == nil{
		return true
	} else {
		return false
	}
}

func (this *List) Length() int {
	cur := this.headNode
	count := 0

	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

// 头插法
func (this *List) Add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = this.headNode
	this.headNode = node
	return node
}

// 尾插法
func (this *List) Append(data Object) {
	node := &Node{Data: data}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for cur.Next != nil{
			cur = cur.Next
		}
		cur.Next = node
	}
}

// 指定位置插入
func (this *List) Insert(index int, data Object)  {
	// 小于0强行头插法
	if index < 0{
		this.Add(data)
	} else if index > this.Length() {
		// 大于链表长度，强行尾插法
		this.Append(data)
	} else {
		pre := this.headNode
		count := 0
		for count < index - 1 {
			pre = pre.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}

}