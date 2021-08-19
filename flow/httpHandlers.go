package flow

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var flows []ChangeListFlowInfo

func ListFlowsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, flows)
	return
}

func init() {
	flows = make([]ChangeListFlowInfo, 0)
	//flows = append(flows, generateRandomData(true))
	//flows = append(flows, generateRandomData(false))
	flowsJson := `[
				{"producers":
					[{"id":"12390","name":"232 Producer","status":"running"}],
				"consumers":
					[{"id":"12081","name":"232 ftest gcp","status":"waiting for producer"},
					{"id":"12887","name":"232 utest gcp","status":"waiting for producer"},
					{"id":"12847","name":"232 casam producer","status":"waiting for producer"},
					{"id":"12059","name":"232 selenium","status":"waiting for producer"}]},
				{"producers":
					[{"id":"13423","name":"Main Producer","status":"completed"}],
				"consumers":
					[{"id":"12281","name":"main ftest gcp","status":"running"},
					{"id":"12318","name":"main utest gcp","status":"running"},
					{"id":"12540","name":"main casam producer","status":"running"}]}]`

	err := json.Unmarshal([]byte(flowsJson), &flows)
	if err != nil {
		log.Fatal(err)
	}
}

func generateRandomData(isProducerRunning bool) ChangeListFlowInfo {
	producers := make([]AutobuildInfo, 0)
	producers = append(producers, getRandomProducerAutobuildInfo(isProducerRunning))
	consumers := make([]AutobuildInfo, 4)
	for i := range consumers {
		consumers[i] = getRandomConsumerAutobuildInfo(!isProducerRunning)
	}
	return ChangeListFlowInfo{
		//PrecheckinInfo: PrecheckinInfo{
		//	Stages:    []string{"validations", "precheckin run", "submission"},
		//	Current:   "",
		//	Completed: true,
		//},
		Producers: producers,
		Consumers: consumers,
	}
}

func getRandomProducerAutobuildInfo(isRunning bool) AutobuildInfo {
	var status string

	if isRunning {
		status = "running"
	} else {
		status = "completed"
	}
	return AutobuildInfo{
		ID:     "12390",
		Name:   "Main Producer",
		Status: status,
	}
}

func getRandomConsumerAutobuildInfo(isRunning bool) AutobuildInfo {
	var status string
	if isRunning {
		status = "running"
	} else {
		status = "waiting for producer"
	}
	return AutobuildInfo{
		ID:     getRandomAutobuildId(),
		Name:   "",
		Status: status,
	}
}

func getRandomAutobuildId() string {
	return strconv.Itoa(12000 + rand.Intn(1000))
}
