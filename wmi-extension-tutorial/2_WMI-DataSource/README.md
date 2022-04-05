# Exercise 2 - The WMI Metric Data Source
---

## Description

For your extension to be able to collect any metrics and have them ingested into Dynatrace, you must define a data source. In this tutorial we're using the WMI data source. This must be a section called `wmi` in your extension.

The purpose of the `wmi` section is define the WMI queries that retrieve your metrics, how often they should run, and how to map their results into metrics and dimensions that Dynatrace can ingest. Groups and subgroups are used to organise data and define shared properties like dimensions and running frequency.

The extension we're builing uses 3 WMI Queries:
```sql
SELECT Name, PercentProcessorTime, PercentIdleTime, PercentUserTime FROM Win32_PerfFormattedData_PerfOS_Processor WHERE Name LIKE '_Total'
```
* extracts CPU Usage, User CPU, and Idle CPU for each of the host's processors (by split by cpu id).

```sql
SELECT Name, BytesTotalPersec, BytesReceivedPersec, BytesSentPersec FROM Win32_PerfFormattedData_Tcpip_NetworkAdapter
```
* extracts the Total, Sent, and Received Bytes per second for each Network Adapter running on the Host

```sql
SELECT Name, BytesTotalPersec, BytesReceivedPersec, BytesSentPersec FROM Win32_PerfFormattedData_Tcpip_NetworkInterface
```
* extracts the Total, Sent, and Received Bytes per second for each Network Interface running on the Host

**Metric best practice**

* Prefixing your metric keys with the name of the extension ensures no clashes will happen in Dynatrace. For this exercise, prefix each metric key with `custom.demo.host-observability`

**Other tips**

* You can identify the Host running the extension through the `this:device.host` passed as a dimension value
* You can add dimensions that are fixed strings using the prefix `const:`


## Tasks
1. Add the `wmi` section to your `extension.yaml`
2. Create two groups called `Host` and `Network` which run every 1 min. Both groups should have a dimension which identifies the host running the extension.
3. Create a subgroup for each WMI query given above and map the columns retrieved to metrics and dimensions (*hint: each query extracts 3 metrics and 1 dimension*)
4. Add a dimension called `network.type` which takes the value "Adapter" or "Interface" depending on the WMI query
5. Package a new version of your extension and upload it
6. Configure it to monitor your Windows host
7. Give it a minute, and validate metric collection.

## Results
You have completed the exercise when all your 6 metrics show up in the "Metrics" menu (*hint: filter by text `custom.demo`*)

![result](img/result.png)

---
## [Next exercise ▶](../3_Metric-Metadata/README.md)

#### [◀ Previous exercise](../1_Basic-Extension/README.md)
