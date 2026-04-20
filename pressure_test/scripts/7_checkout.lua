-- checkout.lua
-- 结算流程压测脚本

wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

product_count = 1000

function request()
   product_id = math.random(1, product_count)
   quantity = math.random(1, 3)

   -- 结算请求参数（模拟表单提交）
   local body = string.format(
      "email=test%d@gomall.com&firstname=Test&lastname=User&country=CN&zipcode=100000&city=Beijing&province=Beijing&street=Test+Street&cardNum=4111111111111111&expirationYear=2027&expirationMonth=12&cvv=123",
      product_id
   )

   return wrk.format("POST", "/checkout/waiting", wrk.headers, body)
end
