import time
import random
import requests
import argparse
import os

BACKEND_URL = "http://localhost:8080/events"

USERS = ["u1", "u2", "u3", "u4"]

SOURCE_IPS = ["10.0.0.1", "10.0.0.2", "10.0.0.3"]
DESTINATION_IPS = ["172.16.0.5", "172.16.0.10", "192.168.1.20"]

SOURCE_NAMES = ["laptop-01", "desktop-02", "mobile-03"]
DESTINATION_NAMES = ["server-auth", "server-db", "server-api"]

PROTOCOLS = ["TCP", "UDP", "HTTP", "HTTPS"]

STATUSES = ["SUCCESS"] * 7 + ["FAIL"] * 3

EVENTS_PER_SECOND = 5
BURST_MODE = False


def generate_event(process_id, user_id=None, status=None):
    """ Generate a single network flow event """

    event = {
        "processId": process_id,
        "userId": user_id if user_id else random.choice(USERS),
        "sourceIp": random.choice(SOURCE_IPS),
        "destinationIp": random.choice(DESTINATION_IPS),
        "sourceName": random.choice(SOURCE_NAMES),
        "destinationName": random.choice(DESTINATION_NAMES),
        "protocol": random.choice(PROTOCOLS),
        "status": status if status else random.choice(STATUSES),
        "timestamp": int(time.time() * 1000)
    }
    return event


def send_event(event):
    """ POST event to backend """
    try:
        resp = requests.post(BACKEND_URL, json=event, timeout=1)
        print(f"> {event} | status={resp.status_code}")
    except Exception as e:
        print(f"!! Error sending event: {e}")


def burst_fail_sequence(process_id):
    """ Send 3 FAILED events for same user """
    user = random.choice(USERS)
    for _ in range(3):
        event = generate_event(process_id, user_id=user, status="FAIL")
        send_event(event)
        time.sleep(0.1)


def normal_event_sequence(process_id):
    """ Send normal random event """
    event = generate_event(process_id)
    send_event(event)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--id", help="Unique process ID for this generator", default=None)
    args = parser.parse_args()

    # If no ID is given, fallback to OS PID
    process_id = args.id or f"proc-{os.getpid()}"

    print("Flow generator started...")
    print(f"Process ID → {process_id}")
    print(f"Backend → {BACKEND_URL}")
    print(f"Rate → {EVENTS_PER_SECOND} events/sec")
    print(f"Burst Mode → {BURST_MODE}")

    while True:
        start = time.time()

        if BURST_MODE and random.random() < 0.2:
            burst_fail_sequence(process_id)
        else:
            normal_event_sequence(process_id)

        elapsed = time.time() - start
        sleep_time = max(0, (1.0 / EVENTS_PER_SECOND) - elapsed)
        time.sleep(sleep_time)


if __name__ == "__main__":
    main()
