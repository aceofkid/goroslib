// Autogenerated with msg-import, do not edit.
package nav_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type OccupancyGrid struct {
	Header std_msgs.Header
	Info   MapMetaData
	Data   []msgs.Int8
}
