# DigitSign

Distributed Lab course

Завдання полягало у власній реалізації алгоритму цифрового підпису. За основу було взято алгоритм RSA. В якості геш-функції використовується власна реалізація SHA-1.

Було реалізовано наступні функції:
+ KeyGen() -> pb, pv, n *big.Int -- повертає значення e, d, n, де (e, n) публічний ключ, а (d, n) особистий ключ. Публічний ключ також може виступати особистим, а особистий - публічним
+ SignFile(filePath string, pv, n *big.Int) -> s *big.Int -- на вхід приймає шлях до підписуваного файлу та приватний ключ (пару e та n); повертає підпис файлу у big.Int
+ SignMessage(filePath string, pv, n *big.Int) -> s *big.Int -- аналогічно до SignFile, але замість шляху до файлу приймає повідомлення, яке має бути підписане
+ VerifyFile(filePath string, signature, pb, n *big.Int) -> bool -- на вхід приймає шлях до підписаного файлу, його сигнатуру та публічний ключ (пару d та n); повертає "Sign verified!" якщо підпис валідний, в іншому випадку - "Sign was corrupted!"
+ VerifyMessage(message string, signature, pb, n *big.Int) -> bool -- аналогічно до VerifyFile тільки для повідомлень

[releases](https://github.com/Zaptoss/Cryptography-for-Developers/releases/tag/digitsign)
