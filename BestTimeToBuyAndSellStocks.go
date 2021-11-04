package main

func maxProfit(prices []int) int {
    min_price := prices[0]
    max_profit := 0


    for i, _ := range(prices) {
        if prices[i] < min_price {
            min_price = prices[i]
        }

        if prices[i] - min_price > max_profit {
            max_profit = prices[i] - min_price
        }
   }

    return max_profit
}
