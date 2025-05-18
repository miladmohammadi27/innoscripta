package repo

type LedgerRepo interface {
	WriteLogs(data []byte) error
}
