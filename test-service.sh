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

payload='{
  "property": "news",
  "section": 1,
  "id": 1
}'

curl -s -H 'Content-Type: application/json' $url -d "$payload" | jq .

curl -s -H 'Content-Type: application/json' -H "X-Forwarded-For: 8.8.8.8, 127.0.0.1" $url -d "$payload" | jq .

curl -s -H 'Content-Type: application/json' -H "X-Real-IP: 9.9.9.9" $url -d "$payload" | jq .

docker-compose -p microservice \
  exec db mysql -u root stats -e 'SELECT * FROM incoming ORDER BY id DESC LIMIT 3'


echo
echo "============================================================================+"
echo "=== Haberdasher Service                                                     |"
echo "============================================================================+"
echo

function get_hd_ip {
	docker inspect \
		microservice_hd_1 | \
	jq -r '.[0].NetworkSettings.Networks["microservice_default"]'.IPAddress
}

hd_url="http://$(get_hd_ip):3000/twirp/haberdasher.HaberdasherService/MakeHat"
curl -s -H 'Content-Type: application/json' $hd_url -d '{"centimeters": 61}' | jq .

docker-compose -p microservice \
  exec db mysql -u root haberdasher -e 'SELECT * FROM hat ORDER BY id DESC LIMIT 5'
