package pkg

import (
	"github.com/huangxinchun/hxcgo/admin/app/proto/adminprivilegeproto"
	"bytes"
	"encoding/gob"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var sessionKey = "admin"

type Session struct {
	ID         uint
	Name       string
	Email      string
	Mobile     string
	Avatar     string
	State      uint8
	GroupID    uint
	LoginAt    string
	LoginIP    string
	RoleIDs    []uint
	ActiveTime string
	iSession   sessions.Session
}

func NewSession(c *gin.Context) *Session {
	sess := sessions.Default(c)
	val := sess.Get(sessionKey)

	s := &Session{
		ID:      0,
		Name:    "",
		Email:   "",
		Mobile:  "",
		Avatar:  "",
		State:   0,
		GroupID: 0,
		LoginAt: "",
		LoginIP: "",
	}

	if val != nil {
		if buf, ok := val.([]byte); ok {
			reader := bytes.NewBuffer(buf)
			dec := gob.NewDecoder(reader)

			dec.Decode(s)
		}
	}

	s.iSession = sess
	return s
}

func (s *Session) Save() error {
	var writer bytes.Buffer

	enc := gob.NewEncoder(&writer)
	err := enc.Encode(s)
	if err != nil {
		return err
	}

	s.iSession.Set(sessionKey, writer.Bytes())
	return s.iSession.Save()
}

//判断用户是否授权访问
func (s *Session) IsGranted(uri string) bool {
	privilege, exists := PrivilegeURITree().Match(uri)
	log.Println(privilege, exists)
	if !exists {
		return false
	}
	for _, rid := range s.RoleIDs {
		idMap, err := roleService.PrivilegeIDMap(rid)
		log.Printf("%#v", idMap)
		if err != nil {
			continue
		}
		if _, ok := idMap[privilege.ID]; ok {
			return true
		}
	}
	return false
}

//获取菜单
func (s *Session) Menus() []*Node {
	menus, err := privilegeService.Menus()
	if err != nil {
		return nil
	}

	idMap := map[uint]bool{}
	for _, rid := range s.RoleIDs {
		privilegeIDs, err := roleService.PrivilegeIDs(rid)
		if err != nil {
			continue
		}

		for _, v := range privilegeIDs {
			idMap[v] = true
		}
	}

	var res []*adminprivilegeproto.Privilege
	for _, menu := range menus {
		if _, ok := idMap[menu.ID]; ok {
			res = append(res, menu)
		}
	}

	return newPrivilegeTree(res)
}
