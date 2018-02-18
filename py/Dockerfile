FROM ubuntu:latest

RUN apt-get update -y
RUN apt-get install -y python-pip python-dev build-essential

COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt
RUN python setup.py
ENTRYPOINT ["python"]
CMD ["main.py"]

EXPOSE 8002