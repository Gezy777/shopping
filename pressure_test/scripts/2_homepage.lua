-- homepage.lua
-- 首页压测脚本

function request()
   return wrk.format("GET", "/")
end
