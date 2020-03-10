package mocks

import "github.com/stretchr/testify/mock"

type FetcherMock struct {
	mock.Mock
}

func (f *FetcherMock) Fetch(s string) (string, error) {
	args := f.Called(s)
	return args.String(0), args.Error(1)
}
