from functools import partial as p
from pathlib import Path

import pytest

from tests.helpers.agent import Agent
from tests.helpers.assertions import has_any_metric_or_dim, has_log_message
from tests.helpers.kubernetes.utils import get_metrics
from tests.helpers.util import wait_for

pytestmark = [pytest.mark.collectd, pytest.mark.processes, pytest.mark.monitor_without_endpoints]

SCRIPT_DIR = Path(__file__).parent.resolve()


def test_processes():
    expected_metrics = get_metrics(SCRIPT_DIR)
    with Agent.run(
        """
    procPath: /proc
    monitors:
      - type: collectd/processes
        collectContextSwitch: true
        processMatch:
          collectd: ".*collectd.*"
    """
    ) as agent:
        assert wait_for(
            p(has_any_metric_or_dim, agent.fake_services, expected_metrics, None), timeout_seconds=60
        ), "timed out waiting for metrics and/or dimensions!"
        assert not has_log_message(agent.output.lower(), "error"), "error found in agent output!"
