#!/usr/bin/env bash
mkdir -p api
cp ../main.go api
sed -i '' 's/func main/func old_main/g' api/main.go
cat api/main.go
./now.sh --prod "$@"
rm -rf api