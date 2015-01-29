package ci

type Scheduler interface {
	Run(BuildStep) <-chan Status
}
