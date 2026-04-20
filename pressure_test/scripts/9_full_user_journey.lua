-- full_user_journey.lua
-- 完整用户旅程压测脚本
-- 模拟: 首页 → 商品搜索 → 商品详情 → 登录 → 添加购物车 → 结算

wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

-- 用户池
user_count = 10000
product_count = 1000

-- 随机选择用户
local user_id = math.random(1, user_count)
local email = string.format("user%d@gomall.com", user_id)

-- 状态机阶段
local phase = "homepage"

-- 统计
local stats = {
   homepage = 0,
   search = 0,
   detail = 0,
   login = 0,
   cart = 0,
   checkout = 0,
   errors = 0
}

function request()
   if phase == "homepage" then
      -- 1. 访问首页
      phase = "search"
      return wrk.format("GET", "/")

   elseif phase == "search" then
      -- 2. 搜索商品
      keywords = {"laptop", "phone", "headphones", "keyboard", "mouse"}
      query = keywords[math.random(1, #keywords)]
      phase = "detail"
      return wrk.format("GET", string.format("/search?query=%s", query))

   elseif phase == "detail" then
      -- 3. 查看商品详情
      product_id = math.random(1, product_count)
      phase = "login"
      return wrk.format("GET", string.format("/product?id=%d", product_id))

   elseif phase == "login" then
      -- 4. 用户登录
      local body = string.format("email=%s&password=password123", email)
      phase = "cart"
      return wrk.format("POST", "/auth/login", wrk.headers, body)

   elseif phase == "cart" then
      -- 5. 添加购物车
      product_id = math.random(1, product_count)
      quantity = math.random(1, 3)
      local body = string.format("productId=%d&productNum=%d", product_id, quantity)
      phase = "checkout"
      return wrk.format("POST", "/cart", wrk.headers, body)

   elseif phase == "checkout" then
      -- 6. 结算
      local body = string.format(
         "email=%s&firstname=Test&lastname=User&country=CN&zipcode=100000&city=Beijing&province=Beijing&street=Test+Street&cardNum=4111111111111111&expirationYear=2027&expirationMonth=12&cvv=123",
         email
      )
      phase = "done"
      return wrk.format("POST", "/checkout/waiting", wrk.headers, body)
   end

   -- 完成后重新开始
   phase = "homepage"
   return wrk.format("GET", "/")
end

function response(status, headers, body)
   if status ~= 200 and status ~= 302 then
      stats.errors = stats.errors + 1
      return
   end

   if phase == "search" then
      stats.homepage = stats.homepage + 1
   elseif phase == "detail" then
      stats.search = stats.search + 1
   elseif phase == "login" then
      stats.detail = stats.detail + 1
   elseif phase == "cart" then
      stats.login = stats.login + 1
   elseif phase == "checkout" then
      stats.cart = stats.cart + 1
   elseif phase == "done" or phase == "homepage" then
      stats.checkout = stats.checkout + 1
   end
end

function done(summary, latency, errors)
   print("\n========== Full Journey Stats ==========")
   print(string.format("Homepage visits:   %d", stats.homepage))
   print(string.format("Search requests:   %d", stats.search))
   print(string.format("Product views:    %d", stats.detail))
   print(string.format("Login attempts:   %d", stats.login))
   print(string.format("Cart operations:  %d", stats.cart))
   print(string.format("Checkout attempts:%d", stats.checkout))
   print(string.format("Errors:           %d", stats.errors))
   print("========================================\n")
end
