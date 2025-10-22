@echo off
goctl api swagger --dir . --filename swagger --api admin.api --yaml
go run .
pause
