package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/celsoRodrigues/dec/pkg/conf"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// DecOptions stores information regarding the decryption assets, and I take advantage
//of genericclioptions to have access to native kubecl flags

type DecOptions struct {
	dir    string   //directory where secrets live
	file   string   //file to decrypt
	args   []string //arguments passed to cmdline
	config conf.ViperConfig
}

// NewDecOptions provides an instance of DecOptions with default values
func NewDecOptions() *DecOptions {
	return &DecOptions{}
}

//NewCmdDec returns the cobra cmd and sets the flags
func NewCmdDec(config conf.ViperConfig) *cobra.Command {
	o := NewDecOptions()
	var secretsDir string

	cmd := &cobra.Command{
		Use:          "dec [filename]",
		Short:        "decrypt the available secrets file",
		Example:      "",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {

			o.args = args
			o.config = config

			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}
	if config.IsSet("secretsdir") {
		secretsDir = config.GetString("secretsdir")

	} else {
		secretsDir = "../secrets"
	}
	cmd.Flags().StringVar(&o.dir, "d", fmt.Sprintf("%v", secretsDir), "secrets directory")
	return cmd
}

// Run is the principal function that drives this plugin
func (o *DecOptions) Run() error {

	if len(o.args) < 1 {

		var files []string
		f, err := o.readSecretsDir(o.dir)
		if err != nil {
			return err

		}

		for _, x := range f {
			files = append(files, x.Name())

		}

		o.prompt(files)

	} else {
		fmt.Printf("You choose %q\n", os.Args[1])
	}

	return nil
}

func (o *DecOptions) readSecretsDir(d string) ([]fs.FileInfo, error) {

	abs, err := filepath.Abs(d)

	if err != nil {
		return []fs.FileInfo{}, err
	}

	if _, err := os.Stat(abs); os.IsNotExist(err) {
		return []fs.FileInfo{}, fmt.Errorf("secrets directory not found in %s", abs)
	}

	files, err := ioutil.ReadDir(abs)
	if err != nil {
		return []fs.FileInfo{}, err
	}
	return files, nil
}

func (o *DecOptions) prompt(files []string) error {
	var err error
	prompt := promptui.Select{
		Label: "Select Your secrets file",
		Items: files,
	}

	_, o.file, err = prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	fmt.Printf("You choose %q\n", o.file)

	return nil
}