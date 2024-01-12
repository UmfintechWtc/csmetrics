# collect-metrics

#### Description
**1. 统计各用户Session会话数**  
**2. 统计各用户启动的进程数**  
**3. 统计各TCP连接状态**


#### Prometheus Metrics
##### Gauge Metrics Name
**session_group**  
example: 
```shell
# HELP session_group get tty order by user
# TYPE session_group gauge
session_group{user="root"} 2
session_group{user="wtc"} 0
```  
**process_group**  
example:
```shell
# HELP process_group get process order by user
# TYPE process_group gauge
process_group{user="chrony"} 1
process_group{user="dbus"} 1
process_group{user="polkitd"} 1
process_group{user="root"} 397
process_group{user="wtc"} 0
process_group{user="zabbix"} 13
```
**netstat_group**  
example:
```shell
# HELP netstat_group get netstat order by state
# TYPE netstat_group gauge
netstat_group{state="ESTABLISHED"} 22
netstat_group{state="LISTEN"} 8
netstat_group{state="TIME_WAIT"} 1
```
##### Counter Metrics Name
**requests_url_total**
example:
```shell
# HELP requests_url_total get requests order by url
# TYPE requests_url_total counter
requests_url_total{url="/metrics"} 11
requests_url_total{url="/metrics/"} 1
requests_url_total{url="/metrics/12323"} 4
requests_url_total{url="/metrics/asdasscxcvcxv"} 1
```
##### Histogram Metrics Name
**requests_delay_with_histogram_bucket**
```shell
# HELP requests_delay_with_histogram Total number of HTTP requests delay with histogram
# TYPE requests_delay_with_histogram histogram
requests_delay_with_histogram_bucket{code="200",le="10"} 1
requests_delay_with_histogram_bucket{code="200",le="20"} 1
requests_delay_with_histogram_bucket{code="200",le="30"} 2
requests_delay_with_histogram_bucket{code="200",le="40"} 4
requests_delay_with_histogram_bucket{code="200",le="50"} 5
requests_delay_with_histogram_bucket{code="200",le="+Inf"} 11
requests_delay_with_histogram_sum{code="200"} 582
requests_delay_with_histogram_count{code="200"} 11
```
##### Summary Metrics Name
```shell
# HELP requests_delay_with_summary Total number of HTTP requests delay with summary
# TYPE requests_delay_with_summary summary
requests_delay_with_summary{code="200",quantile="0.1"} 7
requests_delay_with_summary{code="200",quantile="0.2"} 11
requests_delay_with_summary{code="200",quantile="0.3"} 13
requests_delay_with_summary{code="200",quantile="0.4"} 21
requests_delay_with_summary{code="200",quantile="0.5"} 24
requests_delay_with_summary_sum{code="200"} 470
requests_delay_with_summary_count{code="200"} 11
```
#### Configuration
```yaml
---
server:
  listen: 0.0.0.0
  port: 9099
  periodSeconds: 30s
  shutdown_timeout_ms: 1000
  # debug | release, default: debug
  mode: release
      

metrics:
  # gauage类型指标
  gauge:
    netstat:
      # 扫描频率
      periodSeconds: 15s
    process:
      periodSeconds: 15s
    session:
      periodSeconds: 15s
  histogram:
    delay:
      buckets:
        # 支持切片和线性，若都开启线性优先级高
        linear:
          enabled: True
          range:
            # 边界值: [10 20 30 40 50]
            start: 10
            width: 10
            count: 5
        slice:
          enabled: True
          range:
            # 边界值: [10 20 30 40 50]
            - 10
            - 20
            - 30
            - 40
            - 50
  summary:
    delay:
      median: 
        # 中位数: 度量指标
        10: 10  # example: 采样10%样本，请求小于10ms
        20: 20  # example: 采样20%样本，请求小于20ms
        30: 30  # example: 采样30%样本，请求小于30ms
        40: 40  # example: 采样40%样本，请求小于40ms
        50: 50  # example: 采样50%样本，请求小于50ms
```
#### Installation

1.  xxxx
2.  xxxx
3.  xxxx

#### Instructions

1.  xxxx
2.  xxxx
3.  xxxx

#### Contribution

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request


#### Gitee Feature

1.  You can use Readme\_XXX.md to support different languages, such as Readme\_en.md, Readme\_zh.md
2.  Gitee blog [blog.gitee.com](https://blog.gitee.com)
3.  Explore open source project [https://gitee.com/explore](https://gitee.com/explore)
4.  The most valuable open source project [GVP](https://gitee.com/gvp)
5.  The manual of Gitee [https://gitee.com/help](https://gitee.com/help)
6.  The most popular members  [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
