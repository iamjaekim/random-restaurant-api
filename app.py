#!flask/bin/python
from flask import Flask, request, send_from_directory, render_template
from flask_cors import CORS

import os

web = Flask(__name__, static_url_path='', static_folder='./frontend')
CORS (web)

@web.route('/')
def index():
    return send_from_directory('frontend','index.html')

@web.route('/css/style.css')
def css():
    return send_from_directory('frontend/css','style.css')

@web.route('/js/yelpRecommender.js')
def js():
    return send_from_directory('frontend/js','yelpRecommender.js')

@web.route('/restaurant.html')
def restaurant():
    return send_from_directory('frontend','restaurant.html')

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 3000))
    web.run(host='0.0.0.0', port=port)