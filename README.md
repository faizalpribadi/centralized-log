# centralized-log
Centralized Log of Application With ELK Stack ( Structured Logging )

## Requirements

* Elasticsearch 
* Kibana
* Logstash
* Filebeat
* Application with log files , for example : `app.log`


### Steps Action

1. `./filebeat -e -c filebeat/filebeat.yml -d "publish"`

2. `./logstash -f logstash.conf --config.reload.automatic`

3. `go run main.go`


run this command : 
```javascript
curl --location --request POST 'http://localhost:3000/create' \
--header 'Content-Type: application/json' \
--data-raw '{
	"id": "1",
	"name": "iCal",
	"data": "Hello ELK"
}'
```


### Screenshots

1. discover index on kibana ( logstash-application index name )
![kibana-discover-index](https://github.com/faizalpribadi/centralized-log/blob/master/images/kibana-01.png)

2. discover index data 
![kibana-discover-index](https://github.com/faizalpribadi/centralized-log/blob/master/images/kibana-03.png)

3. show the data of logs ( structured logging )
![kibana-discover-index](https://github.com/faizalpribadi/centralized-log/blob/master/images/kibana-04.png)
