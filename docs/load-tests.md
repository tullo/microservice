# Load Tests

## 1st Iteration

```sh
# ./test-benchmark-service.sh (d15s t4 c400)
Running 15s test @ http://192.168.176.2:3000/twirp/stats.StatsService/Push
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    98.90ms  163.43ms   1.13s    91.93%
    Req/Sec     1.81k     1.09k    4.88k    76.83%
  108351 requests in 15.09s, 17.95MB read
  Socket errors: connect 0, read 0, write 0, timeout 8
  Non-2xx or 3xx responses: 60979
Requests/sec:   7179.14
Transfer/sec:      1.19MB
```

Out of `108351` requests we have `60979` non-2xx or 3xx responses.

APM: Exception stack trace
```
twirp error internal: dial tcp 192.168.176.5:3306: operation was canceled

twirp.go in NewServerHooks.func1 at line 14
stats.twirp.go in callError at line 1020
stats.twirp.go in writeError at line 559
stats.twirp.go in (*statsServiceServer).writeError at line 256
stats.twirp.go in (*statsServiceServer).servePushJSON at line 373
stats.twirp.go in (*statsServiceServer).servePush at line 316
stats.twirp.go in (*statsServiceServer).ServeHTTP at line 299
wrap.go in WrapWithIP.func1 at line 29
```

---

## 2nd Iteration

```sh
# ./test-benchmark-service.sh (d15s t4 c400)
Running 15s test @ http://192.168.240.3:3000/twirp/stats.StatsService/Push
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   150.09ms  190.97ms   1.94s    88.03%
    Req/Sec     0.99k   691.65     2.78k    58.19%
  59292 requests in 15.10s, 6.60MB read
  Socket errors: connect 0, read 0, write 0, timeout 1
  Non-2xx or 3xx responses: 3078
Requests/sec:   3927.27
Transfer/sec:    447.51KB
```

Out of `59292` requests we still have `3078` non-2xx or 3xx responses and stack traces like `twirp error internal: context canceled`

---

## 3rd Iteration

```sh
# ./test-benchmark-service.sh (d60s t4 c100)
Running 1m test @ http://172.22.0.2:3000/twirp/stats.StatsService/Push
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    20.09ms    8.38ms 196.03ms   90.40%
    Req/Sec     1.27k   242.35     1.79k    70.96%
  304527 requests in 1.00m, 31.66MB read
Requests/sec:   5067.03
Transfer/sec:    539.36KB
```

---

## 4rd Iteration

```sh
# ./test-benchmark-service.sh (d60s t4 c100)
Running 1m test @ http://172.26.0.2:3000/twirp/stats.StatsService/Push
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.13ms    5.47ms 108.96ms   86.70%
    Req/Sec     6.39k   594.54     9.31k    77.46%
  1527397 requests in 1.00m, 158.77MB read
Requests/sec:  25434.90
Transfer/sec:      2.64MB
```

We got more (5x) throughput but SQL traces disappeared in APM transactions view.

![stats service load test](images/StatsService-Push-Transactions.png)
