#!/bin/bash

mkdir -p "day$1"

cp tpl/dayX_cmd.go.tpl cmd/day$1_cmd.go
sed -i s/X/$1/g cmd/day$1_cmd.go

cp tpl/dayX.go.tpl day$1/day$1.go
sed -i s/X/$1/g day$1/day$1.go


cp tpl/dayX_test.go.tpl day$1/day$1_test.go
sed -i s/X/$1/g day$1/day$1_test.go


