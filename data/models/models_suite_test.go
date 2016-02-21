package models_test

import (
	"github.com/bbuck/dragon-mud/data"
	"github.com/bbuck/dragon-mud/data/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var (
	testFactory = &TestFactory{}
	old         data.Factory
)

var _ = BeforeSuite(func() {
	old = data.DefaultFactory
	data.DefaultFactory = testFactory
	models.MigrateDatabase()
})

var _ = AfterSuite(func() {
	data.DefaultFactory = old
	testFactory.Cleanup()
})