FROM python:3.9-slim

RUN apt-get update && apt-get install -y cron

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

COPY cron_file /etc/cron.d/cron_file

RUN chmod 0644 /etc/cron.d/cron_file && \
    crontab /etc/cron.d/cron_file

RUN touch /var/log/cron.log

ENV PYTHONPATH /app/src

CMD ["sh", "-c", "/usr/local/bin/python3 /app/src/main.py && cron && tail -f /var/log/cron.log"]