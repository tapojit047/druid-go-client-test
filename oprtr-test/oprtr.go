package util

import (
	"encoding/json"
	"fmt"
	"github.com/grafadruid/go-druid"
	"github.com/hashicorp/go-retryablehttp"
)

func SubmitIngestionTask(d *druid.Client, oldData string) (string, error) {
	task := GetJsonTask(oldData)
	rawMessage := json.RawMessage(task)
	fmt.Println("------------------Rawmessage: ", rawMessage)
	method := "POST"
	path := "druid/indexer/v1/task"

	response, err := SubmitRequest(d, method, path, rawMessage)
	if err != nil {
		return "", err
	}

	taskID, err := GetValueFromClusterResponse(response, "task")
	if err != nil {
		return "", fmt.Errorf("Failed to parse respons of task ingestion request: %v", err)
	}
	err = response.Body.Close()
	if err != nil {
		return "", fmt.Errorf("Failed to close the response body: %v\n", err)
	}
	return fmt.Sprintf("%v", taskID), nil
	return "", nil
}

func CheckDataSourceExistence(d *druid.Client) (bool, error) {
	method := "POST"
	path := "druid/v2/sql"

	data := map[string]interface{}{
		"query": "SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'druid' AND TABLE_NAME = 'kubedb-datasource'",
	}
	//fmt.Println(data)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return false, fmt.Errorf("failed to marshal json response: %v\n", err)
	}
	rawMessage := json.RawMessage(jsonData)
	response, err := SubmitRequest(d, method, path, rawMessage)
	if err != nil {
		return false, err
	}

	exists, err := parseDatasourceExistenceQueryResponse(response)
	if err != nil {
		return false, fmt.Errorf("Failed to parse response of datasource existence request: %v\n", err)
	}

	err = response.Body.Close()
	if err != nil {
		return exists, fmt.Errorf("Failed to close the response body: %v\n", err)
	}
	return exists, nil
}

func RunSelectQuery(d *druid.Client) (string, error) {
	method := "POST"
	path := "druid/v2/sql"

	data := map[string]interface{}{
		"query": "SELECT * FROM \"kubedb-datasource\"",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal query json data: %v", err)
	}
	rawMessage := json.RawMessage(jsonData)
	response, err := SubmitRequest(d, method, path, rawMessage)
	if err != nil {
		return "", err
	}

	id, err := ParseSelectQueryResponse(response, "id")
	if err != nil {
		return "", fmt.Errorf("failed to parse the response body: %v\n", err)
	}
	//fmt.Println(name)
	err = response.Body.Close()
	if err != nil {
		return "", fmt.Errorf("Failed to close the response body: %v\n", err)
	}
	return id.(string), nil
}

func SubmitRequest(d *druid.Client, method, path string, opts interface{}) (*druid.Response, error) {
	res, err := d.NewRequest(method, path, opts)
	if err != nil {
		fmt.Printf("failed to submit API request : %v\n", err)
		return nil, err
	}
	http := retryablehttp.NewClient()
	resp, err := http.Do(res)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	// defer resp.Body.Close()
	response := &druid.Response{resp}
	return response, nil
}

func GetValueFromClusterResponse(res *druid.Response, key string) (interface{}, error) {
	responseBody := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to deserialize the response: %v", err)
	}
	//fmt.Println(responseBody)
	value := responseBody[key]
	return value, nil
}

func ParseSelectQueryResponse(res *druid.Response, key string) (interface{}, error) {
	var responseBody []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to deserialize the response: %v\n", err)
	}
	value := responseBody[0][key]
	return value, nil
}

func parseDatasourceExistenceQueryResponse(res *druid.Response) (bool, error) {
	//responseBody := make(map[string]interface{})
	var responseBody []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return false, fmt.Errorf("failed to deserialize the response: %v", err)
	}
	//fmt.Println(responseBody)
	return len(responseBody) != 0, nil
}

func GetJsonTask(oldData string) string {
	task := `{
		"type": "index_parallel",
	  "spec": {
	    "ioConfig": {
	      "type": "index_parallel",
	      "inputSource": {
	        "type": "inline",
	        "data": "{\"id\": %s, \"name\": \"druid\", \"time\": \"2023-12-20T00:46:58.771Z\"}"
	      },
	      "inputFormat": {
	        "type": "json"
	      }
	    },
	    "tuningConfig": {
	      "type": "index_parallel",
	      "partitionsSpec": {
	        "type": "dynamic"
	      }
	    },
	    "dataSchema": {
	      "dataSource": "dummy-data",
	      "timestampSpec": {
	        "column": "time",
	        "format": "iso"
	      },
	      "dimensionsSpec": {
	        "dimensions": ["id", "name", "time"]
	      },
	      "granularitySpec": {
	        "queryGranularity": "none",
	        "rollup": false,
	        "segmentGranularity": "day"
	      }
	    }
	  }
	}`
	if oldData == "0" {
		task = fmt.Sprintf(task, "1")
	} else {
		task = fmt.Sprintf(task, "0")
	}
	fmt.Println("----------------TASK: ", task)
	return task
}
