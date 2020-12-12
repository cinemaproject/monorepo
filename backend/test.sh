export PYTHONPATH=$(pwd):$PYTHONPATH

pytest unittests --doctest-modules --junitxml=junit/test-results.xml --cov=app --cov-report=xml --cov-report=html
