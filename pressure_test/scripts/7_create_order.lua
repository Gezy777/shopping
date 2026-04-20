-- create_order.lua
-- 创建订单压测脚本

wrk.headers["Content-Type"] = "application/json"
wrk.headers["Authorization"] = "Bearer test-token-123456"

product_count = 1000

function request()
   product_id = math.random(1, product_count)
   quantity = math.random(1, 3)

   local body = string.format([[{
      "items": [
         {
            "product_id": %d,
            "quantity": %d,
            "price": 99.99
         }
      ],
      "address": {
         "name": "Test User",
         "phone": "13800138000",
         "address": "Test Address",
         "city": "Beijing",
         "district": "Haidian"
      },
      "payment_method": "alipay"
   }]], product_id, quantity)

   return wrk.format("POST", "/api/order/create", wrk.headers, body)
end
