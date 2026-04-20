-- user_login.lua
-- 用户登录压测脚本

wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

counter = 0
user_count = 10000

function request()
   counter = counter + 1
   user_id = counter % user_count + 1

   -- GoMall实际登录接口使用表单数据
   local body = string.format("email=user%d@gomall.com&password=password123", user_id)

   return wrk.format("POST", "/auth/login", wrk.headers, body)
end
