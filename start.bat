@echo off
REM 设置Redis的安装路径
set REDIS_HOME=D:\WorkSpcae\middlerware\Redis-x64-5.0.14.1

REM 启动Redis服务器
%REDIS_HOME%\redis-server.exe %REDIS_HOME%\redis.windows.conf
