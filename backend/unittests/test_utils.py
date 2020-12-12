from app import utils
import unittest

class TestUtils(unittest.TestCase):
  def test_get_film_id(self):
    res = utils.get_film_id('t01')
    self.assertEqual(len(res), 10)

  def test_get_film_id(self):
    res = utils.get_person_id('nm0001')
    self.assertEqual(len(res), 14)
