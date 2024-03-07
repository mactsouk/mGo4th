package main

import (
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PORT is the TCP port number the server will listen to
var PORT = ":1234"

var nGoroutines = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "packt",
		Name:      "n_goroutines",
		Help:      "Number of goroutines"})

var nMemory = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "packt",
		Name:      "n_memory",
		Help:      "Memory usage"})

func main() {
	prometheus.MustRegister(nGoroutines)
	prometheus.MustRegister(nMemory)

	const nGo = "/sched/goroutines:goroutines"
	const nMem = "/memory/classes/heap/free:bytes"
	getMetric := make([]metrics.Sample, 2)
	getMetric[0].Name = nGo
	getMetric[1].Name = nMem

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		for {
			for i := 1; i < 4; i++ {
				go func() {
					_ = make([]int, 1000000)
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				}()
			}

			runtime.GC()
			metrics.Read(getMetric)
			goVal := getMetric[0].Value.Uint64()
			memVal := getMetric[1].Value.Uint64()
			time.Sleep(time.Duration(rand.Intn(15)) * time.Second)

			nGoroutines.Set(float64(goVal))
			nMemory.Set(float64(memVal))
		}
	}()

	log.Println("Listening to port", PORT)
	log.Println(http.ListenAndServe(PORT, nil))
}
