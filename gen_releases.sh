#!/bin/sh

tag=$1

generate_release()
{
	GOOS=$1 GOARCH=$2 go build
	zip kafka-offset-mgr_$tag_$1_$2.zip kafka-offset-mgr
}

generate_release linux 386
generate_release linux amd64
generate_release darwin 386
generate_release darwin amd64
generate_release windows 386
generate_release windows amd64

rm kafka-offset-mgr
rm kafka-offset-mgr.exe