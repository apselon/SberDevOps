FROM python:3.9-alpine

WORKDIR /CalcApp

ADD sources .

CMD ["python", "-c", "\"print(__import__(\\"calc\\").add2(2,2))\""]
