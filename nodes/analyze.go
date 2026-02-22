package nodes

import (
	"context"
	"fmt"
	"strings"

	gen "axiom-analytics/gen"
	axiomtextops "axiom-analytics/gen/imports/axiom-text-ops/0.1.0"
)

// Analyze accepts tokenized text and produces a summary report with the
// token count and a human-readable listing of the tokens.
func Analyze(ctx context.Context, input *axiomtextops.TokensResult) (*gen.AnalysisReport, error) {
	return &gen.AnalysisReport{
		Summary:   fmt.Sprintf("Processed %d tokens: %s", input.GetCount(), strings.Join(input.GetTokens(), ", ")),
		WordCount: input.GetCount(),
	}, nil
}
