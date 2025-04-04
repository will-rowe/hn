package handlers

import (
	"context"
	"log"

	report "github.com/will-rowe/hn/api/gen/go/report/v1"
	"github.com/will-rowe/hn/backend/reporting"
)

type ReportHandler struct {
	report.UnimplementedReportServiceServer
	Svc reporting.ReportServiceInterface
}

func NewReportHandler(svc reporting.ReportServiceInterface) report.ReportServiceServer {
	return &ReportHandler{Svc: svc}
}

func (h *ReportHandler) SubmitReport(ctx context.Context, req *report.SubmitReportRequest) (*report.SubmitReportResponse, error) {
	log.Printf("Received report: dataset=%s, data=%s, type=%s", req.DatasetId, req.DataId, req.ViolationType)

	reportID, err := h.Svc.ProcessReport(ctx, req)
	if err != nil {
		return nil, err
	}

	return &report.SubmitReportResponse{
		ReportId: reportID,
		Status:   "RECEIVED",
	}, nil
}
