package sample001

import "fmt"

type MetricsController struct {

}

func (m *MetricsController) recordRequest (requestInfo string) {
	fmt.Printf("Metrics is calling: [%v]", requestInfo)
}
