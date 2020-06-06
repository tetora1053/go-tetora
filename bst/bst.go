// 二分探索木パッケージ
package bst

import (
	"fmt"
)

type node struct {
	Data int
	Parent, Left, Right *node
}

type bst struct {
	root *node
}

// 二分探索木オブジェクト取得
func NewBst() bst {
	return bst{}
}

// rootノード取得
func (bst *bst) GetRoot() *node {
	return bst.root
}

// ノード追加
func (bst *bst) Append(data int) (err bool) {
	if bst.root == nil {
		bst.root = &node{Data: data}
		return true
	}

	cur := bst.root
	var p *node

	for cur != nil {
		p = cur
		if data < cur.Data {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	target := &node{Data: data, Parent: p}
	if data < p.Data {
		p.Left = target
	} else {
		p.Right = target
	}

	return true
}

// 中間順巡回
func InOrder(root *node) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Println(root.Data)
	InOrder(root.Right)
}

// 先行順巡回
func PreOrder(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

// 後行順巡回
func PostOrder(root *node) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Println(root.Data)
}

// 最小値取得
func Min(root *node) int {
	cur := root
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur.Data
}

// 最大値取得
func Max(root *node) int {
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur.Data
}

func Search(root *node, searchVal int) bool {
	if root == nil {
		return false
	}

	if searchVal == root.Data {
		return true
	} else if searchVal < root.Data {
		return Search(root.Left, searchVal)
	} else {
		return Search(root.Right, searchVal)
	}
}

// 値からノードを取得
func GetNodeByVal(root *node, val int) (*node,  bool) {
	if root == nil {
		return nil, false
	}

	if val == root.Data {
		return root, true
	} else if val < root.Data {
		return GetNodeByVal(root.Left, val)
	} else {
		return GetNodeByVal(root.Right, val)
	}
}

// ノード削除
func (bst *bst) Delete(val int) bool {
	root := bst.GetRoot()
	target, ok := GetNodeByVal(root, val)
	if ok == false {
		return false
	}

	var child, delNode *node
	if target.Left == nil || target.Right == nil {
		delNode = target
		if target.Left != nil {
			child = target.Left
		} else {
			child = target.Right
		}
	} else {
		delNode, _ = GetNodeByVal(target.Right, Min(target.Right))
		target.Data = delNode.Data
	}
	if delNode.Parent.Left == delNode {
		delNode.Parent.Left = child
	} else {
		delNode.Parent.Right = child
	}
	if child != nil {
		child.Parent = delNode.Parent
	}
	return true
}



