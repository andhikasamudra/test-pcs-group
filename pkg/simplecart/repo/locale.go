package repo

import "github.com/uptrace/bun"

type Locale struct {
	Key    string `bun:",pk"`
	LangID string
	LangEN string
	RecordTimestamp

	bun.BaseModel `bun:"table:locales"`
}
