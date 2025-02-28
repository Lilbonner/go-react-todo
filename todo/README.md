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
<img width="1136" alt="Снимок экрана 2025-02-28 в 16 35 35" src="https://github.com/user-attachments/assets/27c6e0c2-7751-4dc8-89f3-52c3bd8d680d" />

- deleteTask: Удаляет задачу с подтверждением.
<img width="1136" alt="Снимок экрана 2025-02-28 в 16 37 00" src="https://github.com/user-attachments/assets/716ae33a-d06d-4037-8738-bf921c75c587" />

- toggleTask: Переключает статус задачи.
<img width="1136" alt="Снимок экрана 2025-02-28 в 16 35 42" src="https://github.com/user-attachments/assets/2338cf5a-ed7d-4f36-982c-e92168fa7b32" />



  

`Backend`
#### Task: Модель задачи.

Функции:

- AddTask: Добавляет задачу.
<img width="1136" alt="Снимок экрана 2025-02-28 в 16 16 08" src="https://github.com/user-attachments/assets/ee945127-463b-487a-9f7f-bc9288512f77" />

- GetTasks: Возвращает список задач.
<img width="1136" alt="Снимок экрана 2025-02-28 в 16 17 43" src="https://github.com/user-attachments/assets/25c4d403-eadd-4f7d-b716-492ce55493e8" />

- DeleteTask: Удаляет задачу.
<img width="1136" alt="Снимок экрана 2025-02-28 в 16 28 55" src="https://github.com/user-attachments/assets/bac446fd-d492-48ed-9c11-725870f7a119" />

`Лицензия
MIT`

Автор:
`Daniyar`
#### [https://github.com/Lilbonner]
