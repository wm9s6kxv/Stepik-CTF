# Решение python2-vs-python3

По стандарту мы получаем такой код
```python
import random

print("Guess the number in the range from -10000000 to 10000000000")

while True:
    a = random.randint(1, 100)
    
    b = input(">> ")

    if a == b:
        decoded_binary = "Тут будет флаг"
        print("Good! Flag:", decoded_binary)
        break
    else:
        print("No")
```

Также есть информация, что нужно почитать про различие функций. 

Находим информацию 

https://stackoverflow.com/questions/4960208/python-2-7-getting-user-input-and-manipulating-as-string-without-quotations

В итоге понимаем, что в python2 функция input равносильна eval(input()) в python3. 


Зная исходный код - вводим `a`, в результате в переменной `b` оказывается значение переменной `a` 