package subscribe

import (
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/itskass/headhunter/gather"
)

type Options struct {
	Client        *ethclient.Client
	Delay         time.Duration
	GatherOptions *gather.Options
}
