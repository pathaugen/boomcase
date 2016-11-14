
package main

import (
	//"html/template"
	//"net/http"
	"time"
	
	"appengine"
	"appengine/datastore"
	//"appengine/user"
)




// ========== ========== ========== ========== ==========
// [START greeting_struct]
/*
type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}
*/
// [END greeting_struct]
// ========== ========== ========== ========== ==========



// ========== ========== ========== ========== ==========
// [START case_struct]
type Case struct {
	Name				string
	Overview			string
	Featuring			string
	FrequencyResponse	string
	
	Length				int8	// uint8 = 0-255
	Width				int8	// uint8 = 0-255
	Height				int8	// uint8 = 0-255
	
	Weight				int8	// uint8 = 0-255
	Battery				int8	// uint8 = 0-255
	Notes				string
	
	Price				int16	// uint16 = 0-65,535
	Watts				int16
	Sold				bool	// bool - Mark as sold
	
	BlobKey				string
	
	DateAdded			time.Time
}
// [END case_struct]
// ========== ========== ========== ========== ==========
// ========== ========== ========== ========== ==========
// caseKey returns the key used for all case entries.
func caseKey(ctx appengine.Context) *datastore.Key {
	// The string "default_case" here could be varied to have multiple types of cases.
	return datastore.NewKey(ctx, "Case", "default_case", 0, nil)
}
// ========== ========== ========== ========== ==========



// ========== ========== ========== ========== ==========
// Images for each Case, and marking which one can be customized
// [START image_struct]
type CaseImage struct {
	CaseId				int64	// e.g. (4,767,482,418,036,736) uint64 = 0-18,446,744,073,709,551,615 uint32 = 0-4,294,967,295
	BlobKey				string	//appengine.BlobKey // e.g. (ahBkZXZ-cG1kLWJvb21jYXNlcicLEgRDYXNlIgxkZWZhdWx0X2Nhc2UMCxIEQ2FzZRiAgICAgIC8CAw)
	
	Customizable		bool	// bool - Mark as customizable
	
	DateAdded			time.Time
}
// [END image_struct]
// ========== ========== ========== ========== ==========
// ========== ========== ========== ========== ==========
// imageKey returns the key used for all case entries.
func caseImageKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "CaseImage", "default_caseimage", 0, nil)
}
// ========== ========== ========== ========== ==========



// ========== ========== ========== ========== ==========
// [START driver_struct]
type Driver struct {
	Name				string
	FrequencyResponse	string
	Diameter			int8	// uint8 = 0-255
	Price				int16	// uint16 = 0-65,535
	
	BlobKey				string
	
	DateAdded			time.Time
}
// [END driver_struct]
// ========== ========== ========== ========== ==========
// ========== ========== ========== ========== ==========
// driverKey returns the key used for all case entries.
func driverKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "Driver", "default_driver", 0, nil)
}
// ========== ========== ========== ========== ==========




// ========== ========== ========== ========== ==========
// Return holds values returned from the datastore but keeps the key and Id for reference
// [START return_struct]
/*
type Return struct {
	Key					*datastore.Key
	Id					uint64 // datastore IDs = 4,767,482,418,036,736 -> uint32 = 4,294,967,295 -> uint64 = 18,446,744,073,709,551,615
	Data				interface{}
}
*/
// [END return_struct]
// ========== ========== ========== ========== ==========




