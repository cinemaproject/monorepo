FROM clearlinux/python

RUN mkdir -p /workspace
WORKDIR /workspace

ADD app ./app
ADD static ./static
ADD templates ./templates
COPY embeddings.pickle ./
COPY keys_dict.json ./
COPY requirements.txt ./

RUN swupd bundle-add opencv-python
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

CMD ["python", "/workspace/app/main.py"]
