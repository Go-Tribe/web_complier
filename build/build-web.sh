#!/bin/bash

### 构建要求
### 1. go >= 1.18
### 2. node >= 8

baseDir=`echo $PWD`
projectDir=`echo $baseDir/web_complier`
serverDir=`echo $projectDir/web`
deployDir=`echo $baseDir/deploy`

funcBuildSite() {
    echo 'web module building...'
    cd $serverDir
    npm install --registry=https://registry.npm.taobao.org
    npm run build
    echo 'web module building...finished'
}

funcTouchDir() {
    if [ ! -d "$1" ]; then
        mkdir -p $1
    fi
}

funcCleanBuild() {
    rm -rf $deployDir/web/*

    funcTouchDir $deployDir/web
    funcTouchDir $deployDir/logs

     cp -r $serverDir/.nuxt $deployDir/web/
     cp -r $serverDir/static $deployDir/web/
     cp -r $serverDir/nuxt.config.js $deployDir/web/
     cp -r $serverDir/package.json $deployDir/web/
     cp -r $serverDir/node_modules $deployDir/web/
     cp -r $serverDir/web_complier.js $deployDir/web/
     cp -r $serverDir/ecosystem.config.js $deployDir/web/

     cd $deployDir/web
     npm install --registry=https://registry.npm.taobao.org
     npm run start
}

funcBuildSite
funcCleanBuild

echo 'web done!!!'