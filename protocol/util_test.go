package protocol

import (
	"log"
	"testing"
)

func TestUtilFromBytes8AndToBytes8(t *testing.T) {
	testValues := []uint8{0, 1, 10, 100, 127}
	var testBytes [][]byte
	var testResults []uint8

	for _, val := range testValues {
		tB, err := toBytes8(val)
		if err != nil {
			t.Error(err)
		}
		testBytes = append(testBytes, tB)
	}

	for _, val := range testBytes {
		rInt, err := fromBytes8(val)
		if err != nil {
			t.Error(err)
		}
		testResults = append(testResults, rInt)
	}

	for i, val := range testResults {
		if val != testValues[i] {
			t.Errorf("Incorrect value %v should be %v", val, testValues[i])
		}
	}
	t.Logf("orig: %v\nresult: %v\n", testValues, testResults)
	log.Printf("orig: %v\nresult: %v\n", testValues, testResults)
}

func TestUtilFromBytes8AndToBytes32(t *testing.T) {
	testValues := []uint32{0, 1, 10, 100, 127, 2147483647}
	var testBytes [][]byte
	var testResults []uint32

	for _, val := range testValues {
		tB, err := toBytes32(val)
		if err != nil {
			t.Error(err)
		}
		testBytes = append(testBytes, tB)
	}

	for _, val := range testBytes {
		rInt, err := fromBytes32(val)
		if err != nil {
			t.Error(err)
		}
		testResults = append(testResults, rInt)
	}

	for i, val := range testResults {
		if val != testValues[i] {
			t.Errorf("Incorrect value %v should be %v", val, testValues[i])
		}
	}
	t.Logf("orig: %v\nresult: %v\n", testValues, testResults)
	log.Printf("orig: %v\nresult: %v\n", testValues, testResults)
}

func TestTest(t *testing.T) {
	var a uint8 = 61
	log.Printf("t: %v \n", MsgType(a))
}
