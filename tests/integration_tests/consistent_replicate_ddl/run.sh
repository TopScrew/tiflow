#!/bin/bash

set -eu

CUR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
source $CUR/../_utils/test_prepare
WORK_DIR=$OUT_DIR/$TEST_NAME
CDC_BINARY=cdc.test
SINK_TYPE=$1

rm -rf "$WORK_DIR"
mkdir -p "$WORK_DIR"

stop() {
	# to distinguish whether the test failed in the DML synchronization phase or the DDL synchronization phase
	echo $(mysql -h${DOWN_TIDB_HOST} -P${DOWN_TIDB_PORT} -uroot -e "select count(*) from consistent_replicate_ddl.usertable;")
	stop_tidb_cluster
}

function run() {
	# we only support eventually consistent replication with MySQL sink
	if [ "$SINK_TYPE" != "mysql" ]; then
		return
	fi

	start_tidb_cluster --workdir $WORK_DIR

	cd $WORK_DIR
	run_sql "set @@global.tidb_enable_exchange_partition=on" ${DOWN_TIDB_HOST} ${DOWN_TIDB_PORT}

	run_cdc_server --workdir $WORK_DIR --binary $CDC_BINARY

	SINK_URI="mysql://normal:123456@127.0.0.1:3306/"
	changefeed_id=$(cdc cli changefeed create --sink-uri="$SINK_URI" --config="$CUR/conf/changefeed.toml" 2>&1 | tail -n2 | head -n1 | awk '{print $2}')

	run_sql "CREATE DATABASE consistent_replicate_ddl;" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	go-ycsb load mysql -P $CUR/conf/workload -p mysql.host=${UP_TIDB_HOST} -p mysql.port=${UP_TIDB_PORT} -p mysql.user=root -p mysql.db=consistent_replicate_ddl
	run_sql "create table consistent_replicate_ddl.usertable2 like consistent_replicate_ddl.usertable" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	run_sql "create table consistent_replicate_ddl.usertable3 like consistent_replicate_ddl.usertable" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	check_table_exists "consistent_replicate_ddl.usertable" ${DOWN_TIDB_HOST} ${DOWN_TIDB_PORT}
	check_table_exists "consistent_replicate_ddl.usertable2" ${DOWN_TIDB_HOST} ${DOWN_TIDB_PORT} 120
	check_table_exists "consistent_replicate_ddl.usertable3" ${DOWN_TIDB_HOST} ${DOWN_TIDB_PORT} 120

	sleep 20
	# Inject the failpoint to prevent sink execution, but the global resolved can be moved forward.
	# Then we can apply redo log to reach an eventual consistent state in downstream.
	cleanup_process $CDC_BINARY
	export GO_FAILPOINTS='github.com/pingcap/tiflow/cdc/sink/dmlsink/txn/mysql/MySQLSinkHangLongTime=return(true);github.com/pingcap/tiflow/cdc/sink/ddlsink/mysql/MySQLSinkExecDDLDelay=return(true)'
	run_cdc_server --workdir $WORK_DIR --binary $CDC_BINARY
	run_sql "insert into consistent_replicate_ddl.usertable2 select * from consistent_replicate_ddl.usertable" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	run_sql "ALTER TABLE consistent_replicate_ddl.usertable2 DROP COLUMN FIELD0" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	run_sql "ALTER TABLE consistent_replicate_ddl.usertable2 ADD COLUMN dummy varchar(30)" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	run_sql "insert into consistent_replicate_ddl.usertable3 select * from consistent_replicate_ddl.usertable" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	run_sql "ALTER TABLE consistent_replicate_ddl.usertable3 DROP COLUMN FIELD1" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	# enable this after global resolved ts is decoupled from globalBarrierTs
	# run_sql "create table consistent_replicate_ddl.usertable3 like consistent_replicate_ddl.usertable" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	# run_sql "insert into consistent_replicate_ddl.usertable3 select * from consistent_replicate_ddl.usertable" ${UP_TIDB_HOST} ${UP_TIDB_PORT}
	# run_sql "CREATE table consistent_replicate_ddl.check1(id int primary key);" ${UP_TIDB_HOST} ${UP_TIDB_PORT}

	# to ensure row changed events have been replicated to TiCDC
	sleep 20

	storage_path="file://$WORK_DIR/redo"
	tmp_download_path=$WORK_DIR/cdc_data/redo/$changefeed_id
	current_tso=$(cdc cli tso query --pd=http://$UP_PD_HOST_1:$UP_PD_PORT_1)
	ensure 50 check_redo_resolved_ts $changefeed_id $current_tso $storage_path $tmp_download_path/meta
	cleanup_process $CDC_BINARY

	export GO_FAILPOINTS=''

	cdc redo apply --tmp-dir="$tmp_download_path/apply" \
		--storage="$storage_path" \
		--sink-uri="mysql://normal:123456@127.0.0.1:3306/"
	check_sync_diff $WORK_DIR $CUR/conf/diff_config.toml
}

trap stop EXIT
run $*
check_logs $WORK_DIR
echo "[$(date)] <<<<<< run test case $TEST_NAME success! >>>>>>"
