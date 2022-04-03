package main

import "fmt"

// 887. 鸡蛋掉落
// 给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。

// 已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。

// 每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。

// 请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？

// 示例 1：

// 输入：k = 1, n = 2
// 输出：2
// 解释：
// 鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。
// 否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。
// 如果它没碎，那么肯定能得出 f = 2 。
// 因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。
// 示例 2：

// 输入：k = 2, n = 6
// 输出：3
// 示例 3：

// 输入：k = 3, n = 14
// 输出：4

// 提示：

// 1 <= k <= 100
// 1 <= n <= 104

func superEggDrop1(k int, n int) int {
	//f(t,k)=1+f(t−1,k−1)+f(t−1,k)
	//边界条件为：当 t≥1 的时候 f(t, 1) = t f(t,1)=t，当 k≥1 时，f(1, k) = 1。
	//操作次数是一定不会超过楼层数的，因此 t<n

	if n == 1 {
		return 1
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}

	for i := 1; i <= k; i++ {

		f[1][i] = 1
	}

	ret := -1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			f[i][j] = 1 + f[i-1][j-1] + f[i-1][j]
		}

		if f[i][k] >= n {
			ret = i
			break
		}
	}
	return ret
}

func superEggDrop(k int, n int) int {

	// K ，不断增大 T ，计算出来的 F 的个数已经超过了 N + 1 时
	t := 1

	for calcF(k, t) < n+1 {
		t++
	}

	return t
}

func calcF(k int, t int) int {

	//如果只有 1 个蛋，或只有 1 次机会时，只可以确定出 T + 1 个 F
	//其他情况时，递归。【蛋碎了减 1 个，机会减 1 次】 + 【蛋没碎，机会减 1 次】
	if k == 1 || t == 1 {
		return t + 1
	}
	return calcF(k-1, t-1) + calcF(k, t-1)

}
func main() {

	k := 2
	n := 6

	fmt.Println(superEggDrop(k, n))
	fmt.Println(superEggDrop1(k, n))
}
