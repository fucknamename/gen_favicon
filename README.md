# gen_favicon
Text generation favicon.ico icon

# package
go build -ldflags="-s -w" -o ico.exe main.go && upx -9 ico.exe


# use
ico.exe -t 发  
