package mock

import (
	"github.com/GabbyyLS/data-linter/models"
//	"github.com/GabbyyLS/data-linter/linter/checker"
	"fmt"
)

var (
	CheckerList = getMockCheckerList()
	EmptyCount = 0
	NonEmptyCount = 50
	Step = 15
)

type MockChecker struct {
	Name string
}

func (mockChecker MockChecker) Check() ([]*models.Problem, error) {
	return Problems, nil
}

func getMockCheckerList() []*MockChecker {
	checkerList := []*MockChecker{}

	for i := 0; i < Step; i++ {
//		checker := checker.Checker()
		checkerList = append(checkerList, &MockChecker{Name: fmt.Sprintf("Mock Checker #%d", i)})
	}

	return checkerList
}
