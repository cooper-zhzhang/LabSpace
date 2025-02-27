redis-cli -p 6378 --eval key.lua  key1 value1
redis-cli -p 6378 SCRIPT LOAD "`cat key.lua`
redis-cli -p 6378 evalsha 6d9454f194be48c0c4a948abbcdbb8f64ed5dc46 1 k21 233
