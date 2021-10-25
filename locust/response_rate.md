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

report graph image: ./images/test_one_moderate.png

## Both servers 1000 users
using CORRECT_POST_REQUEST_PERCENTAGE=90 <br/>
using CORRECT_GET_REQUEST_PERCENTAGE = 90 <br/>
for 100 users (about 125 between 110-150) <br/>
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
report graph image: ./images/test_harder_after_moderate.png
