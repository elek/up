// AUTOGENERATED BY pkg/config/gen
// DO NOT EDIT.

package config

func authserviceConfig() []Option {
	return []Option{
		
		{
			Name:        "endpoint",
			Description: "Gateway endpoint URL to return to clients",
			Default:     "",
		},
		{
			Name:        "auth-token",
			Description: "auth security token to validate requests",
			Default:     "",
		},
		{
			Name:        "post-size-limit",
			Description: "maximum size that the incoming POST request body with access grant can be",
			Default:     "4KiB",
		},
		{
			Name:        "allowed-satellites",
			Description: "list of satellite NodeURLs allowed for incoming access grants",
			Default:     "https://www.storj.io/dcs-satellites",
		},
		{
			Name:        "cache-expiration",
			Description: "length of time satellite addresses are cached for",
			Default:     "10m",
		},
		{
			Name:        "kv-backend",
			Description: "key/value store backend url",
			Default:     "",
		},
		{
			Name:        "migration",
			Description: "create or update the database schema, and then continue service startup",
			Default:     "false",
		},
		{
			Name:        "listen-addr",
			Description: "public HTTP address to listen on",
			Default:     ":20000",
		},
		{
			Name:        "listen-addr-tls",
			Description: "public HTTPS address to listen on",
			Default:     ":20001",
		},
		{
			Name:        "drpc-listen-addr",
			Description: "public DRPC address to listen on",
			Default:     ":20002",
		},
		{
			Name:        "drpc-listen-addr-tls",
			Description: "public DRPC+TLS address to listen on",
			Default:     ":20003",
		},
		{
			Name:        "lets-encrypt",
			Description: "use lets-encrypt to handle TLS certificates",
			Default:     "false",
		},
		{
			Name:        "cert-file",
			Description: "server certificate file",
			Default:     "",
		},
		{
			Name:        "key-file",
			Description: "server key file",
			Default:     "",
		},
		{
			Name:        "public-url",
			Description: "public url for the server, for the TLS certificate",
			Default:     "",
		},
		{
			Name:        "delete-unused.run",
			Description: "whether to run unused records deletion chore",
			Default:     "false",
		},
		{
			Name:        "delete-unused.interval",
			Description: "interval unused records deletion chore waits to start next iteration",
			Default:     "24h",
		},
		{
			Name:        "delete-unused.as-of-system-interval",
			Description: "the interval specified in AS OF SYSTEM in unused records deletion chore query as negative interval",
			Default:     "5s",
		},
		{
			Name:        "delete-unused.select-size",
			Description: "batch size of records selected for deletion at a time",
			Default:     "10000",
		},
		{
			Name:        "delete-unused.delete-size",
			Description: "batch size of records to delete from selected records at a time",
			Default:     "1000",
		},
		{
			Name:        "node.id",
			Description: "unique identifier for the node",
			Default:     "",
		},
		{
			Name:        "node.first-start",
			Description: "allow start with empty storage",
			Default:     "",
		},
		{
			Name:        "node.path",
			Description: "path where to store data",
			Default:     "",
		},
		{
			Name:        "node.address",
			Description: "address that the node listens on",
			Default:     ":20004",
		},
		{
			Name:        "node.join",
			Description: "comma delimited list of cluster peers",
			Default:     "",
		},
		{
			Name:        "node.certs-dir",
			Description: "directory for certificates for mutual authentication",
			Default:     "",
		},
		{
			Name:        "node.replication-interval",
			Description: "how often to replicate",
			Default:     "30s",
		},
		{
			Name:        "node.replication-limit",
			Description: "maximum entries returned in replication response",
			Default:     "1000",
		},
		{
			Name:        "node.conflict-backoff.delay",
			Description: "The active time between retries, typically not set",
			Default:     "0ms",
		},
		{
			Name:        "node.conflict-backoff.max",
			Description: "The maximum total time to allow retries",
			Default:     "5m",
		},
		{
			Name:        "node.conflict-backoff.min",
			Description: "The minimum time between retries",
			Default:     "100ms",
		},
		{
			Name:        "node.insecure-disable-tls",
			Description: "",
			Default:     "",
		},
		{
			Name:        "node.backup.enabled",
			Description: "enable backups",
			Default:     "false",
		},
		{
			Name:        "node.backup.endpoint",
			Description: "backup bucket endpoint hostname, e.g. s3.amazonaws.com",
			Default:     "",
		},
		{
			Name:        "node.backup.bucket",
			Description: "bucket name where database backups are stored",
			Default:     "",
		},
		{
			Name:        "node.backup.prefix",
			Description: "database backup object path prefix",
			Default:     "",
		},
		{
			Name:        "node.backup.interval",
			Description: "how often full backups are run",
			Default:     "1h",
		},
		{
			Name:        "node.backup.access-key-id",
			Description: "access key for backup bucket",
			Default:     "",
		},
		{
			Name:        "node.backup.secret-access-key",
			Description: "secret key for backup bucket",
			Default:     "",
		},
		{
			Name:        "node-migration.migration-select-size",
			Description: "page size while performing migration",
			Default:     "1000",
		},
		{
			Name:        "node-migration.source-sql-auth-kv-backend",
			Description: "source key/value store backend (must be sqlauth) url",
			Default:     "",
		},
   }
}
