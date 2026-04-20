-- ping_benchmark.lua
-- Ping接口基准压测脚本

function request()
   return wrk.format("GET", "/ping")
end
