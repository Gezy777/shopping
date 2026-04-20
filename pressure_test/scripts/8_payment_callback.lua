-- payment_callback.lua
-- 支付回调压测脚本

wrk.headers["Content-Type"] = "application/json"

counter = 0

function request()
   counter = counter + 1
   order_id = string.format("ORD%d%06d", os.time(), counter % 1000000)
   payment_id = string.format("PAY%d%06d", os.time(), math.random(1, 1000000))

   local body = string.format([[{
      "order_id": "%s",
      "payment_id": "%s",
      "status": "success",
      "amount": 99.99,
      "paid_at": %d
   }]], order_id, payment_id, os.time())

   return wrk.format("POST", "/api/payment/callback", wrk.headers, body)
end
