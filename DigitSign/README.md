# DigitSign

Distributed Lab course

Завдання полягало у власній реалізації алгоритму цифрового підпису. За основу було взято алгоритм RSA. В якості геш-функції використовується власна реалізація SHA-1.

Було реалізовано наступні функції:
+ KeyGen() -> pb, pv, n *big.Int -- повертає значення e, d, n, де (e, n) публічний ключ, а (d, n) особистий ключ. Публічний ключ також може виступати особистим, а особистий - публічним
+ SignFile(filePath string, pv, n *big.Int) -> s *big.Int -- на вхід приймає шлях до підписуваного файлу та приватний ключ (пару e та n); повертає підпис файлу у big.Int
+ SignMessage(filePath string, pv, n *big.Int) -> s *big.Int -- аналогічно до SignFile, але замість шляху до файлу приймає повідомлення, яке має бути підписане
+ VerifyFile(filePath string, signature, pb, n *big.Int) -> bool -- на вхід приймає шлях до підписаного файлу, його сигнатуру та публічний ключ (пару d та n); повертає "Sign verified!" якщо підпис валідний, в іншому випадку - "Sign was corrupted!"
+ VerifyMessage(message string, signature, pb, n *big.Int) -> bool -- аналогічно до VerifyFile тільки для повідомлень

Для запуску можна завантажити вже зібраний файл з [releases](https://github.com/Zaptoss/Cryptography-for-Developers/releases/tag/digitsign) або завантажити весь вміст DigitSign та виконати go run.

Програма підтримує наступні команди:
+ kg - key generation; this command generate p and q, one of this can used as private key and other as public key, and n - modul
+ sf - sign file; it take your private key, n as hex and file path as string for sign file
+ vf - verify file; it take signature, private key and n as hex and file path as string for verify file
+ sm - sign message; it take your private key, n as hex and message as string for sign file
+ vf - verify message; it take signature, private key and n as hex and file path as string for verify message
+ h - help options
+ q - quit

Приклад генерації ключів:
![image](https://github.com/Zaptoss/Cryptography-for-Developers/assets/57064743/1816a783-f1a3-4568-9958-fd8c54870608)

Приклад підпису та перевірки підпису файлу:
![image](https://github.com/Zaptoss/Cryptography-for-Developers/assets/57064743/f403f32b-ef7b-476a-8864-b541fa8b8f05)

Приклад підпису та перевірки підпису повідомлення(сигнатура була спеціально пошкоджена):
![image](https://github.com/Zaptoss/Cryptography-for-Developers/assets/57064743/d9738b23-dd1c-4a4c-9ea9-2a60c2a0b2eb)
