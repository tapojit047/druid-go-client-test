package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/grafadruid/go-druid"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var druidOpts []druid.ClientOption
	druidOpts = append(druidOpts, druid.WithBasicAuth("admin", "newPassword2"))

	//clientWithTLS, err := httpClientWithTLS()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//druidOpts = append(druidOpts, druid.WithHTTPClient(clientWithTLS))
	d, err := druid.NewClient("http://localhost:8081", druidOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := d.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//fmt.Println(CheckDataSourceExist(d))
	err = UpdateDruidPassword(d, "tapojit")
	if err != nil {
		log.Println("Failed to update password", err.Error())
		return
	}
	//CheckDataSourceExist
	//taskID, err := util.SubmitIngestionTask(d, "0")
	//if err != nil {
	//	fmt.Sprintf("Failed to submit ingestion task: %v", err)
	//	return
	//}
	//fmt.Println(taskID)

	//path := "druid/coordinator/v1/datasources/dummy-data/loadstatus?forceMetadataRefresh==true"
	//response := SubmitRequest(d, "GET", path, nil)
	//
	//status, err := GetValueFromClusterResponse(response, "dummy-data")
	//fmt.Println(status)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer func() {
	//	err := response.Body.Close()
	//	if err != nil {
	//		fmt.Println("Error closing response body:", err)
	//	}
	//}()
	//DruidHealthCheck(d)

	//fmt.Println(CheckDataSourceExist(d))
	//DruidCreateDataSource(d)
}

func UpdateDruidPassword(d *druid.Client, password string) error {
	method := "POST"
	path := "druid-ext/basic-security/authentication/db/basic/users/admin/credentials"

	data := map[string]interface{}{
		"password": password,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	rawMessage := json.RawMessage(jsonData)

	_, err = d.ExecuteRequest(method, path, rawMessage, nil)
	if err != nil {
		fmt.Println("Failed to execute coordinator config update request", err.Error())
		return err
	}
	return nil
}

func CheckDataSourceExist(d *druid.Client) bool {
	method := "POST"
	path := "druid/v2/sql"

	data := map[string]interface{}{
		"query": "SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'druid' AND TABLE_NAME = 'kubedb-datasource'",
	}
	//fmt.Println(data)

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return false
	}
	rawMessage := json.RawMessage(jsonData)
	//response := SubmitRequest(d, method, path, rawMessage)
	var responseBody []map[string]interface{}
	_, err = d.ExecuteRequest(method, path, rawMessage, &responseBody)
	if err != nil {
		fmt.Printf("could not execute request: %s\n", err)
		return false
	}
	fmt.Println(responseBody)
	if len(responseBody) != 0 {
		return true
	}
	return false
	//defer func() {
	//	err := response.Body.Close()
	//	if err != nil {
	//		fmt.Println("Error closing response body:", err)
	//	}
	//}()

	//fmt.Println(response)

	//exists, err := GetValueFromDSQueryResponse(response)
	//if err != nil {
	//	fmt.Printf("Failed to fetch Name in the response, %v\n", err)
	//	return false
	//}
	//return exists
}

//func SubmitRequest(d *druid.Client, method, path string, opts interface{}) *druid.Response {
//	req, err := d.NewRequest(method, path, opts)
//	if err != nil {
//		fmt.Printf("API request submission failed: %v", err)
//		return nil
//	}
//	//http := retryablehttp.NewClient()
//	if err := d.Do(req, nil); err != nil {
//		fmt.Printf("API request submission failed: %v", err)
//		return nil
//	}
//	//resp, err := http.Do(res)
//	if err != nil {
//		fmt.Println("ERRRRRRRRRRRROR")
//		return nil
//	}
//	// defer resp.Body.Close()
//	response := &druid.Response{resp}
//	return response
//}

func GetValueFromQueryResponse(res *druid.Response, key string) (interface{}, error) {
	//responseBody := make(map[string]interface{})
	var responseBody []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to deserialize the response: %v", err)
	}
	//fmt.Println(responseBody)
	value := responseBody[0][key]
	return value, nil
}

func httpClientWithTLS() (*http.Client, error) {
	// get tls cert, clientCA and rootCA for tls config
	clientCA := x509.NewCertPool()
	rootCA := x509.NewCertPool()

	crt, err := tls.X509KeyPair(TLSCertKey, TLSPrivateKeyKey)
	if err != nil {
		fmt.Println(err, "Failed to parse private key pair")
		return nil, err
	}
	clientCA.AppendCertsFromPEM(TLSCACertFileName)
	rootCA.AppendCertsFromPEM(TLSCACertFileName)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{crt},
				ClientAuth:         tls.NoClientCert,
				InsecureSkipVerify: true,
				ClientCAs:          clientCA,
				RootCAs:            rootCA,
				MaxVersion:         tls.VersionTLS13,
			},
		},
	}

	return client, nil
}

func httpClientWithTLSGPT() (*http.Client, error) {
	// get tls cert, clientCA and rootCA for tls config
	// Load the self-signed certificate
	certFile := "/home/appscode/go/src/github.com/tapojit047/druid-go-client-test/ca-cert.pem"
	caCert, err := ioutil.ReadFile(certFile)
	if err != nil {
		fmt.Println("Failed to read certificate:", err)
		return nil, err
	}

	// Create a certificate pool and add the self-signed certificate
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		fmt.Println("Failed to append certificate")
		return nil, err
	}

	// Create a custom TLS configuration with the certificate pool
	tlsConfig := &tls.Config{
		RootCAs: caCertPool, // Trust only certificates in this pool
	}

	// Create a custom HTTP client with the custom TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	return client, nil
}

