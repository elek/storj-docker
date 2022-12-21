// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package config

func authserviceConfig() []Option {
	return []Option{

		{
			Name:        "STORJ_ENDPOINT",
			Description: "Gateway endpoint URL to return to clients",
			Default:     "",
		},
		{
			Name:        "STORJ_AUTH_TOKEN",
			Description: "auth security token to validate requests",
			Default:     "",
		},
		{
			Name:        "STORJ_POSTSIZE_LIMIT",
			Description: "maximum size that the incoming POST request body with access grant can be",
			Default:     "4KiB",
		},
		{
			Name:        "STORJ_ALLOWED_SATELLITES",
			Description: "list of satellite NodeURLs allowed for incoming access grants",
			Default:     "https://www.storj.io/dcs-satellites",
		},
		{
			Name:        "STORJ_CACHE_EXPIRATION",
			Description: "length of time satellite addresses are cached for",
			Default:     "10m",
		},
		{
			Name:        "STORJ_KVBACKEND",
			Description: "key/value store backend url",
			Default:     "",
		},
		{
			Name:        "STORJ_MIGRATION",
			Description: "create or update the database schema, and then continue service startup",
			Default:     "false",
		},
		{
			Name:        "STORJ_LISTEN_ADDR",
			Description: "public HTTP address to listen on",
			Default:     ":20000",
		},
		{
			Name:        "STORJ_LISTEN_ADDR_TLS",
			Description: "public HTTPS address to listen on",
			Default:     ":20001",
		},
		{
			Name:        "STORJ_DRPCLISTEN_ADDR",
			Description: "public DRPC address to listen on",
			Default:     ":20002",
		},
		{
			Name:        "STORJ_DRPCLISTEN_ADDR_TLS",
			Description: "public DRPC+TLS address to listen on",
			Default:     ":20003",
		},
		{
			Name:        "STORJ_LETS_ENCRYPT",
			Description: "use lets-encrypt to handle TLS certificates",
			Default:     "false",
		},
		{
			Name:        "STORJ_CERT_FILE",
			Description: "server certificate file",
			Default:     "",
		},
		{
			Name:        "STORJ_KEY_FILE",
			Description: "server key file",
			Default:     "",
		},
		{
			Name:        "STORJ_PUBLIC_URL",
			Description: "public url for the server, for the TLS certificate",
			Default:     "",
		},
		{
			Name:        "STORJ_DELETE_UNUSED_RUN",
			Description: "whether to run unused records deletion chore",
			Default:     "false",
		},
		{
			Name:        "STORJ_DELETE_UNUSED_INTERVAL",
			Description: "interval unused records deletion chore waits to start next iteration",
			Default:     "24h",
		},
		{
			Name:        "STORJ_DELETE_UNUSED_AS_OF_SYSTEM_INTERVAL",
			Description: "the interval specified in AS OF SYSTEM in unused records deletion chore query as negative interval",
			Default:     "5s",
		},
		{
			Name:        "STORJ_DELETE_UNUSED_SELECT_SIZE",
			Description: "batch size of records selected for deletion at a time",
			Default:     "10000",
		},
		{
			Name:        "STORJ_DELETE_UNUSED_DELETE_SIZE",
			Description: "batch size of records to delete from selected records at a time",
			Default:     "1000",
		},
		{
			Name:        "STORJ_NODE_ID",
			Description: "unique identifier for the node",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_FIRST_START",
			Description: "allow start with empty storage",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_PATH",
			Description: "path where to store data",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_ADDRESS",
			Description: "address that the node listens on",
			Default:     ":20004",
		},
		{
			Name:        "STORJ_NODE_JOIN",
			Description: "comma delimited list of cluster peers",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_CERTS_DIR",
			Description: "directory for certificates for mutual authentication",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_REPLICATION_INTERVAL",
			Description: "how often to replicate",
			Default:     "30s",
		},
		{
			Name:        "STORJ_NODE_REPLICATION_LIMIT",
			Description: "maximum entries returned in replication response",
			Default:     "1000",
		},
		{
			Name:        "STORJ_NODE_CONFLICT_BACKOFF_DELAY",
			Description: "The active time between retries, typically not set",
			Default:     "0ms",
		},
		{
			Name:        "STORJ_NODE_CONFLICT_BACKOFF_MAX",
			Description: "The maximum total time to allow retries",
			Default:     "5m",
		},
		{
			Name:        "STORJ_NODE_CONFLICT_BACKOFF_MIN",
			Description: "The minimum time between retries",
			Default:     "100ms",
		},
		{
			Name:        "STORJ_NODE_INSECURE_DISABLE_TLS",
			Description: "",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_BACKUP_ENABLED",
			Description: "enable backups",
			Default:     "false",
		},
		{
			Name:        "STORJ_NODE_BACKUP_ENDPOINT",
			Description: "backup bucket endpoint hostname, e.g. s3.amazonaws.com",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_BACKUP_BUCKET",
			Description: "bucket name where database backups are stored",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_BACKUP_PREFIX",
			Description: "database backup object path prefix",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_BACKUP_INTERVAL",
			Description: "how often full backups are run",
			Default:     "1h",
		},
		{
			Name:        "STORJ_NODE_BACKUP_ACCESS_KEY_ID",
			Description: "access key for backup bucket",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_BACKUP_SECRET_ACCESS_KEY",
			Description: "secret key for backup bucket",
			Default:     "",
		},
		{
			Name:        "STORJ_NODE_MIGRATION_MIGRATION_SELECT_SIZE",
			Description: "page size while performing migration",
			Default:     "1000",
		},
		{
			Name:        "STORJ_NODE_MIGRATION_SOURCE_SQLAUTH_KVBACKEND",
			Description: "source key/value store backend (must be sqlauth) url",
			Default:     "",
		},
	}
}
