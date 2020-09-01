package tests

import (
	"github.com/google/uuid"
	"sync"
)

func (suite *GeodeTestSuite) TestRegionOps() {
	regionName := "TestRegionOps"
	err := suite.createRegion(regionName, Replicate)
	if err != nil {
		suite.Fail("Failed to create Region %v", err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			k := uuid.New().String()
			err2 := suite.Client.Region(regionName).Put(k, k)
			if err2 != nil {
				suite.Fail("Failed during Put, Error = %v\n", err)
			}
		}(&wg)
	}
	wg.Wait()
}

//func TestBasicOp(t *testing.T) {
//	t.Run("BasicOp", func(t *testing.T) {
//		geodeHome := os.Getenv("GEODE_HOME")
//		if geodeHome == "" {
//			t.Fatalf("Define Environment variable GEODE_HOME")
//		}
//		err := startLocator(geodeHome)
//		if err != nil {
//			t.FailNow()
//		}
//		defer stopCluster(geodeHome)
//
//		ports, err := startServers(geodeHome, 2)
//		if err != nil {
//			t.FailNow()
//		}
//		geodeClient, err := client.NewClient(100)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		regionName := "ReplicatedRegion"
//		err = createRegion(geodeHome, regionName, Replicate)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		for _, p := range ports {
//			err = geodeClient.AddServer("127.0.0.1", p)
//			if err != nil {
//				log.Fatalln(err)
//			}
//		}
//
//		names, err := geodeClient.GetRegionNames()
//		if err != nil {
//			t.FailNow()
//		}
//
//		if len(names) != 1 {
//			t.FailNow()
//		}
//
//		var wg sync.WaitGroup
//		for i := 0; i < 20; i++ {
//			wg.Add(1)
//			go func(wg *sync.WaitGroup, client *client.GeodeClient, regionName string) {
//				k := uuid.New().String()
//				err2 := client.Region(regionName).Put(k, k)
//				if err2 != nil {
//					t.FailNow()
//				}
//			}(&wg, geodeClient, regionName)
//		}
//		wg.Wait()
//
//
//	})
//}
