investigator
============

Integrate all information over a Riak / Riak CS cluster and help
investigation, by processing riak-debug and riak-cs-debug. It was a
pain to seek for root cause between scattered information over
multiple nodes. This tool enables such pain works done in a single
command.

## Example check items

- `app.config` and `vm.args` correctly configured
- The IO scheduler correctly configured at all nodes
- Correct `pb_backlog`, `somaxconn`, `request_pools`
- Parse logs and stores into database, enabling qieries and aggregations
- Find suspicious log line
- Verify kernel logs and parameters
- Verify correct version

# External Depends

- Erlang/OTP (to parse `app.config`)

- Don't need in runtime

 - mattn/go-sqlite3
 - josh/gobert

# License

Apache 2.0
