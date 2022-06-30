module github.com/richbai90/xfer-bin

go 1.18

replace (
	// a dependency of buildx (github.com/tonistiigi/fsutil) relies on an untagged version of Docker on the master branch
	github.com/docker/cli => github.com/docker/cli v20.10.3-0.20220226190722-8667ccd1124c+incompatible
	github.com/docker/docker => github.com/docker/docker v20.10.3-0.20220121014307-40bb9831756f+incompatible

	// See https://github.com/hashicorp/go-plugin/pull/127 and
	// https://github.com/hashicorp/go-plugin/pull/163
	// Also includes branches we haven't PR'd yet: capture-yamux-logs, context-cancellation
	// Tagged from v1.4.4, the porter branch
	github.com/hashicorp/go-plugin => github.com/getporter/go-plugin v1.4.4-porter.1

	// Fixes https://github.com/spf13/viper/issues/761
	github.com/spf13/viper => github.com/getporter/viper v1.7.1-porter.2.0.20210514172839-3ea827168363
)

require (
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.5.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
