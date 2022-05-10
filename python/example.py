import os

from omniversion import Omniversion, show_dashboard

data = Omniversion(os.path.join(os.path.dirname(__file__), "test_omniversion/vectors"))

show_dashboard(data)
