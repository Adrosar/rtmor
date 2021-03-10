@ECHO OFF
setlocal
set "SCRIPTS_DIR=%~dp0"
set "PROJECT_DIR=%SCRIPTS_DIR%.."
set "BUILD_DIR=%PROJECT_DIR%\build"
set "MAIN_DIR=%PROJECT_DIR%\cmd\rtmor"
cd "%MAIN_DIR%"
go build -o "%BUILD_DIR%" && "%BUILD_DIR%\rtmor.exe" %*
cd "%PROJECT_DIR%"
endlocal