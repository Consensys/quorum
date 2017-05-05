package influxdb

import (
	"fmt"
	"log"
	uurl "net/url"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/rcrowley/go-metrics"
)

type reporter struct {
	reg      metrics.Registry
	interval time.Duration

	url      uurl.URL
	database string
	username string
	password string
	tags     map[string]string

	client client.Client
}

// InfluxDB starts a InfluxDB reporter which will post the metrics from the given registry at each d interval.
func InfluxDB(r metrics.Registry, d time.Duration, url, database, username, password string) {
	InfluxDBWithTags(r, d, url, database, username, password, nil)
}

// InfluxDBWithTags starts a InfluxDB reporter which will post the metrics from the given registry at each d interval with the specified tags
func InfluxDBWithTags(r metrics.Registry, d time.Duration, url, database, username, password string, tags map[string]string) {
	u, err := uurl.Parse(url)
	if err != nil {
		log.Printf("unable to parse InfluxDB url %s. err=%v", url, err)
		return
	}

	rep := &reporter{
		reg:      r,
		interval: d,
		url:      *u,
		database: database,
		username: username,
		password: password,
		tags:     tags,
	}
	if err := rep.makeClient(); err != nil {
		log.Printf("unable to make InfluxDB client. err=%v", err)
		return
	}

	rep.run()
}

func (r *reporter) makeClient() (err error) {
	r.client, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     r.url.String(),
		Username: r.username,
		Password: r.password,
	})

	return
}

func (r *reporter) run() {
	intervalTicker := time.Tick(r.interval)
	pingTicker := time.Tick(time.Second * 5)

	for {
		select {
		case <-intervalTicker:
			if err := r.send(); err != nil {
				log.Printf("unable to send metrics to InfluxDB. err=%v", err)
			}
		case <-pingTicker:
			_, _, err := r.client.Ping(0)
			if err != nil {
				log.Printf("got error while sending a ping to InfluxDB, trying to recreate client. err=%v", err)

				if err = r.makeClient(); err != nil {
					log.Printf("unable to make InfluxDB client. err=%v", err)
				}
			}
		}
	}
}

func (r *reporter) send() error {
	bps, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  r.database,
		Precision: "s",
	})
	log.Println("new batch points", bps)

	r.reg.Each(func(name string, i interface{}) {
		now := time.Now()
		log.Println("a point:", i)

		switch metric := i.(type) {
		case metrics.Counter:
			ms := metric.Snapshot()
			pt, _ := client.NewPoint(
				fmt.Sprintf("%s.count", name),
				r.tags,
				map[string]interface{}{
					"value": ms.Count(),
				},
				now,
			)
			bps.AddPoint(pt)
		case metrics.Gauge:
			ms := metric.Snapshot()
			pt, _ := client.NewPoint(
				fmt.Sprintf("%s.gauge", name),
				r.tags,
				map[string]interface{}{
					"value": ms.Value(),
				},
				now,
			)
			bps.AddPoint(pt)
		case metrics.GaugeFloat64:
			ms := metric.Snapshot()
			pt, _ := client.NewPoint(
				fmt.Sprintf("%s.gauge", name),
				r.tags,
				map[string]interface{}{
					"value": ms.Value(),
				},
				now,
			)
			bps.AddPoint(pt)
		case metrics.Histogram:
			ms := metric.Snapshot()
			ps := ms.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999, 0.9999})
			pt, _ := client.NewPoint(
				fmt.Sprintf("%s.histogram", name),
				r.tags,
				map[string]interface{}{
					"count":    ms.Count(),
					"max":      ms.Max(),
					"mean":     ms.Mean(),
					"min":      ms.Min(),
					"stddev":   ms.StdDev(),
					"variance": ms.Variance(),
					"p50":      ps[0],
					"p75":      ps[1],
					"p95":      ps[2],
					"p99":      ps[3],
					"p999":     ps[4],
					"p9999":    ps[5],
				},
				now,
			)
			bps.AddPoint(pt)
		case metrics.Meter:
			ms := metric.Snapshot()
			pt, _ := client.NewPoint(
				fmt.Sprintf("%s.meter", name),
				r.tags,
				map[string]interface{}{
					"count": ms.Count(),
					"m1":    ms.Rate1(),
					"m5":    ms.Rate5(),
					"m15":   ms.Rate15(),
					"mean":  ms.RateMean(),
				},
				now,
			)
			bps.AddPoint(pt)
		case metrics.Timer:
			ms := metric.Snapshot()
			ps := ms.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999, 0.9999})
			pt, _ := client.NewPoint(
				fmt.Sprintf("%s.timer", name),
				r.tags,
				map[string]interface{}{
					"count":    ms.Count(),
					"max":      ms.Max(),
					"mean":     ms.Mean(),
					"min":      ms.Min(),
					"stddev":   ms.StdDev(),
					"variance": ms.Variance(),
					"p50":      ps[0],
					"p75":      ps[1],
					"p95":      ps[2],
					"p99":      ps[3],
					"p999":     ps[4],
					"p9999":    ps[5],
					"m1":       ms.Rate1(),
					"m5":       ms.Rate5(),
					"m15":      ms.Rate15(),
					"meanrate": ms.RateMean(),
				},
				now,
			)
			bps.AddPoint(pt)
		}

	})

	err := r.client.Write(bps)
	log.Println("bps", bps)
	log.Println("writing err:", err)
	return err
}
