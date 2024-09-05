package entities

import "time"

type Entity interface {
    Save() error
    Validate() error
}

type Habit struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Streak int64 `json:"streak"`
}

func (u *Habit) Save() error {
    // logic to save User to the database
    return nil
}

func (u *Habit) Validate() error {
    // validation logic for User
    return nil
}

type Contribution struct {
	Id        string `json:"id"`
	TimeSpent int64  `json:"timeSpent"`
	HabitId   int64 `json:"habitId"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func (p *Contribution) Save() error {
    // logic to save Contribution to the database
    return nil
}

func (p *Contribution) Validate() error {
    // validation logic for Contribution
    return nil
}