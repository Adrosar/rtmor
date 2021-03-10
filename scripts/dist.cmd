@ECHO OFF
setlocal

set "SCRIPTS_DIR=%~dp0"
set "PROJECT_DIR=%SCRIPTS_DIR%.."
set "MAIN_DIR=%PROJECT_DIR%\cmd\rtmor"

cd "%MAIN_DIR%"

set GOOS=linux
set GOARCH=386
go build -o "%PROJECT_DIR%\build\%GOOS%-%GOARCH%\rtmor"

set GOOS=linux
set GOARCH=amd64
go build -o "%PROJECT_DIR%\build\%GOOS%-%GOARCH%\rtmor"

set GOOS=windows
set GOARCH=386
go build -o "%PROJECT_DIR%\build\%GOOS%-%GOARCH%\rtmor.exe"

set GOOS=windows
set GOARCH=amd64
go build -o "%PROJECT_DIR%\build\%GOOS%-%GOARCH%\rtmor.exe"

set GOOS=darwin
set GOARCH=amd64
go build -o "%PROJECT_DIR%\build\%GOOS%-%GOARCH%\rtmor"

cd "%PROJECT_DIR%"
endlocal
