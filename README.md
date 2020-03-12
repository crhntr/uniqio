# uniqio
Use the binary `uniqio` to stream unique lines from stdin to stdout.

## If Using with JQ

Make sure you use the `--buffered` flag on jq.

```
cat some-file.json | jq -r --unbuffered .some_key | uniqio
```
