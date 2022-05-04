import os

from omniversion import Data
from omniversion.samples import show_dashboard

data = Data(os.path.join(os.path.dirname(__file__), "test_omniversion/vectors"))

show_dashboard(data)
