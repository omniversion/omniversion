from omniversion import Omniversion
from .dashboard import show_dashboard

data = Omniversion("/tmp/omniversion")

show_dashboard(data)
