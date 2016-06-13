#!/bin/bash


if [ "$1" == "" ]; then
    version=`date +%s`
else
    version=$1
fi

echo "version: server-$version"

dir=$PWD
fileName="backbone_server-${version}"
releaseFileName="release_server-${version}"
tmpDir="tmp_server-${version}"

echo "gulp build"
#gulp uglify

echo "cp floder"
cp -r ${dir} ../${fileName}

echo "delete git"
rm -rf ../${fileName}/.git

echo "delete node_modules"
rm -rf ../${fileName}/node_modules

echo "delete components"
rm -rf ../${fileName}/public/bower_components
rm -rf ../${fileName}/public/components

echo "delete backbone_*.tar.gz"
rm -rf ../${fileName}.tar.gz

echo "revel package"
revel build github.com/deepglint/${fileName} ../${releaseFileName} &>/dev/null


#echo "build arm"
#GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0  go build -o ../${releaseFileName}/${fileName} ../${fileName}/app/tmp/main.go


echo ${fileName} > ../${releaseFileName}/Version

echo "tar"

# 进入根目录
cd ..

# 创建临时文件夹

if [ ! -d $tmpDir ]; then
    mkdir $tmpDir
fi

# 复制编译后的目录到backbone

cp -r ${releaseFileName} ${tmpDir}/backbone

# 进入临时目录
cd ${tmpDir}


cp ${dir}/Dockerfile ./

# 打包成docker

echo "docker build"

echo "123"|sudo -S docker build -t 192.168.5.46:5000/backbone-server:${version} . 2>&1

echo "123"|sudo -S docker push 192.168.5.46:5000/backbone-server:${version} 2>&1

rm -rf ../${fileName}
rm -rf ../${releaseFileName}
rm -rf ../${tmpDir}

echo "build done"