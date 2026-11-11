package main

import "fmt"

type JobPost struct {
	title string
}

func NewJobPost(title string) *JobPost {
	return &JobPost{
		title: title,
	}
}

func (jp *JobPost) Title() string {
	return jp.title
}

type Observer interface {
	OnJobPosted(job *JobPost)
}

type JobSeeker struct {
	name string
}

var _ Observer = &JobSeeker{}

func NewJobSeeker(name string) *JobSeeker {
	return &JobSeeker{
		name: name,
	}
}

func (js *JobSeeker) OnJobPosted(job *JobPost) {
	fmt.Printf("Hi %s! New job posted: %s\n", js.name, job.Title())
}

type Observable interface {
	Notify(jobPosting *JobPost)
	Attach(observer Observer)
	AddJob(jobPosting *JobPost)
}

type EmploymentAgency struct {
	observers []Observer
}

var _ Observable = &EmploymentAgency{}

func NewEmploymentAgency() *EmploymentAgency {
	return &EmploymentAgency{
		observers: make([]Observer, 0),
	}
}

func (ea *EmploymentAgency) Notify(jobPosting *JobPost) {
	for _, observer := range ea.observers {
		observer.OnJobPosted(jobPosting)
	}
}

func (ea *EmploymentAgency) Attach(observer Observer) {
	ea.observers = append(ea.observers, observer)
}

func (ea *EmploymentAgency) AddJob(jobPosting *JobPost) {
	ea.Notify(jobPosting)
}

func main() {
	johnDoe := NewJobSeeker("John Doe")
	janeDoe := NewJobSeeker("Jane Doe")

	jobPostings := NewEmploymentAgency()
	jobPostings.Attach(johnDoe)
	jobPostings.Attach(janeDoe)

	jobPostings.AddJob(NewJobPost("Software Engineer"))
}
