package keeper_test

import "fmt"

func (s *KeeperTestSuite) TestGetParams() {
	params := s.dKeeper.GetParams(s.Ctx)
	fmt.Println(params)
}

func (s *KeeperTestSuite) TestSetParams() {

}
