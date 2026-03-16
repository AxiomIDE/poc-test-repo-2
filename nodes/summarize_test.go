package nodes

import (
	"testing"

	gen "axiom-analytics/gen"
	axiomtextops "axiom-analytics/gen/imports/christian-axiom-text-ops/0.1.0"
)

func TestSummarize(t *testing.T) {
	var results []*gen.AnalysisReport
	in := make(chan *axiomtextops.TokensResult, 1)
	in <- &axiomtextops.TokensResult{Tokens: []string{"foo", "bar"}, Count: 2}
	close(in)

	err := Summarize(nil, nil, nil, in, func(r *gen.AnalysisReport) error {
		results = append(results, r)
		return nil
	})
	if err != nil {
		t.Fatalf("Summarize returned error: %v", err)
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 frames, got %d", len(results))
	}
	if results[0].Summary != "Token: foo" {
		t.Errorf("unexpected first summary: %q", results[0].Summary)
	}
}
