# 视频记录截图插件

可通过http请求获取到指定视频的截图（默认jpg格式），常用作视频封面。
9cd7b5545a3bb28033a0a0a2b88b02c.jpg

## 插件地址

https://github.com/3201301734/plugin-snapplug.git

## 默认配置

```yaml
snapplug:
    ffmpeg: "ffmpeg"
```
如果ffmpeg无法全局访问，则可修改ffmpeg路径为本地的绝对路径
## API

### `/snapplug/api/video/cover?type=${type}&videoPath=${videoPath}&format=${format}`
#### 参数说明：

| 参数 | 类型 | 说明 |
|:-----|:--|:----|
| type | string | 视频记录类型，flv、mp4 |
| videoPath | string |视频记录地址，取 [m7s record](https://m7s.live/guide/plugins/record.html#api) 插件 /record/api/list?type=flv 接口返回json的Path值 |
| format | string | 适配截图格式, 取值：jpg（默认）、png |


### 示例
例如 m7s（localhost)中record记录 flv/test/video/1675220979.flv,

- 请求 http://localhost:8080/snapplug/api/video/cover?type=flv&videoPath=test/video/1675220979.flv 获取到该视频记录的截图，同时会在视频相同目录下 生成同名图片，例如：1675220979.jpg，再次获取时则直接返回该图片。
- 请求 http://localhost:8080/snapplug/api/video/cover?type=flv&videoPath=test/video/1675220979.flv&format=png 获取到该视频记录的截图，同时会在视频相同目录下 生成同名图片，例如：1675220979.png，再次获取时则直接返回该图片。
  



