-- product_detail.lua
-- 商品详情页压测脚本

product_count = 1000

function request()
   product_id = math.random(1, product_count)
   path = string.format("/product?id=%d", product_id)
   return wrk.format("GET", path)
end
