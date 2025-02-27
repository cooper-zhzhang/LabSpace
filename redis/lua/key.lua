local key = KEYS[1]
local value = ARGV[1]
if redis.call('get', key) == false then
    redis.call('set', key, value)
    return 1
else
    return 0
end
