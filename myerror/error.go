package myerror

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
)

/**
 * 1.一般而言，被调函数f(x)会将调用信息和参数信息作为发生错误时的上下文放在错误信息中并返回给调
 * 用者，调用者需要添加一些错误信息中不包含的信息，比如添加url到html.Parse返回的错误中
 * 2. 重新尝试失败的操作
 * 3. 输出错误信息并结束程序
 * 4. 我们只需要输出错误信息就足够了，不需要中断程序的运行。我们可以通过log包提供函数
 * 5. 直接忽略掉错误
 */

func Err1(url string) (error){
	resp,err :=  http.Get(url)
	fmt.Sprint(resp.Body)
	if err != nil{
		return err
	}
    doc,error :=	html.Parse(resp.Body)
    resp.Body.Close()
    if error != nil{
    	return fmt.Errorf("parsing %s as html: %v",url,error)
	}
	fmt.Sprintln(doc.Data)
    return nil
}