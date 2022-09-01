#!/bin/bash

### 构建要求
### 1. go >= 1.18
### 2. node >= 8
baseDir=`echo $PWD`
projectDir=`echo $baseDir/web_complier`
deployDir=`echo $baseDir/deploy`

funcBuildServer() {
    echo 'server module building...'
    export GOPROXY=https://goproxy.cn
    cd $projectDir
    go mod download
    make buildpro
    echo 'server module building...finished'
}

funcTouchDir() {
    if [ ! -d "$1" ]; then
        mkdir -p $1
    fi
}

funcCleanBuild() {
    rm -rf $deployDir/web_complier/*

    funcTouchDir $deployDir/web_complier
    funcTouchDir $deployDir/logs

    cp $projectDir/apiserver $deployDir/web_complier/
    cp -r $projectDir/configs $deployDir/web_complier/

    cd $deployDir/web_complier
    export DQENV=pro
    pm2 stop apiserver
    pm2 start apiserver
}

funcBuildServer
funcCleanBuild

echo 'server done!!!'