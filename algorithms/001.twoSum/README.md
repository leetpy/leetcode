## 说明
题中说了存在且只有一组解，并且返回的是索引（不是值），这点不要搞错了。

## 思路
这题思路比较简单, 用 map 即可搞定。 遍历数组, 假设当前值为 A, 索引为 a，目标值为 target, 规则如下：

- 如果当前值 A 是 map 的 某个 key, 说明找到解了，直接返回；
- 如果当前值 A 不是 map 的 key, 则在 map中保存 map[target-A] = a;