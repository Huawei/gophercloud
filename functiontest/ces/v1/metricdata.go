package main

import (
	"encoding/json"
	"fmt"

	"github.com/Huawei/gophercloud"
	"github.com/Huawei/gophercloud/functiontest/common"
	"github.com/Huawei/gophercloud/openstack"
	"github.com/Huawei/gophercloud/openstack/ces/v1/metricdata"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	//provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewCESV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get ces client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestBatchQueryMetricDatas(sc)
	//TestGetMetricData(sc)
	//TestGetEventData(sc)
	//TestAddMetricData(sc)
	fmt.Println("main end...")
}

func TestBatchQueryMetricDatas(sc *gophercloud.ServiceClient) {
	opts := metricdata.BatchQueryOpts{
		Metrics: []metricdata.Metric{
			{
				Namespace: "SYS.VPC",
				Dimensions: []map[string]string{
					{
						"name":  "bandwidth_id",
						"value": "ea31a911-dad7-4218-9036-77a7c3a16a45",
					},
				},
				MetricName: "downstream_bandwidth",
			},
		},
		From:   1540526925098,
		To:     1540537725098,
		Period: "1",
		Filter: "average",
	}
	metricdatas, err := metricdata.BatchQuery(sc, opts).ExtractMetricDatas()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test batch query metric data success！")
	p, _ := json.MarshalIndent(metricdatas, "", " ")
	fmt.Println(string(p))
}

/*
func TestGetMetricData(sc *gophercloud.ServiceClient) {
	opts := metricdata.GetOpts{
		Namespace: "SYS.ECS",
		MetricName: "cpu_util",
		From:   1548041969418,
		To:     1548052769418,
		Period: 3600,
		Filter: "average",
		Dim0:"instance_id,070c1ed3-176a-446e-8eff-b116b529b4b7",
	}
	metricdatas, err := metricdata.Get(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Get metric data success！")
	p, _ := json.MarshalIndent(metricdatas, "", " ")
	fmt.Println(string(p))
}

func TestAddMetricData(sc *gophercloud.ServiceClient) {
	opts := metricdata.AddMetricDataOpts{
		{
			Metric:metricdata.MetricInfo{
				Namespace:"MINE.APP",
				Dimensions:[]metricdata.MetricsDimension{
					{
						Name:"instance_id",
						Value:"33328f02-3814-422e-b688-bfdba93d4050",
					},
				},
				MetricName:"cpu_util",
			},
			Ttl:172800,
			CollectTime:1463598260000,
			Value:60,
			Unit:"%",
			Type:"int",
		},
	}
	err := metricdata.AddMetricData(sc, opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Add metric data success！")
}

func TestGetEventData(sc *gophercloud.ServiceClient) {
	opts := metricdata.GetEventDataOpts{
		Namespace: "SYS.ECS",
		From:   1548041969418,
		To:     1548052769418,
		Dim0:"instance_id,070c1ed3-176a-446e-8eff-b116b529b4b7",
		Type:"instance_host_info",
	}
	eventdata, err := metricdata.GetEventData(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Get event data success！")
	p, _ := json.MarshalIndent(eventdata, "", " ")
	fmt.Println(string(p))
}
*/