/**
 * Copyright (c) 2021 Radu Stefan
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 **/

package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/radu-stefan-dt/fleet-simulator/pkg/models"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

type DTClient interface {
	PostMetrics(string)
	PostLogEvent([]byte)
	PostEvent([]byte)
	GetEntityId(selector string) (string, error)
}

type DTClientImpl struct {
	baseURL string
	token   string
}

func doRequest(req *http.Request) []byte {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		util.PrintError(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.PrintError(err)
	}
	if !models.ApiResponses[resp.StatusCode] {
		fmt.Println("Got unexpected response code:", resp.StatusCode, "(", resp.StatusCode, ")")
		fmt.Println(string(body))
		fmt.Println(req.RequestURI)
		fmt.Println(req.URL.Query())
	}

	return body
}

func (dtc DTClientImpl) PostMetrics(data string) {
	payload := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest("POST", dtc.baseURL+models.MetricIngestAPI, payload)
	if err != nil {
		util.PrintError(err)
	}
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Authorization", "Api-Token "+dtc.token)
	doRequest(req)
}

func (dtc DTClientImpl) PostLogEvent(content []byte) {
	payload := bytes.NewBuffer([]byte(content))

	req, err := http.NewRequest("POST", dtc.baseURL+models.LogsIngestAPI, payload)
	if err != nil {
		util.PrintError(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Api-Token "+dtc.token)
	doRequest(req)
}

func (dtc DTClientImpl) PostEvent(content []byte) {
	payload := bytes.NewBuffer([]byte(content))

	req, err := http.NewRequest("POST", dtc.baseURL+models.EventsIngestAPI, payload)
	if err != nil {
		util.PrintError(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Api-Token "+dtc.token)
	doRequest(req)
}

func (dtc DTClientImpl) GetEntityId(selector string) (string, error) {
	req, err := http.NewRequest("GET", dtc.baseURL+models.EntitiesAPI, nil)
	if err != nil {
		util.PrintError(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Api-Token "+dtc.token)

	params := req.URL.Query()
	params.Add("from", "now-10m")
	params.Add("entitySelector", selector)
	req.URL.RawQuery = params.Encode()

	rawResp := doRequest(req)

	var parsedResp models.EntitiesAPIResponse
	json.Unmarshal(rawResp, &parsedResp)
	if parsedResp.TotalCount > 0 {
		return parsedResp.Entities[0].EntityID, nil
	}
	return "", fmt.Errorf("could not find entity")
}

func NewDTClient(tenant, token string) DTClient {
	return &DTClientImpl{
		baseURL: models.Protocol + tenant,
		token:   token,
	}
}
