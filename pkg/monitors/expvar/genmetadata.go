// Code generated by monitor-code-gen. DO NOT EDIT.

package expvar

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "expvar"

var groupSet = map[string]bool{}

const (
	memstatsAlloc                = "memstats.alloc"
	memstatsBuckHashSys          = "memstats.buck_hash_sys"
	memstatsBySizeFrees          = "memstats.by_size.frees"
	memstatsBySizeMallocs        = "memstats.by_size.mallocs"
	memstatsBySizeSize           = "memstats.by_size.size"
	memstatsDebugGc              = "memstats.debug_gc"
	memstatsEnableGc             = "memstats.enable_gc"
	memstatsFrees                = "memstats.frees"
	memstatsGcSys                = "memstats.gc_sys"
	memstatsGccpuFraction        = "memstats.gccpu_fraction"
	memstatsHeapAlloc            = "memstats.heap_alloc"
	memstatsHeapIdle             = "memstats.heap_idle"
	memstatsHeapInuse            = "memstats.heap_inuse"
	memstatsHeapObjects          = "memstats.heap_objects"
	memstatsHeapReleased         = "memstats.heap_released"
	memstatsHeapSys              = "memstats.heap_sys"
	memstatsLastGc               = "memstats.last_gc"
	memstatsLookups              = "memstats.lookups"
	memstatsMCacheInuse          = "memstats.m_cache_inuse"
	memstatsMCacheSys            = "memstats.m_cache_sys"
	memstatsMSpanInuse           = "memstats.m_span_inuse"
	memstatsMSpanSys             = "memstats.m_span_sys"
	memstatsMallocs              = "memstats.mallocs"
	memstatsMostRecentGcPauseEnd = "memstats.most_recent_gc_pause_end"
	memstatsMostRecentGcPauseNs  = "memstats.most_recent_gc_pause_ns"
	memstatsNextGc               = "memstats.next_gc"
	memstatsNumForcedGc          = "memstats.num_forced_gc"
	memstatsNumGc                = "memstats.num_gc"
	memstatsOtherSys             = "memstats.other_sys"
	memstatsPauseTotalNs         = "memstats.pause_total_ns"
	memstatsStackInuse           = "memstats.stack_inuse"
	memstatsStackSys             = "memstats.stack_sys"
	memstatsSys                  = "memstats.sys"
	memstatsTotalAlloc           = "memstats.total_alloc"
)

var metricSet = map[string]monitors.MetricInfo{
	memstatsAlloc:                {Type: datapoint.Gauge},
	memstatsBuckHashSys:          {Type: datapoint.Gauge},
	memstatsBySizeFrees:          {Type: datapoint.Count},
	memstatsBySizeMallocs:        {Type: datapoint.Count},
	memstatsBySizeSize:           {Type: datapoint.Count},
	memstatsDebugGc:              {Type: datapoint.Gauge},
	memstatsEnableGc:             {Type: datapoint.Gauge},
	memstatsFrees:                {Type: datapoint.Count},
	memstatsGcSys:                {Type: datapoint.Gauge},
	memstatsGccpuFraction:        {Type: datapoint.Gauge},
	memstatsHeapAlloc:            {Type: datapoint.Gauge},
	memstatsHeapIdle:             {Type: datapoint.Gauge},
	memstatsHeapInuse:            {Type: datapoint.Gauge},
	memstatsHeapObjects:          {Type: datapoint.Gauge},
	memstatsHeapReleased:         {Type: datapoint.Gauge},
	memstatsHeapSys:              {Type: datapoint.Gauge},
	memstatsLastGc:               {Type: datapoint.Gauge},
	memstatsLookups:              {Type: datapoint.Count},
	memstatsMCacheInuse:          {Type: datapoint.Gauge},
	memstatsMCacheSys:            {Type: datapoint.Gauge},
	memstatsMSpanInuse:           {Type: datapoint.Gauge},
	memstatsMSpanSys:             {Type: datapoint.Gauge},
	memstatsMallocs:              {Type: datapoint.Count},
	memstatsMostRecentGcPauseEnd: {Type: datapoint.Gauge},
	memstatsMostRecentGcPauseNs:  {Type: datapoint.Gauge},
	memstatsNextGc:               {Type: datapoint.Gauge},
	memstatsNumForcedGc:          {Type: datapoint.Count},
	memstatsNumGc:                {Type: datapoint.Count},
	memstatsOtherSys:             {Type: datapoint.Gauge},
	memstatsPauseTotalNs:         {Type: datapoint.Count},
	memstatsStackInuse:           {Type: datapoint.Gauge},
	memstatsStackSys:             {Type: datapoint.Gauge},
	memstatsSys:                  {Type: datapoint.Gauge},
	memstatsTotalAlloc:           {Type: datapoint.Count},
}

var defaultMetrics = map[string]bool{
	memstatsBuckHashSys:         true,
	memstatsFrees:               true,
	memstatsGcSys:               true,
	memstatsGccpuFraction:       true,
	memstatsHeapAlloc:           true,
	memstatsHeapIdle:            true,
	memstatsHeapInuse:           true,
	memstatsHeapObjects:         true,
	memstatsHeapReleased:        true,
	memstatsHeapSys:             true,
	memstatsMCacheInuse:         true,
	memstatsMCacheSys:           true,
	memstatsMSpanInuse:          true,
	memstatsMSpanSys:            true,
	memstatsMallocs:             true,
	memstatsMostRecentGcPauseNs: true,
	memstatsNextGc:              true,
	memstatsNumForcedGc:         true,
	memstatsNumGc:               true,
	memstatsOtherSys:            true,
	memstatsStackInuse:          true,
	memstatsStackSys:            true,
	memstatsSys:                 true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "expvar",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}
