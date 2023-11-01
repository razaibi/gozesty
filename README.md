# GoZesty

This is a Golang based API testing framework.

# Setting Up Tests

In order to setup tests, navigate to **zest_test** folder.

Here you should see the following:

```
- calls
-- sample.yaml
- payloads
-- sample.json
```

## Setting Up Calls

Open up one of the sample API calls like **sample.yaml**. Add additional endpoints as needed.

It will look like this:

```yaml
endpoints:
  - name: "Sample"
    endpoint_url: "https://jsonip.com/"
    request_method: "GET"
    request_headers:
      Content: "en-US"
    request_body: ""
    expected_conditions:
      - expected_status: 200
        expected_response: "sample.json"
```

## Setting Up Payloads

In the payloads folder, you will find the **sample.json** file.

```json
{
    "ip":"XXXX:XXXX:XXXX:XXXX:XXXX:XXXX:XXXX:XXXX",
    "geo-ip":"https://getjsonip.com/#plus",
    "API Help":"https://getjsonip.com/#docs"
}
```

Multiple payloads can be added to this folder.


Run tests using:

```go
go test -v ./...
```