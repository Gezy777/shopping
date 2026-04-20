-- cart_operations.lua
-- 购物车操作压测脚本

wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

product_count = 1000

function request()
   product_id = math.random(1, product_count)
   quantity = math.random(1, 5)

   -- 随机操作: 1=添加购物车 2=查看购物车
   op = math.random(1, 2)

   if op == 1 then
      -- 添加到购物车
      local body = string.format("productId=%d&productNum=%d", product_id, quantity)
      return wrk.format("POST", "/cart", wrk.headers, body)
   else
      -- 查看购物车
      return wrk.format("GET", "/cart", wrk.headers)
   end
end
