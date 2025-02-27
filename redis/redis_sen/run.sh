cd redis/redis_80 
redis-server 6380.conf &
cd .. 

cd redis/redis_81
redis-server 6381.conf &
cd ..

cd redis/redis_82
redis-server 6382.conf &
cd ..

cd redis/redis_83
redis-server 6383.conf &
cd ..

cd redis/redis_84
redis-server 6384.conf &
cd ..

cd redis/redis_85
redis-server 6385.conf &
cd ..

cd redis/redis_86
redis-server 6386.conf &
cd ..

redis-server sentinel/sentinel.conf --sentinel&
redis-server sentinel/sentinel2.conf --sentinel&
redis-server sentinel/sentinel3.conf --sentinel&
#redis-server sentinel/sentinel.conf &


# redis-server 127.0.0.1:6380

# pkill -f "redis-server 127.0.0.1:6380"
