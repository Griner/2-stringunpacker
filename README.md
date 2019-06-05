# 2-stringunpacker

Homework #2
String unpacker

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)
* "qwe\4\5" => "qwe45" (*)
* "qwe\45" => "qwe44444" (*)
* "qwe\\5" => "qwe\\\\\" (*)