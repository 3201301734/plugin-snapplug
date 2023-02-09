# 视频记录截图插件

可通过http请求获取到指定视频的截图（jpg格式），常用作视频封面。

## 插件地址

https://github.com/3201301734/plugin-snapplug.git

## 插件引入
```go

```
## 默认配置

```yaml
snapplug:
    ffmpeg: "ffmpeg"
```
如果ffmpeg无法全局访问，则可修改ffmpeg路径为本地的绝对路径
## API

### `/snapplug/api/video/cover?type=${type}&videoPath=${videoPath}`

type       视频记录类型，flv、mp4

videoPath  视频记录地址，取record插件 /record/api/list?type=flv 接口返回json的Path值,  参见：https://m7s.live/guide/plugins/record.html#api

例如m7s（localhost)中记录 flv/test/video/1675220979.flv,
可以通过http://localhost:8080/snapplug/api/video/cover?type=flv&videoPath=test/video/1675220979.flv 获取到该视频记录的截图，同时会在视频相同目录下 生成同名图片，例如：1675220979.jpg，再次获取时则直接返回该图片。


