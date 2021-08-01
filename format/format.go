package format

import (
	"github.com/dalaomai/vdk/av/avutil"
	"github.com/dalaomai/vdk/format/aac"
	"github.com/dalaomai/vdk/format/flv"
	"github.com/dalaomai/vdk/format/mp4"
	"github.com/dalaomai/vdk/format/rtmp"
	"github.com/dalaomai/vdk/format/rtsp"
	"github.com/dalaomai/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
