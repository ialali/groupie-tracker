package global

import "net/http"

type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    LocationsData
	Dates        DateData
	Relations    RelationsData
}
type Relation struct {
	Id             int                    `json:"id"`
	DatesLocations map[string]interface{} `json:"datesLocations"`
}
type RelationsData struct {
	Index []Relation `json:"index"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type DateData struct {
	Index []Date `json:"index"`
}
type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
	Dates    string   `json:"dates"`
}
type LocationsData struct {
	Index []Location `json:"index"`
}

var Client *http.Client
var Locations LocationsData
var Artists []Artist
var Dates DateData
var Relations RelationsData

func AppendtoStruct() {
	for index := range Locations.Index {
		Artists[index].Locations.Index = append(Artists[index].Locations.Index, Locations.Index[index])
	}

	for index := range Dates.Index {
		Artists[index].Dates.Index = append(Artists[index].Dates.Index, Dates.Index[index])
	}

	for index := range Relations.Index {
		Artists[index].Relations.Index = append(Artists[index].Relations.Index, Relations.Index[index])
	}
}
