package main

import (
	"math"
)

type BTNode struct {
	weight int
	deep   int
	lchild *BTNode
	rchild *BTNode
	parent *BTNode
}

func NewNode(w int) *BTNode {
	return &BTNode{
		weight: w,
		deep:   1,
		lchild: nil,
		rchild: nil,
		parent: nil,
	}
}

func (node *BTNode) Parent() *BTNode {
	return node.parent
}

func (node *BTNode) LChild() *BTNode {
	return node.lchild
}

func (node *BTNode) RChild() *BTNode {
	return node.rchild
}

func (node *BTNode) DeepUpdate() {
	rnBF, lnBF := 0, 0

	if node.lchild != nil {
		rnBF = node.lchild.deep
	}

	if node.rchild != nil {
		lnBF = node.rchild.deep
	}
	node.deep = int(math.Max(float64(rnBF), float64(lnBF))) + 1
}

func (node *BTNode) Deep() int {
	return node.deep
}

func (node *BTNode) BF() int {
	rnDeep, lnDeep := 0, 0

	if node.lchild != nil {
		lnDeep = node.lchild.deep
	}

	if node.rchild != nil {
		rnDeep = node.rchild.deep
	}
	return rnDeep - lnDeep
}

func (node *BTNode) SetWeight(w int) {
	node.weight = w
}

func (node *BTNode) Weight() int {
	return node.weight
}

func (node *BTNode) swap(other *BTNode) {
	other_w := other.weight
	other.SetWeight(node.weight)
	node.weight = other_w
}
