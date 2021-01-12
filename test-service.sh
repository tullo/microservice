#!/bin/bash
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

docker-compose -p microservice -f docker/docker-compose-migrations.yml \
  exec db mysql -u root stats -e 'select * from incoming'
