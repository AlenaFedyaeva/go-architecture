package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Params struct {
	method, uri, bodyO string
	threadsNum         int
	requestNumO, timeO int
	url *url.URL
}

func (p *Params) checkParams() error{
	//должен быть задан один из двух флагов
	if p.timeO > 0 && p.requestNumO > 0 {
		return errors.New("ERR: select one flag: time or numReq.")
	}
	if p.timeO < 0 || p.requestNumO < 0 || p.threadsNum < 0 {
		return errors.New("ERR: time, threadsNum or numReq <0.")
	}
	switch p.method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		return errors.New("ERR: wrong method.")

	}
	url, err := url.ParseRequestURI(p.uri)
	if err != nil{
		return err
	}
	p.url=url
	// if u.Scheme != "" && u.Host != ""{
	// 	return  errors.New("ERR: wrong uri")
	// }

	return nil
}

func parseParams() (*Params, error) {
	p := &Params{}
	flag.StringVar(&p.method, "m", "GET", "http method")
	flag.StringVar(&p.uri, "url", "http://localhost:8085/orders", "http - method")
	flag.StringVar(&p.bodyO, "b", "", "body")

	flag.IntVar(&p.threadsNum, "threads", 5, "Number of threads")

	flag.IntVar(&p.timeO, "time", 50, "Work time")
	flag.IntVar(&p.requestNumO, "numReq", 0, "Number of requests")

	err:=p.checkParams()
	if err != nil {
		return nil, err
	}
	return p,nil
	
}

func main() {

	p, err := parseParams()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	

	req,_:=http.NewRequest(p.method,p.uri,nil)  
	http.DefaultClient.Do(req)



	fmt.Println(p)
}
