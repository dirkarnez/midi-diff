@echo off

set PATH=^
%USERPROFILE%\Downloads\PortableGit\bin;

@REM git gui
@REM pause
@REM set /p commit="Enter commit hash: "
@REM git difftool -d -x "midiffweb.exe" "%commit%" HEAD

git difftool -d -x "midiffweb.exe" fea62ee8ca07185542ca5c9a4df83ff186b199cc HEAD

