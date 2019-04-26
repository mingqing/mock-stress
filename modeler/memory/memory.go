package memory

import (
	"time"
)

// Allocate 用于分配内存
// totalSize 总分配的大小，stepSize 每次分配的步长，单位均为byte
// runDuration 内存需占用多长时间
func Allocate(totalSize, stepSize int64, runDuration time.Duration) {
	var buffer [][]byte

	for i := int64(1); i*stepSize <= totalSize; i++ {
		temp := make([]byte, stepSize)
		for j := range temp {
			temp[j] = 0
		}
		buffer = append(buffer, temp)

		time.Sleep(runDuration)
	}

	buffer = nil
}
