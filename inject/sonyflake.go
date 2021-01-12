package inject

import (
	"os"
	"strconv"

	"github.com/sony/sonyflake"
)

// Sonyflake produces a sonyflake ID generator dependency.
func Sonyflake() *sonyflake.Sonyflake {
	var sfs sonyflake.Settings
	var serverID uint16
	if val, err := strconv.ParseInt(os.Getenv("SERVER_ID"), 10, 16); err == nil {
		serverID = uint16(val)
	}
	if serverID > 0 {
		sfs.MachineID = func() (uint16, error) {
			return serverID, nil
		}
	}

	return sonyflake.NewSonyflake(sfs)
}
