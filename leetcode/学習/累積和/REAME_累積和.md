### 累積和とは

事前に配列の総和を計算しておく事で任意の区間の和をO(1)で計算できるアルゴリズム


- 累積和の作り方
  - 0 から始まる
- 計算方法
  - i から jは 累積和の sum_array[j + 1] - sum_array[i]で求められる

- ページ

https://interviewcat.dev/p/coding-interviewcat/array-string-prefix-sum

- 累積和をする目的を深掘り
  - https://chatgpt.com/c/68f2474c-85e4-8323-b050-15c814aea220

### 例題

正の整数のリストnumsに対してK個の連続する整数の和の最大値を求めよ。

- 使わないケース
  - ループでもいける？
    - i から i + kを足したかずをmaxに入れて更新するとか
  - 計算量が O(n × k)になる
- 累積和
  - 累積和を作る
  - j+1 - iをループしていって最大の値を取得する
  - 計算量 O(n)