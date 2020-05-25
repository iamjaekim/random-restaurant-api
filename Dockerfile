FROM python:3.8

ADD . /api 
WORKDIR /api
RUN pip install -r requirements.txt
EXPOSE 5000
CMD python api.py