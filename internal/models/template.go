package models

var sqlExample = `SELECT setval('action_record_logs_id_seq', (SELECT MAX(id) FROM action_record_logs)+1);`

const IdSeqSetValSqlTpl = `SELECT setval('{{ .sequenceName }}',(SELECT MAX(id) FROM {{ .tableName }})+1);`
