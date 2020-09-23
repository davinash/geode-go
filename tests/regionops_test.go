package tests

import (
	"fmt"
	"github.com/davinash/geode-go/pkg"
	"strings"
	"sync"
)

func (suite *GeodeTestSuite) DoRegionOps(regionName string, regionType RegionType) {
	var wgPut sync.WaitGroup
	for i := 0; i < 20; i++ {
		wgPut.Add(1)
		go func(wg *sync.WaitGroup, keyIndex int) {
			defer wg.Done()
			err2 := suite.Client.Region(regionName).Put(fmt.Sprintf("Key-%d", keyIndex),
				fmt.Sprintf("Val-%d", keyIndex))
			if err2 != nil {
				suite.Fail(fmt.Sprintf("Failed during Put, Error = %v\n", err2))
			}
		}(&wgPut, i)
	}
	wgPut.Wait()

	var wgGet sync.WaitGroup
	for i := 0; i < 20; i++ {
		wgGet.Add(1)
		go func(wg *sync.WaitGroup, keyIndex int) {
			defer wg.Done()
			value, err2 := suite.Client.Region(regionName).Get(fmt.Sprintf("Key-%d", keyIndex))
			if err2 != nil {
				suite.Fail(fmt.Sprintf("Failed during Get, Error = %v", err2))
			}
			if value != fmt.Sprintf("Val-%d", keyIndex) {
				suite.Fail(fmt.Sprintf("Value Mismatch, Expected = %s Actual = %s",
					fmt.Sprintf("Val-%d", keyIndex), value))
			}
		}(&wgGet, i)
	}
	wgGet.Wait()

	//Size
	if regionType == Replicate {
		size, err := suite.Client.Region(regionName).Size()
		if err != nil {
			suite.Fail(fmt.Sprintf("Size Failed, Error = %v", err))
		}

		if 20 != size {
			suite.Fail("size of region should be 20")
		}
	}
	//KeySet
	keySet, err := suite.Client.Region(regionName).KeySet()
	if err != nil {
		suite.Fail("KeySet Failed, Error = %v", err)
	}
	if len(keySet) != 20 {
		suite.Fail(fmt.Sprintf("Did not return all the keys, Expected 20, Actual = %v", len(keySet)))
	}

	// PutIfAbsent
	v, err := suite.Client.Region(regionName).PutIfAbsent("Key-1", "Value-New")
	if err != nil {
		suite.Fail(fmt.Sprintf("PutIfAbsent failed, Error = %v", err))
	}
	if v != "Val-1" {
		suite.Fail(fmt.Sprintf("PutIfAbsent failed, Expected %v ", v))
	}
	v, err = suite.Client.Region(regionName).PutIfAbsent("Key-100", "Value-New-100")
	if err != nil {
		suite.Fail(fmt.Sprintf("PutIfAbsent failed, Error = %v", err))
	}
	if v != nil {
		suite.Fail("PutIfAbsent should return nil")
	}

	//PutAll
	kvs := make([]*pkg.KeyValue, 0)
	for i := 0; i < 10; i++ {
		kvs = append(kvs, &pkg.KeyValue{
			Key:   fmt.Sprintf("Key-PutAll-%d", i),
			Value: fmt.Sprintf("Val-PutAll-%d", i),
		})
	}
	putAll, err := suite.Client.Region(regionName).PutAll(kvs)
	if err != nil {
		suite.Fail(fmt.Sprintf("PutAll failed, Error =  %v", err))
	}
	if len(putAll) != 0 {
		suite.Fail(fmt.Sprintf("PutAll Failed, Error = %v", err))
	}

	// GetAll
	keys := make([]*pkg.Keys, 0)
	for i := 0; i < 10; i++ {
		keys = append(keys, &pkg.Keys{
			Key: fmt.Sprintf("Key-PutAll-%d", i),
		})
	}

	values, err := suite.Client.Region(regionName).GetAll(keys)
	if err != nil {
		suite.Fail("GetAll failed, Error = %v", err)
	}
	for _, v := range values {
		tokens := strings.Split(v.Key.(string), "-")
		expectedValue := fmt.Sprintf("Val-PutAll-%s", tokens[len(tokens)-1])
		if v.Value != expectedValue {
			suite.Fail(fmt.Sprintf("GetAll Failed Expected Value = %v, Actual Value = %v",
				expectedValue, v.Value))
		}
	}
	// Remove
	key := "Key-PutAll-0"
	err = suite.Client.Region(regionName).Remove(key)
	if err != nil {
		suite.Fail(fmt.Sprintf("Removed Failed, Error = %v", err))
	}
	v, err = suite.Client.Region(regionName).Get(key)
	if err != nil {
		suite.Fail(fmt.Sprintf("Get Failed, Error = %v", err))
	}
	if v != nil {
		suite.Fail("Remove failed, Get got the value")
	}

	// Clear
	if regionType == Replicate {
		err = suite.Client.Region(regionName).Clear()
		if err != nil {
			suite.Fail(fmt.Sprintf("Clear failed, Error = %v", err))
		}
	}

	if regionType == Replicate {
		size, err := suite.Client.Region(regionName).Size()
		if err != nil {
			suite.Fail(fmt.Sprintf("Size Failed, Error = %v", err))
		}
		if 0 != size {
			suite.Fail("size of region should be 20")
		}
	}
}
