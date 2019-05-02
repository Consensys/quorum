## Using Splunk
Tessera logs can be interpreted by Splunk to allow for monitoring and analysis.  The general steps to set up Splunk monitoring for a network of Tessera nodes are:

1. If one does not already exist, set up a central Splunk instance (a Receiver) on a separate host.
1. Configure the Tessera hosts to forward their logging info to the Receiver by:
    1. Providing Logback configuration to Tessera as a CLI arg on start-up to specify the format of the logging output (e.g. save to a file).  
      This is achieved by providing an XML/Groovy config file defining the logging-level and Logback Appenders to use, for example:
        ``` xml
        <?xml version="1.0" encoding="UTF-8"?>
           <configuration>            
               <appender name="FILE" class="ch.qos.logback.core.FileAppender">
                   <file>/path/to/file.log</file>
                   <encoder>
                       <pattern>%d{HH:mm:ss.SSS} [%thread] %-5level %logger{36} - %msg%n</pattern>
                   </encoder>
               </appender>    
               
               <logger name="org.glassfish.jersey.internal.inject.Providers" level="ERROR" />
               <logger name="org.hibernate.validator.internal.util.Version" level="ERROR" />
               <logger name="org.hibernate.validator.internal.engine.ConfigurationImpl" level="ERROR" />

               <root level="INFO">
                   <appender-ref ref="FILE"/>
               </root>
           </configuration>
        ```
        
        Further information can be found in the [Logback documentation](https://logback.qos.ch/manual/configuration.html#syntax).
        
    1. Set up Splunk Universal Forwarders on each Tessera host.  These are lightweight Splunk clients that will be used to collect and pass logging data to the central Splunk instance for analysis.
    1. Set up the central Splunk instance to listen and receive logging data from the Universal Forwarders


Further information about setting up Splunk and Universal Forwarders can be found in the Splunk documentation.  The following pages are a good starting point:

* [Consolidate data from multiple hosts](http://docs.splunk.com/Documentation/Forwarder/7.1.2/Forwarder/Consolidatedatafrommultiplehosts)
* [Set up the Universal Forwarder](http://docs.splunk.com/Documentation/Splunk/7.1.2/Forwarding/EnableforwardingonaSplunkEnterpriseinstance#Set_up_the_universal_forwarder)
* [Configure the Universal Forwarder](http://docs.splunk.com/Documentation/Forwarder/7.1.2/Forwarder/Configuretheuniversalforwarder)
*  [Enable a receiver](http://docs.splunk.com/Documentation/Forwarder/7.1.2/Forwarder/Enableareceiver)


## Jersey Web Server Metrics
Simple Jersey web server metrics for a Tessera node can be monitored if desired.  Tessera can store this performance data in a time-series database.  Two open-source database options are available for use, depending on your particular use-case:

* [InfluxDB](https://www.influxdata.com/time-series-platform/influxdb/): For 'push'-style data transmission 
* [Prometheus](https://prometheus.io/): For 'pull'-style data transmission

To set up monitoring requires the installation and configuration of one of these database offerings.  Both databases integrate well with the open source metrics dashboard editor [Grafana](https://grafana.com/) to allow for easy creation of dashboards to visualise the data being captured from Tessera. 

### Using InfluxDB
The [InfluxDB documentation](https://docs.influxdata.com/influxdb) provides all the information needed to get InfluxDB setup and ready to integrate with Tessera.  A summary of the steps is as follows:

1. Download and install InfluxDB
1. Create an InfluxDB database
1. Add configuration details to the `server` section of your Tessera config file to allow Tessera to post metrics data to the InfluxDB host.  An example configuration using InfluxDB's default hostName and port is (truncated for clarity):
    ```json
    "server": {
        "influxConfig": {
            "port": 8086,
            "hostName": "http://localhost",
            "dbName": "tessera_demo",
            "pushIntervalInSecs": 60
        }
    }
    ```
    With `influxConfig` provided, Tessera will collect metrics data and push it to the InfluxDB service periodically based on the value set for `pushIntervalInSecs`
1. You can use the `influx` CLI to query the database and view the data that is being stored

### Using Prometheus
The [Prometheus documentation](https://prometheus.io/docs/) provides all the information needed to get Prometheus setup and ready to integrate with Tessera.  A summary of the steps is as follows:

1. Download and install Prometheus
1. Configure `prometheus.yml` to give the Prometheus instance the necessary information to pull metrics from each of the Tessera nodes.  As Prometheus is pull-based, no additional config needs to be added to Tessera
1. Go to `localhost:9090` (or whatever host and port have been defined for the Prometheus instance) to see the Prometheus UI and view the data that is being stored

### Creating dashboards with Grafana
Once Tessera usage data is being stored in either InfluxDB or Prometheus, Grafana can be used to easily create dashboards to visualise that data.  The [Grafana documentation](http://docs.grafana.org/) provides all the information needed to set up a Grafana instance and integrate it with both of these time-series databases.  A summary of the steps is as follows:

1. Download and install Grafana
1. Create a new Data Source to connect Grafana with your database of choice 
1. Create a new Dashboard
1. Create charts and other elements for the dashboard
