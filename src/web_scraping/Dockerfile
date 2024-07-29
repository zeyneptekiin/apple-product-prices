FROM python:3.9-slim

RUN apt-get update && apt-get install -y cron

WORKDIR /app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .
COPY cron_file /etc/cron.d/cron_file

RUN chmod 0644 /etc/cron.d/cron_file && \
    crontab /etc/cron.d/cron_file

RUN touch /var/log/cron.log

CMD ["sh", "-c", "cron -f"]
