package main

import (
	"fmt"
	"time"
)

// Constants for Miscellaneous parameters
const (
	MaxCommitteesPerSlot      = 64
	TargetCommitteeSize       = 128
	MaxValidatorsPerCommittee = 2048
	ShuffleRoundCount         = 90
	HysteresisQuotient        = 4
	HysteresisDownwardMultiplier = 1
	HysteresisUpwardMultiplier   = 5
)

// Gwei values
const (
	MinDepositAmount      = 1_000_000_000       // 1 Gwei in Wei
	MaxEffectiveBalance   = 32_000_000_000      // 32 Gwei in Wei
	EffectiveBalanceIncrement = 1_000_000_000  // 1 Gwei in Wei
)

// Time parameters
const (
	SecondsPerSlot                  = 12 * time.Second
	SlotsPerEpoch                  = 32
	MinAttestationInclusionDelay    = 1
	MinSeedLookahead                = 1
	MaxSeedLookahead                = 4
	EpochsPerEth1VotingPeriod       = 64
	EpochsPerHistoricalVector       = 65_536
	EpochsPerSlashingsVector        = 8_192
	HistoricalRootsLimit            = 16_777_216
	ValidatorRegistryLimit          = 1_099_511_627_776
	MinEpochsToInactivityPenalty    = 4
)

// Rewards and penalties
const (
	BaseRewardFactor             = 64
	WhistleblowerRewardQuotient  = 512
	ProposerRewardQuotient       = 8
	InactivityPenaltyQuotient    = 67_108_864
	MinSlashingPenaltyQuotient   = 128
	ProportionalSlashingMultiplier = 1
)


func tconfig(){
	fmt.Println("Slot time:", 12*time.Second)

}