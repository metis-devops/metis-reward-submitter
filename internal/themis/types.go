package themis

type MpcAddrType uint64

const (
	CommonMpcType       MpcAddrType = iota
	StateSubmitMpcType              // Special type for state submit
	RewardSubmitMpcType             // Special type for reward submit
	BlobSubmitMpcType
)

type MpcSignType int

const (
	BatchSubmitSignType MpcSignType = iota
	BatchRewardSignType
	CommitEpochToMetisSignType
	ReCommitEpochToMetisSignType
	L1UpdateMpcAddressSignType
	L2UpdateMpcAddressSignType
)
