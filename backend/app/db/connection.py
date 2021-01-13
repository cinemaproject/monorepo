import config

import sqlalchemy
from sqlalchemy import MetaData

instance = sqlalchemy.create_engine(
    config.postgres_url, connect_args={"connect_timeout": 100000}
)


def connect(max_retries, retry_count=0):
    try:
        instance.connect()
    except BaseException as err:
        retry_count += 1
        if retry_count > max_retries:
            raise err
        else:
            connect(max_retries, retry_count)


connect(100)

metadata = MetaData()
