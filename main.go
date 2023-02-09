package snapplug

import (
	"bytes"
	_ "embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"strings"

	. "m7s.live/engine/v4"
	"m7s.live/engine/v4/config"
)

//go:embed default.yaml
var defaultYaml DefaultYaml

type SnapplugConfig struct {
	DefaultYaml
	config.HTTP
	FFmpeg string // ffmpeg的路径
}

var conf = &SnapplugConfig{
	DefaultYaml: defaultYaml,
}

func (p *SnapplugConfig) OnEvent(event any) {
	switch event.(type) {
	case FirstConfig: //插件初始化逻辑
	// case Config: //插件热更新逻辑
	case *Stream: //按需拉流逻辑
	case SEwaitPublish: //由于发布者掉线等待发布者
	case SEpublish: //进入发布状态
	// case SEsubscribe: //订阅者逻辑
	case SEwaitClose: //由于最后一个订阅者离开等待关闭流
	case SEclose: //关闭流
	case UnsubscribeEvent: //订阅者离开
		// case ISubscribe: //订阅者进入
	}
}

// var plugin = InstallPlugin(new(SnapplugConfig))
var _ = InstallPlugin(conf)

func (p *SnapplugConfig) API_video_cover(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	// videoType := query.Get("type")
	// videoPath := query.Get("videoPath")
	var err error
	var imageFile *os.File
	var fi fs.FileInfo

	wd, _ := os.Getwd()
	// fmt.Println("工作目录: " + wd)
	// videoPath := wd + "/record/flv/test/video/1675220979.flv"
	videoPath := fmt.Sprintf("%s/record/%s/%s", wd, query.Get("type"), query.Get("videoPath"))
	imagePath := videoPath[:strings.Index(videoPath, ".")] + ".jpg"

	if fi, err = getVideoCover(videoPath, imagePath); err == nil && fi != nil {
		buff := make([]byte, fi.Size())
		if imageFile, err = os.Open(imagePath); err == nil {
			if _, err = imageFile.Read(buff); err == nil {
				w.Header().Set("Content-Type", "image/jpeg")
				_, err = w.Write(buff)
			}
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getVideoCover(videoPath, imagePath string) (fi os.FileInfo, err error) {
	fi, err = PathExists(imagePath)
	if err != nil {
		fmt.Printf("err0=%s\n", err)
		return fi, err
	}
	if fi == nil {
		cmd := exec.Command(conf.FFmpeg, "-i", videoPath, "-y", "-f", "image2", "-frames", "1", imagePath)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = os.Stderr
		err = cmd.Start()
		if err != nil {
			fmt.Printf("err1=%s\n", err)
			return fi, err
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Printf("err2=%s\n", err) // err2=exit status 1
			return fi, err
		}
		// fmt.Printf("out=%s\n", out.String())
		fi, err = PathExists(imagePath)
	}
	return fi, err
}

func PathExists(path string) (fi fs.FileInfo, err error) {
	fi, err = os.Stat(path)
	if err == nil {
		return fi, nil
	}
	if os.IsNotExist(err) {
		return fi, nil
	}
	return fi, err
}
