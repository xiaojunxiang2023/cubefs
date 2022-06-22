#!/usr/bin/env bash

BranchName=`git rev-parse --abbrev-ref HEAD`
CommitID=`git rev-parse HEAD`
BuildTime=`date +%Y-%m-%d\ %H:%M`
Debug="0"

[[ "-$GOPATH" == "-" ]] && { echo "GOPATH not set"; exit 1 ; }

goflag="-s"

build_sdk=1
build_client=1

help() {
    cat <<EOF

Usage: ./build.sh [ -h | --help ] [ -g ] [ --sdk-only | --client-only ]
    -h, --help              show help info
    -g                      setup Debug="1" goflag="" gccflag="-g"
    --sdk-only              build sdk (libcfssdk.so libempty.so) only
    --client-only           build client (libcfsclient.so and cfs-client) only
EOF
    exit 0
}

ARGS=( "$@" )
for opt in ${ARGS[*]} ; do
    case "$opt" in
        -h|--help)
            help
            ;;
        -g)
            Debug="1"
            goflag=""
            gccflag="-g"
            ;;
    	--sdk-only)
    	    build_sdk=1
    	    build_client=0
	        ;;
	    --client-only)
	        build_sdk=0
	        build_client=1
	        ;;
    esac
done

dir=$(dirname $0)
echo "using Debug=\"${Debug}\""
echo "using goflag=\"${goflag}\""
echo "using gccflag=\"${gccflag}\""
if [[ ${build_sdk} -eq 1 ]]; then
    echo "building sdk (libcfssdk_${CommitID}.so libempty.so) ..."
    go build -ldflags "${goflag} -E main.main -X main.BranchName=${BranchName} -X main.CommitID=${CommitID} -X 'main.BuildTime=${BuildTime}' -X 'main.Debug=${Debug}'" -buildmode=plugin -linkshared -o ${dir}/libcfssdk_${CommitID}.so ${dir}/sdk_fuse.go ${dir}/sdk_bypass.go ${dir}/http.go ${dir}/ump.go
    go build -buildmode=plugin -linkshared -o ${dir}/libempty.so  ${dir}/empty.go
fi
if [[ ${build_client} -eq 1 ]]; then
    echo "building client (cfs-client libcfsclient.so) ..."
    go build -ldflags "${goflag}" -linkshared -o ${dir}/cfs-client ${dir}/main_fuse.go
    gcc -std=c99 ${gccflag} -fPIC -shared -o ${dir}/libcfsclient.so ${dir}/main_bypass.c ${dir}/bypass/ini.c -ldl -lpthread
fi
