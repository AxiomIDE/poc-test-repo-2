package nodes

import (
	"context"
	"fmt"

	"axiom-analytics/axiom"
	gen "axiom-analytics/gen"
	axiomtextops "axiom-analytics/gen/imports/axiom-text-ops/0.1.0"
)

// Summarize emits one AnalysisReport per token in the input stream, providing
// an incremental per-token summary — demonstrating a pipeline node.
func Summarize(ctx context.Context, log axiom.Logger, secrets axiom.Secrets, in <-chan *axiomtextops.TokensResult, emit func(*gen.AnalysisReport) error) error {
	for input := range in {
		for _, token := range input.GetTokens() {
			if err := emit(&gen.AnalysisReport{
				Summary:   fmt.Sprintf("Token: %s", token),
				WordCount: 1,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}
