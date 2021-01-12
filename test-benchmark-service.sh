#!/bin/bash
function get_stats_ip {
	docker inspect \
		microservice_stats_1 | \
	jq -r '.[0].NetworkSettings.Networks["microservice_default"]'.IPAddress
}

url="http://$(get_stats_ip):3000/twirp/stats.StatsService/Push"
image="skandyla/wrk:latest"

docker run --rm --net=host -it -v $PWD:/data $image -d60s -t4 -c100 -s test.lua $url

#docker run --rm --net=host -it -v $PWD:/data $image -d15s -t4 -c400 -s test.lua $url
