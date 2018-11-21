// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package syncer

import (
	"fmt"

	"github.com/pingcap/tidb-enterprise-tools/pkg/gtid"
	"github.com/siddontang/go-mysql/mysql"
)

type opType byte

const (
	null opType = iota
	insert
	update
	del
	ddl
	xid
	flush
	skip // used by Syncer.recordSkipSQLsPos to record global pos, but not execute SQL
)

func (t opType) String() string {
	switch t {
	case insert:
		return "insert"
	case update:
		return "update"
	case del:
		return "delete"
	case ddl:
		return "ddl"
	case xid:
		return "xid"
	case flush:
		return "flush"
	case skip:
		return "skip"
	}

	return ""
}

type job struct {
	tp           opType
	sourceSchema string
	sourceTable  string
	targetSchema string
	targetTable  string
	sql          string
	args         []interface{}
	key          string
	retry        bool
	pos          mysql.Position
	cmdPos       mysql.Position // exactly binlog position of current SQL
	gtidSet      gtid.Set
	ddlExecItem  *DDLExecItem
	ddls         []string
}

func (j *job) String() string {
	// only output some important information, maybe useful in execution.
	return fmt.Sprintf("sql: %s, args: %v, key: %s, last_pos: %s, cmd_pos: %s, gtid:%v", j.sql, j.args, j.key, j.pos, j.cmdPos, j.gtidSet)
}

// TODO zxc: add RENAME TABLE support
func newJob(tp opType, sourceSchema, sourceTable, targetSchema, targetTable, sql string, args []interface{}, key string, pos, cmdPos mysql.Position, currentGtidSet gtid.Set) *job {
	var gs gtid.Set
	if currentGtidSet != nil {
		gs = currentGtidSet.Clone()
	}
	return &job{
		tp:           tp,
		sourceSchema: sourceSchema,
		sourceTable:  sourceTable,
		targetSchema: targetSchema,
		targetTable:  targetTable,
		sql:          sql,
		args:         args,
		key:          key,
		pos:          pos,
		cmdPos:       cmdPos,
		gtidSet:      gs,
		retry:        true,
	}
}

func newDDLJob(ddlInfo *shardingDDLInfo, ddls []string, pos, cmdPos mysql.Position, currentGtidSet gtid.Set, ddlExecItem *DDLExecItem) *job {
	var gs gtid.Set
	if currentGtidSet != nil {
		gs = currentGtidSet.Clone()
	}
	j := &job{
		tp:          ddl,
		ddls:        ddls,
		pos:         pos,
		cmdPos:      cmdPos,
		gtidSet:     gs,
		ddlExecItem: ddlExecItem,
	}

	if ddlInfo != nil {
		j.sourceSchema = ddlInfo.tableNames[0][0].Schema
		j.sourceTable = ddlInfo.tableNames[0][0].Name
		j.targetSchema = ddlInfo.tableNames[1][0].Schema
		j.targetTable = ddlInfo.tableNames[1][0].Name
	}

	return j
}

func newXIDJob(pos, cmdPos mysql.Position, currentGtidSet gtid.Set) *job {
	var gs gtid.Set
	if currentGtidSet != nil {
		gs = currentGtidSet.Clone()
	}
	return &job{
		tp:      xid,
		pos:     pos,
		cmdPos:  cmdPos,
		gtidSet: gs,
	}
}

func newFlushJob() *job {
	return &job{tp: flush}
}

func newSkipJob(pos mysql.Position, currentGtidSet gtid.Set) *job {
	return &job{
		tp:      skip,
		pos:     pos,
		gtidSet: currentGtidSet,
	}
}

// put queues into bucket to monitor them
func queueBucketName(queueID int) string {
	return fmt.Sprintf("q_%d", queueID%defaultBucketCount)
}
