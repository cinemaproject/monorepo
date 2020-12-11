from db.connection import instance as db
from sqlalchemy.sql import text
from utils import get_film_id, get_person_id


def get_film_by_id(id):
  statement = text(
      "select * from films where id = '" + get_film_id(id) + "';")
  return db.execute(statement)


def get_related_people(film_id):
  statement = text(
      "select * from people where id in (select person_id from relations where film_id = '" + get_film_id(film_id) + "');")
  return db.execute(statement)


def get_person_by_id(id):
  statement = text(
      "select * from people where id = '" + get_person_id(id) + "';")
  return db.execute(statement)


def get_related_films(person_id):
  statement = text(
      "select * from films where id in (select film_id from relations where person_id = '" + get_person_id(person_id) + "');")
  return db.execute(statement)
