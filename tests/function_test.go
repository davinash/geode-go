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
	log.Printf(fmt.Sprintf("ExecuteFunctionOnRegion Result = %v", result))
}
