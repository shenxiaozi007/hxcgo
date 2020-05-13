package pkg

import (
	"github.com/huangxinchun/hxcgo/admin/app/proto/adminprivilegeproto"
	"strings"
)

type PrivilegeURINode struct {
	data     *adminprivilegeproto.Privilege
	children map[string]*PrivilegeURINode
}

func PrivilegeURITree() *PrivilegeURINode {
	root := &PrivilegeURINode{
		data:     nil,
		children: make(map[string]*PrivilegeURINode),
	}
	privileges, err := privilegeService.FindAll()

	if err != nil {
		return root
	}

	for _, privilege := range privileges {
		arr := strings.Split(strings.Trim(privilege.URIRule, "/"), "/")
		if len(arr) < 1 {
			break
		}

		parent := root
		for i := 0; i < len(arr); i++ {
			sub := arr[i]
			_, ok := parent.children[sub]
			if !ok {
				parent.children[sub] = &PrivilegeURINode{
					data:     nil,
					children: make(map[string]*PrivilegeURINode),
				}
			}

			parent = parent.children[sub]
		}

		parent.children["*"] = &PrivilegeURINode{
			data:     privilege,
			children: make(map[string]*PrivilegeURINode),
		}

	}

	return root
}

func (pn *PrivilegeURINode) Match(uri string) (*adminprivilegeproto.Privilege, bool) {
	pattern := strings.Split(strings.Trim(uri, "/"), "/")

	parentNode := pn
	var data *adminprivilegeproto.Privilege
	for i := 0; i < len(pattern); i++ {
		n, ok := parentNode.children[pattern[i]]
		if !ok {
			n, ok = parentNode.children["*"]
			if ok {
				data = n.data
			}
			break
		}
		parentNode = n
		if i == len(pattern)-1 {
			n, ok = parentNode.children["*"]
			if ok {
				data = n.data
			}
			break
		}

	}

	return data, data != nil
}
