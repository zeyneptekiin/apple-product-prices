FROM python:3.9-slim

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      cron \
      gcc \
      python3-dev \
 && pip install --upgrade pip

COPY requirements.txt .
RUN pip install -r requirements.txt

WORKDIR /app
COPY . .

RUN echo "*/2 * * * * python /app/src/web_scraping/products_scraper.py >> /var/log/cron.log 2>&1" > /etc/cron.d/cronjob \
 && chmod 0644 /etc/cron.d/cronjob \
 && crontab /etc/cron.d/cronjob

CMD ["cron", "-f"]
