package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/celsoRodrigues/dec/pkg/conf"
	"github.com/celsoRodrigues/dec/pkg/mycrypt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// DecOptions stores information regarding the decryption assets, and I take advantage
//of genericclioptions to have access to native kubecl flags

type DecOptions struct {
	dir        string           //directory where secrets live
	file       string           //file to decrypt
	passphrase string           //key passphrase
	args       []string         //arguments passed to cmdline
	config     conf.ViperConfig //config object
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

	//getting the cluster environment from the working directory path
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	pathSlice := strings.Split(path, "/")
	clusterEnv := strings.ToUpper(strings.ReplaceAll(pathSlice[len(pathSlice)-3], "-", "_"))
	fmt.Println(clusterEnv)

	//get the passphrase assuming the environment variable containing the phrase is GPG_PASSPHRASE_ENV
	//for example GPG_PASSPHRASE_DEV_EUW2
	envName := fmt.Sprintf("GPG_PASSPHRASE_%s", clusterEnv)
	phrase := os.Getenv(envName)
	if len(phrase) < 1 {
		fmt.Printf("environment variable %s is not set ", envName)
		os.Exit(1)
	}
	o.passphrase = phrase

	//check if file not passed as the first argument, if so, use the first arg as the file to decrypt
	if len(o.args) < 1 {

		o.showPrompt()

		p := filepath.Join(o.dir, o.file)
		fmt.Println(p)

		f, err := os.Open(filepath.Join(o.dir, o.file))
		if err != nil {
			log.Fatal(err)
		}
		secbyte, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal("asdasd", err)
		}
		fmt.Println("the file is", string(secbyte))

		mcrypt := mycrypt.NewEncWithOptins("/home/celso/", o.passphrase, "/home/celso/.gnupg/secring.gpg", "/home/celso/.gnupg/pubring.gpg")
		res, err := mcrypt.Dec(string(secbyte))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("result:", string(res))

	} else {
		o.file = os.Args[1]
		fmt.Printf("You choose %q\n", o.file)

		f, err := os.Open(filepath.Join(o.dir, o.file))
		if err != nil {
			log.Fatal(err)
		}

		secbyte, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal("asdasd", err)
		}
		mcrypt := mycrypt.NewEncWithOptins("/home/celso/", o.passphrase, "/home/celso/.gnupg/secring.gpg", "/home/celso/.gnupg/pubring.gpg")
		res, err := mcrypt.Dec(string(secbyte))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("result:", string(res))
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

func (o *DecOptions) showPrompt() error {
	var files []string
	f, err := o.readSecretsDir(o.dir)
	if err != nil {
		return err

	}

	for _, x := range f {
		files = append(files, x.Name())

	}

	o.prompt(files)
	return nil
}
