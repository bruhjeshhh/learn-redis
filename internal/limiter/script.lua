local now = tonumber(redis.call("TIME")[1])
local user = KEYS[1]

local capacity = 6
local refill_rate = 1

local current_tokens = tonumber(redis.call("HGET", user, "tokens"))
local last_refill = tonumber(redis.call("HGET", user, "last_refill"))

-- First request: initialize the bucket
if current_tokens == nil then
    current_tokens = capacity
    last_refill = now
end

-- Calculate how many tokens have accumulated
local elapsed = now - last_refill
current_tokens = math.min(
    current_tokens + elapsed * refill_rate,
    capacity
)

-- Try to consume a token
if current_tokens >= 1 then
    current_tokens = current_tokens - 1

    redis.call("HSET",user,"tokens", current_tokens,"last_refill", now)

    return 1
end

-- No token available
redis.call(
    "HSET",user,"tokens", current_tokens,"last_refill", now
)

return 0
