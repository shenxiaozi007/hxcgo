package pkg

import (
	"github.com/huangxinchun/hxcgo/admin/app/proto/adminprivilegeproto"
)

type Node struct {
	IsLeaf   bool
	Data     *adminprivilegeproto.Privilege
	Children []*Node
}

func PrivilegeTree() ([]*Node, error) {
	privileges, err := privilegeService.FindAll()
	if err != nil {
		return nil, err
	}

	return newPrivilegeTree(privileges), nil
}

func newPrivilegeTree(privileges []*adminprivilegeproto.Privilege) []*Node {
	var root []*Node
	tmpNodes := map[uint]*Node{}
	for _, privilege := range privileges {
		node, ok := tmpNodes[privilege.ID]
		if ok {
			node.Data = privilege
		} else {
			node = &Node{
				IsLeaf:   true,
				Data:     privilege,
				Children: nil,
			}
			tmpNodes[privilege.ID] = node
		}

		if privilege.PID == 0 {
			root = append(root, node)
		} else {
			pNode, ok := tmpNodes[privilege.PID]
			if !ok {
				pNode = &Node{
					IsLeaf:   true,
					Data:     nil,
					Children: nil,
				}
				tmpNodes[privilege.PID] = pNode
			}
			pNode.IsLeaf = false
			pNode.Children = append(pNode.Children, node)
		}
	}

	return root
}
