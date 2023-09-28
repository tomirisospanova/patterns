package main

import "fmt"

type Observer interface {
    Update(message string)
}
type Messenger struct {
    observers []Observer
}

func (m *Messenger) Subscribe(observer Observer) {
    m.observers = append(m.observers, observer)
}

func (m *Messenger) Unsubscribe(observer Observer) {
    for i, obs := range m.observers {
        if obs == observer {
            m.observers = append(m.observers[:i], m.observers[i+1:]...)
            break
        }
    }
}

func (m *Messenger) NotifyObservers(message string) {
    for _, observer := range m.observers {
        observer.Update(message)
    }
}
type User struct {
    name string
}

func NewUser(name string) *User {
    return &User{name: name}
}

func (u *User) Update(message string) {
    fmt.Printf("User %s received message: %s\n", u.name, message)
}
func main() {
    // Создаем мессенджер
    messenger := Messenger{}

    // Создаем пользователей
    user1 := NewUser("Make")
    user2 := NewUser("Sake")
    user3 := NewUser("Take")

    // Подписываем пользователей на мессенджер
    messenger.Subscribe(user1)
    messenger.Subscribe(user2)
    messenger.Subscribe(user3)

    // Мессенджер отправляет сообщение всем подписанным пользователям
    messenger.NotifyObservers("Hello, everyone!")

    // Отписываем пользователя 2 от мессенджера
    messenger.Unsubscribe(user2)

    // Мессенджер отправляет еще одно сообщение
    messenger.NotifyObservers("This is a new message!")
}
