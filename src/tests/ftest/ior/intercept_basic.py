#!/usr/bin/python
"""
  (C) Copyright 2019-2022 Intel Corporation.

  SPDX-License-Identifier: BSD-2-Clause-Patent
"""

import os
from ior_test_base import IorTestBase
from ior_utils import IorCommand, IorMetrics
from general_utils import percent_change


class IorInterceptBasic(IorTestBase):
    # pylint: disable=too-many-ancestors
    """Test class Description: Verify IOR performance with DFUSE + IL is similar to DFS
                               for a single server and single client node.

    :avocado: recursive
    """

    def test_ior_intercept(self):
        """Jira ID: DAOS-3498.

        Test Description:
            Verify IOR performance with DFUSE + IL is similar to DFS.

        Use case:
            Run IOR write + read with DFS.
            Run IOR write + read with DFUSE + IL.
            Verify performance with DFUSE + IL is similar to DFS.

        :avocado: tags=all,daily_regression
        :avocado: tags=hw,small
        :avocado: tags=daosio,dfuse,il,ior,ior_intercept
        :avocado: tags=ior_intercept_basic
        """
        # Run IOR with DFS
        self.ior_cmd.api.update("DFS")
        dfs_out = self.run_ior_with_pool(fail_on_warning=self.log.info)
        dfs_perf = IorCommand.get_ior_metrics(dfs_out)

        # Run IOR with dfuse + IL
        self.ior_cmd.api.update("POSIX")
        dfuse_out = self.run_ior_with_pool(
            intercept=os.path.join(self.prefix, 'lib64', 'libioil.so'),
            fail_on_warning=self.log.info)
        dfuse_perf = IorCommand.get_ior_metrics(dfuse_out)

        # Index of each metric
        max_mib = int(IorMetrics.Max_MiB)
        mean_mib = int(IorMetrics.Mean_MiB)

        # Write and read performance thresholds
        write_x = self.params.get("write_x", self.ior_cmd.namespace, None)
        read_x = self.params.get("read_x", self.ior_cmd.namespace, None)
        if write_x is None or read_x is None:
            self.fail("Failed to get write_x and read_x from config")

        # Verify write performance
        actual_write_x = abs(percent_change(dfs_perf[0][max_mib], dfuse_perf[0][max_mib]))
        self.log.info("Max Write diff = %s", actual_write_x)
        self.assertLessEqual(actual_write_x, write_x, "Max Write outside performance threshold")
        actual_write_x = abs(percent_change(dfs_perf[0][mean_mib], dfuse_perf[0][mean_mib]))
        self.log.info("Mean Write diff = %s", actual_write_x)
        self.assertLessEqual(actual_write_x, write_x, "Mean Write outside performance threshold")

        actual_read_x = abs(percent_change(dfs_perf[1][max_mib], dfuse_perf[1][max_mib]))
        self.log.info("Max Read diff = %s", actual_read_x)
        self.assertLessEqual(actual_read_x, write_x, "Max Read outside performance threshold")
        actual_read_x = abs(percent_change(dfs_perf[1][mean_mib], dfuse_perf[1][mean_mib]))
        self.log.info("Mean Read diff = %s", actual_read_x)
        self.assertLessEqual(actual_read_x, write_x, "Mean Read outside performance threshold")
