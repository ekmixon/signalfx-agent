# Install to Linux using a tar.gz File

:warning: **SignalFx Smart Agent is deprecated. For details, see the [Deprecation Notice](./smartagent-deprecation-notice.md)** :warning:

Install the Smart Agent to Linux host using a compressed
`tar` file.

## Prerequisites

* Kernel version 3.2 or higher
* cap_dac_read_search and cap_sys_ptrace Linux capabilities
* Terminal or a similar command-line interface application
* Program that can uncompress and unbundle a `.tar.gz` file.

## Install the Smart Agent using a tar.gz file

1. Remove collector services such as `collectd`

2. Get the latest `.tar.gz` standalone installation package by navigating to
   [Smart Agent releases](https://github.com/signalfx/signalfx-agent/releases)
   and downloading the following file:

   ```
   signalfx-agent-<latest_version>.tar.gz
   ```

   For example, if the latest version is **5.1.6**, perform the following steps:

   1. In the **releases** section, find the section called **v5.1.6**.
   2. In the **Assets** section, click `signalfx-agent-5.1.6.tar.gz`
   3. The file starts downloading.

3. To uncompress the package, run the following command:

   ```
   tar xzf signalfx-agent-<latest_version>.tar.gz
   ```

   The package expands into the `signalfx-agent` directory.

4. Navigate to the `signalfx-agent` directory:

   ```
   cd signalfx-agent
   ```


5. To ensure that the binaries in the install files use the correct loader for your host, run
the following command:

   ```
   bin/patch-interpreter $(pwd)
   ```

## Configure the installation

Create a configuration file for the agent:

1. In a text editor, create the file called signalfx-agent/agent-config.yaml.
2. In the file, add your hostname and port number:

   ```
   internalStatusHost: <local_hostname>
   internalStatusPort: <local_port>
   collectd:
     configDir: <collectd_config_dir>
   ```

3. Save the file.

> The Smart Agent collects metrics based on the settings in
> `agent-config.yaml`. The `internalStatusHost` and `internalStatusPort`
> properties specify the host and port number of the host that's running the Smart Agent.
> The `collectd.configDir` property specifies the directory where the Smart Agent writes
> `collectd` configuration files.

### Start the Smart Agent

To start the Smart Agent, run this command:

```
signalfx-agent/bin/signalfx-agent -config signalfx-agent/agent-config.yaml > <log_file>
```

> The default log output for the Smart Agent goes to `STDOUT` and `STDERR`.
> To persist log output, direct the log output to `<log_file>`.

### Verify the Smart Agent

To verify that your installation and configuration, perform these steps:

* For infrastructure monitoring, perform these steps:
  1. In SignalFx UI, open the **Infrastructure** built-in dashboard
  2. In the override bar at the top of the back, select **Choose a host**. Select one of your hosts from the dropdown.

  The charts display metrics from your infrastructure.

  To learn more, see [Built-In Dashboards and Charts](https://docs.signalfx.com/en/latest/getting-started/built-in-content/built-in-dashboards.html).

* For Kubernetes monitoring, perform these steps:
  1. In SignalFx, from the main menu select **Infrastructure** > **Kubernetes Navigator** > **Cluster map**.
  2. In the cluster display, find the cluster you installed.
  3. Click the magnification icon to view the nodes in the cluster.

  The detail pane displays details of your cluster and nodes.

  To learn more, see [Getting Around the Kubernetes Navigator](https://docs.signalfx.com/en/latest/integrations/kubernetes/get-around-k8s-navigator.html).

* For APM monitoring, learn how to install, configure, and verify the Smart Agent for Microservices APM (**µAPM**). See
  [Get started with SignalFx µAPM](https://docs.signalfx.com/en/latest/apm/apm-getting-started/apm-index.html).

### Uninstall the Smart Agent

To uninstall the Smart Agent, stop the `signalfx-agent` process (if running),
backup any files as necessary, and delete the `signalfx-agent` directory.
