package nodes_test

import (
	"context"
	"testing"

	"axiom-analytics/axiom"
	axiomtextops "axiom-analytics/gen/imports/axiom-official-axiom-text-ops/0.1.0"
	"axiom-analytics/nodes"
)

// testContext is a testing.T-backed axiom.Context for unit tests.
type testContext struct {
	t          *testing.T
	secretsMap map[string]string
}

func newTestContext(t *testing.T) *testContext {
	return &testContext{t: t, secretsMap: map[string]string{}}
}

type testLogger struct{ t *testing.T }

func (l *testLogger) Debug(msg string, args ...any) { l.t.Logf("DEBUG  %s %v", msg, args) }
func (l *testLogger) Info(msg string, args ...any)  { l.t.Logf("INFO   %s %v", msg, args) }
func (l *testLogger) Warn(msg string, args ...any)  { l.t.Logf("WARN   %s %v", msg, args) }
func (l *testLogger) Error(msg string, args ...any) { l.t.Logf("ERROR  %s %v", msg, args) }

type testSecretMap map[string]string

func (s testSecretMap) Get(name string) (string, bool) { v, ok := s[name]; return v, ok }

func (c *testContext) Log() axiom.Logger      { return &testLogger{c.t} }
func (c *testContext) Secrets() axiom.Secrets { return testSecretMap(c.secretsMap) }
func (c *testContext) Agent() axiom.Agent     { return nil }
func (c *testContext) ExecutionID() string    { return "test-execution-id" }
func (c *testContext) FlowID() string         { return "test-flow-id" }
func (c *testContext) TenantID() string       { return "test-tenant-id" }

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
	ax := newTestContext(t)
	input := &axiomtextops.TokensResult{
		Tokens: []string{"hello", "world"},
		Count:  2,
	}

	got, err := nodes.Analyze(ctx, ax, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.WordCount != 2 {
		t.Errorf("expected word count 2, got %d", got.WordCount)
	}
}
