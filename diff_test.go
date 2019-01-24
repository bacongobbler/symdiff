package symdiff

import (
	"reflect"
	"testing"
)

type fixture struct {
	A string
	B int
}

func TestDiff(t *testing.T) {
	var tests = []struct {
		Name                 string
		FixtureA             fixture
		FixtureB             fixture
		ExpectedResult       fixture
		ShouldError          bool
		ExpectedErrorMessage string
	}{
		{
			Name: "simple: identical",
			FixtureA: fixture{
				A: "hi",
				B: 1,
			},
			FixtureB: fixture{
				A: "hi",
				B: 1,
			},
			ExpectedResult: fixture{},
		},
		{
			Name: "simple: different strings",
			FixtureA: fixture{
				A: "hi",
				B: 1,
			},
			FixtureB: fixture{
				A: "bye",
				B: 1,
			},
			ExpectedResult: fixture{
				A: "hi",
			},
		},
		{
			Name: "simple: different integers",
			FixtureA: fixture{
				A: "hi",
				B: 1,
			},
			FixtureB: fixture{
				A: "hi",
				B: 2,
			},
			ExpectedResult: fixture{
				B: 1,
			},
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			err := Diff(&tc.FixtureA, tc.FixtureB)
			if tc.ShouldError {
				if err == nil {
					t.Error("expected an error to occur, got nil")
				} else if err.Error() != tc.ExpectedErrorMessage {
					t.Error(err)
				}
				// we were just checking if an error would occur, so we can skip the rest
				return
			} else if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(tc.ExpectedResult, tc.FixtureA) {
				t.Errorf("results differ.\nExpected: %v\n Actual: %v", tc.ExpectedResult, tc.FixtureA)
			}
		})
	}
}
