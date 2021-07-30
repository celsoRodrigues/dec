package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// DecOptions stores information regarding the decryption assets
type DecOptions struct {
	configFlags *genericclioptions.ConfigFlags

	genericclioptions.IOStreams
	credits bool
}

// NewDecOptions provides an instance of DecOptions with default values
func NewDecOptions(streams genericclioptions.IOStreams) *DecOptions {
	return &DecOptions{
		configFlags: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

func NewCmdDec(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewDecOptions(streams)

	cmd := &cobra.Command{
		Use:          "dec [path/filename]",
		Short:        "decrypt the available secrets file",
		Example:      "",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {

			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&o.credits, "credits", o.credits, "testflag")
	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

func (o *DecOptions) Run() error {

	var files []string
	f, err := o.readSecretsDir(".")
	if err != nil {
		fmt.Println(err)
	}

	for _, x := range f {
		files = append(files, x.Name())

	}

	prompt := promptui.Select{
		Label: "Select Day",
		Items: files,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	fmt.Printf("You choose %q\n", result)

	return nil
}

func (o *DecOptions) readSecretsDir(d string) ([]fs.FileInfo, error) {

	abs, err := filepath.Abs(d)
	if err != nil {
		return []fs.FileInfo{}, err
	}

	files, err := ioutil.ReadDir(abs)
	if err != nil {
		return []fs.FileInfo{}, err
	}
	return files, nil
}
