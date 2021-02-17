package uiamsdk

import (
	"time"
)

// AuthorizedKey AuthorizedKey
type AuthorizedKey struct {
	ID         uint64         `gorm:"PRIMARY_KEY;NOT NULL"`
	IdentityID string         `gorm:"COLUMN:identity_id;NOT NULL;TYPE:BIGINT()"`
	ReamID     string         `gorm:"COLUMN:realm_id;NOT NULL;TYPE:VARCHAR(36)"`
	Name       string         `gorm:"COLUMN:name;NOT NULL;TYPE:VARCHAR(36)"`
	Scheme     AuthSchemeEnum `gorm:"COLUMN:scheme;NOT NULL;TYPE:VARCHAR(16)"`
	Scopes     string         `gorm:"COLUMN:scopes;NOT NULL;TYPE:VARCHAR(128)"`
	Key        string         `gorm:"COLUMN:auth_key;NOT NULL;TYPE:VARCHAR(36)"`
	Secret     string         `gorm:"COLUMN:secret;NOT NULL;TYPE:VARCHAR(1024)"`
	Remark     string         `gorm:"COLUMN:remark;NOT NULL;TYPE:VARCHAR(255)"`
	ExpriedAt  *time.Time     `gorm:"COLUMN:expried_at;NOT NULL"`
	CreatedAt  time.Time      `gorm:"COLUMN:created_at;NOT NULL"`
}

// TableName TableName
func (AuthorizedKey) TableName() string {
	return "authorized_keys"
}
