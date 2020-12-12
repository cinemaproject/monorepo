from app import utils


def test_get_film_id():
        res = utils.get_film_id('t01')
        assert len(res) == 10


def test_get_person_id():
        res = utils.get_person_id('nm01')
        assert len(res) == 14
