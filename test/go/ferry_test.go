package test

import (
	"testing"

	"github.com/Shopify/ghostferry"
	"github.com/Shopify/ghostferry/testhelpers"
	"github.com/stretchr/testify/suite"
)

type FerryTestSuite struct {
	*testhelpers.GhostferryUnitTestSuite

	ferry *ghostferry.Ferry
}

func (t *FerryTestSuite) SetupTest() {
	t.GhostferryUnitTestSuite.SetupTest()
}

func (t *FerryTestSuite) TearDownTest() {
	_, err := t.Ferry.TargetDB.Exec("SET GLOBAL read_only = OFF")
	t.Require().Nil(err)
}

func (t *FerryTestSuite) TestReadOnlyDatabaseFailsInitialization() {
	_, err := t.Ferry.TargetDB.Exec("SET GLOBAL read_only = ON")
	t.Require().Nil(err)

	ferry := testhelpers.NewTestFerry().Ferry // make new ferry that re-uses the same targetDB as t.Ferry
	err = ferry.Initialize()
	t.Require().Equal("@@read_only must be OFF on target db", err.Error())

	_, err = t.Ferry.TargetDB.Exec("SET GLOBAL read_only = OFF")
	t.Require().Nil(err)

	ferry = testhelpers.NewTestFerry().Ferry
	err = ferry.Initialize()
	t.Require().Nil(err)
}

func TestFerryTestSuite(t *testing.T) {
	suite.Run(t, &FerryTestSuite{GhostferryUnitTestSuite: &testhelpers.GhostferryUnitTestSuite{}})
}
