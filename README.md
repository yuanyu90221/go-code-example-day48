# go-code-problem-day48

## 題目描述

將給定的 linked list 向右 shift k 次, 每次 shift 時, 每個節點均向右位移一格, 而結尾節點會變成新的頭節點. 最後請 return shift 過後的 linked list.

Time complexity 的要求為 O(n)
Space complexity 的要求為 O(1)
Constraints:

0 <= k < maximum possible integer
Example 1:

Input: Linked List = 1 -> 2 -> 3 -> 4 -> 5, k = 2
Output: 4 -> 5 -> 1 -> 2 -> 3
Example 2:

Input: Linked List = 1 -> 2 -> 3, k = 3
Output: 1 -> 2 -> 3
Example 3:

Input: Linked List = 1 -> 2 -> 3, k = 4
Output: 3 -> 1 -> 2

## 題目分析

給定一個linkedList L向右 shift k次 每次shift時 每個結點都向右位移一格

最後找出shift過後的 linked list

要求要 時間複雜度 O(n)
空間複雜度O(1)

觀察可以發現

假設 linked list長度為 n

shift k次後的 head

就是 原本linked list的第n-k +1 個元素 

也就是只要先找出 linked list的長度

走n-k步 然後把 原head指標 指到 n-k元素 並且把 n-k的next指向原本的head即可以

## 時間複雜度分析：

找出 linkedlist 長度 需要n step

走n-k步 需要 n-k step

把 head跟 n-k的 next值交換 需要 3個 step

總共 2n-k +3 step  所以是 O(n)

## 空間複雜度分析：

整個演算法裡

需要記住 目前的指標 還有 head的指標

因此是 定值 所以 假設是 c 所以是 O(1)