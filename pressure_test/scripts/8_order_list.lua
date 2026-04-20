-- order_list.lua
-- 订单列表压测脚本

function request()
   return wrk.format("GET", "/order")
end
