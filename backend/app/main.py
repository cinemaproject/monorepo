import config
from db import text_search, items
from utils import resultproxy_to_dict
import os
from flask_cors import CORS, cross_origin
import insightface
import urllib
import urllib.request
import cv2
import json
import numpy as np
import uuid
import os
from face_searcher import face_searcher
from flask import Flask, request, render_template, send_file


if not os.path.exists(config.images_dir):
    os.makedirs(config.images_dir)


print("Preparing model")
model = insightface.app.FaceAnalysis(
    rec_name='arcface_r100_v1', det_name='retinaface_mnet025_v2')
ctx_id = -1
model.prepare(ctx_id=ctx_id, nms=0.4)
searcher = face_searcher()
app = Flask(__name__)


@app.route('/images/<path:path>')
@cross_origin()
def images(path):
    print(path)
    return send_file(os.path.join(config.images_dir, path), mimetype="image/JPG")


def url_to_image(url):
    resp = urllib.request.urlopen(url)
    image = np.asarray(bytearray(resp.read()), dtype="uint8")
    image = cv2.imdecode(image, cv2.IMREAD_COLOR)
    return image


@app.route('/people/search')
@cross_origin()
def search_people():
    name = request.args.get('name')
    people = resultproxy_to_dict(text_search.get_people_by_name(name))
    result = {'people': people}
    return json.dumps(result)


@app.route('/people/<id>')
@cross_origin()
def get_person(id):
    person = resultproxy_to_dict(items.get_person_by_id(id))
    films = resultproxy_to_dict(items.get_related_films(id))
    result = {'person': person, 'films': films}
    return json.dumps(result)


@app.route('/films/search')
@cross_origin()
def search_films():
    title = request.args.get('title')
    films = resultproxy_to_dict(text_search.get_films_by_title(title))
    result = {'films': films}
    return json.dumps(result)


@app.route('/films/<id>')
@cross_origin()
def get_film_info(id):
    film = resultproxy_to_dict(items.get_film_by_id(id))
    people = resultproxy_to_dict(items.get_related_people(id))
    result = {'film': film, 'people': people}
    return json.dumps(result)


@app.route('/actors_recognition', methods=['POST'])
@cross_origin()
def recognize_actors():
    image_undecoded = request.get_data()
    img = cv2.imdecode(np.frombuffer(image_undecoded, np.uint8), -1)
    img_out = img.copy()
    faces = model.get(img)
    faces_arr = []
    img_name = str(uuid.uuid4())
    for idx, face in enumerate(faces):
        face_d = {}
        face_d['index'] = idx
        face_d['age'] = face.age
        gender = 'Male'
        if face.gender == 0:
            gender = 'Female'
        face_d['gender'] = gender
        face_search_res = searcher.search_face(face.embedding, 1)
        face_d['search_results'] = face_search_res
        bboxes = face.bbox.astype(np.int)
        face_name = img_name + '_' + str(idx) + '.jpg'
        face_d['face_image_url'] = 'http://127.0.0.1:5000/images/' + face_name
        cv2.imwrite('images/' + face_name,
                    img[int(bboxes[1]):int(bboxes[3]), int(bboxes[0]):int(bboxes[2])])
        face_d['bbox'] = str(bboxes.flatten())
        img_out = cv2.rectangle(
            img_out, (bboxes[0], bboxes[1]), (bboxes[2], bboxes[3]), (0, 255, 0), 2)
        img_out = cv2.putText(img_out, str(face.age), (int(
            bboxes[0]-10), int(bboxes[1]-10)), cv2.FONT_HERSHEY_SIMPLEX, 1.1, (0, 255, 0), 2)
        img_out = cv2.putText(img_out, str(gender), (int(
            bboxes[0]+40), int(bboxes[1]-10)), cv2.FONT_HERSHEY_SIMPLEX, 1.1, (0, 0, 255), 2)
        img_out = cv2.putText(img_out, str(idx), (int(
            bboxes[0]-40), int(bboxes[1]-10)), cv2.FONT_HERSHEY_SIMPLEX, 1.1, (0, 0, 255), 2)
        faces_arr.append(face_d)
    cv2.imwrite(os.path.join(config.images_dir, img_name + '.jpg'), img_out)
    resp = {"faces": faces_arr,
            'result_image': 'http://127.0.0.1:5000/images/' + img_name + '.jpg'}
    return json.dumps(resp, ensure_ascii=False), 200


if __name__ == '__main__':
    print("Starting server on port 5000")
    app.run(host="0.0.0.0", port=5000)
