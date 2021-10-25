import hashlib
import random
import string

from locust import HttpUser, task

CORRECT_POST_REQUEST_PERCENTAGE = 90
CORRECT_GET_REQUEST_PERCENTAGE = 90
RAW_STRING_MAX_LENGTH = 20
RAW_STRING_MIN_VALID_LENGTH = 8
stored_hashes = []


def __get_valid_raw_string():
    string_length = random.randint(
        RAW_STRING_MIN_VALID_LENGTH, RAW_STRING_MAX_LENGTH)
    return __get_random_string(string_length)


def __get_invalid_raw_string():
    string_length = random.randint(0, RAW_STRING_MIN_VALID_LENGTH - 1)
    return __get_random_string(string_length)


def __get_random_string(length=8):
    return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))


def __get_stored_hash():
    return random.choice(stored_hashes)


def __get_not_stored_hash():
    return __get_hash_from_string(__get_valid_raw_string())


def __get_hash_from_string(string):
    m = hashlib.sha256()
    m.update(string.encode())
    return m.hexdigest()


def __get_random_percentage():
    return random.random() * 100


def get_random_hash():

    percentage = __get_random_percentage()

    return __get_stored_hash() if percentage < CORRECT_GET_REQUEST_PERCENTAGE and len(stored_hashes) > 0 else __get_not_stored_hash()


def get_random_raw_string():
    percentage = __get_random_percentage()

    return __get_valid_raw_string() if percentage < CORRECT_POST_REQUEST_PERCENTAGE else __get_invalid_raw_string()


class CustomUser(HttpUser):
    # wait_time = between(1, 5)

    @task
    def get_sha_node(self):
        encoded = get_random_hash()
        self.client.get(f"/node/sha256/?encoded={encoded}")

    @task(10)
    def set_sha_node(self):
        decoded = get_random_raw_string()
        data = {"raw_string": decoded}
        response = self.client.post("/node/sha256/", json=data)
        encoded = response.json().get("encoded", None)
        if encoded:
            stored_hashes.append(encoded)
    @task
    def get_sha_go(self):
        encoded = get_random_hash()
        self.client.get(f"/go/sha256/?encoded={encoded}")

    @task(10)
    def set_sha_go(self):
        decoded = get_random_raw_string()
        data = {"raw_string": decoded}
        response = self.client.post("/node/sha256/", json=data)
        encoded = response.json().get("encoded", None)
        if encoded:
            stored_hashes.append(encoded)
