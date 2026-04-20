-- product_search.lua
-- 商品搜索/列表压测脚本

function request()
   -- 随机搜索关键词
   keywords = {"laptop", "phone", "headphones", "keyboard", "mouse", "monitor", "camera", "tablet"}
   query = keywords[math.random(1, #keywords)]
   path = string.format("/search?query=%s", query)
   return wrk.format("GET", path)
end
