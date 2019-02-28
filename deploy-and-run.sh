#!/bin/bash

RASPBERRY_PI=${1}

main() {
	if [ -z ${1} ] ; then
		usage
	fi
	utility=${HOME}/go/bin/linux_arm/rpiinfo
	GOOS=linux GOARCH=arm go install
	if [ $? -eq 0 ] ; then
		scp ${utility} ${RASPBERRY_PI}:bin
		ssh ${RASPBERRY_PI} /home/pi/bin/rpiinfo
	fi
}

usage() {
	echo "Usage: $(basename ${0}) <pi>"
	exit 23
}

main $@
