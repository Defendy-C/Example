package main

import "fmt"

type Color int
const (
	none Color = iota
	black
	red
)

type rBNode struct {
	w int
	data interface{}
	left *rBNode
	right *rBNode
	parent *rBNode
	color Color
}

func new(w int, data interface{})*rBNode {
	return &rBNode{data:data, w:w, left:nil, right:nil, parent:nil, color:none}
}

func (r *rBNode)isLeft() bool {
	rp := r.parent
	if rp != nil && rp.left == r {
		return true
	}

	return false
}

func (r *rBNode)HasLeft()bool  {
	if r.left != nil {
		return true
	}
	return false
}

func (r *rBNode)HasRight()bool  {
	if r.right != nil {
		return true
	}
	fmt.Println(r, r.right)
	return false
}

func (r *rBNode)GetSibling() *rBNode {
	if r.isLeft() {
		return r.parent.right
	} else {
		return r.parent.left
	}
}




