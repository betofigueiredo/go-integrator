from src.utils.datetimes import Datetimes
from src.utils.general import General
from src.utils.ids import Ids


class Utils:
    def __init__(self) -> None:
        self.datetimes = Datetimes()
        self.general = General()
        self.ids = Ids()


utils = Utils()
