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

package models

const (
	Protocol        = "https://"
	MetricIngestAPI = "/api/v2/metrics/ingest"
	LogsIngestAPI   = "/api/v2/logs/ingest"
	EventsIngestAPI = "/api/v2/events/ingest"
	EntitiesAPI     = "/api/v2/entities"
)

var ApiResponses = map[int]bool{
	200: true,
	201: true,
	202: true,
	204: true,
}

type EventIngest struct {
	EventType      string            `json:"eventType"`
	Title          string            `json:"title"`
	StartTime      int64             `json:"startTime"`
	EndTime        int64             `json:"endTime"`
	EntitySelector string            `json:"entitySelector"`
	Properties     map[string]string `json:"properties"`
}

type EntitiesAPIResponse struct {
	TotalCount int `json:"totalCount"`
	Entities   []MonitoredEntity
}

type MonitoredEntity struct {
	EntityID    string `json:"entityId"`
	DisplayName string `json:"displayName"`
}
