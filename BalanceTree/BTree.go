package main

import (
	"fmt"
)

type BTree struct {
	head *BTNode
}

func NewTree() *BTree {
	root := NewNode(-1)
	return &BTree{head: root}
}

func leftRotate(node *BTNode) {
	rn := node.rchild
	rln := rn.lchild
	if node.parent.lchild == node {
		node.parent.lchild = rn
	} else {
		node.parent.rchild = rn
	}
	rn.parent = node.parent
	node.parent = rn
	rn.lchild = node
	node.rchild = rln
	if rln != nil {
		rln.parent = node
	}
	node.DeepUpdate()
	node.Parent().DeepUpdate()
}

func rightRotate(node *BTNode) {
	ln := node.lchild
	lrn := ln.rchild

	if node.parent.lchild == node {
		node.parent.lchild = ln
	} else {
		node.parent.rchild = ln
	}

	ln.parent = node.parent
	node.parent = ln
	ln.rchild = node
	node.lchild = lrn
	if lrn != nil {
		lrn.parent = node
	}
	node.DeepUpdate()
	node.Parent().DeepUpdate()
}

// 尽量用LL型，或RR型调整，因为RL型或LR型可能会有子树非平衡的风险，具体参考3, 2, 1, 4, 5, 6, 7, 10, 9, 8 建树后删除1，3的案例

func leftBalance(node *BTNode) {
	ln := node.lchild
	lnBF := ln.BF()
	if lnBF < 0 {
		rightRotate(node)
	} else {
		leftRotate(ln)
		rightRotate(node)
	}
	node.Parent().DeepUpdate()
}

func rightBalance(node *BTNode) {
	rn := node.rchild
	rnBF := rn.BF()
	if rnBF >= 0 {
		leftRotate(node)
	} else {
		rightRotate(rn)
		rn.DeepUpdate()
		leftRotate(node)
	}
}

func (tree *BTree) balance(node *BTNode) {
	cur := node
	par := node.parent
	par.DeepUpdate()
	BF := par.BF()
	for par != tree.head {
		if par.lchild == cur {
			if BF > 1 || BF < -1 {
				leftBalance(par)
				break
			}
		} else {
			if BF > 1 || BF < -1 {
				rightBalance(par)
				break
			}
		}
		cur = par
		par = par.parent
		par.DeepUpdate()
		BF = par.BF()
	}
}

func (tree *BTree) Add(w int) error {
	dir := 0
	var pre *BTNode
	pre = nil
	cur := tree.head
	for cur != nil && cur.weight != w {
		pre = cur
		if w > cur.weight {
			cur = cur.rchild
			dir = 1
		} else {
			cur = cur.lchild
			dir = 0
		}
	}

	if cur != nil {
		return fmt.Errorf("this value have already exists")
	} else {
		cur = &BTNode{weight: w, deep: 1, parent: pre, lchild: nil, rchild: nil}
		if dir == 1 {
			pre.rchild = cur
		} else {
			pre.lchild = cur
		}
	}
	tree.balance(cur)
	return nil
}

func (tree *BTree) Find(w int) *BTNode {
	cur := tree.head
	for cur != nil && cur.weight != w {
		if cur.weight > w {
			cur = cur.lchild
		} else {
			cur = cur.rchild
		}
	}

	return cur
}

func delNode(node *BTNode) *BTNode {
	par := node.parent
	cur := node

	var dir int
	if par.lchild == cur {
		dir = -1
	} else {
		dir = 1
	}

	par = cur.parent

	if cur.lchild != nil && cur.rchild != nil {
		BF := cur.BF()
		tmp := cur
		pre := tmp
		if BF >= 0 {
			cur = tmp.rchild
			for tmp != nil {
				pre = tmp
				tmp = tmp.lchild
			}
		} else {
			tmp = tmp.lchild
			for tmp != nil {
				pre = tmp
				tmp = cur.rchild
			}
			tmp = pre
			cur.swap(tmp)
			cur = tmp
			par = cur.parent
		}
	}

	if cur.rchild == nil && cur.lchild == nil {
		if dir == -1 {
			par.lchild = nil
		} else {
			par.rchild = nil
		}

	} else if cur.rchild != nil && cur.lchild == nil {
		if dir == -1 {
			par.lchild = cur.rchild
		} else {
			par.rchild = cur.rchild
		}
		cur.rchild.parent = par
	} else if cur.rchild == nil && cur.lchild != nil {
		if dir == -1 {
			par.lchild = cur.lchild
		} else {
			par.rchild = cur.lchild
		}
		cur.lchild.parent = par
	}

	par.DeepUpdate()
	return par
}

func (tree *BTree) Del(w int) error {
	// 查找节点
	cur := tree.Find(w)
	if cur == nil {
		return fmt.Errorf("this value is not exists")
	}
	// 删除节点。
	cur = delNode(cur)
	// 平衡
	flag := 0
	for cur != tree.head {
		BF := cur.BF()
		if BF > 1 {
			flag = 1
		} else if BF < -1 {
			flag = -1
		}

		if flag != 0 {
			var tmp *BTNode
			if flag == -1 {
				tmp = cur.lchild
			} else {
				tmp = cur.rchild
			}
			tree.balance(tmp)
			break
		}
		cur = cur.parent
	}

	return nil
}

func (tree BTree) Traverse() {

}
