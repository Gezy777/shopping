#!/bin/bash
# run_all_tests.sh - 执行所有压测脚本

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
BASE_URL="${1:-http://localhost:8888}"
DURATION="${2:-30s}"
REPORT_DIR="$(dirname "$SCRIPT_DIR")/results/$(date '+%Y%m%d_%H%M%S')"

# 创建结果目录
mkdir -p "$REPORT_DIR"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  GoMall 全链路压测${NC}"
echo -e "${GREEN}  时间: $(date '+%Y-%m-%d %H:%M:%S')${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "目标地址: $BASE_URL"
echo "压测时长: $DURATION"
echo "结果目录: $REPORT_DIR"
echo ""

# 检查服务状态
check_service() {
    echo -e "${YELLOW}[检查]${NC} 检测服务状态..."
    if curl -s --connect-timeout 3 "$BASE_URL/ping" > /dev/null 2>&1; then
        echo -e "${GREEN}✓${NC} 服务正常"
    else
        echo -e "${RED}✗${NC} 服务无响应，是否启动？"
        exit 1
    fi
}

# 运行单项压测
run_test() {
    local name=$1
    local script=$2
    local threads=$3
    local connections=$4
    local output_file=$5

    echo ""
    echo -e "${YELLOW}[测试]${NC} $name"
    echo "----------------------------------------"

    wrk -t"$threads" -c"$connections" -d"$DURATION" --latency -s "$SCRIPT_DIR/$script" "$BASE_URL" 2>&1 | tee "$REPORT_DIR/$output_file"

    echo -e "${GREEN}[完成]${NC} $name"
}

# 执行检查
check_service

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  阶段1: 基础基准测试${NC}"
echo -e "${GREEN}========================================${NC}"

run_test "Ping接口基准" "1_ping_benchmark.lua" 8 200 "01_ping_benchmark.log"
run_test "首页压测" "2_homepage.lua" 8 200 "02_homepage.log"

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  阶段2: 核心链路压测${NC}"
echo -e "${GREEN}========================================${NC}"

run_test "商品搜索" "3_product_list.lua" 8 200 "03_product_search.log"
run_test "商品详情" "4_product_detail.lua" 8 200 "04_product_detail.log"
run_test "用户登录" "5_user_login.lua" 8 200 "05_user_login.log"
run_test "购物车操作" "6_cart_operations.lua" 8 200 "06_cart_operations.log"
run_test "结算流程" "7_checkout.lua" 8 200 "07_checkout.log"
run_test "订单列表" "8_order_list.lua" 8 200 "08_order_list.log"

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  阶段3: 全链路压测${NC}"
echo -e "${GREEN}========================================${NC}"

run_test "完整用户旅程" "9_full_user_journey.lua" 8 200 "09_full_journey.log"
run_test "混合流量压测" "10_mixed_traffic.lua" 8 200 "10_mixed_traffic.log"

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  阶段4: 高并发测试${NC}"
echo -e "${GREEN}========================================${NC}"

run_test "高并发Ping (500并发)" "1_ping_benchmark.lua" 16 500 "11_high_ping.log"
run_test "高并发首页 (300并发)" "2_homepage.lua" 16 300 "12_high_home.log"

echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  压测完成${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "结果目录: $REPORT_DIR"
echo "完成时间: $(date '+%Y-%m-%d %H:%M:%S')"

# 生成汇总报告
echo ""
echo "生成汇总报告..."
{
   echo "# GoMall 全链路压测汇总"
   echo ""
   echo "测试时间: $(date '+%Y-%m-%d %H:%M:%S')"
   echo "目标地址: $BASE_URL"
   echo ""
   echo "## 测试结果"
   echo ""
   ls -la "$REPORT_DIR"/*.log | while read f; do
      echo "### $(basename $f)"
      echo '```'
      head -20 "$f"
      echo '```'
      echo ""
   done
} > "$REPORT_DIR/汇总报告.md"

echo "汇总报告已生成: $REPORT_DIR/汇总报告.md"
