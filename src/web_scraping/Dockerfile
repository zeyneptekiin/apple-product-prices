FROM python:3.9-slim

RUN apt-get update && apt-get install -y cron

WORKDIR /app

COPY requirements.txt ./
COPY products_scraper.py ./
COPY my_cron_file /etc/cron.d/my_cron_file

RUN pip install --no-cache-dir -r requirements.txt

RUN chmod 0644 /etc/cron.d/my_cron_file && \
    crontab /etc/cron.d/my_cron_file

CMD ["sh", "-c", "cron -f"]
