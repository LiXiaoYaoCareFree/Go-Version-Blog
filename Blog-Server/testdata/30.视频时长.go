package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

// BoxHeader 信息头
type BoxHeader struct {
	Size       uint32
	FourccType [4]byte
	Size64     uint64
}

func main() {
	dirPath := "\"F:\\迁移文件夹\\视频录制\\屏幕录制 2025-03-14 232700.mp4\""
	//dirPath := "H:\\gvb_video\\gvb_server"
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var s uint32
	var videoList []string
	for _, entry := range dir {
		videoList = append(videoList, entry.Name())
	}
	sort.Slice(videoList, func(i, j int) bool {
		n1, _ := strconv.Atoi(strings.Split(videoList[i], ".")[0])
		n2, _ := strconv.Atoi(strings.Split(videoList[j], ".")[0])
		return n1 < n2
	})

	for _, name := range videoList {
		filePath := path.Join(dirPath, name)
		file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println(err)
			break
		}
		duration, err := GetMP4Duration(file)
		if err != nil {
			fmt.Println(err)
			break
		}
		s += duration
		fmt.Println(name, ResolveTime(duration))
	}
	fmt.Printf("视频总时长 %s  视频总个数 %d\n", ResolveTime(s), len(videoList))

}

// GetMP4Duration 获取视频时长，以秒计
func GetMP4Duration(reader io.ReaderAt) (lengthOfTime uint32, err error) {
	var info = make([]byte, 0x10)
	var boxHeader BoxHeader
	var offset int64 = 0
	// 获取moov结构偏移
	for {
		_, err = reader.ReadAt(info, offset)
		if err != nil {
			return
		}
		boxHeader = getHeaderBoxInfo(info)
		fourccType := getFourccType(boxHeader)
		if fourccType == "moov" {
			break
		}
		// 有一部分mp4 mdat尺寸过大需要特殊处理
		if fourccType == "mdat" {
			if boxHeader.Size == 1 {
				offset += int64(boxHeader.Size64)
				continue
			}
		}
		offset += int64(boxHeader.Size)
	}
	// 获取moov结构开头一部分
	moovStartBytes := make([]byte, 0x100)
	_, err = reader.ReadAt(moovStartBytes, offset)
	if err != nil {
		return
	}
	// 定义timeScale与Duration偏移
	timeScaleOffset := 0x1C
	durationOffest := 0x20
	timeScale := binary.BigEndian.Uint32(moovStartBytes[timeScaleOffset : timeScaleOffset+4])
	Duration := binary.BigEndian.Uint32(moovStartBytes[durationOffest : durationOffest+4])
	lengthOfTime = Duration / timeScale
	return
}

// getHeaderBoxInfo 获取头信息
func getHeaderBoxInfo(data []byte) (boxHeader BoxHeader) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &boxHeader)
	return
}

// getFourccType 获取信息头类型
func getFourccType(boxHeader BoxHeader) (fourccType string) {
	fourccType = string(boxHeader.FourccType[:])
	return
}

func ResolveTime(seconds uint32) string {
	duration := time.Duration(seconds) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds = uint32(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
