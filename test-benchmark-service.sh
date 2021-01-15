#!/bin/bash
echo
echo "============================================================================+"
echo "=== Stats Service                                                           |"
echo "============================================================================+"
echo

function get_stats_ip {
	docker inspect \
		microservice_stats_1 | \
	jq -r '.[0].NetworkSettings.Networks["microservice_default"]'.IPAddress
}

url="http://$(get_stats_ip):3000/twirp/stats.StatsService/Push"
image="skandyla/wrk:latest"

docker run --rm --net=host -it -v $PWD:/data $image -d60s -t4 -c100 -s test.lua $url
#docker run --rm --net=host -it -v $PWD:/data $image -d15s -t4 -c400 -s test.lua $url

echo
echo "============================================================================+"
echo "=== Haberdasher Service                                                     |"
echo "============================================================================+"
echo

function get_hd_ip {
	docker inspect \
		microservice_hd_1 | jq -r '.[0].NetworkSettings.Networks["microservice_default"]'.IPAddress
}

url="http://$(get_hd_ip):3000/twirp/tullo.microservice.haberdasher.HaberdasherService/MakeHat"
$(go env GOPATH)/bin/hey \
	-h2 \
	-T 'application/json' \
	-d '{"centimeters": 59}' \
	-m POST \
	-z 15s \
	$url
