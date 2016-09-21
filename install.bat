@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end
#冒号和ok之间不应该有空格，但是放在一起总是会被wordpress转成一个表情
: ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

gofmt -w src

go install server/studygolang

:end
echo finished
pause