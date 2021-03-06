package main

import (
	"github.com/beaker/client/api"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newTaskCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task <command>",
		Short: "Manage tasks",
	}
	cmd.AddCommand(newTaskInspectCommand())
	cmd.AddCommand(newTaskLogsCommand())
	return cmd
}

func newTaskInspectCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect <task...>",
		Short: "Display detailed information about one or more tasks",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var tasks []*api.Task
			for _, id := range args {
				info, err := beaker.Task(id).Get(ctx)
				if err != nil {
					return err
				}

				tasks = append(tasks, info)
			}
			return printJSON(tasks)
		},
	}
}

func newTaskLogsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "logs <task>",
		Short: "Fetch logs for the most recent execution of a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			task, err := beaker.Task(args[0]).Get(ctx)
			if err != nil {
				return err
			}

			if len(task.Executions) == 0 {
				return errors.Errorf("task has no executions")
			}

			// Most recent execution is last.
			return printExecutionLogs(task.Executions[len(task.Executions)-1].ID)
		},
	}
}
