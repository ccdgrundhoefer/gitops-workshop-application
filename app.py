import os
import requests
import time

def request_url():
    # Hole die URL aus der Umgebungsvariable
    url = os.getenv("TARGET_URL")
    if url is None:
        print("Umgebungsvariable TARGET_URL nicht gesetzt.")
        return
    
    try:
        response = requests.get(url)
        print(f"Statuscode: {response.status_code}")
    except Exception as e:
        print(f"Fehler beim Aufrufen der URL: {e}")

def main():
    # Intervall in Sekunden, wie oft die URL aufgerufen werden soll
    interval = 1  # Zum Beispiel alle 60 Sekunden
    
    while True:
        request_url()
        time.sleep(interval)

if __name__ == "__main__":
    main()
