FROM python:3.6-buster

USER root

RUN apt-get update &\
    apt-get install -y git libpq-dev libcurl-dev

RUN pip install scikit-build

RUN rm -rf /opt/test
RUN git clone https://github.com/cinemaproject/backend.git /opt/test/

WORKDIR /opt/test/

# RUN pip3 install pycopy-pwd
RUN pip install -r requirements.txt

ENTRYPOINT python app.py