//func UpdateCoordinatorDynamicConfig(d *druid.Client) {
//	method := "POST"
//	path := "druid/coordinator/v1/config"
//
//	data := map[string]interface{}{
//		"millisToWaitBeforeDeleting": 500,
//	}
//
//	jsonData, err := json.Marshal(data)
//	if err != nil {
//		fmt.Printf("could not marshal json: %s\n", err)
//	}
//	rawMessage := json.RawMessage(jsonData)
//	response := SubmitRequest(d, method, path, rawMessage)
//	defer func() {
//		err := response.Body.Close()
//		if err != nil {
//			fmt.Println("Error closing response body:", err)
//		}
//	}()
//}

//func RunSelectQuery(d *druid.Client) string {
//	method := "POST"
//	path := "druid/v2/sql"
//
//	data := map[string]interface{}{
//		"query": "SELECT * FROM \"dummy-data\"",
//	}
//	jsonData, err := json.Marshal(data)
//	if err != nil {
//		fmt.Printf("could not marshal json: %s\n", err)
//		return ""
//	}
//	rawMessage := json.RawMessage(jsonData)
//	response := SubmitRequest(d, method, path, rawMessage)
//	defer func() {
//		err := response.Body.Close()
//		if err != nil {
//			fmt.Println("Error closing response body:", err)
//		}
//	}()
//
//	name, err := GetValueFromQueryResponse(response, "name")
//	if err != nil {
//		fmt.Printf("Failed to fetch Name in the response, %v\n", err)
//		return ""
//	}
//	//fmt.Println(name)
//	return name.(string)
//}

func DruidHealthCheck(d *druid.Client) {
	health, _, err := d.Common().Health()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Health: ", *health)
}

//
//func DruidCreateDataSource(d *druid.Client) {
//	//dataSourceExist := false
//	oldData := "0"
//	if CheckDataSourceExist(d) {
//		//fmt.Println("YES")
//		oldData = RunSelectQuery(d)
//		//dataSourceExist = true
//	}
//	//fmt.Println(oldData)
//
//	//startTime := time.Now()
//	//// Create datasource and insert data
//	method := "POST"
//	path := "druid/indexer/v1/task"
//
//	specContent, err := os.ReadFile(fmt.Sprintf("spec-updated%s.json", oldData))
//	if err != nil {
//		log.Fatal(err)
//	}
//	rawMessage := json.RawMessage(specContent)
//
//	response := SubmitRequest(d, method, path, rawMessage)
//	defer func() {
//		err := response.Body.Close()
//		if err != nil {
//			fmt.Println("Error closing response body:", err)
//		}
//	}()
//
//	taskID, err := GetValueFromClusterResponse(response, "task")
//	if err != nil {
//		fmt.Printf("Failed to fetch taskID in the response\n, %v", err)
//		return
//	}
//	//fmt.Printf("Ingestion task: %s has been submitted\n", taskID)
//
//	///// Get task status
//	for true {
//		method = "GET"
//		path = fmt.Sprintf("druid/indexer/v1/task/%s/status", taskID)
//		response = SubmitRequest(d, "GET", path, nil)
//		//fmt.Println(*response)
//		statusRes, err := GetValueFromClusterResponse(response, "status")
//		if err != nil {
//			fmt.Printf("Unable to parse response and get status of task: %v\n", err)
//			return
//		}
//		statusMap := statusRes.(map[string]interface{})
//		status := statusMap["status"].(string)
//		if status == "SUCCESS" {
//			fmt.Printf("Ingestion task: %s has succeed\n", taskID)
//			break
//		} else {
//			//fmt.Printf("Waiting for ingestion task: %s to succeed...\n", taskID)
//		}
//		time.Sleep(2 * time.Second)
//	}
//// Check if segment is loaded
//for true {
//	// Check if segment loaded
//	//path = "druid/coordinator/v1/datasources/dummy-data/loadstatus?forceMetadataRefresh==true"
//	//response = SubmitRequest(d, "GET", path, nil)
//	//status, err := GetValueFromClusterResponse(response, "dummy-data")
//	//fmt.Println(status)
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//
//	//fmt.Println(status)
//	//if status == "100.0" {
//	//	fmt.Sprintf("datasource load successful\n")
//	//} else {
//	//	fmt.Println("Data load unsuccessful\n")
//	//	time.Sleep(2 * time.Second)
//	//	continue
//	//}
//	//fmt.Println(dataSourceExist)
//	if !dataSourceExist {
//		time.Sleep(5 * time.Second)
//	}
//
//	name := RunSelectQuery(d)
//	//fmt.Println(oldData, " ", name)
//	if name != oldData {
//		break
//	}
//	//fmt.Sprintf("lodaing data and datasource...\n")
//
//	if err != nil {
//		fmt.Sprintf("Unable to parse response and get status of datasource\n")
//		return
//	}
//	time.Sleep(3 * time.Second)
//}
//endTime := time.Now()
//duration := endTime.Sub(startTime)
//fmt.Printf("Data Insertion succeeded in %v seconds", duration)
//}

func GetValueFromDSQueryResponse(res *druid.Response) (bool, error) {
	//responseBody := make(map[string]interface{})
	var responseBody []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return false, fmt.Errorf("failed to deserialize the response: %v", err)
	}
	//fmt.Println(responseBody)
	return len(responseBody) != 0, nil
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

func GetSegmentIDFromResponse(res *druid.Response) (string, error) {
	responseBody := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to deserialize the response: %v", err)
	}
	fmt.Println(responseBody)
	taskid := responseBody["task"].(string)
	if taskid == "" {
		fmt.Printf("no task id found in the response")
		return "", fmt.Errorf("no task id found in the response")
	}

	return taskid, nil
}
