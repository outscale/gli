# Waiting for a condition

octl can wait for a specific condition. It will run repeatedly until the condition is met or a timeout is reached.

Conditions are jq queries executed on the octl call result. The query is expected to return a single result.

octl will loop if the query returns:
* `false` or `"false"`,
* `0` or `"0"`,
* `""` or a null value.

waitfor works on calls returning a single entry:
```shell
octl iaas vm describe i-foo --waitfor '.State=="running"'
```
or returning a list:
```shell
octl iaas vm list --waitfor 'length==6'
```

By default, waitfor runs the query every 10 seconds (configured by `--waitfor-interval`) and timeouts at 10 minutes (configured by `--waitfor-timeout`).
