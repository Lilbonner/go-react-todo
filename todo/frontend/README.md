## To-Do List App
#### Простое приложение для управления задачами. Создавайте, отмечайте выполненными и удаляйте задачи с поддержкой приоритетов и сроков.

Функции
- ✅ Добавление задач (название, дата, время, приоритет)

- ✅ Отметка задач как выполненных

- ✅ Удаление задач с подтверждением

- ✅ Автосохранение между запусками

- ✅ Сортировка по приоритету

### Технологии
- Frontend: Vue 3, TypeScript, SweetAlert2

- Backend: Go, Wails

- Хранение данных: JSON

### Установка
Установите зависимости:

- Go (1.20+)

- Node.js (18+)

- Wails CLI:


``` bash 
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Клонируйте и запустите:

```bash
git clone https://github.com/your-username/todo-list-app.git
cd todo-list-app/frontend
npm install
cd ..
wails dev 
```
### Компоненты
`Frontend`
#### App.vue: Основной компонент с формой и списком задач.

Функции:

- addTask: Добавляет задачу.

- deleteTask: Удаляет задачу с подтверждением.

- toggleTask: Переключает статус задачи.

`Backend`
#### Task: Модель задачи.

Функции:

- AddTask: Добавляет задачу.

- GetTasks: Возвращает список задач.

- DeleteTask: Удаляет задачу.

`Лицензия
MIT`

Автор:
`Daniyar`
#### [https://github.com/Lilbonner]