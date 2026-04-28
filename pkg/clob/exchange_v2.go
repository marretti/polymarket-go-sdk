package clob

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Mainnet Polygon — CTF Exchange v2 (see Polymarket contracts / v2 migration).
const (
	exchangeV2Address        = "0xE111180000d2663C0091e4f400237545B87B996B"
	exchangeV2NegRiskAddress = "0xe2222d279d744050d28e00520010520000310F59"
)

func verifyingContractV2(negRisk bool) string {
	if negRisk {
		return exchangeV2NegRiskAddress
	}
	return exchangeV2Address
}

func parseBuilderCodeString(s string) (common.Hash, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return common.Hash{}, nil
	}
	if !strings.HasPrefix(s, "0x") {
		s = "0x" + s
	}
	b, err := hexutil.Decode(s)
	if err != nil {
		return common.Hash{}, fmt.Errorf("builder code: %w", err)
	}
	if len(b) != 32 {
		return common.Hash{}, fmt.Errorf("builder code: want 32 bytes, got %d", len(b))
	}
	var h common.Hash
	copy(h[:], b)
	return h, nil
}
