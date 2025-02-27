# kill -9 `ps -elf | grep "redis" | awk '{print $4}' `

redis-cli -h 127.0.0.1 -p 6381 shutdown
redis-cli -h 127.0.0.1 -p 6382 shutdown
redis-cli -h 127.0.0.1 -p 6383 shutdown

redis-cli -h 127.0.0.1 -p 6480 shutdown
redis-cli -h 127.0.0.1 -p 6481 shutdown
redis-cli -h 127.0.0.1 -p 6482 shutdown
