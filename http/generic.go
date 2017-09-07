package http

import (
	"net/http"
	"fmt"
	"time"
	"log"
	"huanggh.site/learning/common/util"
	"huanggh.site/learning/api-gateway/conf"
	"strconv"
)

func init(){
	go func(){
		fmt.Println("start http server...")
		http.HandleFunc("/",GenericHandle)
		err:=http.ListenAndServe("localhost:8080",nil)
		if err!=nil{
			log.Fatalln("start http server fail,reason:",err)
		}
	}()
}

// api-gateway统一入口
func GenericHandle(w http.ResponseWriter,r *http.Request){
	// 访问次数统计
	conf.ACMute.Lock()
	conf.AccessCount +=1
	conf.ACMute.Unlock()

	if conf.AccessCount == 9 {
		conf.StopChan <- "访问次数10，不能再提供服务！"
	}


	now:=time.Now()

	// request params
	method :=r.Method
	//scheme :=r.URL.Scheme
	url :=r.URL.Path
	ip:=util.GetHttpAccessIP(r)
	cookies :=r.Cookies()

	if conf.AccessCount ==1 {
		w.Header().Add("Set-Cookie",fmt.Sprintf("huanggh.site-auth=%s;path=/;expires=%s",strconv.Itoa(now.Nanosecond()),now.Add(20*time.Second)))
	} else 	if len(cookies) == 0 {
		/*
		c := &http.Cookie{Name:"huanggh.site-auth",Value:strconv.Itoa(now.Nanosecond()),MaxAge:30000}
		r.AddCookie(c)
		*/
		w.Write([]byte("没有登录态!"))
		return
	} else {

		if conf.AccessCount == 5{
			expires:=time.Date(1970,1,1,1,1,1,1,time.Local).Format(time.RFC1123)
			w.Header().Add("Set-Cookie",fmt.Sprintf("huanggh.site-auth=%s;path=/;expires=%s;MaxAge=-1","DELETED",expires))
		}
	}

	fmt.Printf("http request,ip:%s,method:%s,path:%s\n",ip,method,url)
	fmt.Println(r)


	w.Write([]byte(now.Format("2006-01-02 15:04:05.999999")))

	//
	// 认证
	// 鉴权
	//
	// 限流、黑白名单、灰度发布
	// 路由、负载均衡、重试
	// 调用链追踪


	// content-type:json/file/[]byte/




}
