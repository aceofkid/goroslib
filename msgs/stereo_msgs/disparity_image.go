// Autogenerated with msg-import, do not edit.
package stereo_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/sensor_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type DisparityImage struct {
	msgs.Package `ros:"stereo_msgs"`
	Header       std_msgs.Header
	Image        sensor_msgs.Image
	F            float32
	T            float32
	ValidWindow  sensor_msgs.RegionOfInterest
	MinDisparity float32
	MaxDisparity float32
	DeltaD       float32
}
