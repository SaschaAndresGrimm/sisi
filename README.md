# Sisi
**Si**mple **SI**mplon API CLI

## Motivation
DECTRIS EIGER(2) detectors are controlled via a rest-like API ([SIMPLON API](https://media.dectris.com/210607-DECTRIS-SIMPLON-API-Manual_EIGER2-chip-based_detectros.pdf)). _Sisi_ provides a simple stand-alone CLI written in golang that can be [cross-compiled](./build.sh) and used as standalone executable on different platforms:

* [Linux](./bin/sisi-linux-386)
* [Windows](./bin/sisi-windows-386.exe)
* [OSX](./bin/sisi-darwin-amd64)

### Real motivation
I have a terribly hard time remembering CURL commands... See the examples for comparison. 

## Usage
Tell Sisi if she should _get_, _set_, or _do_ anything for you. Be friendly and say 'please'.

	Usage:
	sisi <ip> get <module> <param> <key> [-a <api> -t <timeout> please]
	sisi <ip> set <module> <param> <key> <value> [-a <api> -t <timeout> please]
	sisi <ip> do <module> <task> [-a <api> -t <timeout> please]
	sisi -h | --help

	Options:
	-a <api>      API version [default: 1.8.0]
	-t <timeout>  request timeout in seconds [default: 2]
	-h --help     Show this help screen.
	please        just being friendly, optionally

### Examples
**enable the filewriter interface**
```
sisi 169.254.254.1 set filewriter config mode enabled
```
result
```
PUT http://169.254.254.1/filewriter/api/1.8.0/config/mode {"value":"enabled"} 200 OK 302.285724ms
```

the same in CURL
```
curl -X PUT -H "Content-Type: application/json" -d "{\"value\":\"enabled\"}" http://169.254.254.1/filewriter/api/1.6.0/config/mode
```

**set the number of images to acquire**
```
sisi 169.254.254.1 set detector config nimages 5
```
result
```
PUT http://169.254.254.1/detector/api/1.8.0/config/nimages {"value":5} 200 OK 317.301363ms
  "nimages"
```

the same in CURL
```
curl -X PUT -H "Content-Type: application/json" -d "{\"value\":5}"
http://169.254.254.1/detector/api/1.6.0/config/count_time
```

**arm detector**
```
sisi 169.254.254.1 do detector arm
```
result
```
PUT http://169.254.254.1/detector/api/1.8.0/command/arm  200 OK 669.908552ms
  "sequence id":14
```
the same in CURL
```
curl -X PUT http://169.254.254.1/detector/api/1.6.0/command/arm
```