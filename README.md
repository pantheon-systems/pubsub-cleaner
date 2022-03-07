PubSub Cleaner
--------------

[![Unsupported](https://img.shields.io/badge/Pantheon-Unsupported-yellow?logo=pantheon&color=FFDC28)](https://pantheon.io/docs/oss-support-levels#unsupported)

Helps you clean up subscriptions on google PubSub

Installing
==========

Use `go get` to install

```
go get -u github.com/pantheon-systems/pubsub-cleaner
```

Running
=======

Clean out subscriptions on a topic:

```
pubsub-cleaner topic sometopic --project myproject --keep subname --no-op
```

-	`--project` is the GCE project.
-	`--keep` is a string match on subscriber names that you want to keep.
-	`--no-op` won't actually delete anything, just output what would be deleted

Authentication
==============

This tool relies on Google Application Default Credentials, and honors the environment variables used by that.

read more here: https://developers.google.com/identity/protocols/application-default-credentials

Help
====

The program supports the standard `-h` and `--help` flags
