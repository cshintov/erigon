package jsonrpc

import (
    "encoding/json"
    "testing"

    "github.com/erigontech/erigon-lib/common/hexutil"
)

func TestTraceResult_MarshalJSON_DefaultGasUsedZero(t *testing.T) {
    tr := TraceResult{ // GasUsed is nil on purpose
        Output: hexutil.Bytes{},
    }
    b, err := json.Marshal(tr)
    if err != nil {
        t.Fatalf("marshal error: %v", err)
    }
    // Expect gasUsed to be a hex string "0x0" not null
    got := string(b)
    wantFragment := `"gasUsed":"0x0"`
    if !contains(got, wantFragment) {
        t.Fatalf("unexpected json: %s (want fragment %s)", got, wantFragment)
    }
}

func TestCreateTraceResult_MarshalJSON_DefaultGasUsedZero(t *testing.T) {
    cr := CreateTraceResult{ // GasUsed is nil on purpose
        Code: hexutil.Bytes{},
    }
    b, err := json.Marshal(cr)
    if err != nil {
        t.Fatalf("marshal error: %v", err)
    }
    got := string(b)
    wantFragment := `"gasUsed":"0x0"`
    if !contains(got, wantFragment) {
        t.Fatalf("unexpected json: %s (want fragment %s)", got, wantFragment)
    }
}

// contains is a minimal substring check to avoid importing extra deps
func contains(s, sub string) bool {
    return len(s) >= len(sub) && (s == sub || indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
    for i := 0; i+len(sub) <= len(s); i++ {
        if s[i:i+len(sub)] == sub {
            return i
        }
    }
    return -1
}

