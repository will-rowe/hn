package reporting_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	report "github.com/will-rowe/hn/api/gen/go/report/v1"
	"github.com/will-rowe/hn/backend/reporting"
)

func TestProcessReport_ValidInput(t *testing.T) {
	service := reporting.NewReportService()

	req := &report.SubmitReportRequest{
		DatasetId:     "ds1",
		DataId:        "data1",
		MediaType:     report.MediaType_MEDIA_TYPE_TEXT,
		ViolationType: report.ViolationType_VIOLATION_TYPE_PII,
		Description:   "contains name and email",
	}

	id, err := service.ProcessReport(context.Background(), req)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestProcessReport_InvalidInput(t *testing.T) {
	service := reporting.NewReportService()

	cases := []struct {
		name string
		req  *report.SubmitReportRequest
	}{
		{"NilRequest", nil},
		{"MissingDatasetID", &report.SubmitReportRequest{
			DataId: "data", Description: "desc", MediaType: report.MediaType_MEDIA_TYPE_TEXT, ViolationType: report.ViolationType_VIOLATION_TYPE_PII,
		}},
		{"MissingDataID", &report.SubmitReportRequest{
			DatasetId: "ds", Description: "desc", MediaType: report.MediaType_MEDIA_TYPE_TEXT, ViolationType: report.ViolationType_VIOLATION_TYPE_PII,
		}},
		{"MissingDescription", &report.SubmitReportRequest{
			DatasetId: "ds", DataId: "data", MediaType: report.MediaType_MEDIA_TYPE_TEXT, ViolationType: report.ViolationType_VIOLATION_TYPE_PII,
		}},
		{"MissingMediaType", &report.SubmitReportRequest{
			DatasetId: "ds", DataId: "data", Description: "desc", ViolationType: report.ViolationType_VIOLATION_TYPE_PII,
		}},
		{"MissingViolationType", &report.SubmitReportRequest{
			DatasetId: "ds", DataId: "data", Description: "desc", MediaType: report.MediaType_MEDIA_TYPE_TEXT,
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			id, err := service.ProcessReport(context.Background(), tc.req)
			assert.Error(t, err)
			assert.Empty(t, id)
		})
	}
}

/*
102007             11046 ns/op
PASS
ok      github.com/will-rowe/hn/backend/reporting       2.370s
*/
func BenchmarkProcessReport(b *testing.B) {
	svc := reporting.NewReportService()
	req := &report.SubmitReportRequest{
		DatasetId:     "ds",
		DataId:        "data",
		Description:   "desc",
		MediaType:     report.MediaType_MEDIA_TYPE_TEXT,
		ViolationType: report.ViolationType_VIOLATION_TYPE_PII,
	}

	for i := 0; i < b.N; i++ {
		_, _ = svc.ProcessReport(context.Background(), req)
	}
}
