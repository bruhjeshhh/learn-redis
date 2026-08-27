local now=os.time()
local window=now-60

redis.call("ZREMRANGEBYSCORE", KEYS[1], "-inf", window)
if redis.call("ZCARD",KEYS[1])>=10 then 
    return 0
end

redis.call("ZADD", KEYS[1],now,ARGV[1])
return 1


