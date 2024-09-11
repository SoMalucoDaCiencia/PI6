package src

import (
	"encoding/binary"
	"math"
	"time"
)

func GetTimestamps(year, month, day int) (time.Time, time.Time) {
	start := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, time.Month(month), day, 23, 59, 59, 999999999, time.UTC)
	return start, end
}

func FloatsAsUUID(f1, f2 float64) string {
	var uuid []byte
	uuid = binary.BigEndian.AppendUint64(uuid, math.Float64bits(f1))
	uuid = binary.BigEndian.AppendUint64(uuid, math.Float64bits(f2))
	return string((uuid[:]))
}

func FloatsFromUUID(uuid []byte) (float64, float64) {
	f1 := math.Float64frombits(binary.BigEndian.Uint64(uuid[:8]))
	f2 := math.Float64frombits(binary.BigEndian.Uint64(uuid[8:16]))
	return f1, f2
}
