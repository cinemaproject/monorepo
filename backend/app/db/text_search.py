from db.connection import instance as db
from sqlalchemy.sql import text


def get_people_by_name(name):
    statement = text(
        "select * from people where lower(primary_name) LIKE lower('%" + name + "%');")
    return db.execute(statement)


def get_films_by_title(name):
    statement = text(
        "select * from films where lower(title) LIKE lower('%" + name + "%');")
    return db.execute(statement)
