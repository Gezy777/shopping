-- mixed_traffic.lua
-- 混合流量压测脚本
-- 按真实比例模拟: 60%浏览 + 20%详情 + 10%购物车 + 7%下单 + 3%其他

product_count = 1000

function request()
   -- 按比例选择场景
   r = math.random(1, 100)

   if r <= 60 then
      -- 60%: 首页或搜索
      if math.random(1, 2) == 1 then
         return wrk.format("GET", "/")
      else
         keywords = {"laptop", "phone", "headphones", "keyboard"}
         query = keywords[math.random(1, #keywords)]
         return wrk.format("GET", string.format("/search?query=%s", query))
      end

   elseif r <= 80 then
      -- 20%: 商品详情
      product_id = math.random(1, product_count)
      return wrk.format("GET", string.format("/product?id=%d", product_id))

   elseif r <= 90 then
      -- 10%: 购物车
      product_id = math.random(1, product_count)
      if math.random(1, 2) == 1 then
         -- 查看购物车
         return wrk.format("GET", "/cart")
      else
         -- 添加购物车
         body = string.format("productId=%d&productNum=%d", product_id, math.random(1, 3))
         return wrk.format("POST", "/cart", wrk.headers, body)
      end

   elseif r <= 97 then
      -- 7%: 结算
      email = string.format("user%d@gomall.com", math.random(1, 10000))
      body = string.format(
         "email=%s&firstname=Test&lastname=User&country=CN&zipcode=100000&city=Beijing&province=Beijing&street=Test+Street&cardNum=4111111111111111&expirationYear=2027&expirationMonth=12&cvv=123",
         email
      )
      return wrk.format("POST", "/checkout/waiting", wrk.headers, body)

   else
      -- 3%: 订单/其他
      if math.random(1, 2) == 1 then
         return wrk.format("GET", "/order")
      else
         return wrk.format("GET", "/checkout")
      end
   end
end
