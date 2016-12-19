#!/bin/bash

# usage
# ssh user@host "bash -s" < install.sh "version"
# ssh ubuntu@192.168.4.32 "bash -s" < install.sh "v0.1.1"

if [ "$1" == "" ]; then
    version=`date +%s`
else
    version=$1
fi

echo "version: $version"

# # 部署的目标目录
# targetDir="/data/workspace/default/backbone"

# # 当前目录
# dir=$PWD
# # 编译临时目录
# tmpDir="tmp_deploy_${version}"
# # 下载的目标文件
# tarFile="backbone_sensor-${version}.tar.gz"

# # 检测是否存在临时目录
# if [ ! -d $tmpDir ]; then
#     mkdir $tmpDir
# else
#     rm -rf $tmpDir/*
# fi

# # 进入临时目录
# cd $tmpDir

# # 下载文件
# wget "http://192.168.5.46:9003/${tarFile}"

# # 如果下载文件失败，退出
# if [ ! -f "${tarFile}" ]; then
#     echo "${tarFile} download fail"
#     exit 1
# fi

# # 解压
# tar -zxvf $tarFile &>/dev/null

# # cd $dir

# # delete target
# echo "ubuntu" | sudo -S rm -rf $targetDir

# # copy backbone 
# echo "ubuntu" | sudo -S cp -r backbone $targetDir

# cd $dir

# #删除临时目录
# rm -rf $tmpDir

# # restart
# echo "ubuntu" | sudo -S reboot
# echo "ubuntu" | sudo -S supervisorctl stop backbone 2>&1
# echo "ubuntu" | sudo -S supervisorctl start backbone 2>&1
echo "dg2015" | sudo -S docker rm -f backbone
echo "dg2015" | sudo docker pull 192.168.5.46:5000/backbone-server:${version} 
echo "dg2015" | sudo -S docker run -i -t -d --restart=always --name=backbone --net=hosts 192.168.5.46:5000/backbone-server:${version} bash run.sh
