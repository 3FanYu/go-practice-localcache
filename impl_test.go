package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LocalCacheSuite struct {
	suite.Suite
	cache Cache
}

// type LocalCacheMock struct{
// 	mock.Mock
// }

func TestLocalCacheSuite(t *testing.T) {
	suite.Run(t, new(LocalCacheSuite))
}

func (suite *LocalCacheSuite) SetupTest() {
	suite.cache = New()
}

func (suite *LocalCacheSuite) TestLocalCache() {

	tests := []struct {
		key    string
		data   interface{}
		expect interface{}
	}{
		{key: "key1", data: 1, expect: 1},
		{key: "key2", data: "1", expect: "1"},
		{key: "key3", data: true, expect: true},
		{key: "key4", data: 1.99, expect: 1.99},
		{key: "key5", data: []int{1, 2, 3, 4}, expect: []int{1, 2, 3, 4}},
		{key: "key6", data: map[string]string{"name": "Jeff"}, expect: map[string]string{"name": "Jeff"}},
	}

	for _, tc := range tests {
		suite.cache.Set(tc.key, tc.data)

		cd, _ := suite.cache.Get(tc.key)
		require.Equal(suite.T(), tc.expect, cd, "should be equal")
	}
}

func (suite *LocalCacheSuite) TestCacheOverwrite() {
	expect := 2
	key := "key1"
	suite.cache.Set(key, 1)
	suite.cache.Set(key, expect)
	cd, _ := suite.cache.Get(key)

	require.Equal(suite.T(), expect, cd, "should overwrite cache")
}

func (suite *LocalCacheSuite) TestCacheExpiration() {
	expiredIn = 0 * time.Second
	key := "key1"
	cache := New()
	cache.Set(key, 1)
	time.Sleep(1 * time.Second)
	cd, _ := cache.Get(key)
	require.Equal(suite.T(), nil, cd, "should be nil")
}
