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

package fleet

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/radu-stefan-dt/fleet-simulator/pkg/constants"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/models"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

type Fleet interface {
	GetId() int
	GetLocation() string
	GetAvailableCars() int
	GetBusyCars() int
	GetTotalCars() int
	GetCustomerQueue() int
	GetTaxis() []Taxi
	UpdateCars(int, int, int)
	MakeCarBusy()
	UpdateQueue(int)
	RegisterTaxi(Taxi)
	InitialiseFleet()
	ToMintDimensions() string
	ToMintData() string
	CreateTrafficInfoEvent() []byte
	CreateCustomerRequestEvent() []byte
}

type fleetImpl struct {
	id            int
	location      string
	carsAvailable int
	carsBusy      int
	carsTotal     int
	customerQueue int
	taxis         []Taxi
}

func (f fleetImpl) GetId() int {
	return f.id
}
func (f fleetImpl) GetLocation() string {
	return f.location
}
func (f fleetImpl) GetTotalCars() int {
	return f.carsTotal
}
func (f fleetImpl) GetAvailableCars() int {
	return f.carsAvailable
}
func (f fleetImpl) GetBusyCars() int {
	return f.carsBusy
}
func (f fleetImpl) GetCustomerQueue() int {
	return f.customerQueue
}
func (f fleetImpl) GetTaxis() []Taxi {
	return f.taxis
}
func (f fleetImpl) ToMintDimensions() string {
	return fmt.Sprintf("fleet.id=\"%d\",fleet.location=\"%s\"", f.GetId(), f.GetLocation())
}
func (f fleetImpl) ToMintData() string {
	var sb strings.Builder
	dimensions := f.ToMintDimensions()
	sb.WriteString(fmt.Sprintf("%s%s,%s %d\n", constants.MetricPrefix, "fleet.cars.available", dimensions, f.GetAvailableCars()))
	sb.WriteString(fmt.Sprintf("%s%s,%s %d\n", constants.MetricPrefix, "fleet.cars.busy", dimensions, f.GetBusyCars()))
	sb.WriteString(fmt.Sprintf("%s%s,%s %d\n", constants.MetricPrefix, "fleet.cars.total", dimensions, f.GetTotalCars()))
	sb.WriteString(fmt.Sprintf("%s%s,%s %d\n", constants.MetricPrefix, "fleet.queue", dimensions, f.GetCustomerQueue()))
	return sb.String()
}
func (f fleetImpl) CreateTrafficInfoEvent() []byte {
	eventRaw := models.EventIngest{
		EventType:      "CUSTOM_INFO",
		Title:          "Received updated traffic information",
		StartTime:      time.Now().UTC().UnixMilli(),
		EndTime:        time.Now().UTC().UnixMilli(),
		EntitySelector: fmt.Sprintf("type(easytaxis:smart_fleet),FleetID(%d)", f.GetId()),
		Properties: map[string]string{
			"FleetID":        fmt.Sprintf("%d", f.GetId()),
			"Location":       f.GetLocation(),
			"TrafficDetails": "No accidents. Maintain normal route operation",
		},
	}
	eventEncoded, _ := json.Marshal(eventRaw)
	return eventEncoded
}
func (f fleetImpl) CreateCustomerRequestEvent() []byte {
	eventRaw := models.EventIngest{
		EventType:      "CUSTOM_INFO",
		Title:          "New customer booking requested",
		StartTime:      time.Now().UTC().UnixMilli(),
		EndTime:        time.Now().UTC().UnixMilli(),
		EntitySelector: fmt.Sprintf("type(easytaxis:smart_fleet),FleetID(%d)", f.GetId()),
		Properties: map[string]string{
			"Destination": fmt.Sprintf("%f x %f", util.RandomCoord(), util.RandomCoord()),
		},
	}
	eventEncoded, _ := json.Marshal(eventRaw)
	return eventEncoded
}

func (f *fleetImpl) MakeCarBusy() {
	if f.carsAvailable-1 >= 0 {
		f.carsAvailable--
		f.carsBusy++
	}
}
func (f *fleetImpl) UpdateCars(available, busy, total int) {
	f.carsAvailable = available
	f.carsBusy = busy
	f.carsTotal = total
}
func (f *fleetImpl) UpdateQueue(q int) {
	f.customerQueue = q
}
func (f *fleetImpl) RegisterTaxi(t Taxi) {
	f.taxis = append(f.taxis, t)
}
func (f *fleetImpl) InitialiseFleet() {
	for i := 0; i < f.GetTotalCars(); i++ {
		var class string
		switch {
		case i%3 == 0:
			class = "limo"
		case i%3 == 1:
			class = "executive"
		default:
			class = "casual"
		}
		tID := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(89_999_999) + 10_000_000
		reg := util.GenerateRegNumber()
		t := NewTaxi(tID, class, f.GetId(), reg)
		f.RegisterTaxi(t)
		time.Sleep(time.Nanosecond) // ensures next random seed is different
	}
}

func NewFleet(id int, location string, carsTotal int) Fleet {
	return &fleetImpl{
		id:            id,
		location:      location,
		carsAvailable: carsTotal,
		carsBusy:      0,
		carsTotal:     carsTotal,
		customerQueue: 0,
		taxis:         []Taxi{},
	}
}
