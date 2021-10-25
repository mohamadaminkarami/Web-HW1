## Introduction
  Some test to indicate load effect on backends. 

## Both servers 100 users
using CORRECT_POST_REQUEST_PERCENTAGE=90 <br/>
using CORRECT_GET_REQUEST_PERCENTAGE = 90 <br/>
for 100 users (about 160 RPS) <br/>
requests that should fail (10 percent) <br/>
requests that actually fail (10 percent) <br/>
| NAME        | CPU           | MEM USAGE  |
| ------------- |:-------------:| -----:|
| web-hw1_client_1      | 4.5% | 12MiB |
| web-hw1_go_backend_1      |   <1%    |  8MiB  |
| web-hw1_node_backend_1 | 12%    |  27MiB   |
| web-hw1_redis_1 | 1%    |  increasing by time   |

![report graph image](https://github.com/mohamadaminkarami/Web-HW1/blob/main/locust/images/test_one_moderate.png)

## Both servers 1000 users
using CORRECT_POST_REQUEST_PERCENTAGE=90 <br/>
using CORRECT_GET_REQUEST_PERCENTAGE = 90 <br/>
for 100 users (about 125 RPS between 110-150) <br/>
requests that should fail (10 percent) <br/>
requests that actually fail (11 percent) <br/>
| NAME        | CPU           | MEM USAGE  |
| ------------- |:-------------:| -----:|
| web-hw1_client_1      | 4.5% | 27MiB |
| web-hw1_go_backend_1      |   <1%    |  8MiB  |
| web-hw1_node_backend_1 | 12%    |  40MiB   |
| web-hw1_redis_1 | 1%    |  increasing by time   |

with more load cpu usage didn't seem to change but memory usage increased for node backend and nginx. go seems to be operating more stable. <br/>

response time increased greatly from around 700 ms to 8000 ms.
you can see the jump in the report.

![report graph image](https://github.com/mohamadaminkarami/Web-HW1/blob/main/locust/images/test_harder_after_moderate.png)

## Both servers 100 users replication mode
2 node backends 3 go backends with same resources as before.
using CORRECT_POST_REQUEST_PERCENTAGE=90 <br/>
using CORRECT_GET_REQUEST_PERCENTAGE = 90 <br/>
for 100 users (about 300 RPS) <br/>
requests that should fail (10 percent) <br/>
requests that actually fail (10 percent) <br/>
| NAME        | CPU           | MEM USAGE  |
| ------------- |:-------------:| -----:|
| web-hw1_client_1      | 10% | 8MiB |
| web-hw1_go_backend_1      |   <1%    |  7MiB  |
| web-hw1_go_backend_2      |   <1%    |  7MiB  |
| web-hw1_go_backend_3      |   <1%    |  7MiB  |
| web-hw1_node_backend_1 | 12%    |  25MiB   |
| web-hw1_node_backend_2 | 12%    |  25MiB   |
| web-hw1_redis_1 | 1.5%    |  increasing by time   |

with more backend replicas cpu usage and memory usage stays nearly the same but the RPS handled is about twice and response time is a bit better. (600ms) <br/>

![report graph image](https://github.com/mohamadaminkarami/Web-HW1/blob/main/locust/images/test_replication.png)
