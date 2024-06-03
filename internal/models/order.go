package models

import (
	"time"

	"github.com/sev-2/raiden"
)

type Order struct {
	raiden.ModelBase
	Id        int64     `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"authenticated" write:"authenticated"`
}
