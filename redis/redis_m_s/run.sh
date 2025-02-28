cd master
redis-server 6380.conf &
cd -
cd slave_81
redis-server 6381.conf &
cd -
cd slave_81
redis-server 6382.conf &
cd -



# redis-server 127.0.0.1:6380

# pkill -f "redis-server 127.0.0.1:6380"
