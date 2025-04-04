# HN - Reporting Service

A service to allow licensees to report individual pieces of data (text, image, audio, video, animation) that may be in violation of laws or regulations (e.g. contains PII). The service accepts structured violation reports and returns a report ID for tracking and audit purposes.

## Assumptions

- The service runs as a single instance (no load balancing or distributed coordination needed).
- No persistence is required — reports are processed and acknowledged, but not stored.
- The backend is implemented in Go, using gRPC and gRPC-Gateway to expose both REST and gRPC interfaces.
- A valid Bearer token is required to access the endpoint (currently a static token for development/testing).
- Protobuf definitions are managed via [Buf](https://buf.build) and used to generate server and client code.
- Mock interfaces are generated using [Mockery](https://github.com/vektra/mockery).

## Installing

We use [mise](https://mise.jdx.dev/installing-mise.html) to manage the dependencies and tasks for running this project.

1. Install mise

```sh
curl https://mise.run | sh
```

2. Clone this repo, trust the mise config and install the dependencies

```sh
git clone https://github.com/will-rowe/hn && cd hn
mise trust .
mise install
```

3. Generate the required files

```sh
mise tasks run generate
```

4. Run the linting and tests:

```sh
mise tasks run test
```

5. Start the server with Docker Compose (with an initial run of the e2e):

```sh
mise tasks run start
```

6. Run the full end-to-end test (starts the server, runs the e2e test, then exits):

```sh
mise tasks run e2e
```

## API

The API exposes a single endpoint:

```http
POST /v1/reports
Authorization: Bearer <token>
Content-Type: application/json
```

### Request

```json
{
  "dataset_id": "abc123",
  "data_id": "def456",
  "media_type": "MEDIA_TYPE_TEXT",
  "violation_type": "VIOLATION_TYPE_PII",
  "description": "This contains personally identifiable information."
}
```

### Response

```json
{
  "report_id": "report-uuid",
  "status": "RECEIVED"
}
```

The API is defined via Protobuf (`report.proto`) and includes both REST and gRPC bindings.

## Testing and performance

### Unit tests

- All handler and service logic is unit tested using [testify](https://github.com/stretchr/testify).
- The `ReportServiceInterface` is mocked using [Mockery](https://github.com/vektra/mockery).
- Tests cover validation, correct response behavior, and error cases.
- Run with:

```sh
mise tasks run test
```

### Latency Benchmarks

Use Go’s built-in benchmarking tools to measure performance of critical functions like `ProcessReport()`:

```go
func BenchmarkProcessReport(b *testing.B) {
    svc := NewReportService()
    req := &report.ReportRequest{
        DatasetId: "ds1", DataId: "data1", Description: "Contains PII",
        MediaType: report.MediaType_TEXT, ViolationType: report.ViolationType_PII,
    }

    for i := 0; i < b.N; i++ {
        _, _ = svc.ProcessReport(context.Background(), req)
    }
}
```

Run benchmarks with:

```sh
go test -bench=. ./backend/reporting
```

You can use `pprof` for deeper CPU and memory profiling.

### End to end

You can run the server locally and send a report using `curl`:

```sh
go run main.go
```

Then, in another terminal:

```sh
curl -X POST http://localhost:8080/v1/reports \
  -H "Authorization: Bearer testtoken" \
  -H "Content-Type: application/json" \
  -d '{
        "dataset_id": "abc123",
        "data_id": "def456",
        "media_type": "MEDIA_TYPE_TEXT",
        "violation_type": "VIOLATION_TYPE_PII",
        "description": "Contains a full name and email address"
      }'
```

A valid JSON response with `report_id` and `status` should be returned.
