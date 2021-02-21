@ECHO OFF
setlocal
set "SCRIPTS_DIR=%~dp0"
set "PROJECT_DIR=%SCRIPTS_DIR%.."
cd "%PROJECT_DIR%\cmd\rtmor"

set GOOS=linux
set GOARCH=386
go build -o "%PROJECT_DIR%\build\linux-386\rtmor"

set GOOS=linux
set GOARCH=amd64
go build -o "%PROJECT_DIR%\build\linux-amd64\rtmor"

set GOOS=windows
set GOARCH=386
go build -o "%PROJECT_DIR%\build\windows-386\rtmor.exe"

set GOOS=windows
set GOARCH=amd64
go build -o "%PROJECT_DIR%\build\windows-amd64\rtmor.exe"

set GOOS=darwin
set GOARCH=amd64
go build -o "%PROJECT_DIR%\build\darwin-amd64\rtmor"

cd "%PROJECT_DIR%"
endlocal
