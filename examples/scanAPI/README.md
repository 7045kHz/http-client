# Go Example Benchmarking Client v0.1.1
 
This example creates a client that will take JSON as a configuration source, and output critical connection details to summary and body output files.  Information is timestamped at each stage so you can see response times from a client perspective.

## About Original Source Location
> See main page for information on the client package. This section is unique to 7045kHz: 
 
## About TLS Settings
Client is hard coded to Ignore Verify (for now). Test HTTP Server reports TLS version value. Table for decode is as follows:


|TLS Version|Value Displayed in output body file|
|:--------------:|:---------------:|
|1.0|769|
|1.1|770|
|1.2|771|
|1.3|772|

## Client Build Installation

```
go build benchmarking.go
```
## Build Installation for Optional HTTP Server.
>This server does nothing more than report back details about the client connecting to it. It's used to test how a HTTP server would process your clients connection, including payload.

```
cd server
go build server.go
```
## Client Setup
### NTLMv2 or Standard
In the getHttpClient() function change ``SetNtlm(false)`` to ``SetNtlm(true)`` to use NTLMv2.

### Primary configuration file
Create a ``config.json`` file to define the core test name, and pointer to the file that contains a list of endpoints to run agains.

```
cat config.json
{
    "Name":"Github Web and API",
    "Description":"Testing API and Web Client Side Performance",
    "EndpointFile":"apibench.json"
}
```
### Endpoint configuration file
Create a json file, named as you like. In this example use  ``apibench.json``.  to define the core test name, and pointer to the file that contains a list of endpoints to run agains.

Each endpoint gets it's own configuration, from Name, Description, Url, Method, Optional PayloadFile, and Headers.  The output from the connection will be sent to SummaryOutputFile, and BodyOutputFile.  

**About values and data saved**

1. Set the URL clear when simple, URL Encoded for passing options.  Example: https://localhost:10443/api?name="momo" should be defined as  https%3A%2F%2Flocalhost%3A10443%2Fapi%3Fname%3Dmomo 

2.  each run the SummaryOutputFile, and BodyOutputFile files will be overwritten. If you want to keep them, please copy to an alternate location before re-running.

```
{  
	{
		"Name":"Benchmark: web github",
		"Description":"Testing Webside Client Side Performance",
		"Url":"https://www.github.com",
		"Method":"GET",
		"PayloadFile":"",
		"SummaryOutputFile":"apibench-web-summary.out",
		"BodyOutputFile":"apibench-web-body.out",
		"Headers":[
			{"Header_Key":"Content-Type","Header_Value":"text/html;charset=utf-8"}

		]
    }
}
```
#### Example of multiple API/Client connections
```
cat config.json
{
   "Events":[ 

    {
    "Name":"Benchmark: api github",
    "Description":"Testing API Client Side Performance",
    "Url":"https://api.github.com",
    "Method":"GET",
    "PayloadFile":"",
    "SummaryOutputFile":"apibench-api-summary.out",
    "BodyOutputFile":"apibench-api-body.out",
    "Headers":[{"Header_Key":"Content-Type","Header_Value":"text/html;charset=utf-8"}]
    },
    {
    "Name":"Benchmark: web github",
    "Description":"Testing Webside Client Side Performance",
    "Url":"https://www.github.com",
    "Method":"GET",
    "PayloadFile":"",
    "SummaryOutputFile":"apibench-web-summary.out",
    "BodyOutputFile":"apibench-web-body.out",
    "Headers":[
        {"Header_Key":"Content-Type","Header_Value":"text/html;charset=utf-8"}

    ]
    },    {
        "Name":"TLS Localhost",
        "Description":"Testing API Client Side Performance",
        "Url":"https%3A%2F%2Flocalhost%3A10443%2Fapi%3Fname%3Dmomo",
        "Method":"POST",
        "PayloadFile":".\\localhost_post.json",
        "SummaryOutputFile":"local-api-summary.out",
        "BodyOutputFile":"local-api-body.out",
        "Headers":[{"Header_Key":"Content-Type","Header_Value":"application/json;charset=utf-8"}]
    }
]
 
}
```
## Example Output
### Summary File
>Note in the summary file the key timestamps to pay attention to are those with EVENT in them. These timestamps are taken directly before and after the client connection is made.

```
2021-07-31 23:59:42.0373167 +0000 UTC Testing Event EndPoint: Localhost
=================================================================================================
Name:             Localhost
Description:      Testing API Client Side Performance
URL:              http://localhost:9090
Method:           POST
Summary File:     local-api-summary.out
Body Output File: local-api-body.out
Payload File: .\localhost_post.json
Custom Headers: [1]
=================================================================================================
2021-07-31 23:59:42.0396102 +0000 UTC EVENT [Localhost] Client Sending Url: http://localhost:9090
2021-07-31 23:59:42.0537697 +0000 UTC EVENT [Localhost] Data Returned from Url: http://localhost:9090
2021-07-31 23:59:42.0537697 +0000 UTC Status Code: 200
2021-07-31 23:59:42.0537697 +0000 UTC Status: 200 OK
2021-07-31 23:59:42.0537697 +0000 UTC Creating and logging body content in local-api-body.out

```
### Body File
>Note, one the example test server under the server directory, the body output is minimum. On a real server such as api.github.com, or www.github.com the output can be significant.  The body file will contain a complete dump of the http.Body content.

```
2021-07-31 23:59:42.0564115 +0000 UTC Event Test: Localhost
2021-07-31 23:59:42.0579831 +0000 UTC TARGET URL: http://localhost:9090
2021-07-31 23:59:42.0579831 +0000 UTC RESPONSE Status Code: 200
2021-07-31 23:59:42.0579831 +0000 UTC RESPONSE Headers: map[Content-Length:[362] Content-Type:[text/plain; charset=utf-8] Date:[Sat, 31 Jul 2021 23:59:42 GMT]]
2021-07-31 23:59:42.0579831 +0000 UTC RESPONSE Status: 200 OK
2021-07-31 23:59:42.0579831 +0000 UTC RESPONSE Body: URL: /
Request at 2021-07-31 19:59:42.0500991 -0400 EDT m=+835.767634001
Method: POST
Content Length: 32
Body: "{\"Name\":\"SOMETHING BASIC\"}"
TLS: <nil>
Host: localhost:9090
URL User: 
Remove Address: [::1]:62011
Request URI: /
Response: <nil>
Content-Type:[application/json;charset=utf-8]
Accept-Encoding:[gzip]
User-Agent:[Test-Computer]
Content-Length:[32]

2021-07-31 23:59:42.0579831 +0000 UTC Closing Test: Localhost

```