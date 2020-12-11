from sqlalchemy.ext.declarative import DeclarativeMeta
import json


def get_film_id(id):
    """Pad film ID to be 10 characters in length"""
    return "{:<10}".format(str(id))


def get_person_id(id):
    """Pad person ID to be 10 characters in length"""
    return "{:<14}".format(str(id))


def resultproxy_to_dict(resultproxy):
    d, a = {}, []
    for rowproxy in resultproxy:
        # rowproxy.items() returns an array like [(key0, value0), (key1, value1)]
        for column, value in rowproxy.items():
            # build up the dictionary
            d = {**d, **{column: value}}
        a.append(d)
    return a
