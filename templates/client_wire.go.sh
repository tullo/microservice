#!/bin/bash
cd $(dirname $(dirname $(readlink -f $0)))

## list all services
schemas=$(ls -d rpc/* | xargs -n1 basename)

function render_wire {
	echo "package client"
	echo
	echo "import ("
	echo -e "\t\"github.com/google/wire\""
	for schema in $schemas; do
		echo -e "\t\"${MODULE}/client/${schema}\""
	done
	echo ")"
	echo
	echo "// Inject produces a wire.ProviderSet with our RPC clients."
	echo "var Inject = wire.NewSet("
	for schema in $schemas; do
		echo -e "\t${schema}.New,"
	done
	echo ")"
}

render_wire > client/wire.go
echo "~ client/wire.go"
