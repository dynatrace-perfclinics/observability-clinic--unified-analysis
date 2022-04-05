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

package simulator

import (
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/fleet"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/rest"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

func StartSimulation(dtc rest.DTClient, numFleets int, numTaxis int, verbose bool) error {
	idBase := 100_000
	var fleets []fleet.Fleet

	for i := 0; i < numFleets; i++ {
		f := fleet.NewFleet(
			idBase+i,
			// rand.New(rand.NewSource(time.Now().UnixNano())).Intn(899_999)+100_000,
			util.Locations()[i],
			numTaxis,
		)
		f.InitialiseFleet()
		fleets = append(fleets, f)
	}

	for _, f := range fleets {
		go sendFleetMetrics(dtc, f, verbose)
		go sendTaxiMetrics(dtc, f, verbose)
		go sendFleetEvents(dtc, f)
		go sendTaxiEvents(dtc, f)
	}

	select {}
}

func sendFleetMetrics(dtc rest.DTClient, f fleet.Fleet, verbose bool) {
	log.SetHandler(rest.New(dtc))
	log.WithFields(log.Fields{"fleet.id": f.GetId()}).Info("sending fleet metrics")
	entityID, _ := dtc.GetEntityId(fmt.Sprintf("type(easytaxis:smart_fleet),FleetID(%d)", f.GetId()))
	for {
		log.WithFields(log.Fields{
			"fleet.id":      f.GetId(),
			"custom.device": entityID,
		}).Info("sending fleet metrics")
		mintData := f.ToMintData()
		dtc.PostMetrics(mintData)
		fmt.Println(time.Now().Format("02.01.2006 - 15:04:05"), ": Sent fleet metrics for fleet", f.GetId())
		if verbose {
			fmt.Println(mintData)
		}
		time.Sleep(2 * time.Minute)
	}
}
func sendTaxiMetrics(dtc rest.DTClient, f fleet.Fleet, verbose bool) {
	log.SetHandler(rest.New(dtc))
	for {
		for _, t := range f.GetTaxis() {
			entityID, _ := dtc.GetEntityId(fmt.Sprintf("type(easytaxis:smart_taxi),TaxiID(%d)", t.GetId()))
			log.WithFields(log.Fields{
				"fleet.id":      f.GetId(),
				"taxi.id":       t.GetId(),
				"custom.device": entityID,
			}).Info("sending taxi metrics")
			mintData := t.ToMintData()
			dtc.PostMetrics(mintData)
			fmt.Println(time.Now().Format("02.01.2006 - 15:04:05"), ": Sent taxi metrics for taxi", t.GetId())
			if verbose {
				fmt.Println(mintData)
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
func sendFleetEvents(dtc rest.DTClient, f fleet.Fleet) {
	for {
		dtc.PostEvent(f.CreateTrafficInfoEvent())
		time.Sleep(10 * time.Minute)
	}
}
func sendTaxiEvents(dtc rest.DTClient, f fleet.Fleet) {
	for {
		for _, t := range f.GetTaxis() {
			dtc.PostEvent(f.CreateCustomerRequestEvent())
			dtc.PostEvent(t.CreateAcceptCustomerEvent())
		}
		time.Sleep(5 * time.Minute)
	}
}
