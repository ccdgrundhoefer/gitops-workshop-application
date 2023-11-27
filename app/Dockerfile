FROM registry.access.redhat.com/ubi8/python-39:1-154 as builder

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY app.py .

COPY templates templates

EXPOSE 5000

CMD ["python3", "-m", "app"]
