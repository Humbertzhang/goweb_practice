A restful go web service practice
Written with net/http, encoding/json and gorm

~~Signup~~<br>
~~Signin~~ <br>
Write Post <br>
Write Comments for Post <br>

# Note
## go run *.go 
当main.go运行时需要使用同一个package下其他函数时，使用这个命令来运行

## Json
JSON结构体属性首字母必须要大写，否则不会被反射检测到，则无法完成预期功能
Db也同理，否则不会被gorm的AutoMigrate创建

## Response Writer
 Header returns the header map that will be sent by
 WriteHeader. Changing the header after a call to
 WriteHeader (or Write) has no effect unless the modified
 headers were declared as trailers by setting the
 "Trailer" header before the call to WriteHeader (see example).
 To suppress implicit response headers, set their value to nil.
 From https://golang.org/pkg/net/http/#ResponseWriter   
