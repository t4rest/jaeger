#!/usr/bin/env bash


# go get github.com/golang/mock/gomock
# go get github.com/golang/mock/mockgen

mockgen gitlab-il.cyren.io/apollo/subscription-dispatcher/dao DAO > ./dao/mock_dao/dao.go

mockgen gitlab-il.cyren.io/apollo/subscription-dispatcher/events Publisher > ./events/mock_publish/pub.go

mockgen gitlab-il.cyren.io/apollo/subscription-dispatcher/cache Cacher > ./cache/mock_cache/cache.go
