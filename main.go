package main // import "code.nkcmr.net/mempwgen"

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const nwords = 2
const nextra = 2

var symbols = []string{
	"#", "@", "-", "_",
}

func randset() ([]uint64, error) {
	s := make([]uint64, nwords+nextra)
	for i := range s {
		var max *big.Int
		switch i {
		case 0, 1:
			max = big.NewInt(int64(len(words)))
		case 2:
			max = big.NewInt(int64(1000))
		case 3:
			max = big.NewInt(int64(len(symbols)))
		}
		ri, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		s[i] = ri.Uint64()
	}
	return s, nil
}

func main() {
	_ = mainCommand().Execute()
}

func runWithError(f func(*cobra.Command, []string) error) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := f(cmd, args); err != nil {
			fmt.Fprintf(os.Stderr, "mempwgen: error: %s\n", err.Error())
			os.Exit(1)
		}
	}
}

func mainCommand() *cobra.Command {
	var a args
	cmd := &cobra.Command{
		Use:   "mempwgen",
		Short: "generate more memorable secure passwords",
		Run: runWithError(func(cmd *cobra.Command, args []string) error {
			if err := a.validate(); err != nil {
				return errors.Wrap(err, "invalid arguments")
			}
			passwords, err := mempwgen(a)
			if err != nil {
				return errors.WithStack(err)
			}
			for _, pw := range passwords {
				fmt.Println(pw)
			}
			return nil
		}),
	}
	cmd.Flags().IntVarP(&a.nPasswords, "count", "n", 20, "number of passwords to generate")
	cmd.Flags().IntVarP(&a.minLen, "min-len", "m", 20, "minimum password length")
	cmd.Flags().IntVarP(&a.maxLen, "max-len", "M", math.MaxInt, "maximum password length")
	return cmd
}

type args struct {
	nPasswords     int
	minLen, maxLen int
}

func (a args) validate() error {
	if a.nPasswords < 1 {
		return fmt.Errorf("count must be >= 1")
	}
	if a.maxLen < 8 {
		return fmt.Errorf("max-len must be >= 8")
	}
	if a.minLen > 35 {
		// this is because the dictionary has only so many words that are longer
		// and since we use a "brute force" approach to generate the passwords
		// it can start becoming much more work to generate the passwords
		return fmt.Errorf("min-len must be <= 35")
	}
	if a.maxLen <= a.minLen {
		return fmt.Errorf("max-len must be > min-len")
	}
	return nil
}

func mempwgen(a args) ([]string, error) {
	passwords := []string{}
	for len(passwords) < a.nPasswords {
		minLen := a.minLen
		maxLen := a.maxLen
		const nwords = 2
		const idxint = nwords
		const idxsym = idxint + 1
		s, err := randset()
		if err != nil {
			return nil, fmt.Errorf("failed to generate random set")
		}
		selectedWords := [2]string{}
		for i := range selectedWords {
			selectedWords[i] = words[s[i]]
		}
		randint := s[idxint]
		if randint == 666 {
			// ew...
			continue
		}
		sym := symbols[s[idxsym]]
		pw := selectedWords[0] + strconv.FormatInt(int64(randint), 10) + sym + selectedWords[1]
		if pwlen := len(pw); pwlen >= minLen && pwlen <= maxLen {
			passwords = append(passwords, pw)
		}
	}
	return passwords, nil
}
