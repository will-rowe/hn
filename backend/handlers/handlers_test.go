package handlers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	report "github.com/will-rowe/hn/api/gen/go/report/v1"
	"github.com/will-rowe/hn/backend/reporting/mocks"
)

func TestSubmitReport(t *testing.T) {
	mockSvc := new(mocks.MockReportServiceInterface)
	handler := &ReportHandler{Svc: mockSvc}

	req := &report.SubmitReportRequest{
		DatasetId:     "dataset123",
		DataId:        "data456",
		MediaType:     report.MediaType_MEDIA_TYPE_TEXT,
		ViolationType: report.ViolationType_VIOLATION_TYPE_PII,
		Description:   "Contains email address",
	}

	mockSvc.On("ProcessReport", mock.Anything, req).Return("test-report-id", nil)

	resp, err := handler.SubmitReport(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "test-report-id", resp.ReportId)
	assert.Equal(t, "RECEIVED", resp.Status)

	mockSvc.AssertExpectations(t)
}
