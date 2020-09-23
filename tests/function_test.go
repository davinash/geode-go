package tests

import (
	"fmt"
	"log"
)

func (suite *GeodeTestSuite) TestFunction_OnRegion() {
	regionName := "TestFunction_OnRegion"
	err := suite.createRegion(regionName, Replicate)
	if err != nil {
		suite.Fail("Failed to create Region %v", err)
	}
	result, err := suite.Client.FunctionService().ExecuteFunctionOnRegion("MyFunction",
		regionName, nil, nil)
	if err != nil {
		suite.Fail(fmt.Sprintf("ExecuteFunctionOnRegion Failed, Error = %v", err))
	}
	var expectedValues = map[int]bool{
		1:   false,
		2:   false,
		3:   false,
		4:   false,
		100: false,
	}
	if len(result) != len(expectedValues) {
		suite.FailNow(fmt.Sprintf("Mismatch in result length, Expected = %d, Actual = %d",
			len(expectedValues), len(result)))
	}
	for _, v := range result {
		value := v.(int32)
		_, ok := expectedValues[int(value)]
		if !ok {
			suite.Fail(fmt.Sprintf("Unexpected Value Received %v", v))
		}
		expectedValues[int(value)] = true
	}
	for _, value := range expectedValues {
		if value == false {
			suite.Fail("Did not received all the values")
		}
	}
}

func (suite *GeodeTestSuite) TestFunction_OnRegionThrowsException() {
	regionName := "TestFunction_OnRegionThrowsException"
	err := suite.createRegion(regionName, Replicate)
	if err != nil {
		suite.Fail("Failed to create Region %v", err)
	}
	_, err = suite.Client.FunctionService().ExecuteFunctionOnRegion("MyFunctionException",
		regionName, nil, nil)
	if err == nil {
		suite.Fail("ExecuteFunctionOnRegion should have Failed")
	}
}

func (suite *GeodeTestSuite) TestFunction_OnMember() {
	result, err := suite.Client.FunctionService().ExecuteFunctionOnMember("MyMemberFunction",
		[]string{"server-1", "server-2"}, nil)
	if err != nil {
		suite.FailNow(fmt.Sprintf("ExecuteFunctionOnMember Failed, Error = %v", err))
	}
	log.Printf("-----> %v", result)

	var expectedValues = map[string]bool{
		"server-1": false,
		"server-2": false,
	}
	if len(result) != len(expectedValues) {
		suite.FailNow(fmt.Sprintf("Mismatch in result length, Expected = %d, Actual = %d",
			len(expectedValues), len(result)))
	}
	for _, v := range result {
		_, ok := expectedValues[v.(string)]
		if !ok {
			suite.Fail(fmt.Sprintf("Unexpected Value Received %v", v))
		}
		expectedValues[v.(string)] = true
	}
	for _, value := range expectedValues {
		if value == false {
			suite.Fail("Did not received all the values")
		}
	}
}

//func (suite *GeodeTestSuite) TestFunction_OnGroup() {
//	result, err := suite.Client.FunctionService().ExecuteFunctionOnGroup("MyMemberFunction",
//		[]string{"MyGroup"}, nil)
//	if err != nil {
//		suite.FailNow(fmt.Sprintf("ExecuteFunctionOnGroup Failed, Error = %v", err))
//	}
//	var expectedValues = []string{"server-0", "Success", "server-1", "Success", "server-2", "Success"}
//
//	if len(result) != len(expectedValues) {
//		suite.FailNow(fmt.Sprintf("Mismatch in result length, Expected = %d, Actual = %d",
//			len(expectedValues), len(result)))
//	}
//	for i, v := range result {
//		if v.(string) != expectedValues[i] {
//			suite.FailNow(fmt.Sprintf("Value Mistmatch, Expected = %v, Actual = %v",
//				expectedValues[i], v))
//		}
//	}
//}
