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
	var expectedValues = []int32{1, 2, 3, 4, 100}
	if len(result) != len(expectedValues) {
		suite.FailNow(fmt.Sprintf("Mismatch in result length, Expected = %d, Actual = %d",
			len(expectedValues), len(result)))
	}
	for i, v := range result {
		if v.(int32) != expectedValues[i] {
			suite.FailNow(fmt.Sprintf("Value Mistmatch, Expected = %v, Actual = %v",
				expectedValues[i], v))
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

	var expectedValues = []string{"server-1", "Success", "server-2", "Success"}
	if len(result) != len(expectedValues) {
		suite.FailNow(fmt.Sprintf("Mismatch in result length, Expected = %d, Actual = %d",
			len(expectedValues), len(result)))
	}
	for i, v := range result {
		if v.(string) != expectedValues[i] {
			suite.FailNow(fmt.Sprintf("Value Mistmatch, Expected = %v, Actual = %v",
				expectedValues[i], v))
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
