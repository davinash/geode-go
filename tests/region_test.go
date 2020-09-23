package tests

func (suite *GeodeTestSuite) TestRegionOps() {
	regionName := "TestRegionOpsR"
	err := suite.createRegion(regionName, Replicate)
	if err != nil {
		suite.Fail("Failed to create Region %v", err)
	}
	suite.DoRegionOps(regionName, Replicate)

	regionName = "TestRegionOpsP"
	err = suite.createRegion(regionName, Partition)
	if err != nil {
		suite.Fail("Failed to create Region %v", err)
	}
	suite.DoRegionOps(regionName, Partition)

	regionNames, err := suite.Client.GetRegionNames()
	if err != nil {
		suite.Fail("GetRegionNames Failed, Error = %v", err)
	}
	if len(regionNames) != 2 {
		suite.Fail("GetRegionNames should return 2 regions, Actual = %v", len(regionNames))
	}
}
