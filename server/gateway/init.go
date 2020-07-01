package gateway

var (
	CliManager *CliMgr
)

func init() {
	CliManager = new(CliMgr)
	CliManager.CliMap = make(map[int64]*Cli)
}
