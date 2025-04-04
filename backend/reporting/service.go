package reporting

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	report "github.com/will-rowe/hn/api/gen/go/report/v1"
)

type ReportServiceInterface interface {
	ProcessReport(ctx context.Context, req *report.SubmitReportRequest) (string, error)
}

type reportService struct{}

func NewReportService() ReportServiceInterface {
	return &reportService{}
}

func (s *reportService) ProcessReport(ctx context.Context, req *report.SubmitReportRequest) (string, error) {
	// Step 1: Validate request
	if err := validateReportRequest(req); err != nil {
		return "", fmt.Errorf("invalid report request: %w", err)
	}

	// Step 2: Log the structured report
	log.Printf("Report received — DatasetID: %s, DataID: %s, Violation: %s, MediaType: %s",
		req.GetDatasetId(), req.GetDataId(), req.GetViolationType().String(), req.GetMediaType().String())

	// Step 3: Optionally enrich or notify here
	// - Enrichment based on dataset_id
	// - Asynchronous queue / event notification
	// - Store in DB or trigger alerting pipeline

	// Step 4: Return a generated UUID as the report ID
	reportID := uuid.New().String()
	return reportID, nil
}

func validateReportRequest(req *report.SubmitReportRequest) error {
	if req == nil {
		return errors.New("request is nil")
	}
	if req.GetDatasetId() == "" {
		return errors.New("dataset_id is required")
	}
	if req.GetDataId() == "" {
		return errors.New("data_id is required")
	}
	if req.GetDescription() == "" {
		return errors.New("description is required")
	}
	if req.MediaType == report.MediaType_MEDIA_TYPE_UNSPECIFIED {
		return errors.New("media_type is unspecified")
	}
	if req.ViolationType == report.ViolationType_VIOLATION_TYPE_UNSPECIFIED {
		return errors.New("violation_type is unspecified")
	}
	return nil
}
