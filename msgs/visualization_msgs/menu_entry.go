// Autogenerated with msg-import, do not edit.
package visualization_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type MenuEntry struct {
	msgs.Package `ros:"visualization_msgs"`
	Id           uint32
	ParentId     uint32
	Title        string
	Command      string
	CommandType  uint8
}
