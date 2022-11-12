#!/bin/bash
# source: https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

platforms=("linux/amd64" "linux/386" "windows/386" "darwin/amd64")
package="cmd/main/main.go"
package_name="sisi"
	
for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o "bin/"$output_name $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting...'
		exit 1
	fi
done