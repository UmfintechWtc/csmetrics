# collect-metrics

## Description
**1. 统计各用户Session会话数**  
**2. 统计各用户启动的进程数**  
**3. 统计各TCP连接状态**
**4. 统计URL访问次数**

## Prometheus Metrics
### Gauge Metrics Name
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
### Counter Metrics Name 
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
### Histogram Metrics Name 
**requests_delay_with_histogram_bucket [Just Demo]**  
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
### Summary Metrics Name
**requests_delay_with_summary [Just Demo]**  
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

## Configuration
### Path
**deploy/xconf/collect-metrics.yml**
### log config
```yaml
server:
  logrus_config:
    log_level: debug    # 日志级别，支持 debug | info | warning | error | fatal | panic , 默认: debug
    report_caller: true # 回调记录，支持 true | false , 默认: true
    stdout: true        # 日志控制台标准输出，支持 true | false , 默认: true
    lumberjack_logger:  # 当 stdout 为 false 时配置生效。日志落地文件
      filename: xxx.log # 日志路径，default: ./logs/collect-metrics.log
      max_size: 1       # 日志大小，default: 100MB
      max_backups: 1    # 日志备份个数，default: 5
      max_age: 1        # 历史日志保留时长(days)，default: 7
      compress: true    # 是否压缩，支持 true | false , default: true
```
## Development building
### OS Platform
```shell
NAME="CentOS Linux"
VERSION="7 (Core)"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="7"
PRETTY_NAME="CentOS Linux 7 (Core)"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:7"
HOME_URL="https://www.centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"

CENTOS_MANTISBT_PROJECT="CentOS-7"
CENTOS_MANTISBT_PROJECT_VERSION="7"
REDHAT_SUPPORT_PRODUCT="centos"
REDHAT_SUPPORT_PRODUCT_VERSION="7"
```

### Kernel 
```shell
Linux k8s-master-1 3.10.0-1160.el7.x86_64 #1 SMP Mon Oct 19 16:18:59 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
```

### [Go Version](https://go.dev/dl/)
```shell
go version go1.20 linux/amd64
```
## Installation and Usage

### Build
```shell
go build
```

### Help
```shell
[root@k8s-master-1:collect-metrics-with-go_130]# ./collect-metrics --help
Usage of ./collect-metrics:
  -config string
    	配置文件 (default "./deploy/xconf/collect-metrics.yml")
```

### Running
```shell
./collect-metrics -config /path/to/config.yml
```