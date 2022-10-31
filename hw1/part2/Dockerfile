FROM python:3.9-alpine

WORKDIR /WebAppWithDB

ADD App/flask-db-demo/requirements.txt .
ADD App/flask-db-demo/run.py .
ADD App/flask-db-demo/app/ ./app
ADD App/flask-db-demo/migrations/ ./migrations

RUN                                                              \
    apk add postgresql-libs &&                                   \
    apk add --virtual .build-deps gcc musl-dev postgresql-dev && \
    pip3 install -r requirements.txt &&                                     \
    apk --purge del .build-deps

ENV DATABASE_URL="postgresql://adam:password@aidb/app_db"
CMD ["flask", "db", "upgrade"]
CMD ["python", "run.py"]
