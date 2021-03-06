package libknot

import (
	"testing"
)

// TestDataToFromCtlEmpty makes sure that the consistency of the data between the C and Go structures remains good with nil values
func TestDataToFromCtlEmpty(t *testing.T) {
	testData := KnotCtlData{}

	t.Log("KnotCtlData.ToCtl() : Converting testData (type KnotCtlData) to testCtl (type C.knot_ctl_data_t)")
	testCtl := testData.ToCtl()

	t.Log("KnotCtlData.FromCtl(<testCtl>) : Converting testCtl (type C.knot_ctl_data_t) to testStruct (type KnotCtlData)")
	var dataStruct KnotCtlData
	dataStruct.FromCtl(testCtl)

	if dataStruct != testData {
		t.Errorf("Excepted %s from testData, got %s", testData, dataStruct)
	}
}

// TestDataToFromCtlEmpty makes sure that the consistency of the data between the C and Go structures remains good with non empty values
func TestDataToFromCtl(t *testing.T) {
	testData := KnotCtlData{
		Command: "cmd",
		Flags:   "flags",
		Error:   "err",
		Section: "sec",
		Item:    "item",
		Id:      "id",
		Zone:    "zone",
		Owner:   "own",
		Ttl:     "ttl",
		Type:    "type",
		Data:    "data",
		Filter:  "filter",
	}

	t.Log("KnotCtlData.ToCtl() : Converting testData (type KnotCtlData) to testCtl (type C.knot_ctl_data_t)")
	testCtl := testData.ToCtl()

	t.Log("KnotCtlData.FromCtl(<testCtl>) : Converting testCtl (type C.knot_ctl_data_t) to testStruct (type KnotCtlData)")
	var dataStruct KnotCtlData
	dataStruct.FromCtl(testCtl)

	if dataStruct != testData {
		t.Errorf("Excepted %s from testData, got %s", testData, dataStruct)
	}
}
