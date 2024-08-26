# gen_favicon
generation text to favicon.ico  

# package
go build -ldflags="-s -w" -o ico.exe main.go && upx -9 ico.exe  
go build -ldflags="-s -w" && upx -9 gen_favicon.exe && del ico.exe && ren gen_favicon.exe ico.exe  

# use
ico.exe -t=发 -b=ff9900 -f=000000  

![alt text](favicon.ico)