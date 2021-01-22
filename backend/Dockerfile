FROM clearlinux/python:3.8.2

RUN mkdir -p /workspace
WORKDIR /workspace

ADD app ./app
COPY requirements.txt ./

RUN pip install --upgrade pip
RUN pip install wheel numpy
RUN pip install -r requirements.txt

CMD ["python", "/workspace/app/main.py"]
