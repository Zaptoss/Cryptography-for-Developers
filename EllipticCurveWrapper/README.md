# EllipticCurveWrapper

Distributed Lab course

Реалізовано бібліотеку-обгортку для бібліотеки crypto/elliptic, яка працює з еліптичними кривими, для більш зручного використання.

Бібліотека працює з кривою: P224

Перелік рералізованих функцій та структур:
+ ECPoint - структура з двома полями X та Y типу big.Int
+ BasePointGGen\() -> ECPoint - Генерує випадкову точку G
+ ECPointGen\(x, y *big.Int) -> ECPoint - Створює точку з переданими координатами
+ IsOnCurveCheck(point ECPoint) -> bool - Перевіряє належність точки до кривої
+ AddECPoint\(a, b ECPoint) -> ECPoint - Складає дві точки
+ DoubleECPoint\(a ECPoint) -> ECPoint - Подвоює точку\(множить на 2)
+ ScalarMult\(k *big.Int, a ECPoint) -> ECPoint - Множить точку на скаляр
+ ECPoint2String(point ECPoint) -> string - Перетворює точку на її текстове представлення
+ String2ECPoint(s string) -> ECPoint - Перетворює текстове представлення на точку
+ PrintECPoint(point ECPoint) - Виводить координати точки у термінал

Для запуску коду треба завантажити теку EllipticCurveWrapper та виконати в терміналі команду go run .

На екрані можна бачити 2 точки та перевірку їх рівності\(тестовий приклад з завдання)

Для власного використання у файлі main.go у функції main можна прибрати зайве та використовувати.

Наразі існує проблема із: go get github.com/Zaptoss/Cryptography-for-Developers/EllipticCurveWrapper; а тому такий метод не пропонується. Дана проблема буде вирішена найближчим часом. Всі функції були перевірені та працюють коректно.
