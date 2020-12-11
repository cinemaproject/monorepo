import config

import sqlalchemy
from sqlalchemy import MetaData

instance = sqlalchemy.create_engine(config.postgres_url)
instance.connect()

metadata = MetaData()
