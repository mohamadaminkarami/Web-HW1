import hashlib
import random
import string

from locust import HttpUser, task


def get_random_string(length=8):
    return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))


def hash_string(string):
    m = hashlib.sha256()
    m.update(string.encode())
    return m.hexdigest()


class CustomUser(HttpUser):
    # wait_time = between(1, 5)

    @task
    def get_sha_node(self):
        encoded = hash_string(get_random_string())
        self.client.get(f"/node/sha?encoded={encoded}")

    @task(10)
    def set_sha_node(self):
        decoded = get_random_string()
        self.client.post("/node/sha", {
            "raw_string": decoded
        })

    @task
    def get_sha_go(self):
        encoded = hash_string(get_random_string())
        self.client.get(f"/go/sha?encoded={encoded}")

    @task(10)
    def set_sha_go(self):
        decoded = get_random_string()
        self.client.post("/go/sha", {
            "decoded": decoded
        })
