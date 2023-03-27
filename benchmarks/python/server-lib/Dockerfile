FROM python:3.9-alpine
WORKDIR /python-docker

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY main.py .

CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]