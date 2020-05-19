package main

import (
	"bytes"
	"container/list"
	"fmt"
)

const null = (1 << 31) * -1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	f := t.format()
	f.update(0)
	row := []*treeFormat{f}
	buf := newPrintBuf()
	for len(row) > 0 {
		for _, i := range row {
			buf.RepeatTo(' ', i.p)
			buf.Write(fmt.Sprintf(" %d ", i.n.Val))
		}
		buf.ln()

		newRow := []*treeFormat{}
		for _, i := range row {
			if i.l != nil {
				l := i.l
				newRow = append(newRow, l)
				buf.RepeatTo(' ', l.p+l.nw-1)
				buf.RepeatTo('_', i.p)
				buf.WriteByte('/')
			}
			if i.r != nil {
				r := i.r
				newRow = append(newRow, r)
				buf.RepeatTo(' ', i.p+i.nw-1)
				buf.WriteByte('\\')
				buf.RepeatTo('_', r.p+1)
			}
		}
		buf.ln()
		row = newRow
	}
	return buf.String()
}

func (t *TreeNode) format() *treeFormat {
	f := &treeFormat{
		n:  t,
		nw: len(fmt.Sprintf(" %d ", t.Val)),
	}
	if t.Left != nil {
		f.l = t.Left.format()
		f.lw = f.l.w
	}
	if t.Right != nil {
		f.r = t.Right.format()
		f.rw = f.r.w
	}
	f.w = f.nw + f.lw + f.rw
	f.p = f.lw
	return f
}

func SliceToTree(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   nums[0],
		Left:  nil,
		Right: nil,
	}
	q := list.New()
	q.PushBack(root)

	for i := 1; i < len(nums) && q.Len() > 0; {
		e := q.Front()
		q.Remove(e)
		n := e.Value.(*TreeNode)
		if i < len(nums) && nums[i] != null {
			n.Left = &TreeNode{
				Val:   nums[i],
				Left:  nil,
				Right: nil,
			}
			q.PushBack(n.Left)
		}
		i++
		if i < len(nums) && nums[i] != null {
			n.Right = &TreeNode{
				Val:   nums[i],
				Left:  nil,
				Right: nil,
			}
			q.PushBack(n.Right)
		}
		i++
	}
	return root
}

func TreeToSlice(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		n := e.Value.(*TreeNode)
		if n == nil {
			res = append(res, null)
		} else {
			res = append(res, n.Val)
			if n.Left != nil || n.Right != nil {
				q.PushBack(n.Left)
				q.PushBack(n.Right)
			}
		}
	}
	return res
}

func TreeBFS(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		n := e.Value.(*TreeNode)
		res = append(res, n.Val)

		if n.Left != nil {
			q.PushBack(n.Left)
		}
		if n.Right != nil {
			q.PushBack(n.Right)
		}
	}
	return res
}

func PreorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stk := make([]*TreeNode, 0)
	stk = append(stk, root)

	for len(stk) > 0 {
		n := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		res = append(res, n.Val)

		if n.Left != nil {
			stk = append(stk, n.Left)
		}
		if n.Right != nil {
			stk = append(stk, n.Right)
		}
	}
	return res
}

type treeFormat struct {
	n  *TreeNode
	l  *treeFormat
	r  *treeFormat
	lw int
	rw int
	nw int
	w  int
	p  int
}

func (f *treeFormat) update(p int) {
	f.p = f.p + p
	if f.l != nil {
		f.l.update(p)
	}
	if f.r != nil {
		f.r.update(f.w - f.rw + p)
	}
}

type printBuf struct {
	buf *bytes.Buffer
	idx int
}

func newPrintBuf() *printBuf {
	return &printBuf{
		buf: bytes.NewBuffer(nil),
	}
}

func (p *printBuf) Write(s string) {
	p.buf.WriteString(s)
	p.idx += len(s)
}

func (p *printBuf) WriteByte(b byte) {
	p.buf.WriteByte(b)
	p.idx++
}

func (p *printBuf) ln() {
	p.buf.WriteByte('\n')
	p.idx = 0
}

func (p *printBuf) RepeatTo(b byte, pos int) {
	for p.idx < pos {
		p.buf.WriteByte(b)
		p.idx++
	}
}

func (p *printBuf) String() string {
	return p.buf.String()
}
