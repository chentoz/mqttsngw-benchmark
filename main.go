package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/GaryBoone/GoStats/stats"
	common "github.com/chentoz/mqttsngw-benchmark/common"
	mqttsn "github.com/chentoz/mqttsngw-benchmark/mqttsn"
)

func main() {

	var (
		broker   = flag.String("broker", "www.spidersens.cn:31337", "MQTT broker endpoint as scheme://host:port")
		topic    = flag.String("topic", "BL", "MQTT topic for outgoing messages")
		username = flag.String("username", "", "MQTT username (empty if auth disabled)")
		password = flag.String("password", "", "MQTT password (empty if auth disabled)")
		qos      = flag.Int("qos", 3, "QoS for published messages")
		size     = flag.Int("size", 100, "Size of the messages payload (bytes)")
		count    = flag.Int("count", 100, "Number of messages to send per client")
		clients  = flag.Int("clients", 10, "Number of clients to start")
		format   = flag.String("format", "text", "Output format: text|json")
		quiet    = flag.Bool("quiet", false, "Suppress logs while running")
	)

	flag.Parse()
	if *clients < 1 {
		log.Fatalf("Invalid arguments: number of clients should be > 1, given: %v", *clients)
	}

	if *count < 1 {
		log.Fatalf("Invalid arguments: messages count should be > 1, given: %v", *count)
	}

	resCh := make(chan *common.RunResults)
	start := time.Now()
	for i := 0; i < *clients; i++ {
		if !*quiet {
			log.Println("Starting client ", i)
		}
		//c := &Client{
		c := &mqttsn.UDPClient{
			ID:         i,
			BrokerURL:  *broker,
			BrokerUser: *username,
			BrokerPass: *password,
			MsgTopic:   *topic,
			MsgSize:    *size,
			MsgCount:   *count,
			MsgQoS:     byte(*qos),
			Quiet:      *quiet,
		}
		go c.Run(resCh)
	}

	// collect the results
	results := make([]*common.RunResults, *clients)
	for i := 0; i < *clients; i++ {
		results[i] = <-resCh
	}
	totalTime := time.Now().Sub(start)
	totals := calculateTotalResults(results, totalTime, *clients)

	// print stats
	printResults(results, totals, *format)
}

func calculateTotalResults(results []*common.RunResults, totalTime time.Duration, sampleSize int) *common.TotalResults {
	totals := new(common.TotalResults)
	totals.TotalRunTime = totalTime.Seconds()

	msgTimeMeans := make([]float64, len(results))
	msgsPerSecs := make([]float64, len(results))
	runTimes := make([]float64, len(results))
	bws := make([]float64, len(results))

	totals.MsgTimeMin = results[0].MsgTimeMin
	for i, res := range results {
		totals.Successes += res.Successes
		totals.Failures += res.Failures
		totals.TotalMsgsPerSec += res.MsgsPerSec

		if res.MsgTimeMin < totals.MsgTimeMin {
			totals.MsgTimeMin = res.MsgTimeMin
		}

		if res.MsgTimeMax > totals.MsgTimeMax {
			totals.MsgTimeMax = res.MsgTimeMax
		}

		msgTimeMeans[i] = res.MsgTimeMean
		msgsPerSecs[i] = res.MsgsPerSec
		runTimes[i] = res.RunTime
		bws[i] = res.MsgsPerSec
	}
	totals.Ratio = float64(totals.Successes) / float64(totals.Successes+totals.Failures)
	totals.AvgMsgsPerSec = stats.StatsMean(msgsPerSecs)
	totals.AvgRunTime = stats.StatsMean(runTimes)
	totals.MsgTimeMeanAvg = stats.StatsMean(msgTimeMeans)
	// calculate std if sample is > 1, otherwise leave as 0 (convention)
	if sampleSize > 1 {
		totals.MsgTimeMeanStd = stats.StatsSampleStandardDeviation(msgTimeMeans)
	}

	return totals
}

func printResults(results []*common.RunResults, totals *common.TotalResults, format string) {
	switch format {
	case "json":
		jr := common.JSONResults{
			Runs:   results,
			Totals: totals,
		}
		data, err := json.Marshal(jr)
		if err != nil {
			log.Fatalf("Error marshalling results: %v", err)
		}
		var out bytes.Buffer
		json.Indent(&out, data, "", "\t")

		fmt.Println(string(out.Bytes()))
	default:
		for _, res := range results {
			fmt.Printf("======= CLIENT %d =======\n", res.ID)
			fmt.Printf("Ratio:               %.3f (%d/%d)\n", float64(res.Successes)/float64(res.Successes+res.Failures), res.Successes, res.Successes+res.Failures)
			fmt.Printf("Runtime (s):         %.3f\n", res.RunTime)
			fmt.Printf("Msg time min (ms):   %.3f\n", res.MsgTimeMin)
			fmt.Printf("Msg time max (ms):   %.3f\n", res.MsgTimeMax)
			fmt.Printf("Msg time mean (ms):  %.3f\n", res.MsgTimeMean)
			fmt.Printf("Msg time std (ms):   %.3f\n", res.MsgTimeStd)
			fmt.Printf("Bandwidth (msg/sec): %.3f\n\n", res.MsgsPerSec)
		}
		fmt.Printf("========= TOTAL (%d) =========\n", len(results))
		fmt.Printf("Total Ratio:                 %.3f (%d/%d)\n", totals.Ratio, totals.Successes, totals.Successes+totals.Failures)
		fmt.Printf("Total Runtime (sec):         %.3f\n", totals.TotalRunTime)
		fmt.Printf("Average Runtime (sec):       %.3f\n", totals.AvgRunTime)
		fmt.Printf("Msg time min (ms):           %.3f\n", totals.MsgTimeMin)
		fmt.Printf("Msg time max (ms):           %.3f\n", totals.MsgTimeMax)
		fmt.Printf("Msg time mean mean (ms):     %.3f\n", totals.MsgTimeMeanAvg)
		fmt.Printf("Msg time mean std (ms):      %.3f\n", totals.MsgTimeMeanStd)
		fmt.Printf("Average Bandwidth (msg/sec): %.3f\n", totals.AvgMsgsPerSec)
		fmt.Printf("Total Bandwidth (msg/sec):   %.3f\n", totals.TotalMsgsPerSec)
	}
	return
}
