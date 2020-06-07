package project

import (
	"context"
	"fmt"
	"github.com/winchci/winch/v2/actions"
)

type Manager struct {
	actions *actions.Manager
}

func NewManager() *Manager {
	return &Manager{
		actions: actions.NewManager(),
	}
}

func (m *Manager) Execute(ctx context.Context, p *Project) error {
	err := m.preloadActions(ctx, p)
	if err != nil {
		return err
	}

	err = m.executeSteps(ctx, p.Steps)
	if err != nil {
		return err
	}

	for _, job := range p.Jobs {
		err = m.executeJob(ctx, job)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) executeJob(ctx context.Context, j *Job) error {
	err := m.executeSteps(ctx, j.Steps)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) ExecuteJob(ctx context.Context, p *Project, jobName string) error {
	if j, ok := p.Jobs[jobName]; ok {
		err := m.preloadActions(ctx, p)
		if err != nil {
			return err
		}


		return m.executeJob(ctx, j)
	}

	return fmt.Errorf("project does not contain a jb named %s", jobName)
}

func (m *Manager) executeSteps(ctx context.Context, steps []*Step) error {
	for _, step := range steps {
		err := m.executeStep(ctx, step)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) executeStep(ctx context.Context, step *Step) error {
	fmt.Println("executing step ", step.Name, step.Use, step.Run)
	return nil
}

func (m *Manager) preloadActions(ctx context.Context, p *Project) error {
	err := m.preloadActionsForSteps(ctx, p.Steps)
	if err != nil {
		return err
	}

	for _, job := range p.Jobs {
		err = m.preloadActionsForSteps(ctx, job.Steps)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) preloadActionsForSteps(ctx context.Context, steps []*Step) error {
	for _, step := range steps {
		if len(step.Use) > 0 {
			_, err := m.actions.Load(ctx, actions.ParseActionRef(step.Use))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
