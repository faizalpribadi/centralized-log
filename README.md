# centralized-log
Centralized Log of Application With ELK Stack ( Structured Logging )

## Requirements

* Elasticsearch - `7.7.0 version - already installed`
* Kibana - `7.7.0 version - already installed`
* Logstash - `7.7.0 version - already installed`
* Filebeat - `7.7.0 version - already installed`
* Application with log files , for example : `app.log`


### Steps Action

1. `./filebeat -e -c filebeat/filebeat.yml -d "publish"`

2. `./logstash -f logstash/logstash.conf --config.reload.automatic`

3. `go run main.go`


run this command : 
```javascript
curl --location --request POST 'http://localhost:3000/create' \
--header 'Content-Type: application/json' \
--data-raw '{
	"id": "1",
	"name": "ical",
	"data": {
		"enabled": false
	}
}'
```


### Screenshots

1. The logs on kibana dashboard
![logs](https://github.com/faizalpribadi/centralized-log/blob/master/images/kibana-logs.png)
