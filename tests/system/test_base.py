from gelfbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Gelfbeat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        gelfbeat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("gelfbeat is running"))
        exit_code = gelfbeat_proc.kill_and_wait()
        assert exit_code == 0
