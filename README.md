# elliptic_curve

Запуск програми go run main.go

Мною була використана бібліотека crypto/elliptic для роботи з еліптичними кривими, далі буде описана робота головних функцій.

1.CreateCurve - функція яка створює певний вид(залежно від того який ми подамо в вигляді рядка)еліптичної кривої і повертає об'єкт типу інтерфейсу elliptic.Curve, дана бібліотека підтримує 4 види еліптичних кривих(P224, P256, P384, P521)

2.BasePointGGet в залежності від того яку еліптичну криву ми подамо як аргумент для функції повертає базову точку цієї еліптичної кривої

Варто зазначити, що в багатьох функціях я додав додатковий аргумент на вхід c elliptic.Curve, так як є різні еліптичні криві і результат додавання одинакових точок в різних еліптичних кривих дасть різний результат, також цей інтефрейс оперує в багатьох методах які потрібні для виконання наприклад додавання, скалярного множення і іншого, тому я вважаю за доцільне додавання цього окремого аргументу.

3.ScalarMult яка використовує метод ScalarMult інтерфейсу elliptic.Curve множить точку на число, також варто зазначити що для використання бібліотечного методу потрібно скаляр перетворити з *big.Int в байти.

4.Для серіалізації і десеріалізації я вирішив використати json, тобто в випадку функції ECPointToString я вирішив перетворювати точку в рядок за допомогою json.Marshal і відповідно за допомогою json.Unmarshal в функції StringToECPoint я перетворюю рядок в точку.

В функції main виконується розв'язування рівняння яке було подане в методичних вказівках, а також перевіряється правильність роботи функій, перша з яких перетворює точку в рядок а інша яка перетворює рядок в точку(мається на увазі в тип ECPoint)