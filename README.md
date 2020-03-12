# uniqio
Use the binary `uniqio` to stream unique lines from stdin to stdout.

## If you're using `uniqio` with `jq`,

then make sure you use the `--buffered` flag on jq.

```
cat some-file.json | jq -r --unbuffered .some_key | uniqio
```
