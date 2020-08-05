# Quorum Profiling

[Quorum Profiling](https://github.com/QuorumEngineering/quorum-test) is a custom toolset which could be used to benchmark transaction throughput and network statistics on any existing network using the Jmeter & TPS monitoring tool profiles or could be used to spin up a entire quorum network from scratch in AWS and benchmark the network for TPS, CPU/Memory usage. The various scenarios of usage is explained [here](https://github.com/QuorumEngineering/quorum-test)

## Metrics Visualisation

The tool executes the stress test profile selected and then collects the following metrics:

 * CPU/memory usage metrics for both `Quorum` & `tessera`
 * Transaction & Block count
 * Transaction processing speed
 * `Jmeter` test execution metrics
 
 These metrics can be stored in an InfluxDB or Prometheus time-series database for analysis. Both databases integrate well with the open source dashboard editor Grafana to allow for easy creation of dashboards to visualise the data being captured from the profiling tool. Sample dashboards below:
 
### Sample Network Dashboard 

![Quorum Network Dashboard](images/quorumDashboard.jpeg) 
 
### Sample JMeter Dashboard

![JMeter Dashboard](images/quorumDashboard.jpeg) 
