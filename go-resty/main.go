package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	GetData()
	PostData()
	PostStructData()
	RetryData()
	DumpData()
	TraceData()
}

func GetData() {
	resp, err := resty.New().R().
		SetHeader("Authorization", "Bearer 1234567890").
		Get("https://httpbin.org/get")
	if err != nil {
		slog.Error("get data", "error", err)
	}
	fmt.Println(resp.String())
}

func PostData() {
	resp, err := resty.New().R().
		SetDebug(true).
		SetHeader("Authorization", "Bearer 1234567890").
		SetBody(map[string]interface{}{
			"name": "John",
			"age":  30,
		}).
		Post("https://httpbin.org/post")
	if err != nil {
		slog.Error("get data", "error", err)
	}
	fmt.Println(resp.String())
}

type RequestStruct struct {
	Types   string `json:"types"`
	Content string `json:"content"`
	Count   int    `json:"count"`
}

type ResponseStruct struct {
	Data    string `json:"data"`
	Headers struct {
		AcceptEncoding string `json:"Accept-Encoding"`
		Authorization  string `json:"Authorization"`
		ContentLength  string `json:"Content-Length"`
		ContentType    string `json:"Content-Type"`
		Host           string `json:"Host"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceId   string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	Json struct {
		Content string `json:"content"`
		Count   int    `json:"count"`
		Types   string `json:"types"`
	} `json:"json"`
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

func PostStructData() {
	var response ResponseStruct
	resp, err := resty.New().R().
		SetDebug(true).
		SetHeader("Authorization", "Bearer 1234567890").
		SetBody(RequestStruct{
			Types:   "alert",
			Content: "This is a test alert",
			Count:   30,
		}).
		SetResult(&response).
		Post("https://httpbin.org/post")
	if err != nil {
		slog.Error("get data", "error", err)
	}
	if resp.IsSuccess() {
		fmt.Println(response)
	}
}

func RetryData() {
	resp, err := resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(100*time.Millisecond).
		R().
		SetHeader("Authorization", "Bearer 1234567890").
		Get("https://httpbin.org/get")
	if err != nil {
		slog.Error("get data", "error", err)
	}
	fmt.Println(resp.String())
}

func DumpData() {
	resp, err := resty.New().R().
		SetDebug(true).
		SetHeader("Authorization", "Bearer 1234567890").
		SetBody(map[string]interface{}{
			"name": "John",
			"age":  30,
		}).
		Post("https://httpbin.org/post")
	if err != nil {
		slog.Error("get data", "error", err)
	}

	fmt.Println("statusCode", resp.StatusCode())
	fmt.Println("status", resp.Status())
	fmt.Println("proto", resp.Proto())
	fmt.Println("header", resp.Header())
	fmt.Println("isError", resp.IsError())
	fmt.Println("isSuccess", resp.IsSuccess())
}

func TraceData() {
	resp, err := resty.New().R().
		SetDebug(true).
		EnableTrace().
		SetHeader("Authorization", "Bearer 1234567890").
		SetBody(map[string]interface{}{
			"name": "John",
			"age":  30,
		}).
		Post("https://httpbin.org/post")
	if err != nil {
		slog.Error("get data", "error", err)
	}

	traceInfo := resp.Request.TraceInfo()
	fmt.Println("traceInfo", traceInfo)
	fmt.Println("DNSLookup:", traceInfo.DNSLookup)
	fmt.Println("ConnTime:", traceInfo.ConnTime)
	fmt.Println("TCPConnTime:", traceInfo.TCPConnTime)
	fmt.Println("TLSHandshake:", traceInfo.TLSHandshake)
	fmt.Println("ServerTime:", traceInfo.ServerTime)
	fmt.Println("ResponseTime:", traceInfo.ResponseTime)
	fmt.Println("TotalTime:", traceInfo.TotalTime)
	fmt.Println("IsConnReused:", traceInfo.IsConnReused)
	fmt.Println("IsConnWasIdle:", traceInfo.IsConnWasIdle)
	fmt.Println("ConnIdleTime:", traceInfo.ConnIdleTime)
	fmt.Println("RequestAttempt:", traceInfo.RequestAttempt)
	fmt.Println("RemoteAddr:", traceInfo.RemoteAddr.String())
}
