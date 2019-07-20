# GELFbeat
Have you ever needed to take GELF-encoded (gzipped and/or chunked) UDP input and send it to Graylog? Logstash? Elasticsearch? Something else that takes a Beats output? GELFbeat is for you.

## Acknowledgements
Thanks to:
* Tim Tkachenko for [gelf-go](https://github.com/timtkachenko/gelf-go) - this project wouldnt have come together without his 
* Elastic for [libbeat](https://github.com/elastic/beats/tree/master/libbeat)

## Support
This is a tool we're using internally, but it's not an offically supported item. We're happy to take pull requests, but the main reason we're releasing this is that it might help other folks in the same pickle. Support on this tool is best-effort; we make no guarantees or warranties that this will work in your environment.

## Using GELFbeat
Beacuse this uses `libbeat` 7.2, many of the same processors, outputs, etc. apply here - GELFbeat only takes an IP address to listen on, and a port for configuration. You'll have to look at the [Filebeat documentation](https://www.elastic.co/guide/en/beats/filebeat/7.2/index.html) to figure out what works for other parts of the configuration file, specifically processors and outputs.

Here's an example `gelfbeat.yml`:

```
---
fields_under_root: true
gelfbeat:
  listen: 127.0.0.1
  port: 10000
processors:
 - decode_json_fields:
     fields: ['message']
     target: json
     overwrite_keys: true
     target: ""
fields:
  environment: prod
output.logstash:
  hosts:
  - graylog.tls.zone:25044
  ssl:
    enabled: true
    supported_protocols:
    - TLSv1.2
```

Then, run the tool with `gelfbeat -c gelfbeat.yml -e` and send GELF-encoded messages to `127.0.0.1:10000`.

## A systemd unit file
This assumes all of the paths below exist. Toss this in `/etc/systemd/system/gelfbeat.service`:

```
[Unit]
Description=Gelfbeat sends log files to Logstash or directly to Elasticsearch.
Documentation=https://github.com/threatstack/gelfbeat
Wants=network-online.target
After=network-online.target

[Service]
Environment="BEAT_LOG_OPTS=-e"
Environment="BEAT_CONFIG_OPTS=-c /etc/gelfbeat/gelfbeat.yml"
Environment="BEAT_PATH_OPTS=-path.home /usr/share/gelfbeat -path.config /etc/gelfbeat -path.data /var/lib/gelfbeat -path.logs /var/log/gelfbeat"
ExecStart=/usr/share/gelfbeat/bin/gelfbeat $BEAT_LOG_OPTS $BEAT_CONFIG_OPTS $BEAT_PATH_OPTS
Restart=always

[Install]
WantedBy=multi-user.target
```
