#gin 实现的博客系统
##特点
###前后端分离



#参考资料
1.https://studygolang.com/articles/12408

#接口列表(v1)
##用户 user
1.注册 register
2.登录 login
3.注销 logout
4.修改 edit

##文章




//string到int  
int,err:=strconv.Atoi(string)  
//string到int64  
int64, err := strconv.ParseInt(string, 10, 64)  
//int到string  
string:=strconv.Itoa(int)  
//int64到string  
string:=strconv.FormatInt(int64,10)
//string到float32(float64)
float,err := strconv.ParseFloat(string,32/64)
//float到string
string := strconv.FormatFloat(float32, 'E', -1, 32)


