# Test task

[Specification](task.md)

## Tests

```go test ./...```

## Description

package `types` - types for all requests and responses with marshalling tags; marshal, unmarshal and format change tests

package `mock` - mock server with predefined responses passed as constructor argument

package `test` - some testing utils

package `service` - customizable client, service for polling all defined APIs, some test cases