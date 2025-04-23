cd w6/cmd/_5

beego: https://github.com/beego/beego

go install github.com/beego/bee/v2@latest

bee new web_bp

cd web_bp  
go mod tidy

custom_controller from _5/beego_examples.go -> web_bp/  controllers/default.go

add route for custom_controller in web_bp/routers/router.go  
example in _5/beego.go  

cd ..

bee api user  
cd user  
bee run -downdoc=true -gendoc=true  

go mod tidy  

http://localhost:8080/swagger/ check in browser

gin: https://github.com/gin-gonic/gin

go run gin.go

http://localhost:8080/user/kbtu in browser

http://localhost:8080/admin?key=123 in browser

http://localhost:8080/admin?user_key=123 in browser

http://localhost:8080/user/kbtu