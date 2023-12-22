INSERT INTO kubedb-dummy
SELECT
TIME_PARSE("timestamp") AS __time, *
FROM TABLE(
EXTERN(
'{"type":"inline","data":"\"id\": 27, \"name\": \"Tapojit\", \"time\": \"2015-09-12T00:46:58.771Z\""}',
'{"type":"json"}',
'[{"name":"timestamp","type":"string"},{"name":"name","type":"string"},{"name":"id","type":"long"}]'
)
)
PARTITIONED BY DAY