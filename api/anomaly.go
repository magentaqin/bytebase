package api

import (
	"context"
	"encoding/json"
)

// AnomalyType is the type of a task.
type AnomalyType string

const (
	AnomalyDatabaseBackupPolicyViolation AnomalyType = "bb.anomaly.database.backup.policy-violation"
	AnomalyDatabaseBackupMissing         AnomalyType = "bb.anomaly.database.backup.missing"
	AnomalyDatabaseConnection            AnomalyType = "bb.anomaly.database.connection"
	AnomalyDatabaseSchemaDrift           AnomalyType = "bb.anomaly.database.schema.drift"
)

type AnomalyDatabaseBackupPolicyViolationPayload struct {
	EnvironmentId          int                      `json:"environmentId,omitempty"`
	ExpectedBackupSchedule BackupPlanPolicySchedule `json:"expectedSchedule,omitempty"`
	ActualBackupSchedule   BackupPlanPolicySchedule `json:"actualSchedule,omitempty"`
}

type AnomalyDatabaseBackupMissingPayload struct {
	ExpectedBackupSchedule BackupPlanPolicySchedule `json:"expectedSchedule,omitempty"`
	// Time of last successful backup created
	LastBackupTs int64 `json:"lastBackupTs,omitempty"`
}

type AnomalyDatabaseConnectionPayload struct {
	// Connection failure detail
	Detail string `json:"detail,omitempty"`
}

type AnomalyDatabaseSchemaDriftPayload struct {
	// The schema version corresponds to the expected schema
	Version string `json:"version,omitempty"`
	// The expected latest schema stored in the migration history table
	Expect string `json:"expect,omitempty"`
	// The actual schema dumped from the database
	Actual string `json:"actual,omitempty"`
}

type Anomaly struct {
	ID int `jsonapi:"primary,anomaly"`

	// Standard fields
	CreatorId int
	Creator   *Principal `jsonapi:"attr,creator"`
	CreatedTs int64      `jsonapi:"attr,createdTs"`
	UpdaterId int
	Updater   *Principal `jsonapi:"attr,updater"`
	UpdatedTs int64      `jsonapi:"attr,updatedTs"`

	// Related fields
	InstanceId int `jsonapi:"attr,instanceId"`
	DatabaseId int `jsonapi:"attr,databaseId"`

	// Domain specific fields
	Type    AnomalyType `jsonapi:"attr,type"`
	Payload string      `jsonapi:"attr,payload"`
}

type AnomalyUpsert struct {
	// Standard fields
	CreatorId int

	// Related fields
	InstanceId int
	DatabaseId int

	// Domain specific fields
	Type    AnomalyType `jsonapi:"attr,type"`
	Payload string      `jsonapi:"attr,payload"`
}

type AnomalyFind struct {
	// Standard fields
	RowStatus *RowStatus

	// Related fields
	DatabaseId int
	Type       *AnomalyType
}

func (find *AnomalyFind) String() string {
	str, err := json.Marshal(*find)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

type AnomalyArchive struct {
	DatabaseId int
	Type       AnomalyType
}

type AnomalyService interface {
	// UpsertActiveAnomaly would update the existing active anomaly if both database id and type match, otherwise create a new one.
	UpsertActiveAnomaly(ctx context.Context, upsert *AnomalyUpsert) (*Anomaly, error)
	FindAnomalyList(ctx context.Context, find *AnomalyFind) ([]*Anomaly, error)
	ArchiveAnomaly(ctx context.Context, archive *AnomalyArchive) error
}
