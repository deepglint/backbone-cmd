# 如何开始
---------


## 介绍

格灵深瞳的LIBRA-T2产品可通过API方式输出事件、设备状态等信息。

轻量化的Sensor端可输出非业务型API供二次开发使用；业务API通过后端Smart Server进行访问。

第三方平台通过API获取到事件时间后可从NVR中调取对应时间的视频回放或者对应Sensor的视频直播。

产品的API接口具有跨网的灵活性和一致性，对外开发的API接口具备WEB说明文档。

<br /><br />

## API 用户验证

设备中的API在浏览器中调用的时候会返回错误值，表示Auth拒绝，必须使用BasicAuth来进行访问。

我们设置的用户名和密码是admin和admin。

在HTTP请求时, 首先将admin:admin这个字符串进行base64编码为:

```
YWRtaW46YWRtaW4= 
```

在HTTP请求的Header中增加一个Key Value就可以请求成功了:

```
 Key:Authorization
 Val:Basic YWRtaW4
```
<br /><br />
## 示例

下面通过一个简单的示例来演示LIBRA-T2的API使用：

这个示例展示了通过uptime这个API来获取设备已经运行的时间等信息，请求的URL如下：

```
 http://[IP]:[PORT]/api/get/uptime
```

返回值示例如下（详见uptime API 文档）：

```
 09:03:49 up 1 day

```
<br /><br />
## 一种典型的产品部署案例图

![LibraT2](public/md/LibraT2.png)