package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	steam "github.com/galexrt/go-steam"
	"github.com/galexrt/srcds_exporter/models"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	debug          bool
	connectTimeout string
)

var (
	log          = logrus.New()
	metricUpdate = make(chan models.Status)
)

func init() {
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.StringVar(&connectTimeout, "timeout", "15s", "Connection timeout")
}

func main() {
	log.Out = os.Stdout
	flag.Parse()
	if debug {
		log.Level = logrus.DebugLevel
	}
	steam.SetLog(log)
	addr := os.Getenv("ADDR")
	serverIdentification = addr
	pass := os.Getenv("RCON_PASSWORD")
	if addr == "" || pass == "" {
		fmt.Println("Please set ADDR & RCON_PASSWORD.")
		return
	}
	go func() {
		//manageMetrics()
	}()
	go func() {
		for {
			con, err := steam.Connect(addr, &steam.ConnectOptions{
				RCONPassword: pass,
				Timeout:      connectTimeout,
			})
			if err != nil {
				fmt.Println(err)
				time.Sleep(1 * time.Second)
				continue
			}
			defer con.Close()
			for {
				resp, err := con.Send("status")
				if err != nil {
					log.Error(err)
					break
				}
				log.Debug("Read status command output")
				//metricUpdate <- *parser.Parse(resp)
				fmt.Printf("RESP: %v\n", resp)

				time.Sleep(4 * time.Second)
			}
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*metricsAddr, nil))
}

func manageMetrics() {
	first := true
	for {
		status := <-metricUpdate
		if first {
			initMetrics(status)
			first = false
		}
		log.Debugln("manageMetrics: Received metrics update")
		updateMetrics(status)
	}
}
