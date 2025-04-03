@echo off
setlocal

echo ==============================
echo 📦 Preparando commit para o Git...
echo ==============================

cd /d D:\Desenvolvimento_React\MSD

:: Adiciona todos os arquivos modificados, novos e deletados
git add .

:: Define mensagem de commit automática com data/hora
set COMMIT_MSG=Atualização em %date% %time%
git commit -m "%COMMIT_MSG%"

:: Faz push para a branch principal (main)
git push origin main

echo.
echo ✅ Código enviado com sucesso!
pause
endlocal