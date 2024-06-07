# TikTok弹幕信息流抓取

项目灵感来源于：https://github.com/qtgolang/SunnyNet

项目proto来源:https://github.com/isaackogan/TikTokLive

![20240606192223.jpg](image%2FREADME%2F20240606192223.jpg)
![20240606192150.jpg](image%2FREADME%2F20240606192150.jpg)
具体实现操作.通过代理本地10808端口
浏览器安装SwitchyOmega插件，设置代理为10808端口
然后如果已经无需科学上网环境可直接运行软件访问tk直播间即可
否则话SetGlobalProxy设置你的上游代理地址.具体可参看main.go.
写的匆忙.忙着扣取tk的proto文件.只实现了弹幕文本的消息抓取.期待日后更新.
后期可能会考虑做成https://github.com/jwwsjlm/douyinLive 类似这样.抓取之后推送给ws的客户端.使得其他软件也可接入
