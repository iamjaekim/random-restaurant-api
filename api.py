#!flask/bin/python
from flask import Flask, abort, jsonify, request, make_response
from flask_cors import CORS
import requests
import json
import zipcodes
import random
import os

api_key=os.environ['YELP'] 
app = Flask(__name__)
CORS (app)

def zipValidation (zipCode) :
    #print (zipcodes.is_real (zipCode))
    return zipcodes.is_real (str(zipCode))

@app.route('/')
def index():
    return "Hello, World!"

@app.route('/api/stores/<int:zipCode>', methods=['GET','OPTIONS'])
def get_restaurants(zipCode):

    if (zipValidation(zipCode) is not True):
        abort(404,description="invalid zipcode")
    else:
        input = {'location':zipCode, 'open_now':'true'}
        try:   
            api_key = api_key
            headers = {'Authorization': 'Bearer %s' % api_key}
            yelp_url='https://api.yelp.com/v3/businesses/search'
            req=requests.get(yelp_url, params=input, headers=headers)
            load = json.loads(req.text)
            picked_biz = load['businesses']
            return jsonify(picked_biz) 
        except Exception :
            abort(404,description="Yelp API Error")

@app.route('/api/stores/single/<string:storeId>', methods=['GET','OPTIONS'])
def get_restaurant(storeId):
    input = {'id':storeId}
    print (input)
    try:   
        api_key = api_key
        headers = {'Authorization': 'Bearer %s' % api_key}
        yelp_url='https://api.yelp.com/v3/businesses/'+storeId
        req=requests.get(yelp_url, headers=headers)
        load = json.loads(req.text)
        return jsonify(load) 
    except Exception :
        abort(404,description="Yelp API Error")

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 5000))
    app.run(host='0.0.0.0', port=port)