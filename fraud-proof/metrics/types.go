package metrics

import "github.com/mantlenetworkio/mantle/metrics"

type (
	Size struct {
		metrics.Namespace
	}
	Balance struct {
		metrics.Namespace
	}
	Fee struct {
		metrics.Namespace
	}
	Index struct {
		metrics.Namespace
	}
	Alert struct {
		metrics.Namespace
	}
)

const (
	AssertionSize       = "assertion_size"
	RollupSize          = "rollup_size"
	ProposerBalance     = "proposer_balance"
	ValidatorBalance    = "validator_balance"
	ProposerRollupFee   = "proposer_rollup_fee"
	ValidatorVerifyFee  = "validator_verify_fee"
	ProposerConfirmFee  = "proposer_confirm_fee"
	BatchIndex          = "batch_index"
	AssertionIndex      = "assertion_index"
	VerifiedIndex       = "verified_index"
	AlertChallengeStart = "challenge_start"
	AlertChallengeEnd   = "challenge_end"
)

func (size *Size) LabelAssertionSize() string {
	return size.Label(AssertionSize)
}

func (size *Size) LabelRollupSize() string {
	return size.Label(RollupSize)
}

func (balance *Balance) LabelProposerBalance() string {
	return balance.Label(ProposerBalance)
}

func (balance *Balance) LabelValidatorBalance() string {
	return balance.Label(ValidatorBalance)
}

func (fee *Fee) LabelProposerRollupFee() string {
	return fee.Label(ProposerRollupFee)
}

func (fee *Fee) LabelValidatorVerifyFee() string {
	return fee.Label(ValidatorVerifyFee)
}

func (fee *Fee) LabelProposerConfirmFee() string {
	return fee.Label(ProposerConfirmFee)
}

func (index *Index) LabelBatchIndex() string {
	return index.Label(BatchIndex)
}

func (index *Index) LabelAssertionIndex() string {
	return index.Label(AssertionIndex)
}

func (index *Index) LabelVerifiedIndex() string {
	return index.Label(VerifiedIndex)
}

func (alert *Alert) LabelAlertChallengeStart() string {
	return alert.Label(AlertChallengeStart)
}

func (alert *Alert) LabelAlertChallengeEnd() string {
	return alert.Label(AlertChallengeEnd)
}

var (
	NameSize    = new(Size)
	NameBalance = new(Balance)
	NameFee     = new(Fee)
	NameIndex   = new(Index)
	NameAlert   = new(Alert)
)

func initName() {
	NameSize.Init("Size")
	NameBalance.Init("Balance")
	NameFee.Init("Fee")
	NameIndex.Init("Index")
	NameAlert.Init("Alert")
}
