- 類似のコードを C で実装したらコンパイルエラーなのか気になった

# Link

- Rustのライフタイム注釈は「いつ」「なぜ」必要なのか #初心者向け - Qiita
- https://qiita.com/plotter/items/693b7a9d3a756f82441d

```Rust
// 2つの文字列参照を受け取り、printして、第一引数だけを返す関数
fn func1(x: &str, y: &str) -> &str {
    println!("{x} {y}");
    x
}

// 2つの文字列参照を受け取り、printして、第二引数だけを返す関数
fn func2(x: &str, y: &str) -> &str {
    println!("{x} {y}");
    y
}
```

```Rust
fn main() {
    let s1: &str;
    let s2: &str;
    let hello = "Hello".to_string();
    {
        let world = "World!".to_string();
        s1 = func1(&hello, &world);      // <-- (1)
        s2 = func2(&hello, &world);      // <-- (2)
    }                                    // <-- (3)
    println!("{}", s1);                  // <-- (4)
    println!("{}", s2);                  // <-- (5)
}
```
