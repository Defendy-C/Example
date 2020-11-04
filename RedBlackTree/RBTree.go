package main

import "fmt"

type rBTree struct {
	root *rBNode
	size int
	height int
}

func (r *rBTree)rightRotate(n *rBNode) {
	nl := n.left
	np := n.parent
	nlr := nl.right

	// change nl
	n.color, nl.color = nl.color, n.color
	if np != nil {
		if np.left == n {
			np.left = nl
		} else {
			np.right = nl
		}
	}
	nl.parent = np

	//change n
	nl.right = n
	n.parent = nl

	// change nlr
	n.left = nlr
	if nlr != nil {
		nlr.parent = n
	}

	if n == r.root {
		r.root = nl
	}

}

func (r *rBTree)leftRotate(n *rBNode) {
	np := n.parent
	nr := n.right
	nrl := n.right.left

	// change nr.
	n.color, nr.color = nr.color, n.color
	if np != nil {
		if np.left == n {
			np.left = nr
		} else {
			np.right = nr
		}
	}
	nr.parent = np

	// change n
	nr.left = n
	n.parent = nr

	// change nrl
	n.right = nrl
	if nrl != nil {
		nrl.parent = n
	}

	if r.root == n {
		r.root = nr
	}
}

func (r *rBTree)colorReverse(n *rBNode)  {
	n.right.color, n.left.color = black, black
	if n.color != none {
		n.color = red
	} else {
		r.height++
	}
}

func (r *rBTree)nodeBalance(n *rBNode) *rBNode {
	// 子节点都为红色
	isleft := n.isLeft()
	if n.HasRight() && n.right.color == red {
		if n.HasLeft() && n.left.color == red {
			r.colorReverse(n)
		}
	}

	if n.color == red {
		// 当前节点为左节点且为红色，其右节点为红色
		if isleft {
			if n.HasRight() && n.right.color == red {
				r.leftRotate(n)
				n = n.parent
			}

			// // 当前节点为左节点且为红色，其左节点为红色
			if n.HasLeft() && n.left.color == red {
				r.rightRotate(n.parent)
				r.colorReverse(n)
			}
			// 当前节点为右节点且为红色，其左节点为红色
		} else {
			if n.HasLeft() && n.left.color == red {
				r.rightRotate(n)
				n = n.parent
			}
			if n.HasRight() && n.right.color == red {
				//// 当前节点为右节点且为红色，其右节点为红色
				r.leftRotate(n.parent)
				r.colorReverse(n)
			}
		}
	}

	return n
}

func (r *rBTree)balance(n *rBNode) {
	for n != nil {
		n = r.nodeBalance(n)
		n = n.parent
	}
}

func New() *rBTree {
	return &rBTree{size:0, height:0, root:nil}
}

func (r *rBTree)Add(data interface{}, w int)*rBNode {
	node := new(w, data)
	if r.root == nil {
		r.root = node
		r.height++
	} else {
		var pre *rBNode
		cur := r.root
		for cur != nil {
			pre = cur
			if cur.w > w {
				cur = cur.left
			} else {
				cur = cur.right
			}
		}
		if pre.w > w {
			pre.left = node
		} else {
			pre.right = node
		}
		node.parent = pre
		node.color = red
		r.balance(pre)
	}

	r.size++
	return node
}

func (r *rBTree)delInNotLeaf(n *rBNode) {
	np := n.parent
	// 根节点情况
	if np == nil {
		r.root = nil
		r.height--
		return
	}

	// 非根节点情况
	// 叶子节点且为红色时，不做改动

	// 叶子节点且为黑色
	if n.color == black {
		curn := n
		var s, sl, sr *rBNode
		for ;curn != r.root; np = curn.parent {
			s = curn.GetSibling()
			sl = s.left
			sr = s.right
			// 兄节点为黑色
			if s.color == black {
				if sl == nil && sr == nil || sl.color == black && sr.color == black {
					// 兄子节点都为黑，没有子节点默认为黑
					s.color = red
					if np.color == red {
						np.color = black
						break
					} else {
						curn = curn.parent
					}
				} else if sl != nil && sl.color == red {
					// 兄的左子节点为红
					if !s.isLeft() {
						// 兄为右节点
						r.rightRotate(s)
						r.leftRotate(np)
						s.color = black
					} else {
						// 兄为左节点
						r.rightRotate(s.parent)
						sl.color = black
					}
					break
				} else if sr != nil && sr.color == red {
					// 兄的右子节点为红
					if !s.isLeft() {
						// 兄为右节点
						r.leftRotate(s.parent)
						sr.color = black
					} else {
						// 兄为左节点
						r.leftRotate(s)
						s = sr
						r.rightRotate(np)
						s.color = black
					}
					break
				}
			} else {
				// 兄为红色
				if s.isLeft() {
					r.rightRotate(np)
				} else {
					r.leftRotate(np)
				}
			}
		}
		if curn == r.root {
			r.height--
		}
	}

	// 直接删除n
	np = n.parent
	if np.left == n {
		np.left = nil
	} else {
		np.right = nil
	}
}

func (r *rBTree)delInOneLeaf(n *rBNode) {
	np := n.parent
	hasLeft := n.HasLeft()
	var nc *rBNode
	if hasLeft {
		nc = n.left
	} else {
		nc = n.right
	}

	// 根节点情况
	if np == nil {
		if hasLeft {
			r.root = n.left
		} else {
			r.root = n.right
		}
		r.root.color = none
		r.root.parent = nil
	} else {
		// 非根节点情况
		nc.parent = n.parent
		nc.color = black
		if n.isLeft() {
			np.left = nc
		} else {
			np.right = nc
		}
	}
}

func (r *rBTree)delInAllLeaves(n *rBNode) {
	nr := n.right

	cur := nr
	for cur.left != nil {
		cur = cur.left
	}

	n.w, cur.w = cur.w, n.w
	n.data, cur.data = cur.data, n.data
	if cur.HasRight() {
		r.delInOneLeaf(cur)
	} else {
		r.delInNotLeaf(cur)
	}
}

func (r *rBTree)Del(w int)(n *rBNode)  {
	n = r.Get(w)

	if n == nil {
		return
	}

	hasRight := n.HasRight()
	hasLeft := n.HasLeft()

	if !hasLeft && !hasRight {
		// 叶子节点情况
		r.delInNotLeaf(n)
	} else if hasLeft && hasRight {
		// 两个子节点都不为空
		r.delInAllLeaves(n)
	} else {
		// 一个子节点的情况
		r.delInOneLeaf(n)
	}
	r.size--
	return
}

func (r *rBTree)Get(w int)*rBNode {
	cur := r.root
	for cur != nil {
		if w > cur.w {
			cur = cur.right
		} else if w < cur.w {
			cur = cur.left
		} else {
			return cur
		}
	}

	return nil
}

func (r *rBTree)Mod()*rBNode  {
	return nil
}

func main() {
	var arr = []int{12,1,9,2,0,11,7,19,4,15,18,5,14,13,10,16,6,3,8,17}
	t := New()

	for _,v:=range arr {
		t.Add('A' + v, v)
	}

	for _, v := range arr {
		t.Del(v)
	}

	fmt.Println(t.size, t.height)
}
