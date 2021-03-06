/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package defaults

import (
	"github.com/tricksterproxy/trickster/pkg/cache/evictionmethods"
	"github.com/tricksterproxy/trickster/pkg/cache/providers"
)

const (
	// DefaultLogFile is the default disk location for log files.
	// we use an empty string to indicate log to console
	DefaultLogFile = ""
	// DefaultLogLevel is the default level for logging
	DefaultLogLevel = "INFO"

	// DefaultProxyListenPort is the default port that the HTTP frontend will listen on
	DefaultProxyListenPort = 8480
	// DefaultProxyListenAddress is the default address that the HTTP frontend will listen on
	DefaultProxyListenAddress = ""

	// DefaultMetricsListenPort is the default port that the HTTP metrics endpoint will listen on
	DefaultMetricsListenPort = 8481
	// DefaultMetricsListenAddress is the default address that the HTTP metrics endpoint will listen on
	DefaultMetricsListenAddress = ""

	// 8482 is reserved for mockster, allowing the default TLS port to end with 3

	// DefaultTLSProxyListenPort is the default port that the TLS frontend endpoint will listen on
	DefaultTLSProxyListenPort = 8483
	// DefaultTLSProxyListenAddress is the default address that the TLS frontend endpoint will listen on
	DefaultTLSProxyListenAddress = ""

	// DefaultReloadPort is the default port that the Reload endpoint will listen on
	DefaultReloadPort = 8484
	// DefaultReloadAddress is the default address that the Reload endpoint will listen on
	DefaultReloadAddress = "127.0.0.1"
	// DefaultDrainTimeoutMS is the default time that is allowed for an old configuration's requests to drain
	// before its resources are closed
	DefaultDrainTimeoutMS = 30000
	// DefaultRateLimitMS is the default Rate Limit time for Config Reloads
	DefaultRateLimitMS = 3000

	// DefaultTracerProvider is the default distributed tracer exporter implementation
	DefaultTracerProvider = "none"

	// DefaultTracerServiceName is the default service name under which traces are registered
	DefaultTracerServiceName = "trickster"

	// DefaultCacheProvider is the default cache providers for any defined cache
	DefaultCacheProvider = "memory"
	// DefaultCacheProviderID is the default cache providers ID for any defined cache
	// and should align with DefaultCacheProvider
	DefaultCacheProviderID = providers.Memory

	// DefaultTimeseriesTTLMS is the default Cache TTL for Time Series Objects
	DefaultTimeseriesTTLMS = 21600000
	// DefaultFastForwardTTLMS is the default Cache TTL for Time Series Fast Forward Objects
	DefaultFastForwardTTLMS = 15000
	// DefaultMaxTTLMS is the default Maximum TTL of any cache object
	DefaultMaxTTLMS = 86400000
	// DefaultRevalidationFactor is the default Cache Object Freshness Lifetime to TTL multiplier
	DefaultRevalidationFactor = 2
	// DefaultRedisClientType is the default Redis Client Type
	DefaultRedisClientType = "standard"
	// DefaultRedisProtocol is the default Redis Client protocol
	DefaultRedisProtocol = "tcp"
	// DefaultRedisEndpoint is the default Redis Client endpoint
	DefaultRedisEndpoint = "redis:6379"
	// DefaultBBoltFile is the default bbolt Cache filename
	DefaultBBoltFile = "trickster.db"
	// DefaultBBoltBucket is the default bbolt Cache bucket name
	DefaultBBoltBucket = "trickster"
	// DefaultCacheIndexReap is the default Cache Index Reap interval (in milliseconds)
	DefaultCacheIndexReap = 3000
	// DefaultCacheIndexFlush is the default Cache Index Flush interval (in milliseconds)
	DefaultCacheIndexFlush = 5000
	// DefaultCacheMaxSizeBytes is the default Max Cache Size in Bytes
	DefaultCacheMaxSizeBytes = 536870912
	// DefaultMaxSizeBackoffBytes is the default Max Cache Backoff Size in Bytes
	DefaultMaxSizeBackoffBytes = 16777216
	// DefaultMaxSizeObjects is the default Max Cache Object Count
	DefaultMaxSizeObjects = 0
	// DefaultMaxSizeBackoffObjects is the default Max Cache Backoff Object Count
	DefaultMaxSizeBackoffObjects = 100
	// DefaultMaxObjectSizeBytes is the default Max Size of any Cache Object
	DefaultMaxObjectSizeBytes = 524288
	// DefaultBackendTRF is the default Timeseries Retention Factor for Time Series-based Backends
	DefaultBackendTRF = 1024
	// DefaultBackendTEM is the default Timeseries Eviction Method for Time Series-based Backends
	DefaultBackendTEM = evictionmethods.EvictionMethodOldest
	// DefaultBackendTEMName is the default Timeseries Eviction Method name for Time Series-based Backends
	DefaultBackendTEMName = "oldest"
	// DefaultBackendTimeoutMS is the default Upstream Request Timeout for Backends
	DefaultBackendTimeoutMS = 180000
	// DefaultBackendCacheName is the default Cache Name for Backends
	DefaultBackendCacheName = "default"
	// DefaultBackendNegativeCacheName is the default Negative Cache Name for Backends
	DefaultBackendNegativeCacheName = "default"
	// DefaultTracingConfigName is the default Tracing Config Name for Backends
	DefaultTracingConfigName = "default"
	// DefaultBackfillToleranceMS is the default Backfill Tolerance setting for Backends
	DefaultBackfillToleranceMS = 0
	// DefaultKeepAliveTimeoutMS is the default Keep Alive Timeout for Backends' upstream client pools
	DefaultKeepAliveTimeoutMS = 300000
	// DefaultMaxIdleConns is the default number of Idle Connections in Backends' upstream client pools
	DefaultMaxIdleConns = 20
	// DefaultHealthCheckPath is the default value (noop) for Backends' Health Check Path
	DefaultHealthCheckPath = "-"
	// DefaultHealthCheckQuery is the default value (noop) for Backends' Health Check Query Parameters
	DefaultHealthCheckQuery = "-"
	// DefaultHealthCheckVerb is the default value (noop) for Backends' Health Check Verb
	DefaultHealthCheckVerb = "-"
	// DefaultConfigHandlerPath is the default value for the Trickster Config Printout Handler path
	DefaultConfigHandlerPath = "/trickster/config"
	// DefaultPingHandlerPath is the default value for the Trickster Config Ping Handler path
	DefaultPingHandlerPath = "/trickster/ping"
	// DefaultReloadHandlerPath defines the default path for the Reload Handler
	DefaultReloadHandlerPath = "/trickster/config/reload"
	// DefaultHealthHandlerPath defines the default path for the Health Handler
	DefaultHealthHandlerPath = "/trickster/health"
	// DefaultMaxRuleExecutions is the default value for the number of allowed Rule executions per Request
	DefaultMaxRuleExecutions = 16
	// DefaultPprofServerName defines the default Pprof Server Name
	DefaultPprofServerName = "both"
	// DefaultForwardedHeaders defines which class of 'Forwarded' headers are attached to upstream requests
	DefaultForwardedHeaders = "standard"
)

// DefaultCompressableTypes returns a list of types that Trickster should compress before caching
func DefaultCompressableTypes() []string {
	return []string{
		"text/html",
		"text/javascript",
		"text/css",
		"text/plain",
		"text/xml",
		"text/json",
		"application/json",
		"application/javascript",
		"application/xml",
	}
}
