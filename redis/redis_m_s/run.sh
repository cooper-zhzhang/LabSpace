redis-server 6380.conf &
redis-server slave_81/6381.conf &
redis-server slave_82/6382.conf &



# redis-server 127.0.0.1:6380

# pkill -f "redis-server 127.0.0.1:6380"