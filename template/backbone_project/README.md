#Deepglint-Backbone

格灵深瞳Libra-T Web Platform

## 部署

对于本项目，Master分支并不是最终的Release分支，Backbone分为Sensor版和Server版，具体的代码分支以及运行部署方式如下所述

### SERVER版本

#### 分支：server_release

#### 编译方式

1. 更新code base 
	
	```
	http://192.168.5.46:9005/pull
	```
2. 编译 
	
	```
	http://192.168.5.46:9005/build?tag={{version}}
	```
	
3. 部署
	
	```
	# 登录服务器，执行
	sudo docker run -t -i -d --net=host --restart=always --name=backbone 192.168.5.46:5000/backbone-server:{{version}} bash run.sh
	```

### SENSOR版本

### 分支：sensor_release

#### 编译方式

1. 更新code base 
	
	```
	http://192.168.5.46:9003/pull
	```
2. 编译 
	
	```
	http://192.168.5.46:9003/build?tag={{version}}
	```
	
3. tar包地址

	```
	http://192.168.5.46:9003/backbone_sensor-{{version}}.tar.gz
	```
	
4. 运行
	
	```
	tar -zxvf backbone_sensor-{{version}}.tar.gz
	cd release
	sudo bash run.sh
	```
	

