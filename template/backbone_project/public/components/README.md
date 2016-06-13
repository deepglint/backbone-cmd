描述：
js依赖：
css依赖：
子组件：


## advancedConfig

组件描述：服务设置

## algorithm

组件描述：视觉设置

## algorithmesetview

组件描述：视觉设置
子模块：`algorithm`

## apiDetail

描述：API详情查看
js依赖：bootstrap
css依赖：bootstrap
子组件：无

## apiGuide

描述：API指导手册
js依赖：common
css依赖：common
子组件：apiGuideSearch，apiDetail

## apiGuideSearch

描述：API搜索
js依赖：common
css依赖：common
子组件：无

## bgmodeldisplay

描述：三维场景训练效果图
js依赖：THREE,TrackballControls,stats
css依赖：common
子组件：无

## bgmodeleset

描述：场景设置
js依赖：common
css依赖：common
子组件：["bgmodeldisplay", "bgmodeltrain", "floormodeltrain", "floormodelslider"]

## bgmodeltrain

描述：背景模型训练
js依赖：common
css依赖：common
子组件：无

## changePwd

描述：修改密码
js依赖：common
css依赖：common
子组件：无

## config

描述：服务器设置集合
js依赖：common
css依赖：common
子组件：["systemInfoView", "networkConfig",  "bgmodelset", "timeconfig", "eventConfig", "algorithmsetview"]

## developerSetting

描述：二次开发设置
js依赖：common
css依赖：common
子组件：无

## deviceDescription

描述：设备描述
js依赖：common
css依赖：common
子组件：无

## eventConfig

描述：事件设置
js依赖：bootstrap, moment
css依赖：bootstrap
子组件：["timeFrameConfig", "timeFrameAwareConfig", "latchConfig", "overstepConfig", "regionConfig"]

## floormodelslider

描述：地面模型地面高度设置
js依赖：common
css依赖：common
子组件：无

## floormodeltrain

描述：地面模型训练
js依赖：common
css依赖：common
子组件：无

## hlsplayer

描述：位置
js依赖：common
css依赖：common
子组件：无
引用：无

## iframe

描述：展示iframe页面
js依赖：common
css依赖：common
子组件：无
引用：无

## ipFrom

描述：ip输入
js依赖：common
css依赖：common
子组件：无
引用：无

## latchConfig

描述：加钞间箱门识别点添加
js依赖：randomColor
css依赖：common
子组件：无
引用：无

## localConfig

描述：本地服务设置
js依赖：common
css依赖：common
子组件：["systemInfoView", "networkConfig", "bgmodelset", "timeconfig"]

## login

描述：登录
js依赖：common
css依赖：common
子组件：无

## markdown

描述：markdow阅读器
js依赖：marked
css依赖：common
子组件：无

## navbar

描述：导航条
js依赖：common
css依赖：common
子组件：无

## networkConfig

描述：网络设置
js依赖：jquery.inputmask
css依赖：common
子组件：无

## overstepConfig

描述：自定义越线设置
js依赖：randomColor
css依赖：common
子组件：["timeFrameConfig"]
引用：无

## PCOnly

描述：PC访问引导
js依赖：common
css依赖：common
子组件：无

## regionConfig

描述：区域入侵/区域人数异常设置
js依赖：common
css依赖：common
子组件：无
引用：无

## rtmpPlayer

描述：实时监控
js依赖：swfobject
css依赖：common
子组件：无

## serverConfig

描述：服务器设置
js依赖：common
css依赖：common
子组件：["config"]

## sidebar

描述：左侧导航
js依赖：common
css依赖：common
子组件：无

## storageConfig

描述：存储设置
js依赖：common
css依赖：common
子组件：无

## systemInfo

描述：系统信息展示
js依赖：common
css依赖：common
子组件：无

## systemInfoView

描述：系统设置
js依赖：common
css依赖：common
子组件：["systemInfo", "systemStatus", "rtmpPlayer", "deviceDescription"]

## systemStatus

描述：系统状态检测
js依赖：common
css依赖：common
子组件：无

## timeConfig

描述：时间设置
js依赖：common
css依赖：common
子组件：["vueDatetimePicker"]

## timeFrameAwareConfig

描述：时间框架感知设置
js依赖：common
css依赖：common
子组件：无
引用：无

## timeFrameConfig

描述：时间框架设置
js依赖：common
css依赖：common
子组件：无
引用：无


