// Autogenerated with msg-import, do not edit.
package trajectory_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"time"
)

type MultiDOFJointTrajectoryPoint struct {
	msgs.Package  `ros:"trajectory_msgs"`
	Transforms    []geometry_msgs.Transform
	Velocities    []geometry_msgs.Twist
	Accelerations []geometry_msgs.Twist
	TimeFromStart time.Duration
}
