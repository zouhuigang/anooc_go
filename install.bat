@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end
#ð�ź�ok֮�䲻Ӧ���пո񣬵��Ƿ���һ�����ǻᱻwordpressת��һ������
: ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

gofmt -w src

go install server/studygolang

:end
echo finished
pause