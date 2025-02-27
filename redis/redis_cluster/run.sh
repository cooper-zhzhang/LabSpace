cd master1
redis-server redis.conf &
cd .. 

cd master2
redis-server redis.conf &
cd ..

cd master3
redis-server redis.conf &
cd ..

cd slave1
redis-server redis.conf &
cd ..

cd slave2
redis-server redis.conf &
cd ..

cd slave3
redis-server redis.conf &
cd ..
