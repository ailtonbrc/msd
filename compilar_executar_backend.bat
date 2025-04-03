@echo off
setlocal

echo ==============================
echo    Compilando o backend Go...
echo ==============================

cd /d D:\Desenvolvimento_React\MSD\server
go build -o clinica_server.exe ./cmd/api

if %errorlevel% neq 0 (
    echo Erro na compilação!
    pause
    exit /b %errorlevel%
)

echo ==============================
echo Executando o servidor...
echo ==============================

clinica_server.exe

endlocal
pause