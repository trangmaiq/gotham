package sqlcon

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ParseConnectionOptions parses value for max_conns, max_idle_conns, max_conn_lifetime from DSNs.
// It also returns the URI without those query parameters.
func ParseConnectionOptions(dsn string) (maxConns int, maxIdleConns int, maxConnLifetime time.Duration, cleanedDSN string) {
	maxConns = maxParallelism() * 2
	maxIdleConns = maxParallelism()
	maxConnLifetime = time.Duration(0)
	cleanedDSN = dsn

	var parts = strings.Split(dsn, "?")
	if len(parts) != 2 {
		// TODO: Replace this println with logger later
		println("sql_max_connections", maxConns)
		println("sql_max_idle_connections", maxIdleConns)
		println("sql_max_connection_lifetime", maxConnLifetime)
		println("DEBUG | No SQL connection options have been defined, falling back to default connection options")
		return
	}

	query, err := url.ParseQuery(parts[1])
	if err != nil {
		println("sql_max_connections", maxConns)
		println("sql_max_idle_connections", maxIdleConns)
		println("sql_max_connection_lifetime", maxConnLifetime)
		println("ERROR | ", err.Error())
		println("WARN | Unable to parse SQL DSN query, falling back to default connection options.")
		return
	}

	if v := query.Get("max_conns"); v != "" {
		s, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			println("ERROR | ", err.Error())
			fmt.Printf(`WARN | SQL DSN query parameter "max_conns" value %v could not be parsed to int, falling back to default value %d`, v, maxConns)
		} else {
			maxIdleConns = int(s)
		}
		query.Del("max_idle_conns")
	}

	if v := query.Get("max_conn_lifetime"); v != "" {
		s, err := time.ParseDuration(v)
		if err != nil {
			println("ERROR | ", err.Error())
			fmt.Printf(`WARN | SQL DSN query parameter "max_conn_lifetime" value %v could not be parsed to int, falling back to default value %d`, v, maxConnLifetime)

		} else {
			maxConnLifetime = s
		}
		query.Del("max_conn_lifetime")
	}
	cleanedDSN = fmt.Sprintf("%s?%s", parts[0], query.Encode())

	return
}
