package nodes_test

import (
	"context"
	"testing"

	"axiom-analytics/axiom"
	axiomtextops "axiom-analytics/gen/imports/axiom-text-ops/0.1.0"
	"axiom-analytics/nodes"
)

type testLogger struct{ t *testing.T }

func (l *testLogger) Debug(msg string, args ...any) { l.t.Logf("DEBUG  %s %v", msg, args) }
func (l *testLogger) Info(msg string, args ...any)  { l.t.Logf("INFO   %s %v", msg, args) }
func (l *testLogger) Warn(msg string, args ...any)  { l.t.Logf("WARN   %s %v", msg, args) }
func (l *testLogger) Error(msg string, args ...any) { l.t.Logf("ERROR  %s %v", msg, args) }

var _ axiom.Logger = (*testLogger)(nil)

// TESTS — delete this block when done ─────────────────────────────────────────
// Tests are required to publish this package. The publish pipeline runs your
// tests as a quality gate — a package will not be published if tests fail or
// do not meet the minimum requirements.
//
// Requirements checked before publishing:
//   - At least one test per node
//   - All tests must pass
//   - Output fields must be meaningfully asserted — not just error-checked
//
// The generated test below is a starting point. Replace the TODO comment with
// real assertions that verify your node returns correct data for known inputs.
// Think: given a specific input, what should the output fields contain?
//
// Run your tests locally at any time:
//   axiom test

func TestAnalyze(t *testing.T) {
	ctx := context.Background()
	log := &testLogger{t}
	input := &axiomtextops.TokensResult{
		Tokens: []string{"hello", "world"},
		Count:  2,
	}

	got, err := nodes.Analyze(ctx, log, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.WordCount != 2 {
		t.Errorf("expected word count 2, got %d", got.WordCount)
	}
}
