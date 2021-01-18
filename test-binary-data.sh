#!/bin/bash
echo
echo "============================================================================+"
echo "=== Stats Service                                                           |"
echo "============================================================================+"
echo

function get_stats_ip {
	docker inspect \
		microservice_stats_1 | jq -r '.[0].NetworkSettings.Networks["microservice_default"]'.IPAddress
}

url="http://$(get_stats_ip):3000/twirp/tullo.microservice.stats.StatsService/Push"
echo "~ url: $url"
payload='
property: "news",
section: 1,
id: 1
'
echo $payload | protoc --deterministic_output --encode tullo.microservice.stats.PushRequest rpc/stats/stats.proto \
	| curl -s -H 'Content-Type: application/protobuf' --data-binary @- $url \
	| protoc --decode tullo.microservice.stats.PushResponse rpc/stats/stats.proto

echo
echo "============================================================================+"
echo "=== Haberdasher Service                                                     |"
echo "============================================================================+"
echo

function get_hd_ip {
	docker inspect \
		microservice_hd_1 | jq -r '.[0].NetworkSettings.Networks["microservice_default"]'.IPAddress
}

hd_url="http://$(get_hd_ip):3000/twirp/tullo.microservice.haberdasher.HaberdasherService/MakeHat"
echo "~ url: $hd_url"
echo '~ protoc decoded response:'
echo 'centimeters:53' \
	| protoc --deterministic_output --encode tullo.microservice.haberdasher.Size rpc/haberdasher/haberdasher.proto \
	| curl -s -H 'Content-Type: application/protobuf' --data-binary @- $hd_url \
	| protoc --decode tullo.microservice.haberdasher.Hat rpc/haberdasher/haberdasher.proto

echo
echo "============================================================================+"
echo "=== Haberdasher Service - step by step                                      |"
echo "============================================================================+"
echo

echo 'centimeters:53' | protoc --deterministic_output --encode tullo.microservice.haberdasher.Size rpc/haberdasher/haberdasher.proto > /tmp/binary.data
echo "~ url: $hd_url"
echo '~ protoc encoded request:'
hexdump -C /tmp/binary.data
curl -s -o /tmp/resp.data -H 'Content-Type: application/protobuf' --data-binary @/tmp/binary.data $hd_url
echo
echo '~ protoc decoded response:'
hexdump -C /tmp/resp.data
rm /tmp/resp.data /tmp/binary.data
