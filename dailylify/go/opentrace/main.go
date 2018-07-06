package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"math/rand"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"sourcegraph.com/sourcegraph/appdash"
	"net"
	"net/url"
	"sourcegraph.com/sourcegraph/appdash/traceapp"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href="/home">Click here to start a request</a>`))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Request started"))

	span := opentracing.StartSpan("/home")
	defer span.Finish()

	asyncReq,_:=http.NewRequest("GET","http://localhost:8080/async",nil)
	err := span.Tracer().Inject(span.Context(),
		opentracing.TextMap,
		opentracing.HTTPHeadersCarrier(asyncReq.Header))
	if err != nil{
		log.Fatal("could not injext span context into header:%v",err)
	}
	_,err = http.Get("http://localhost:8080/service")
	if err != nil{
		ext.Error.Set(span,true)
		span.LogEventWithPayload("get service error:%v",err)
	}
	time.Sleep(time.Duration(rand.Intn(200))*time.Millisecond)
	w.Write([]byte("request done!"))
}
func serviceHandler(w http.ResponseWriter, r *http.Request) {
	var sp opentracing.Span
	opName := r.URL.Path

	wireContext,err:=opentracing.GlobalTracer().Extract(
		opentracing.TextMap,
		opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil{
		sp = opentracing.StartSpan(opName)
	}else {
		sp = opentracing.StartSpan(opName,opentracing.ChildOf(wireContext))
	}
	defer sp.Finish()

	http.Get("http://localhost:8080/db")
	time.Sleep(time.Duration(rand.Intn(200))*time.Millisecond)
}
func dbHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(200))*time.Millisecond)
}

func main() {
	store := appdash.NewMemoryStore()
	l,err:= net.ListenTCP("tcp",&net.TCPAddr{
		IP:net.IPv4(127,0,0,1),Port:0})
	if err != nil{
		log.Fatal(err)
	}
	collectorPort := l.Addr().(*net.TCPAddr).Port
	collectorAdd := fmt.Sprintf(":%d",collectorPort)

	cs := appdash.NewServer(l,appdash.NewLocalCollector(store))
	go cs.Start()

	appdashPort :=8700
	appdashURLStr := fmt.Sprintf("http://localhost:%d",appdashPort)
	appdashURL,err := url.Parse(appdashURLStr)
	if err !=nil{
		log.Fatal("Error parsing %s:%s",appdashURLStr,err)
	}
	fmt.Printf("too see your traces go to %s/traces\n",appdashURL)

	tapp,err := traceapp.New(nil,appdashURL)
	if err !=nil{
		log.Fatal(err)
	}
	tapp.Store = store
	tapp.Queryer = store

	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",appdashPort),tapp))
	}()
	tracer := appdashot.NewTracer(appdash.NewRemoteCollector(collectorAdd))
	opentracing.InitGlobalTracer(tracer)

	port := 8080
	addr := fmt.Sprintf(":%d",port)
	mux := http.NewServeMux()

	mux.HandleFunc("/",indexHandler)
	mux.HandleFunc("/home",homeHandler)
	mux.HandleFunc("/async",serviceHandler)
	mux.HandleFunc("/service",serviceHandler)
	mux.HandleFunc("/db",dbHandler)
	fmt.Printf("go to http://localhost:%d/home to start a request!\n",port)
	log.Fatal(http.ListenAndServe(addr,mux))
}
