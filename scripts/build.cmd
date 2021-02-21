@ECHO OFF
setlocal
set "SCRIPTS_DIR=%~dp0"
set "PROJECT_DIR=%SCRIPTS_DIR%\.."
set "BUILD_DIR=%PROJECT_DIR%\build"
cd "%PROJECT_DIR%\cmd\rtmor"
go build -o "%BUILD_DIR%"
cd "%PROJECT_DIR%"
endlocal
