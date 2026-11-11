package main

import (
	"fmt"
	"iter"
	"slices"
)

type RadioStation struct {
	frequency float32
}

func NewRadioStation(frequency float32) *RadioStation {
	return &RadioStation{frequency: frequency}
}

func (rs *RadioStation) Frequency() float32 {
	return rs.frequency
}

type StationList struct {
	stations []*RadioStation
}

func NewStationList() *StationList {
	return &StationList{
		stations: make([]*RadioStation, 0),
	}
}

func (sl *StationList) AddStation(station *RadioStation) {
	sl.stations = append(sl.stations, station)
}

func (sl *StationList) RemoveStation(station *RadioStation) {
	_ = slices.DeleteFunc(sl.stations, func(rs *RadioStation) bool {
		return rs.Frequency() == station.Frequency()
	})
}

func (sl *StationList) All() iter.Seq[*RadioStation] {
	return func(yield func(*RadioStation) bool) {
		for _, e := range sl.stations {
			if !yield(e) {
				return
			}
		}
	}
}

func main() {
	stationList := NewStationList()

	stationList.AddStation(NewRadioStation(89))
	stationList.AddStation(NewRadioStation(101))
	stationList.AddStation(NewRadioStation(102))
	stationList.AddStation(NewRadioStation(103.2))

	for rs := range stationList.All() {
		fmt.Println(rs.Frequency())
	}

	stationList.RemoveStation(NewRadioStation(89))

	fmt.Println(stationList.stations)
}
