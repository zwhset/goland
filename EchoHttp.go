package main

/*
	单个go返回一个http json的时间为python gunicon的6倍左右。
	ab -n 10000 -c 100 $url

	# aliyun
	Concurrency Level:      100
	Time taken for tests:   4.051 seconds
	Complete requests:      1000
	Failed requests:        0
	Total transferred:      149000 bytes
	HTML transferred:       32000 bytes
	Requests per second:    246.87 [#/sec] (mean)
	Time per request:       405.079 [ms] (mean)
	Time per request:       4.051 [ms] (mean, across all concurrent requests)
	Transfer rate:          35.92 [Kbytes/sec] received

	Connection Times (ms)
				  min  mean[+/-sd] median   max
	Connect:        5   73 212.4     21    1895
	Processing:     5  189 501.5     35    3751
	Waiting:        5  184 501.6     34    3751
	Total:         10  262 553.0     78    3772

	# local
	Concurrency Level:      100
	Time taken for tests:   0.148 seconds
	Complete requests:      1000
	Failed requests:        0
	Total transferred:      149000 bytes
	HTML transferred:       32000 bytes
	Requests per second:    6747.96 [#/sec] (mean)
	Time per request:       14.819 [ms] (mean)
	Time per request:       0.148 [ms] (mean, across all concurrent requests)
	Transfer rate:          981.88 [Kbytes/sec] received

	Connection Times (ms)
				  min  mean[+/-sd] median   max
	Connect:        2    7   2.0      7      11
	Processing:     1    7   2.2      7      15
	Waiting:        1    5   2.0      5      12
	Total:          8   14   2.7     15      20

	# python 没开协程以及只用了粗版运行
	# py2
	Concurrency Level:      100
	Time taken for tests:   1.379 seconds
	Complete requests:      1000
	Failed requests:        0
	Total transferred:      190000 bytes
	HTML transferred:       43000 bytes
	Requests per second:    725.21 [#/sec] (mean)
	Time per request:       137.891 [ms] (mean)
	Time per request:       1.379 [ms] (mean, across all concurrent requests)
	Transfer rate:          134.56 [Kbytes/sec] received

	Connection Times (ms)
				  min  mean[+/-sd] median   max
	Connect:        0    0   1.0      0       5
	Processing:     2  130  24.6    137     151
	Waiting:        2  130  24.6    136     150
	Total:          7  131  23.7    137     151

	# py3
	Concurrency Level:      100
	Time taken for tests:   1.376 seconds
	Complete requests:      1000
	Failed requests:        0
	Total transferred:      189000 bytes
	HTML transferred:       43000 bytes
	Requests per second:    726.85 [#/sec] (mean)
	Time per request:       137.579 [ms] (mean)
	Time per request:       1.376 [ms] (mean, across all concurrent requests)
	Transfer rate:          134.16 [Kbytes/sec] received

	Connection Times (ms)
				  min  mean[+/-sd] median   max
	Connect:        0    1   1.3      0       7
	Processing:     5  130  22.5    136     144
	Waiting:        3  130  22.5    135     144
	Total:         12  131  21.3    136     145

	# gunicorn python2
	Concurrency Level:      100
	Time taken for tests:   0.648 seconds
	Complete requests:      1000
	Failed requests:        0
	Total transferred:      195000 bytes
	HTML transferred:       43000 bytes
	Requests per second:    1543.80 [#/sec] (mean)
	Time per request:       64.775 [ms] (mean)
	Time per request:       0.648 [ms] (mean, across all concurrent requests)
	Transfer rate:          293.98 [Kbytes/sec] received

	Connection Times (ms)
				  min  mean[+/-sd] median   max
	Connect:        0    1   1.2      0       9
	Processing:     6   61  10.7     64      86
	Waiting:        6   60  10.9     63      82
	Total:         11   62   9.7     64      87

	# gunicorn -k gevent python2
	Concurrency Level:      100
	Time taken for tests:   0.712 seconds
	Complete requests:      1000
	Failed requests:        0
	Total transferred:      195000 bytes
	HTML transferred:       43000 bytes
	Requests per second:    1405.27 [#/sec] (mean)
	Time per request:       71.160 [ms] (mean)
	Time per request:       0.712 [ms] (mean, across all concurrent requests)
	Transfer rate:          267.61 [Kbytes/sec] received

	Connection Times (ms)
				  min  mean[+/-sd] median   max
	Connect:        0    1   1.7      0      10
	Processing:     2   64  41.1     56     194
	Waiting:        2   64  41.2     55     194
	Total:          2   65  41.1     56     203
*/

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func echo(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("echo error"))
		return
	}

	result := Result{0, "successed"}

	if data, err := json.Marshal(result); err == nil {
		writelen, err := w.Write(data)
		if err != nil || writelen != len(data) {
			log.Println(err, "write len:", writelen)
		}
	}
}

func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Run EchoServer 0.0.0.0:9090")
}
