package main

import (
	"collect-metrics/collector"
	"collect-metrics/common"
	"collect-metrics/config"
	cmlog "collect-metrics/log"
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func dorisHandler(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := collector.Doris(config)
		if err != nil {
			cmlog.Error(err.Error())
			return
		}
		// get restart & tablet metrics
		promhttp.Handler().ServeHTTP(w, r)
	}
}

var (
	configPath = flag.String("config", common.COLLECT_METRICS_CONFIG_PATH, "use configuration")
)

func main() {
	flag.Parse()
	globalConfig, err := config.LoadInternalConfig(*configPath)
	if err != nil {
		cmlog.Error(err.Error())
	}

	// 启动http服务
	mux := http.NewServeMux()
	mux.HandleFunc("/doris", dorisHandler(globalConfig))
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", globalConfig.Internal.Server.Listen, globalConfig.Internal.Server.Port),
		Handler: mux,
	}
	fmt.Println(
		fmt.Sprintf(
			"Server is running on http://%s:%d",
			globalConfig.Internal.Server.Listen,
			globalConfig.Internal.Server.Port,
		),
	)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
